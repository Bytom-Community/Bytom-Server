package rpcpb;

import static io.grpc.MethodDescriptor.generateFullMethodName;
import static io.grpc.stub.ClientCalls.asyncBidiStreamingCall;
import static io.grpc.stub.ClientCalls.asyncClientStreamingCall;
import static io.grpc.stub.ClientCalls.asyncServerStreamingCall;
import static io.grpc.stub.ClientCalls.asyncUnaryCall;
import static io.grpc.stub.ClientCalls.blockingServerStreamingCall;
import static io.grpc.stub.ClientCalls.blockingUnaryCall;
import static io.grpc.stub.ClientCalls.futureUnaryCall;
import static io.grpc.stub.ServerCalls.asyncBidiStreamingCall;
import static io.grpc.stub.ServerCalls.asyncClientStreamingCall;
import static io.grpc.stub.ServerCalls.asyncServerStreamingCall;
import static io.grpc.stub.ServerCalls.asyncUnaryCall;
import static io.grpc.stub.ServerCalls.asyncUnimplementedStreamingCall;
import static io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall;

/**
 */
@javax.annotation.Generated(
    value = "by gRPC proto compiler (version 1.13.0-SNAPSHOT)",
    comments = "Source: rpc/pb/rpc.proto")
public final class ApiServiceGrpc {

  private ApiServiceGrpc() {}

  public static final String SERVICE_NAME = "rpcpb.ApiService";

  // Static method descriptors that strictly reflect the proto.
  @io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
  @java.lang.Deprecated // Use {@link #getGetStateMethod()} instead. 
  public static final io.grpc.MethodDescriptor<rpcpb.Rpc.NonParamsRequest,
      rpcpb.Rpc.GetStateResponse> METHOD_GET_STATE = getGetStateMethodHelper();

  private static volatile io.grpc.MethodDescriptor<rpcpb.Rpc.NonParamsRequest,
      rpcpb.Rpc.GetStateResponse> getGetStateMethod;

  @io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
  public static io.grpc.MethodDescriptor<rpcpb.Rpc.NonParamsRequest,
      rpcpb.Rpc.GetStateResponse> getGetStateMethod() {
    return getGetStateMethodHelper();
  }

