// Code generated by protoc-gen-ts_proto. DO NOT EDIT.
// versions:
//   protoc-gen-ts_proto  v2.6.1
//   protoc               v5.27.3
// source: projections/delivery/service.proto
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
import { DeleteRequest, DeleteResponse } from "./delete";
import {
    GetDataListRequest,
    GetDataListResponse,
    GetDataRequest,
    GetDataResponse,
} from "./get_data";
import {
    GetViewDataListRequest,
    GetViewDataListResponse,
    GetViewDataRequest,
    GetViewDataResponse,
} from "./get_view_data";
import { UpsertRequest, UpsertResponse } from "./upsert";

export type ServiceService = typeof ServiceService;
export const ServiceService = {
    getData: {
        path: "/freym.projections.delivery.Service/GetData",
        requestStream: false,
        responseStream: false,
        requestSerialize: (value: GetDataRequest) =>
            Buffer.from(GetDataRequest.encode(value).finish()),
        requestDeserialize: (value: Buffer) => GetDataRequest.decode(value),
        responseSerialize: (value: GetDataResponse) =>
            Buffer.from(GetDataResponse.encode(value).finish()),
        responseDeserialize: (value: Buffer) => GetDataResponse.decode(value),
    },
    getViewData: {
        path: "/freym.projections.delivery.Service/GetViewData",
        requestStream: false,
        responseStream: false,
        requestSerialize: (value: GetViewDataRequest) =>
            Buffer.from(GetViewDataRequest.encode(value).finish()),
        requestDeserialize: (value: Buffer) => GetViewDataRequest.decode(value),
        responseSerialize: (value: GetViewDataResponse) =>
            Buffer.from(GetViewDataResponse.encode(value).finish()),
        responseDeserialize: (value: Buffer) => GetViewDataResponse.decode(value),
    },
    getDataList: {
        path: "/freym.projections.delivery.Service/GetDataList",
        requestStream: false,
        responseStream: false,
        requestSerialize: (value: GetDataListRequest) =>
            Buffer.from(GetDataListRequest.encode(value).finish()),
        requestDeserialize: (value: Buffer) => GetDataListRequest.decode(value),
        responseSerialize: (value: GetDataListResponse) =>
            Buffer.from(GetDataListResponse.encode(value).finish()),
        responseDeserialize: (value: Buffer) => GetDataListResponse.decode(value),
    },
    getViewDataList: {
        path: "/freym.projections.delivery.Service/GetViewDataList",
        requestStream: false,
        responseStream: false,
        requestSerialize: (value: GetViewDataListRequest) =>
            Buffer.from(GetViewDataListRequest.encode(value).finish()),
        requestDeserialize: (value: Buffer) => GetViewDataListRequest.decode(value),
        responseSerialize: (value: GetViewDataListResponse) =>
            Buffer.from(GetViewDataListResponse.encode(value).finish()),
        responseDeserialize: (value: Buffer) => GetViewDataListResponse.decode(value),
    },
    upsert: {
        path: "/freym.projections.delivery.Service/Upsert",
        requestStream: false,
        responseStream: false,
        requestSerialize: (value: UpsertRequest) =>
            Buffer.from(UpsertRequest.encode(value).finish()),
        requestDeserialize: (value: Buffer) => UpsertRequest.decode(value),
        responseSerialize: (value: UpsertResponse) =>
            Buffer.from(UpsertResponse.encode(value).finish()),
        responseDeserialize: (value: Buffer) => UpsertResponse.decode(value),
    },
    delete: {
        path: "/freym.projections.delivery.Service/Delete",
        requestStream: false,
        responseStream: false,
        requestSerialize: (value: DeleteRequest) =>
            Buffer.from(DeleteRequest.encode(value).finish()),
        requestDeserialize: (value: Buffer) => DeleteRequest.decode(value),
        responseSerialize: (value: DeleteResponse) =>
            Buffer.from(DeleteResponse.encode(value).finish()),
        responseDeserialize: (value: Buffer) => DeleteResponse.decode(value),
    },
} as const;

