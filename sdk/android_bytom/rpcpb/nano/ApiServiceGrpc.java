package rpcpb.nano;

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

import java.io.IOException;

/**
 */
@javax.annotation.Generated(
    value = "by gRPC proto compiler (version 1.13.0-SNAPSHOT)",
    comments = "Source: rpc/pb/rpc.proto")
public final class ApiServiceGrpc {

  private ApiServiceGrpc() {}

  public static final String SERVICE_NAME = "rpcpb.ApiService";

  // Static method descriptors that strictly reflect the proto.
  private static final int ARG_IN_METHOD_GET_STATE = 0;
  private static final int ARG_OUT_METHOD_GET_STATE = 1;
  @io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
  @java.lang.Deprecated // Use {@link #getGetStateMethod()} instead. 
  public static final io.grpc.MethodDescriptor<rpcpb.nano.Rpc.NonParamsRequest,
      rpcpb.nano.Rpc.GetStateResponse> METHOD_GET_STATE = getGetStateMethodHelper();

  private static volatile io.grpc.MethodDescriptor<rpcpb.nano.Rpc.NonParamsRequest,
      rpcpb.nano.Rpc.GetStateResponse> getGetStateMethod;

  @io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
  public static io.grpc.MethodDescriptor<rpcpb.nano.Rpc.NonParamsRequest,
      rpcpb.nano.Rpc.GetStateResponse> getGetStateMethod() {
    return getGetStateMethodHelper();
  }

  private static io.grpc.MethodDescriptor<rpcpb.nano.Rpc.NonParamsRequest,
      rpcpb.nano.Rpc.GetStateResponse> getGetStateMethodHelper() {
    io.grpc.MethodDescriptor<rpcpb.nano.Rpc.NonParamsRequest, rpcpb.nano.Rpc.GetStateResponse> getGetStateMethod;
    if ((getGetStateMethod = ApiServiceGrpc.getGetStateMethod) == null) {
      synchronized (ApiServiceGrpc.class) {
        if ((getGetStateMethod = ApiServiceGrpc.getGetStateMethod) == null) {
          ApiServiceGrpc.getGetStateMethod = getGetStateMethod = 
              io.grpc.MethodDescriptor.<rpcpb.nano.Rpc.NonParamsRequest, rpcpb.nano.Rpc.GetStateResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(
                  "rpcpb.ApiService", "GetState"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.nano.NanoUtils.<rpcpb.nano.Rpc.NonParamsRequest>marshaller(
                  new NanoFactory<rpcpb.nano.Rpc.NonParamsRequest>(ARG_IN_METHOD_GET_STATE)))
              .setResponseMarshaller(io.grpc.protobuf.nano.NanoUtils.<rpcpb.nano.Rpc.GetStateResponse>marshaller(
                  new NanoFactory<rpcpb.nano.Rpc.GetStateResponse>(ARG_OUT_METHOD_GET_STATE)))
              .build();
        }
      }
    }
    return getGetStateMethod;
  }
  private static final int ARG_IN_METHOD_CREATE_KEY = 2;
  private static final int ARG_OUT_METHOD_CREATE_KEY = 3;
  @io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
  @java.lang.Deprecated // Use {@link #getCreateKeyMethod()} instead. 
  public static final io.grpc.MethodDescriptor<rpcpb.nano.Rpc.CreateKeyRequest,
      rpcpb.nano.Rpc.CreateKeyResponse> METHOD_CREATE_KEY = getCreateKeyMethodHelper();

  private static volatile io.grpc.MethodDescriptor<rpcpb.nano.Rpc.CreateKeyRequest,
      rpcpb.nano.Rpc.CreateKeyResponse> getCreateKeyMethod;

  @io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
  public static io.grpc.MethodDescriptor<rpcpb.nano.Rpc.CreateKeyRequest,
      rpcpb.nano.Rpc.CreateKeyResponse> getCreateKeyMethod() {
    return getCreateKeyMethodHelper();
  }

