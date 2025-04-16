// Code generated by protoc-gen-ts_proto. DO NOT EDIT.
// versions:
//   protoc-gen-ts_proto  v2.7.0
//   protoc               v5.29.3
// source: sync/management/service.proto
import {
    type CallOptions,
    ChannelCredentials,
    Client,
    type ClientOptions,
    type ClientReadableStream,
    type ClientUnaryCall,
    Metadata,
    type ServiceError,
    type UntypedServiceImplementation,
    type handleServerStreamingCall,
    type handleUnaryCall,
    makeGenericClientConstructor,
} from "@grpc/grpc-js";
import {
    CreateLeaseRequest,
    CreateLeaseResponse,
    DropLeaseRequest,
    DropLeaseResponse,
    KeepLeaseRequest,
    KeepLeaseResponse,
} from "./lease";
import {
    LockRequest,
    LockResponse,
    RLockRequest,
    RLockResponse,
    RUnlockRequest,
    RUnlockResponse,
    UnlockRequest,
    UnlockResponse,
} from "./lock";
import { GetPeerNodesRequest, GetPeerNodesResponse } from "./peer_nodes";

export type ServiceService = typeof ServiceService;
export const ServiceService = {
    createLease: {
        path: "/freym.sync.management.Service/CreateLease",
        requestStream: false,
        responseStream: false,
        requestSerialize: (value: CreateLeaseRequest) =>
            Buffer.from(CreateLeaseRequest.encode(value).finish()),
        requestDeserialize: (value: Buffer) => CreateLeaseRequest.decode(value),
        responseSerialize: (value: CreateLeaseResponse) =>
            Buffer.from(CreateLeaseResponse.encode(value).finish()),
        responseDeserialize: (value: Buffer) => CreateLeaseResponse.decode(value),
    },
    keepLease: {
        path: "/freym.sync.management.Service/KeepLease",
        requestStream: false,
        responseStream: false,
        requestSerialize: (value: KeepLeaseRequest) =>
            Buffer.from(KeepLeaseRequest.encode(value).finish()),
        requestDeserialize: (value: Buffer) => KeepLeaseRequest.decode(value),
        responseSerialize: (value: KeepLeaseResponse) =>
            Buffer.from(KeepLeaseResponse.encode(value).finish()),
        responseDeserialize: (value: Buffer) => KeepLeaseResponse.decode(value),
    },
    dropLease: {
        path: "/freym.sync.management.Service/DropLease",
        requestStream: false,
        responseStream: false,
        requestSerialize: (value: DropLeaseRequest) =>
            Buffer.from(DropLeaseRequest.encode(value).finish()),
        requestDeserialize: (value: Buffer) => DropLeaseRequest.decode(value),
        responseSerialize: (value: DropLeaseResponse) =>
            Buffer.from(DropLeaseResponse.encode(value).finish()),
        responseDeserialize: (value: Buffer) => DropLeaseResponse.decode(value),
    },
    getPeerNodes: {
        path: "/freym.sync.management.Service/GetPeerNodes",
        requestStream: false,
        responseStream: true,
        requestSerialize: (value: GetPeerNodesRequest) =>
            Buffer.from(GetPeerNodesRequest.encode(value).finish()),
        requestDeserialize: (value: Buffer) => GetPeerNodesRequest.decode(value),
        responseSerialize: (value: GetPeerNodesResponse) =>
            Buffer.from(GetPeerNodesResponse.encode(value).finish()),
        responseDeserialize: (value: Buffer) => GetPeerNodesResponse.decode(value),
    },
    lock: {
        path: "/freym.sync.management.Service/Lock",
        requestStream: false,
        responseStream: false,
        requestSerialize: (value: LockRequest) => Buffer.from(LockRequest.encode(value).finish()),
        requestDeserialize: (value: Buffer) => LockRequest.decode(value),
        responseSerialize: (value: LockResponse) =>
            Buffer.from(LockResponse.encode(value).finish()),
        responseDeserialize: (value: Buffer) => LockResponse.decode(value),
    },
    unlock: {
        path: "/freym.sync.management.Service/Unlock",
        requestStream: false,
        responseStream: false,
        requestSerialize: (value: UnlockRequest) =>
            Buffer.from(UnlockRequest.encode(value).finish()),
        requestDeserialize: (value: Buffer) => UnlockRequest.decode(value),
        responseSerialize: (value: UnlockResponse) =>
            Buffer.from(UnlockResponse.encode(value).finish()),
        responseDeserialize: (value: Buffer) => UnlockResponse.decode(value),
    },
    rLock: {
        path: "/freym.sync.management.Service/RLock",
        requestStream: false,
        responseStream: false,
        requestSerialize: (value: RLockRequest) => Buffer.from(RLockRequest.encode(value).finish()),
        requestDeserialize: (value: Buffer) => RLockRequest.decode(value),
        responseSerialize: (value: RLockResponse) =>
            Buffer.from(RLockResponse.encode(value).finish()),
        responseDeserialize: (value: Buffer) => RLockResponse.decode(value),
    },
    rUnlock: {
        path: "/freym.sync.management.Service/RUnlock",
        requestStream: false,
        responseStream: false,
        requestSerialize: (value: RUnlockRequest) =>
            Buffer.from(RUnlockRequest.encode(value).finish()),
        requestDeserialize: (value: Buffer) => RUnlockRequest.decode(value),
        responseSerialize: (value: RUnlockResponse) =>
            Buffer.from(RUnlockResponse.encode(value).finish()),
        responseDeserialize: (value: Buffer) => RUnlockResponse.decode(value),
    },
} as const;

