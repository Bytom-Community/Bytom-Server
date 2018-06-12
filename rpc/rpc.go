package rpc

import (
	"context"
	"net"
	"net/http"

	"github.com/bytom/chaincache"
	"github.com/bytom/rpc/pb"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	log "github.com/sirupsen/logrus"
	cmn "github.com/tendermint/tmlibs/common"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Rpc struct {
	rpcServer *grpc.Server
	apiListen string
	rpcListen string
}

func NewRpc(chainCache *chaincache.ChainCache) *Rpc {
	rpcServer := grpc.NewServer()
	rpc := &Rpc{
		rpcServer: rpcServer,
	}

	api := &ApiService{
		rpcServer:  rpc,
		chainCache: chainCache,
	}

	rpcpb.RegisterApiServiceServer(rpcServer, api)
	reflection.Register(rpcServer)
	return rpc
}

func (r *Rpc) Start(api_addr, rpc_addr string) error {
	r.apiListen = api_addr
	r.rpcListen = rpc_addr

	// rpc server
	go func() {
		log.WithField("RPC start on address:", rpc_addr).Info("Rpc listen")
		listener, err := net.Listen("tcp", rpc_addr)
		if err != nil {
			cmn.Exit(cmn.Fmt("RPC Failed to listen tcp port: %v, err:%v", rpc_addr, err))
		}
		if err = r.rpcServer.Serve(listener); err != nil {
			cmn.Exit(cmn.Fmt("RPC Failed to serve tcp port: %v, err:%v", rpc_addr, err))
		}
	}()

	// gateway server
	go func() {
		log.WithField("API start on address:", api_addr).Info("API listen")
		ctx := context.Background()
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()

		mux := runtime.NewServeMux()
		opts := []grpc.DialOption{grpc.WithInsecure()}
		err := rpcpb.RegisterApiServiceHandlerFromEndpoint(ctx, mux, rpc_addr, opts)
		if err != nil {
			cmn.Exit(cmn.Fmt("RPC Failed to register api service, err:%v", err))
		}

		err = http.ListenAndServe(api_addr, mux)
		if err != nil {
			cmn.Exit(cmn.Fmt("API Failed to serve tcp port: %v, err:%v", api_addr, err))
		}
	}()

	return nil
}

func (r *Rpc) Stop() {
	log.WithField("RPC stop on address:", r.rpcListen).Info("Rpc listen")
	log.WithField("API stop on address:", r.apiListen).Info("API listen")
	r.rpcServer.Stop()
}
