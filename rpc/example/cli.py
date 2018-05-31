#!/usr/bin/python

from __future__ import print_function

import grpc

import rpc_pb2 
import rpc_pb2_grpc


def run():
    channel = grpc.insecure_channel('localhost:9889')
    stub = rpc_pb2_grpc.ApiServiceStub(channel)
    response = stub.GetState(rpc_pb2.NonParamsRequest())
    print("client received: ", response)


if __name__ == '__main__':
    run()
