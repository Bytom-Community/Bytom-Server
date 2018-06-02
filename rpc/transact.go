package rpc

import (
	"context"
	"encoding/json"
	"strings"
	"time"

	"github.com/bytom/account"
	"github.com/bytom/blockchain/txbuilder"
	"github.com/bytom/errors"
	"github.com/bytom/net/http/reqid"
)

var (
	defaultTxTTL    = 5 * time.Minute
	defaultBaseRate = float64(100000)
)

func (s *ApiService) BuildTranstraction(ctx context.Context, req *rpcpb.BuildTranstractionRequest) (*rpcpb.BuildTranstractionResponse, error) {
	subctx := reqid.NewSubContext(ctx, reqid.New())

	buildReqs := &BuildRequest{
		TTL:       req.Ttl,
		TimeRange: req.TimeRange,
	}

	actions := []map[string]interface{}{}
	for i := range req.Actions {
		m := map[string]interface{}{}
		m["type"] = req.Actions[i].Type
		m["asset_id"] = req.Actions[i].AssetID
		m["asset_alias"] = req.Actions[i].AssetAlias
		m["account_id"] = req.Actions[i].AccountID
		m["account_alias"] = req.Actions[i].AccountAlias
		actions = append(actions, m)
	}
	buildReqs.Actions = actions

	tmpl, err := s.buildSingle(subctx, buildReqs)
	if err != nil {
		return nil, err
	}

	return &rpcpb.BuildTranstractionResponse{Template: tmpl}, nil
}

func (s *ApiService) buildSingle(ctx context.Context, req *BuildRequest) (*txbuilder.Template, error) {
	err := s.filterAliases(ctx, req)
	if err != nil {
		return nil, err
	}

	if onlyHaveSpendActions(req) {
		return nil, errors.New("transaction only contain spend actions, didn't have output actions")
	}

	spendActions := []txbuilder.Action{}
	actions := make([]txbuilder.Action, 0, len(req.Actions))
	for i, act := range req.Actions {
		typ, ok := act["type"].(string)
		if !ok {
			return nil, errors.WithDetailf(errBadActionType, "no action type provided on action %d", i)
		}
		decoder, ok := s.actionDecoder(typ)
		if !ok {
			return nil, errors.WithDetailf(errBadActionType, "unknown action type %q on action %d", typ, i)
		}

		// Remarshal to JSON, the action may have been modified when we
		// filtered aliases.
		b, err := json.Marshal(act)
		if err != nil {
			return nil, err
		}
		action, err := decoder(b)
		if err != nil {
			return nil, errors.WithDetailf(errBadAction, "%s on action %d", err.Error(), i)
		}

		if typ == "spend_account" {
			spendActions = append(spendActions, action)
		} else {
			actions = append(actions, action)
		}
	}
	actions = append(account.MergeSpendAction(spendActions), actions...)

	ttl := req.TTL.Duration
	if ttl == 0 {
		ttl = defaultTxTTL
	}
	maxTime := time.Now().Add(ttl)

	tpl, err := txbuilder.Build(ctx, req.Tx, actions, maxTime, req.TimeRange)
	if errors.Root(err) == txbuilder.ErrAction {
		// append each of the inner errors contained in the data.
		var Errs string
		for _, innerErr := range errors.Data(err)["actions"].([]error) {
			Errs = Errs + "<" + innerErr.Error() + ">"
		}
		err = errors.New(err.Error() + "-" + Errs)
	}
	if err != nil {
		return nil, err
	}

	// ensure null is never returned for signing instructions
	if tpl.SigningInstructions == nil {
		tpl.SigningInstructions = []*txbuilder.SigningInstruction{}
	}
	return tpl, nil
}

func onlyHaveSpendActions(req *BuildRequest) bool {
	count := 0
	for _, m := range req.Actions {
		if actionType := m["type"].(string); strings.HasPrefix(actionType, "spend") {
			count++
		}
	}

	return count == len(req.Actions)
}

func (s *ApiService) actionDecoder(action string) (func([]byte) (txbuilder.Action, error), bool) {
	var decoder func([]byte) (txbuilder.Action, error)
	switch action {
	case "control_address":
		decoder = txbuilder.DecodeControlAddressAction
	case "control_program":
		decoder = txbuilder.DecodeControlProgramAction
	case "control_receiver":
		decoder = txbuilder.DecodeControlReceiverAction
	case "issue":
		decoder = s.wallet.AssetReg.DecodeIssueAction
	case "retire":
		decoder = txbuilder.DecodeRetireAction
	case "spend_account":
		decoder = s.wallet.AccountMgr.DecodeSpendAction
	case "spend_account_unspent_output":
		decoder = s.wallet.AccountMgr.DecodeSpendUTXOAction
	default:
		return nil, false
	}
	return decoder, true
}
