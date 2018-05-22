all:
	protoc -I/usr/local/include -I. \
      -I${GOPATH}/src \
      -I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
      --go_out=plugins=grpc:. \
      --grpc-gateway_out=logtostderr=true:. \
      ./rpc.proto
clean:
	rm *.go
