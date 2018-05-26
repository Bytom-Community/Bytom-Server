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
	//var alias string = "alice"
	//var password string = "alice"
	//{
	//	resp, err := cli.CreateKey(context.Background(), &rpcpb.CreateKeyRequest{Alias: alias, Password: password})
	//	if err != nil {
	//		panic(err)
	//	}
	//	log.Println("CreateKey: ",resp.String())
	//}
	//{
	//	resp, err := cli.ListKey(context.Background(), &rpcpb.NonParamsRequest{})
	//	if err != nil {
	//		panic(err)
	//	}
	//	log.Println("ListKey: ",resp.GetXpub())
	//}
	{
		_, err := cli.DeleteKey(context.Background(),
			&rpcpb.DeleteKeyRequest{Password:"alice",Xpub:"ba54fb0abe0d644b7d368936dcde425a268a8c3e3d90adeefdf0852aecb26a5a945e4800dc34712ec5393a4e65defa441361832c51d2f905d89b31d0f8b453c6"})
		if err != nil {
			panic(err)
		}
	}
}