  private static io.grpc.MethodDescriptor<rpcpb.nano.Rpc.CreateKeyRequest,
      rpcpb.nano.Rpc.CreateKeyResponse> getCreateKeyMethodHelper() {
    io.grpc.MethodDescriptor<rpcpb.nano.Rpc.CreateKeyRequest, rpcpb.nano.Rpc.CreateKeyResponse> getCreateKeyMethod;
    if ((getCreateKeyMethod = ApiServiceGrpc.getCreateKeyMethod) == null) {
      synchronized (ApiServiceGrpc.class) {
        if ((getCreateKeyMethod = ApiServiceGrpc.getCreateKeyMethod) == null) {
          ApiServiceGrpc.getCreateKeyMethod = getCreateKeyMethod = 
              io.grpc.MethodDescriptor.<rpcpb.nano.Rpc.CreateKeyRequest, rpcpb.nano.Rpc.CreateKeyResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(
                  "rpcpb.ApiService", "CreateKey"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.nano.NanoUtils.<rpcpb.nano.Rpc.CreateKeyRequest>marshaller(
                  new NanoFactory<rpcpb.nano.Rpc.CreateKeyRequest>(ARG_IN_METHOD_CREATE_KEY)))
              .setResponseMarshaller(io.grpc.protobuf.nano.NanoUtils.<rpcpb.nano.Rpc.CreateKeyResponse>marshaller(
                  new NanoFactory<rpcpb.nano.Rpc.CreateKeyResponse>(ARG_OUT_METHOD_CREATE_KEY)))
              .build();
        }
      }
    }
    return getCreateKeyMethod;
  }
  private static final int ARG_IN_METHOD_LIST_KEY = 4;
  private static final int ARG_OUT_METHOD_LIST_KEY = 5;
  @io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
  @java.lang.Deprecated // Use {@link #getListKeyMethod()} instead. 
  public static final io.grpc.MethodDescriptor<rpcpb.nano.Rpc.NonParamsRequest,
      rpcpb.nano.Rpc.ListKeyResponse> METHOD_LIST_KEY = getListKeyMethodHelper();

  private static volatile io.grpc.MethodDescriptor<rpcpb.nano.Rpc.NonParamsRequest,
      rpcpb.nano.Rpc.ListKeyResponse> getListKeyMethod;

  @io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
  public static io.grpc.MethodDescriptor<rpcpb.nano.Rpc.NonParamsRequest,
      rpcpb.nano.Rpc.ListKeyResponse> getListKeyMethod() {
    return getListKeyMethodHelper();
  }

  private static io.grpc.MethodDescriptor<rpcpb.nano.Rpc.NonParamsRequest,
      rpcpb.nano.Rpc.ListKeyResponse> getListKeyMethodHelper() {
    io.grpc.MethodDescriptor<rpcpb.nano.Rpc.NonParamsRequest, rpcpb.nano.Rpc.ListKeyResponse> getListKeyMethod;
    if ((getListKeyMethod = ApiServiceGrpc.getListKeyMethod) == null) {
      synchronized (ApiServiceGrpc.class) {
        if ((getListKeyMethod = ApiServiceGrpc.getListKeyMethod) == null) {
          ApiServiceGrpc.getListKeyMethod = getListKeyMethod = 
              io.grpc.MethodDescriptor.<rpcpb.nano.Rpc.NonParamsRequest, rpcpb.nano.Rpc.ListKeyResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(
                  "rpcpb.ApiService", "ListKey"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.nano.NanoUtils.<rpcpb.nano.Rpc.NonParamsRequest>marshaller(
                  new NanoFactory<rpcpb.nano.Rpc.NonParamsRequest>(ARG_IN_METHOD_LIST_KEY)))
              .setResponseMarshaller(io.grpc.protobuf.nano.NanoUtils.<rpcpb.nano.Rpc.ListKeyResponse>marshaller(
                  new NanoFactory<rpcpb.nano.Rpc.ListKeyResponse>(ARG_OUT_METHOD_LIST_KEY)))
              .build();
        }
      }
    }
    return getListKeyMethod;
  }
  private static final int ARG_IN_METHOD_DELETE_KEY = 6;
  private static final int ARG_OUT_METHOD_DELETE_KEY = 7;
  @io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
  @java.lang.Deprecated // Use {@link #getDeleteKeyMethod()} instead. 
  public static final io.grpc.MethodDescriptor<rpcpb.nano.Rpc.DeleteKeyRequest,
      com.google.protobuf.nano.Empty> METHOD_DELETE_KEY = getDeleteKeyMethodHelper();

  private static volatile io.grpc.MethodDescriptor<rpcpb.nano.Rpc.DeleteKeyRequest,
      com.google.protobuf.nano.Empty> getDeleteKeyMethod;

  @io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
  public static io.grpc.MethodDescriptor<rpcpb.nano.Rpc.DeleteKeyRequest,
      com.google.protobuf.nano.Empty> getDeleteKeyMethod() {
    return getDeleteKeyMethodHelper();
  }