export interface ServiceServer extends UntypedServiceImplementation {
    createLease: handleUnaryCall<CreateLeaseRequest, CreateLeaseResponse>;
    keepLease: handleUnaryCall<KeepLeaseRequest, KeepLeaseResponse>;
    dropLease: handleUnaryCall<DropLeaseRequest, DropLeaseResponse>;
    getPeerNodes: handleServerStreamingCall<GetPeerNodesRequest, GetPeerNodesResponse>;
    lock: handleUnaryCall<LockRequest, LockResponse>;
    unlock: handleUnaryCall<UnlockRequest, UnlockResponse>;
    rLock: handleUnaryCall<RLockRequest, RLockResponse>;
    rUnlock: handleUnaryCall<RUnlockRequest, RUnlockResponse>;
}

export interface ServiceClient extends Client {
    createLease(
        request: CreateLeaseRequest,
        callback: (error: ServiceError | null, response: CreateLeaseResponse) => void
    ): ClientUnaryCall;
    createLease(
        request: CreateLeaseRequest,
        metadata: Metadata,
        callback: (error: ServiceError | null, response: CreateLeaseResponse) => void
    ): ClientUnaryCall;
    createLease(
        request: CreateLeaseRequest,
        metadata: Metadata,
        options: Partial<CallOptions>,
        callback: (error: ServiceError | null, response: CreateLeaseResponse) => void
    ): ClientUnaryCall;
    keepLease(
        request: KeepLeaseRequest,
        callback: (error: ServiceError | null, response: KeepLeaseResponse) => void
    ): ClientUnaryCall;
    keepLease(
        request: KeepLeaseRequest,
        metadata: Metadata,
        callback: (error: ServiceError | null, response: KeepLeaseResponse) => void
    ): ClientUnaryCall;
    keepLease(
        request: KeepLeaseRequest,
        metadata: Metadata,
        options: Partial<CallOptions>,
        callback: (error: ServiceError | null, response: KeepLeaseResponse) => void
    ): ClientUnaryCall;
    dropLease(
        request: DropLeaseRequest,
        callback: (error: ServiceError | null, response: DropLeaseResponse) => void
    ): ClientUnaryCall;
    dropLease(
        request: DropLeaseRequest,
        metadata: Metadata,
        callback: (error: ServiceError | null, response: DropLeaseResponse) => void
    ): ClientUnaryCall;
    dropLease(
        request: DropLeaseRequest,
        metadata: Metadata,
        options: Partial<CallOptions>,
        callback: (error: ServiceError | null, response: DropLeaseResponse) => void
    ): ClientUnaryCall;
    getPeerNodes(
        request: GetPeerNodesRequest,
        options?: Partial<CallOptions>
    ): ClientReadableStream<GetPeerNodesResponse>;
    getPeerNodes(
        request: GetPeerNodesRequest,
        metadata?: Metadata,
        options?: Partial<CallOptions>
    ): ClientReadableStream<GetPeerNodesResponse>;
    lock(
        request: LockRequest,
        callback: (error: ServiceError | null, response: LockResponse) => void
    ): ClientUnaryCall;
    lock(
        request: LockRequest,
        metadata: Metadata,
        callback: (error: ServiceError | null, response: LockResponse) => void
    ): ClientUnaryCall;
    lock(
        request: LockRequest,
        metadata: Metadata,
        options: Partial<CallOptions>,
        callback: (error: ServiceError | null, response: LockResponse) => void
    ): ClientUnaryCall;
    unlock(
        request: UnlockRequest,
        callback: (error: ServiceError | null, response: UnlockResponse) => void
    ): ClientUnaryCall;
    unlock(
        request: UnlockRequest,
        metadata: Metadata,
        callback: (error: ServiceError | null, response: UnlockResponse) => void
    ): ClientUnaryCall;
    unlock(
        request: UnlockRequest,
        metadata: Metadata,
        options: Partial<CallOptions>,
        callback: (error: ServiceError | null, response: UnlockResponse) => void
    ): ClientUnaryCall;
    rLock(
        request: RLockRequest,
        callback: (error: ServiceError | null, response: RLockResponse) => void
    ): ClientUnaryCall;
    rLock(
        request: RLockRequest,
        metadata: Metadata,
        callback: (error: ServiceError | null, response: RLockResponse) => void
    ): ClientUnaryCall;
    rLock(
        request: RLockRequest,
        metadata: Metadata,
        options: Partial<CallOptions>,
        callback: (error: ServiceError | null, response: RLockResponse) => void
    ): ClientUnaryCall;
    rUnlock(
        request: RUnlockRequest,
        callback: (error: ServiceError | null, response: RUnlockResponse) => void
    ): ClientUnaryCall;
    rUnlock(
        request: RUnlockRequest,
        metadata: Metadata,
        callback: (error: ServiceError | null, response: RUnlockResponse) => void
    ): ClientUnaryCall;
    rUnlock(
        request: RUnlockRequest,
        metadata: Metadata,
        options: Partial<CallOptions>,
        callback: (error: ServiceError | null, response: RUnlockResponse) => void
    ): ClientUnaryCall;
}

export const ServiceClient = makeGenericClientConstructor(
    ServiceService,
    "freym.sync.management.Service"
) as unknown as {
    new (
        address: string,
        credentials: ChannelCredentials,
        options?: Partial<ClientOptions>
    ): ServiceClient;
    service: typeof ServiceService;
    serviceName: string;
};
