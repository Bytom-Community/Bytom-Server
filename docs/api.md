# wallet接口规范


## Overview

| Protocol | URL | Format | Version |
|-------|:------------:|:------------:|:------------:|
| gRPC |  http://localhost:9888 | Protobuf| V1 |
| HTTP |http://localhost:9889 | Json | V1 |

正确请求返回 example
``` bash
curl -d '' http://127.0.0.1:9888/v1/state
{"status":"OK"}
```
错误请求返回 example
``` bash
curl -d '{"alias":"alice", "password":"alice"}' http://localhost:9888/v1/create-key
{"error":"duplicate key alias","code":2}
```
## API

#### 查询钱包资产列表
| URL | Method | Version |
|-------|:------------:|:------------:|
| /v1/list-assets | POST | V1 |

**Parameters:**
``` json
{
    account_id string
}
```
**Returns:**
``` base
assets: [
    {
        id string
    }
]
```

#### 查询资产余额
| URL | Method | Version |
|-------|:------------:|:------------:|
| /v1/list-balances | POST | V1 |

**Parameters:**
``` json
{
    account_id string
    asset_id   string
}
```

**Returns:**
``` base
assets: [
    {
        amount int
    }
]
```


#### 查询资产交易记录
| URL | Method | Version |
|-------|:------------:|:------------:|
| /v1/list-balances | POST | V1 |

**Parameters:**
``` json
{
    account_id string
    asset_id   string
}
```

**Returns:**
``` base
transactions: [
    {
        "tx_id":        string
        "block_time":   int64
        "block_hash":   string
        "block_height": int
        "block_index":  int
        "block_transactions_count": int
        "inputs": [
            {
                "type":     string
                "asset_id": string
                "asset_alias": string
                "amount":   int64
                "address":  string
                "spent_output_id": string
                "account_id":      string
                "account_alias":   string
            }
        ],
        "outputs": [
            {
                "type":     string
                "id":       string
                "position": int
                "asset_id": string
                "asset_alias": string
                "amount":      int64
                "account_id":  string
                "account_alias": string
                "address":     string
            }
        ],
        "status_fail": bool
    }
]
```


#### 发送转帐交易
| URL | Method | Version |
|-------|:------------:|:------------:|
| /v1/submit-transaction | POST | V1 |

**Parameters:**
``` json
{
    raw_transaction string
}
```

**Returns:**
``` base
id: {
    tx_id string
}
```

#### 预估gas费用
| URL | Method | Version |
|-------|:------------:|:------------:|
| /v1/estimate-transaction-gas | POST | V1 |

**Parameters:**
``` json
{
    account_id  string
    asset_id    string
    address     string
    amount      int64
}
```

**Returns:**
``` base
gas: {
    total_neu   int64
    storage_neu int64
    vm_neu      int64
}
```