export interface ServiceServer extends UntypedServiceImplementation {
    getData: handleUnaryCall<GetDataRequest, GetDataResponse>;
    getViewData: handleUnaryCall<GetViewDataRequest, GetViewDataResponse>;
    getDataList: handleUnaryCall<GetDataListRequest, GetDataListResponse>;
    getViewDataList: handleUnaryCall<GetViewDataListRequest, GetViewDataListResponse>;
    upsert: handleUnaryCall<UpsertRequest, UpsertResponse>;
    delete: handleUnaryCall<DeleteRequest, DeleteResponse>;
}

export interface ServiceClient extends Client {
    getData(
        request: GetDataRequest,
        callback: (error: ServiceError | null, response: GetDataResponse) => void
    ): ClientUnaryCall;
    getData(
        request: GetDataRequest,
        metadata: Metadata,
        callback: (error: ServiceError | null, response: GetDataResponse) => void
    ): ClientUnaryCall;
    getData(
        request: GetDataRequest,
        metadata: Metadata,
        options: Partial<CallOptions>,
        callback: (error: ServiceError | null, response: GetDataResponse) => void
    ): ClientUnaryCall;
    getViewData(
        request: GetViewDataRequest,
        callback: (error: ServiceError | null, response: GetViewDataResponse) => void
    ): ClientUnaryCall;
    getViewData(
        request: GetViewDataRequest,
        metadata: Metadata,
        callback: (error: ServiceError | null, response: GetViewDataResponse) => void
    ): ClientUnaryCall;
    getViewData(
        request: GetViewDataRequest,
        metadata: Metadata,
        options: Partial<CallOptions>,
        callback: (error: ServiceError | null, response: GetViewDataResponse) => void
    ): ClientUnaryCall;
    getDataList(
        request: GetDataListRequest,
        callback: (error: ServiceError | null, response: GetDataListResponse) => void
    ): ClientUnaryCall;
    getDataList(
        request: GetDataListRequest,
        metadata: Metadata,
        callback: (error: ServiceError | null, response: GetDataListResponse) => void
    ): ClientUnaryCall;
    getDataList(
        request: GetDataListRequest,
        metadata: Metadata,
        options: Partial<CallOptions>,
        callback: (error: ServiceError | null, response: GetDataListResponse) => void
    ): ClientUnaryCall;
    getViewDataList(
        request: GetViewDataListRequest,
        callback: (error: ServiceError | null, response: GetViewDataListResponse) => void
    ): ClientUnaryCall;
    getViewDataList(
        request: GetViewDataListRequest,
        metadata: Metadata,
        callback: (error: ServiceError | null, response: GetViewDataListResponse) => void
    ): ClientUnaryCall;
    getViewDataList(
        request: GetViewDataListRequest,
        metadata: Metadata,
        options: Partial<CallOptions>,
        callback: (error: ServiceError | null, response: GetViewDataListResponse) => void
    ): ClientUnaryCall;
    upsert(
        request: UpsertRequest,
        callback: (error: ServiceError | null, response: UpsertResponse) => void
    ): ClientUnaryCall;
    upsert(
        request: UpsertRequest,
        metadata: Metadata,
        callback: (error: ServiceError | null, response: UpsertResponse) => void
    ): ClientUnaryCall;
    upsert(
        request: UpsertRequest,
        metadata: Metadata,
        options: Partial<CallOptions>,
        callback: (error: ServiceError | null, response: UpsertResponse) => void
    ): ClientUnaryCall;
    delete(
        request: DeleteRequest,
        callback: (error: ServiceError | null, response: DeleteResponse) => void
    ): ClientUnaryCall;
    delete(
        request: DeleteRequest,
        metadata: Metadata,
        callback: (error: ServiceError | null, response: DeleteResponse) => void
    ): ClientUnaryCall;
    delete(
        request: DeleteRequest,
        metadata: Metadata,
        options: Partial<CallOptions>,
        callback: (error: ServiceError | null, response: DeleteResponse) => void
    ): ClientUnaryCall;
}

export const ServiceClient = makeGenericClientConstructor(
    ServiceService,
    "freym.projections.delivery.Service"
) as unknown as {
    new (
        address: string,
        credentials: ChannelCredentials,
        options?: Partial<ClientOptions>
    ): ServiceClient;
    service: typeof ServiceService;
    serviceName: string;
};
