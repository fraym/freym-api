// Code generated by protoc-gen-ts_proto. DO NOT EDIT.
// versions:
//   protoc-gen-ts_proto  v2.6.1
//   protoc               v5.29.3
// source: deployments/management/service.proto
import {
    type CallOptions,
    ChannelCredentials,
    Client,
    type ClientOptions,
    type ClientUnaryCall,
    Metadata,
    type ServiceError,
    type UntypedServiceImplementation,
    type handleUnaryCall,
    makeGenericClientConstructor,
} from "@grpc/grpc-js";
import { ActivateDeploymentRequest, ActivateDeploymentResponse } from "./activate";
import { ConfirmDeploymentRequest, ConfirmDeploymentResponse } from "./confirm";
import { CreateDeploymentRequest, CreateDeploymentResponse } from "./create";
import { GetDeploymentRequest, GetDeploymentResponse } from "./deployment";
import {
    RollbackDeploymentByIdRequest,
    RollbackDeploymentRequest,
    RollbackDeploymentResponse,
} from "./rollback";

export type ServiceService = typeof ServiceService;
export const ServiceService = {
    createDeployment: {
        path: "/freym.deployments.management.Service/CreateDeployment",
        requestStream: false,
        responseStream: false,
        requestSerialize: (value: CreateDeploymentRequest) =>
            Buffer.from(CreateDeploymentRequest.encode(value).finish()),
        requestDeserialize: (value: Buffer) => CreateDeploymentRequest.decode(value),
        responseSerialize: (value: CreateDeploymentResponse) =>
            Buffer.from(CreateDeploymentResponse.encode(value).finish()),
        responseDeserialize: (value: Buffer) => CreateDeploymentResponse.decode(value),
    },
    activateDeployment: {
        path: "/freym.deployments.management.Service/ActivateDeployment",
        requestStream: false,
        responseStream: false,
        requestSerialize: (value: ActivateDeploymentRequest) =>
            Buffer.from(ActivateDeploymentRequest.encode(value).finish()),
        requestDeserialize: (value: Buffer) => ActivateDeploymentRequest.decode(value),
        responseSerialize: (value: ActivateDeploymentResponse) =>
            Buffer.from(ActivateDeploymentResponse.encode(value).finish()),
        responseDeserialize: (value: Buffer) => ActivateDeploymentResponse.decode(value),
    },
    confirmDeployment: {
        path: "/freym.deployments.management.Service/ConfirmDeployment",
        requestStream: false,
        responseStream: false,
        requestSerialize: (value: ConfirmDeploymentRequest) =>
            Buffer.from(ConfirmDeploymentRequest.encode(value).finish()),
        requestDeserialize: (value: Buffer) => ConfirmDeploymentRequest.decode(value),
        responseSerialize: (value: ConfirmDeploymentResponse) =>
            Buffer.from(ConfirmDeploymentResponse.encode(value).finish()),
        responseDeserialize: (value: Buffer) => ConfirmDeploymentResponse.decode(value),
    },
    rollbackDeployment: {
        path: "/freym.deployments.management.Service/RollbackDeployment",
        requestStream: false,
        responseStream: false,
        requestSerialize: (value: RollbackDeploymentRequest) =>
            Buffer.from(RollbackDeploymentRequest.encode(value).finish()),
        requestDeserialize: (value: Buffer) => RollbackDeploymentRequest.decode(value),
        responseSerialize: (value: RollbackDeploymentResponse) =>
            Buffer.from(RollbackDeploymentResponse.encode(value).finish()),
        responseDeserialize: (value: Buffer) => RollbackDeploymentResponse.decode(value),
    },
    rollbackDeploymentById: {
        path: "/freym.deployments.management.Service/RollbackDeploymentById",
        requestStream: false,
        responseStream: false,
        requestSerialize: (value: RollbackDeploymentByIdRequest) =>
            Buffer.from(RollbackDeploymentByIdRequest.encode(value).finish()),
        requestDeserialize: (value: Buffer) => RollbackDeploymentByIdRequest.decode(value),
        responseSerialize: (value: RollbackDeploymentResponse) =>
            Buffer.from(RollbackDeploymentResponse.encode(value).finish()),
        responseDeserialize: (value: Buffer) => RollbackDeploymentResponse.decode(value),
    },
    getDeployment: {
        path: "/freym.deployments.management.Service/GetDeployment",
        requestStream: false,
        responseStream: false,
        requestSerialize: (value: GetDeploymentRequest) =>
            Buffer.from(GetDeploymentRequest.encode(value).finish()),
        requestDeserialize: (value: Buffer) => GetDeploymentRequest.decode(value),
        responseSerialize: (value: GetDeploymentResponse) =>
            Buffer.from(GetDeploymentResponse.encode(value).finish()),
        responseDeserialize: (value: Buffer) => GetDeploymentResponse.decode(value),
    },
} as const;