  private static io.grpc.MethodDescriptor<rpcpb.nano.Rpc.DeleteKeyRequest,
      com.google.protobuf.nano.Empty> getDeleteKeyMethodHelper() {
    io.grpc.MethodDescriptor<rpcpb.nano.Rpc.DeleteKeyRequest, com.google.protobuf.nano.Empty> getDeleteKeyMethod;
    if ((getDeleteKeyMethod = ApiServiceGrpc.getDeleteKeyMethod) == null) {
      synchronized (ApiServiceGrpc.class) {
        if ((getDeleteKeyMethod = ApiServiceGrpc.getDeleteKeyMethod) == null) {
          ApiServiceGrpc.getDeleteKeyMethod = getDeleteKeyMethod = 
              io.grpc.MethodDescriptor.<rpcpb.nano.Rpc.DeleteKeyRequest, com.google.protobuf.nano.Empty>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(
                  "rpcpb.ApiService", "DeleteKey"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.nano.NanoUtils.<rpcpb.nano.Rpc.DeleteKeyRequest>marshaller(
                  new NanoFactory<rpcpb.nano.Rpc.DeleteKeyRequest>(ARG_IN_METHOD_DELETE_KEY)))
              .setResponseMarshaller(io.grpc.protobuf.nano.NanoUtils.<com.google.protobuf.nano.Empty>marshaller(
                  new NanoFactory<com.google.protobuf.nano.Empty>(ARG_OUT_METHOD_DELETE_KEY)))
              .build();
        }
      }
    }
    return getDeleteKeyMethod;
  }