  private static io.grpc.MethodDescriptor<rpcpb.Rpc.NonParamsRequest,
      rpcpb.Rpc.GetStateResponse> getGetStateMethodHelper() {
    io.grpc.MethodDescriptor<rpcpb.Rpc.NonParamsRequest, rpcpb.Rpc.GetStateResponse> getGetStateMethod;
    if ((getGetStateMethod = ApiServiceGrpc.getGetStateMethod) == null) {
      synchronized (ApiServiceGrpc.class) {
        if ((getGetStateMethod = ApiServiceGrpc.getGetStateMethod) == null) {
          ApiServiceGrpc.getGetStateMethod = getGetStateMethod = 
              io.grpc.MethodDescriptor.<rpcpb.Rpc.NonParamsRequest, rpcpb.Rpc.GetStateResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(
                  "rpcpb.ApiService", "GetState"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
                  rpcpb.Rpc.NonParamsRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
                  rpcpb.Rpc.GetStateResponse.getDefaultInstance()))
                  .build();
          }
        }
     }
     return getGetStateMethod;
  }
  @io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
  @java.lang.Deprecated // Use {@link #getCreateKeyMethod()} instead. 
  public static final io.grpc.MethodDescriptor<rpcpb.Rpc.CreateKeyRequest,
      rpcpb.Rpc.CreateKeyResponse> METHOD_CREATE_KEY = getCreateKeyMethodHelper();

  private static volatile io.grpc.MethodDescriptor<rpcpb.Rpc.CreateKeyRequest,
      rpcpb.Rpc.CreateKeyResponse> getCreateKeyMethod;

  @io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
  public static io.grpc.MethodDescriptor<rpcpb.Rpc.CreateKeyRequest,
      rpcpb.Rpc.CreateKeyResponse> getCreateKeyMethod() {
    return getCreateKeyMethodHelper();
  }

  private static io.grpc.MethodDescriptor<rpcpb.Rpc.CreateKeyRequest,
      rpcpb.Rpc.CreateKeyResponse> getCreateKeyMethodHelper() {
    io.grpc.MethodDescriptor<rpcpb.Rpc.CreateKeyRequest, rpcpb.Rpc.CreateKeyResponse> getCreateKeyMethod;
    if ((getCreateKeyMethod = ApiServiceGrpc.getCreateKeyMethod) == null) {
      synchronized (ApiServiceGrpc.class) {
        if ((getCreateKeyMethod = ApiServiceGrpc.getCreateKeyMethod) == null) {
          ApiServiceGrpc.getCreateKeyMethod = getCreateKeyMethod = 
              io.grpc.MethodDescriptor.<rpcpb.Rpc.CreateKeyRequest, rpcpb.Rpc.CreateKeyResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(
                  "rpcpb.ApiService", "CreateKey"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
                  rpcpb.Rpc.CreateKeyRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
                  rpcpb.Rpc.CreateKeyResponse.getDefaultInstance()))
                  .build();
          }
        }
     }
     return getCreateKeyMethod;
  }
  @io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
  @java.lang.Deprecated // Use {@link #getListKeyMethod()} instead. 
  public static final io.grpc.MethodDescriptor<rpcpb.Rpc.NonParamsRequest,
      rpcpb.Rpc.ListKeyResponse> METHOD_LIST_KEY = getListKeyMethodHelper();

  private static volatile io.grpc.MethodDescriptor<rpcpb.Rpc.NonParamsRequest,
      rpcpb.Rpc.ListKeyResponse> getListKeyMethod;

  @io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
  public static io.grpc.MethodDescriptor<rpcpb.Rpc.NonParamsRequest,
      rpcpb.Rpc.ListKeyResponse> getListKeyMethod() {
    return getListKeyMethodHelper();
  }

  private static io.grpc.MethodDescriptor<rpcpb.Rpc.NonParamsRequest,
      rpcpb.Rpc.ListKeyResponse> getListKeyMethodHelper() {
    io.grpc.MethodDescriptor<rpcpb.Rpc.NonParamsRequest, rpcpb.Rpc.ListKeyResponse> getListKeyMethod;
    if ((getListKeyMethod = ApiServiceGrpc.getListKeyMethod) == null) {
      synchronized (ApiServiceGrpc.class) {
        if ((getListKeyMethod = ApiServiceGrpc.getListKeyMethod) == null) {
          ApiServiceGrpc.getListKeyMethod = getListKeyMethod = 
              io.grpc.MethodDescriptor.<rpcpb.Rpc.NonParamsRequest, rpcpb.Rpc.ListKeyResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(
                  "rpcpb.ApiService", "ListKey"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
                  rpcpb.Rpc.NonParamsRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
                  rpcpb.Rpc.ListKeyResponse.getDefaultInstance()))
                  .build();
          }
        }
     }
     return getListKeyMethod;
  }
  @io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
  @java.lang.Deprecated // Use {@link #getDeleteKeyMethod()} instead. 
  public static final io.grpc.MethodDescriptor<rpcpb.Rpc.DeleteKeyRequest,
      com.google.protobuf.Empty> METHOD_DELETE_KEY = getDeleteKeyMethodHelper();

  private static volatile io.grpc.MethodDescriptor<rpcpb.Rpc.DeleteKeyRequest,
      com.google.protobuf.Empty> getDeleteKeyMethod;

  @io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
  public static io.grpc.MethodDescriptor<rpcpb.Rpc.DeleteKeyRequest,
      com.google.protobuf.Empty> getDeleteKeyMethod() {
    return getDeleteKeyMethodHelper();
  }

  private static io.grpc.MethodDescriptor<rpcpb.Rpc.DeleteKeyRequest,
      com.google.protobuf.Empty> getDeleteKeyMethodHelper() {
    io.grpc.MethodDescriptor<rpcpb.Rpc.DeleteKeyRequest, com.google.protobuf.Empty> getDeleteKeyMethod;
    if ((getDeleteKeyMethod = ApiServiceGrpc.getDeleteKeyMethod) == null) {
      synchronized (ApiServiceGrpc.class) {
        if ((getDeleteKeyMethod = ApiServiceGrpc.getDeleteKeyMethod) == null) {
          ApiServiceGrpc.getDeleteKeyMethod = getDeleteKeyMethod = 
              io.grpc.MethodDescriptor.<rpcpb.Rpc.DeleteKeyRequest, com.google.protobuf.Empty>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(
                  "rpcpb.ApiService", "DeleteKey"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
                  rpcpb.Rpc.DeleteKeyRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
                  .build();
          }
        }
     }
     return getDeleteKeyMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static ApiServiceStub newStub(io.grpc.Channel channel) {
    return new ApiServiceStub(channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static ApiServiceBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    return new ApiServiceBlockingStub(channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static ApiServiceFutureStub newFutureStub(
      io.grpc.Channel channel) {
    return new ApiServiceFutureStub(channel);
  }

  /**
   */
  public static abstract class ApiServiceImplBase implements io.grpc.BindableService {

    /**
     * <pre>
     * state
     * </pre>
     */
    public void getState(rpcpb.Rpc.NonParamsRequest request,
        io.grpc.stub.StreamObserver<rpcpb.Rpc.GetStateResponse> responseObserver) {
      asyncUnimplementedUnaryCall(getGetStateMethodHelper(), responseObserver);
    }

    /**
     * <pre>
     * Key
     * </pre>
     */
    public void createKey(rpcpb.Rpc.CreateKeyRequest request,
        io.grpc.stub.StreamObserver<rpcpb.Rpc.CreateKeyResponse> responseObserver) {
      asyncUnimplementedUnaryCall(getCreateKeyMethodHelper(), responseObserver);
    }

    /**
     */
    public void listKey(rpcpb.Rpc.NonParamsRequest request,
        io.grpc.stub.StreamObserver<rpcpb.Rpc.ListKeyResponse> responseObserver) {
      asyncUnimplementedUnaryCall(getListKeyMethodHelper(), responseObserver);
    }

    /**
     */
    public void deleteKey(rpcpb.Rpc.DeleteKeyRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      asyncUnimplementedUnaryCall(getDeleteKeyMethodHelper(), responseObserver);
    }

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return io.grpc.ServerServiceDefinition.builder(getServiceDescriptor())
          .addMethod(
            getGetStateMethodHelper(),
            asyncUnaryCall(
              new MethodHandlers<
                rpcpb.Rpc.NonParamsRequest,
                rpcpb.Rpc.GetStateResponse>(
                  this, METHODID_GET_STATE)))
          .addMethod(
            getCreateKeyMethodHelper(),
            asyncUnaryCall(
              new MethodHandlers<
                rpcpb.Rpc.CreateKeyRequest,
                rpcpb.Rpc.CreateKeyResponse>(
                  this, METHODID_CREATE_KEY)))
          .addMethod(
            getListKeyMethodHelper(),
            asyncUnaryCall(
              new MethodHandlers<
                rpcpb.Rpc.NonParamsRequest,
                rpcpb.Rpc.ListKeyResponse>(
                  this, METHODID_LIST_KEY)))
          .addMethod(
            getDeleteKeyMethodHelper(),
            asyncUnaryCall(
              new MethodHandlers<
                rpcpb.Rpc.DeleteKeyRequest,
                com.google.protobuf.Empty>(
                  this, METHODID_DELETE_KEY)))
          .build();
    }
  }

  /**
   */
  public static final class ApiServiceStub extends io.grpc.stub.AbstractStub<ApiServiceStub> {
    private ApiServiceStub(io.grpc.Channel channel) {
      super(channel);
    }

    private ApiServiceStub(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected ApiServiceStub build(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      return new ApiServiceStub(channel, callOptions);
    }

    /**
     * <pre>
     * state
     * </pre>
     */
    public void getState(rpcpb.Rpc.NonParamsRequest request,
        io.grpc.stub.StreamObserver<rpcpb.Rpc.GetStateResponse> responseObserver) {
      asyncUnaryCall(
          getChannel().newCall(getGetStateMethodHelper(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * Key
     * </pre>
     */
    public void createKey(rpcpb.Rpc.CreateKeyRequest request,
        io.grpc.stub.StreamObserver<rpcpb.Rpc.CreateKeyResponse> responseObserver) {
      asyncUnaryCall(
          getChannel().newCall(getCreateKeyMethodHelper(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void listKey(rpcpb.Rpc.NonParamsRequest request,
        io.grpc.stub.StreamObserver<rpcpb.Rpc.ListKeyResponse> responseObserver) {
      asyncUnaryCall(
          getChannel().newCall(getListKeyMethodHelper(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void deleteKey(rpcpb.Rpc.DeleteKeyRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      asyncUnaryCall(
          getChannel().newCall(getDeleteKeyMethodHelper(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   */
  public static final class ApiServiceBlockingStub extends io.grpc.stub.AbstractStub<ApiServiceBlockingStub> {
    private ApiServiceBlockingStub(io.grpc.Channel channel) {
      super(channel);
    }

    private ApiServiceBlockingStub(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected ApiServiceBlockingStub build(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      return new ApiServiceBlockingStub(channel, callOptions);
    }

    /**
     * <pre>
     * state
     * </pre>
     */
    public rpcpb.Rpc.GetStateResponse getState(rpcpb.Rpc.NonParamsRequest request) {
      return blockingUnaryCall(
          getChannel(), getGetStateMethodHelper(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Key
     * </pre>
     */
    public rpcpb.Rpc.CreateKeyResponse createKey(rpcpb.Rpc.CreateKeyRequest request) {
      return blockingUnaryCall(
          getChannel(), getCreateKeyMethodHelper(), getCallOptions(), request);
    }

    /**
     */
    public rpcpb.Rpc.ListKeyResponse listKey(rpcpb.Rpc.NonParamsRequest request) {
      return blockingUnaryCall(
          getChannel(), getListKeyMethodHelper(), getCallOptions(), request);
    }

    /**
     */
    public com.google.protobuf.Empty deleteKey(rpcpb.Rpc.DeleteKeyRequest request) {
      return blockingUnaryCall(
          getChannel(), getDeleteKeyMethodHelper(), getCallOptions(), request);
    }
  }

  /**
   */
  public static final class ApiServiceFutureStub extends io.grpc.stub.AbstractStub<ApiServiceFutureStub> {
    private ApiServiceFutureStub(io.grpc.Channel channel) {
      super(channel);
    }

    private ApiServiceFutureStub(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected ApiServiceFutureStub build(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      return new ApiServiceFutureStub(channel, callOptions);
    }

    /**
     * <pre>
     * state
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<rpcpb.Rpc.GetStateResponse> getState(
        rpcpb.Rpc.NonParamsRequest request) {
      return futureUnaryCall(
          getChannel().newCall(getGetStateMethodHelper(), getCallOptions()), request);
    }

    /**
     * <pre>
     * Key
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<rpcpb.Rpc.CreateKeyResponse> createKey(
        rpcpb.Rpc.CreateKeyRequest request) {
      return futureUnaryCall(
          getChannel().newCall(getCreateKeyMethodHelper(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<rpcpb.Rpc.ListKeyResponse> listKey(
        rpcpb.Rpc.NonParamsRequest request) {
      return futureUnaryCall(
          getChannel().newCall(getListKeyMethodHelper(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.google.protobuf.Empty> deleteKey(
        rpcpb.Rpc.DeleteKeyRequest request) {
      return futureUnaryCall(
          getChannel().newCall(getDeleteKeyMethodHelper(), getCallOptions()), request);
    }
  }

  private static final int METHODID_GET_STATE = 0;
  private static final int METHODID_CREATE_KEY = 1;
  private static final int METHODID_LIST_KEY = 2;
  private static final int METHODID_DELETE_KEY = 3;

  private static final class MethodHandlers<Req, Resp> implements
      io.grpc.stub.ServerCalls.UnaryMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ServerStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ClientStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.BidiStreamingMethod<Req, Resp> {
    private final ApiServiceImplBase serviceImpl;
    private final int methodId;

    MethodHandlers(ApiServiceImplBase serviceImpl, int methodId) {
      this.serviceImpl = serviceImpl;
      this.methodId = methodId;
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public void invoke(Req request, io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        case METHODID_GET_STATE:
          serviceImpl.getState((rpcpb.Rpc.NonParamsRequest) request,
              (io.grpc.stub.StreamObserver<rpcpb.Rpc.GetStateResponse>) responseObserver);
          break;
        case METHODID_CREATE_KEY:
          serviceImpl.createKey((rpcpb.Rpc.CreateKeyRequest) request,
              (io.grpc.stub.StreamObserver<rpcpb.Rpc.CreateKeyResponse>) responseObserver);
          break;
        case METHODID_LIST_KEY:
          serviceImpl.listKey((rpcpb.Rpc.NonParamsRequest) request,
              (io.grpc.stub.StreamObserver<rpcpb.Rpc.ListKeyResponse>) responseObserver);
          break;
        case METHODID_DELETE_KEY:
          serviceImpl.deleteKey((rpcpb.Rpc.DeleteKeyRequest) request,
              (io.grpc.stub.StreamObserver<com.google.protobuf.Empty>) responseObserver);
          break;
        default:
          throw new AssertionError();
      }
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public io.grpc.stub.StreamObserver<Req> invoke(
        io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        default:
          throw new AssertionError();
      }
    }
  }

  private static volatile io.grpc.ServiceDescriptor serviceDescriptor;

  public static io.grpc.ServiceDescriptor getServiceDescriptor() {
    io.grpc.ServiceDescriptor result = serviceDescriptor;
    if (result == null) {
      synchronized (ApiServiceGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .addMethod(getGetStateMethodHelper())
              .addMethod(getCreateKeyMethodHelper())
              .addMethod(getListKeyMethodHelper())
              .addMethod(getDeleteKeyMethodHelper())
              .build();
        }
      }
    }
    return result;
  }
}