export interface ServiceServer extends UntypedServiceImplementation {
    createDeployment: handleUnaryCall<CreateDeploymentRequest, CreateDeploymentResponse>;
    activateDeployment: handleUnaryCall<ActivateDeploymentRequest, ActivateDeploymentResponse>;
    confirmDeployment: handleUnaryCall<ConfirmDeploymentRequest, ConfirmDeploymentResponse>;
    rollbackDeployment: handleUnaryCall<RollbackDeploymentRequest, RollbackDeploymentResponse>;
    rollbackDeploymentById: handleUnaryCall<
        RollbackDeploymentByIdRequest,
        RollbackDeploymentResponse
    >;
    getDeployment: handleUnaryCall<GetDeploymentRequest, GetDeploymentResponse>;
}

export interface ServiceClient extends Client {
    createDeployment(
        request: CreateDeploymentRequest,
        callback: (error: ServiceError | null, response: CreateDeploymentResponse) => void
    ): ClientUnaryCall;
    createDeployment(
        request: CreateDeploymentRequest,
        metadata: Metadata,
        callback: (error: ServiceError | null, response: CreateDeploymentResponse) => void
    ): ClientUnaryCall;
    createDeployment(
        request: CreateDeploymentRequest,
        metadata: Metadata,
        options: Partial<CallOptions>,
        callback: (error: ServiceError | null, response: CreateDeploymentResponse) => void
    ): ClientUnaryCall;
    activateDeployment(
        request: ActivateDeploymentRequest,
        callback: (error: ServiceError | null, response: ActivateDeploymentResponse) => void
    ): ClientUnaryCall;
    activateDeployment(
        request: ActivateDeploymentRequest,
        metadata: Metadata,
        callback: (error: ServiceError | null, response: ActivateDeploymentResponse) => void
    ): ClientUnaryCall;
    activateDeployment(
        request: ActivateDeploymentRequest,
        metadata: Metadata,
        options: Partial<CallOptions>,
        callback: (error: ServiceError | null, response: ActivateDeploymentResponse) => void
    ): ClientUnaryCall;
    confirmDeployment(
        request: ConfirmDeploymentRequest,
        callback: (error: ServiceError | null, response: ConfirmDeploymentResponse) => void
    ): ClientUnaryCall;
    confirmDeployment(
        request: ConfirmDeploymentRequest,
        metadata: Metadata,
        callback: (error: ServiceError | null, response: ConfirmDeploymentResponse) => void
    ): ClientUnaryCall;
    confirmDeployment(
        request: ConfirmDeploymentRequest,
        metadata: Metadata,
        options: Partial<CallOptions>,
        callback: (error: ServiceError | null, response: ConfirmDeploymentResponse) => void
    ): ClientUnaryCall;
    rollbackDeployment(
        request: RollbackDeploymentRequest,
        callback: (error: ServiceError | null, response: RollbackDeploymentResponse) => void
    ): ClientUnaryCall;
    rollbackDeployment(
        request: RollbackDeploymentRequest,
        metadata: Metadata,
        callback: (error: ServiceError | null, response: RollbackDeploymentResponse) => void
    ): ClientUnaryCall;
    rollbackDeployment(
        request: RollbackDeploymentRequest,
        metadata: Metadata,
        options: Partial<CallOptions>,
        callback: (error: ServiceError | null, response: RollbackDeploymentResponse) => void
    ): ClientUnaryCall;
    rollbackDeploymentById(
        request: RollbackDeploymentByIdRequest,
        callback: (error: ServiceError | null, response: RollbackDeploymentResponse) => void
    ): ClientUnaryCall;
    rollbackDeploymentById(
        request: RollbackDeploymentByIdRequest,
        metadata: Metadata,
        callback: (error: ServiceError | null, response: RollbackDeploymentResponse) => void
    ): ClientUnaryCall;
    rollbackDeploymentById(
        request: RollbackDeploymentByIdRequest,
        metadata: Metadata,
        options: Partial<CallOptions>,
        callback: (error: ServiceError | null, response: RollbackDeploymentResponse) => void
    ): ClientUnaryCall;
    getDeployment(
        request: GetDeploymentRequest,
        callback: (error: ServiceError | null, response: GetDeploymentResponse) => void
    ): ClientUnaryCall;
    getDeployment(
        request: GetDeploymentRequest,
        metadata: Metadata,
        callback: (error: ServiceError | null, response: GetDeploymentResponse) => void
    ): ClientUnaryCall;
    getDeployment(
        request: GetDeploymentRequest,
        metadata: Metadata,
        options: Partial<CallOptions>,
        callback: (error: ServiceError | null, response: GetDeploymentResponse) => void
    ): ClientUnaryCall;
}

export const ServiceClient = makeGenericClientConstructor(
    ServiceService,
    "freym.deployments.management.Service"
) as unknown as {
    new (
        address: string,
        credentials: ChannelCredentials,
        options?: Partial<ClientOptions>
    ): ServiceClient;
    service: typeof ServiceService;
    serviceName: string;
};