  private static final class NanoFactory<T extends com.google.protobuf.nano.MessageNano>
      implements io.grpc.protobuf.nano.MessageNanoFactory<T> {
    private final int id;

    NanoFactory(int id) {
      this.id = id;
    }

    @java.lang.Override
    public T newInstance() {
      Object o;
      switch (id) {
      case ARG_IN_METHOD_GET_STATE:
        o = new rpcpb.nano.Rpc.NonParamsRequest();
        break;
      case ARG_OUT_METHOD_GET_STATE:
        o = new rpcpb.nano.Rpc.GetStateResponse();
        break;
      case ARG_IN_METHOD_CREATE_KEY:
        o = new rpcpb.nano.Rpc.CreateKeyRequest();
        break;
      case ARG_OUT_METHOD_CREATE_KEY:
        o = new rpcpb.nano.Rpc.CreateKeyResponse();
        break;
      case ARG_IN_METHOD_LIST_KEY:
        o = new rpcpb.nano.Rpc.NonParamsRequest();
        break;
      case ARG_OUT_METHOD_LIST_KEY:
        o = new rpcpb.nano.Rpc.ListKeyResponse();
        break;
      case ARG_IN_METHOD_DELETE_KEY:
        o = new rpcpb.nano.Rpc.DeleteKeyRequest();
        break;
      case ARG_OUT_METHOD_DELETE_KEY:
        o = new com.google.protobuf.nano.Empty();
        break;
      default:
        throw new AssertionError();
      }
      @java.lang.SuppressWarnings("unchecked")
      T t = (T) o;
      return t;
    }
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
    public void getState(rpcpb.nano.Rpc.NonParamsRequest request,
        io.grpc.stub.StreamObserver<rpcpb.nano.Rpc.GetStateResponse> responseObserver) {
      asyncUnimplementedUnaryCall(getGetStateMethodHelper(), responseObserver);
    }

    /**
     * <pre>
     * Key
     * </pre>
     */
    public void createKey(rpcpb.nano.Rpc.CreateKeyRequest request,
        io.grpc.stub.StreamObserver<rpcpb.nano.Rpc.CreateKeyResponse> responseObserver) {
      asyncUnimplementedUnaryCall(getCreateKeyMethodHelper(), responseObserver);
    }

    /**
     */
    public void listKey(rpcpb.nano.Rpc.NonParamsRequest request,
        io.grpc.stub.StreamObserver<rpcpb.nano.Rpc.ListKeyResponse> responseObserver) {
      asyncUnimplementedUnaryCall(getListKeyMethodHelper(), responseObserver);
    }

    /**
     */
    public void deleteKey(rpcpb.nano.Rpc.DeleteKeyRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.nano.Empty> responseObserver) {
      asyncUnimplementedUnaryCall(getDeleteKeyMethodHelper(), responseObserver);
    }

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return io.grpc.ServerServiceDefinition.builder(getServiceDescriptor())
          .addMethod(
            getGetStateMethodHelper(),
            asyncUnaryCall(
              new MethodHandlers<
                rpcpb.nano.Rpc.NonParamsRequest,
                rpcpb.nano.Rpc.GetStateResponse>(
                  this, METHODID_GET_STATE)))
          .addMethod(
            getCreateKeyMethodHelper(),
            asyncUnaryCall(
              new MethodHandlers<
                rpcpb.nano.Rpc.CreateKeyRequest,
                rpcpb.nano.Rpc.CreateKeyResponse>(
                  this, METHODID_CREATE_KEY)))
          .addMethod(
            getListKeyMethodHelper(),
            asyncUnaryCall(
              new MethodHandlers<
                rpcpb.nano.Rpc.NonParamsRequest,
                rpcpb.nano.Rpc.ListKeyResponse>(
                  this, METHODID_LIST_KEY)))
          .addMethod(
            getDeleteKeyMethodHelper(),
            asyncUnaryCall(
              new MethodHandlers<
                rpcpb.nano.Rpc.DeleteKeyRequest,
                com.google.protobuf.nano.Empty>(
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
    public void getState(rpcpb.nano.Rpc.NonParamsRequest request,
        io.grpc.stub.StreamObserver<rpcpb.nano.Rpc.GetStateResponse> responseObserver) {
      asyncUnaryCall(
          getChannel().newCall(getGetStateMethodHelper(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * Key
     * </pre>
     */
    public void createKey(rpcpb.nano.Rpc.CreateKeyRequest request,
        io.grpc.stub.StreamObserver<rpcpb.nano.Rpc.CreateKeyResponse> responseObserver) {
      asyncUnaryCall(
          getChannel().newCall(getCreateKeyMethodHelper(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void listKey(rpcpb.nano.Rpc.NonParamsRequest request,
        io.grpc.stub.StreamObserver<rpcpb.nano.Rpc.ListKeyResponse> responseObserver) {
      asyncUnaryCall(
          getChannel().newCall(getListKeyMethodHelper(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void deleteKey(rpcpb.nano.Rpc.DeleteKeyRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.nano.Empty> responseObserver) {
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
    public rpcpb.nano.Rpc.GetStateResponse getState(rpcpb.nano.Rpc.NonParamsRequest request) {
      return blockingUnaryCall(
          getChannel(), getGetStateMethodHelper(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Key
     * </pre>
     */
    public rpcpb.nano.Rpc.CreateKeyResponse createKey(rpcpb.nano.Rpc.CreateKeyRequest request) {
      return blockingUnaryCall(
          getChannel(), getCreateKeyMethodHelper(), getCallOptions(), request);
    }

    /**
     */
    public rpcpb.nano.Rpc.ListKeyResponse listKey(rpcpb.nano.Rpc.NonParamsRequest request) {
      return blockingUnaryCall(
          getChannel(), getListKeyMethodHelper(), getCallOptions(), request);
    }

    /**
     */
    public com.google.protobuf.nano.Empty deleteKey(rpcpb.nano.Rpc.DeleteKeyRequest request) {
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
    public com.google.common.util.concurrent.ListenableFuture<rpcpb.nano.Rpc.GetStateResponse> getState(
        rpcpb.nano.Rpc.NonParamsRequest request) {
      return futureUnaryCall(
          getChannel().newCall(getGetStateMethodHelper(), getCallOptions()), request);
    }

    /**
     * <pre>
     * Key
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<rpcpb.nano.Rpc.CreateKeyResponse> createKey(
        rpcpb.nano.Rpc.CreateKeyRequest request) {
      return futureUnaryCall(
          getChannel().newCall(getCreateKeyMethodHelper(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<rpcpb.nano.Rpc.ListKeyResponse> listKey(
        rpcpb.nano.Rpc.NonParamsRequest request) {
      return futureUnaryCall(
          getChannel().newCall(getListKeyMethodHelper(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.google.protobuf.nano.Empty> deleteKey(
        rpcpb.nano.Rpc.DeleteKeyRequest request) {
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
          serviceImpl.getState((rpcpb.nano.Rpc.NonParamsRequest) request,
              (io.grpc.stub.StreamObserver<rpcpb.nano.Rpc.GetStateResponse>) responseObserver);
          break;
        case METHODID_CREATE_KEY:
          serviceImpl.createKey((rpcpb.nano.Rpc.CreateKeyRequest) request,
              (io.grpc.stub.StreamObserver<rpcpb.nano.Rpc.CreateKeyResponse>) responseObserver);
          break;
        case METHODID_LIST_KEY:
          serviceImpl.listKey((rpcpb.nano.Rpc.NonParamsRequest) request,
              (io.grpc.stub.StreamObserver<rpcpb.nano.Rpc.ListKeyResponse>) responseObserver);
          break;
        case METHODID_DELETE_KEY:
          serviceImpl.deleteKey((rpcpb.nano.Rpc.DeleteKeyRequest) request,
              (io.grpc.stub.StreamObserver<com.google.protobuf.nano.Empty>) responseObserver);
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
