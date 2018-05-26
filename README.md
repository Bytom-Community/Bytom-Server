Bytom RPC 开发者指南
===============
[![Build Status](https://travis-ci.org/Bytom/bytom.svg)](https://travis-ci.org/Bytom/bytom)
[![AGPL v3](https://img.shields.io/badge/license-AGPL%20v3-brightgreen.svg)](./LICENSE)

## 说明
bytom RPC版本将逐渐废弃掉api模块和bytomcli模块

目前sdk支持：golang、python、java、android

HTTP协议端口：9888

RPC协议端口：9889

## 安装依赖
#### 安装protobuf 3.4.0
[protobuf download](https://github.com/google/protobuf/releases)

#### 安装java和android版本插件
[安装protoc-gen-grpc-java](https://github.com/grpc/grpc-java/tree/master/compiler)

## 下载bytom
``` bash
$ git clone https://github.com/Bytom-Community/Bytom-Mobile-Wallet.git $GOPATH/src/github.com/bytom
$ cd $GOPATH/src/github.com/bytom
```


## 生成rpc server
生成rpc server文件：rpc.pb.go rpc.pb.gw.go
``` bash
$ cd rpc/pb
$ make clean; make
```

## 生成SDK
目前sdk支持 golang、python、java、android
``` bash
$ cd $GOPATH/src/github.com/bytom
$ make clean
$ make sdk
```

``` bash
 $ tree sdk
sdk
├── android_bytom
│   └── rpcpb
│       └── nano
│           └── ApiServiceGrpc.java
├── go_bytom
│   └── rpc
│       └── pb
│           └── rpc.pb.go
├── java_bytom
│   └── rpcpb
│       └── ApiServiceGrpc.java
└── py_bytom
    └── rpc
        └── pb
            └── rpc_pb2.py
```

## 安装编译bytomd
``` bash
$ make bytomd    # build bytomd
```

## 初始化并启动

```bash
$ cd ./cmd/bytomd
$ ./bytomd init --chain_id testnet
$ ./bytomd node -web.closed
```

## 使用SDK RPC协议创建/遍历/删除key
``` golang
package main

import (
	"google.golang.org/grpc"
	"log"
	"context"

	"github.com/bytom/rpc/pb"
)

func main() {
	address := "0.0.0.0:9889"

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	cli := rpcpb.NewApiServiceClient(conn)
	{
		resp, err := cli.GetState(context.Background(), &rpcpb.NonParamsRequest{})
		if err != nil {
			panic(err)
		}
		log.Println("GetState: ",resp.String())
	}
	var alias string = "alice"
	var password string = "alice"
	{
		resp, err := cli.CreateKey(context.Background(), &rpcpb.CreateKeyRequest{Alias: alias, Password: password})
		if err != nil {
			panic(err)
		}
		log.Println("CreateKey: ",resp.String())
	}
	{
		resp, err := cli.ListKey(context.Background(), &rpcpb.NonParamsRequest{})
		if err != nil {
			panic(err)
		}
		log.Println("ListKey: ",resp.GetXpub())
	}
	{
		_, err := cli.DeleteKey(context.Background(),
			&rpcpb.DeleteKeyRequest{Password:"alice",Xpub:"ba54fb0abe0d644b7d368936dcde425a268a8c3e3d90adeefdf0852aecb26a5a945e4800dc34712ec5393a4e65defa441361832c51d2f905d89b31d0f8b453c6"})
		if err != nil {
			panic(err)
		}
	}
}

```

## 使用http协议创建/遍历/删除key
``` bash
curl -d '{"alias":"alice", "password":"alice"}' http://localhost:9888/v1/create-key
curl -d '' http://localhost:9888/v1/list-keys
curl -d '{"password":"alice","xpub":"779a574304c4dcf61f9eb52a6eb4f6057a60c400190f2790d1af30caafd385424a4bcc8f87166dd32ec9366d04cf00b0e30930549fb9559db48945bf58de62d6"}' \
 http://localhost:9888/v1/delete-keys
```