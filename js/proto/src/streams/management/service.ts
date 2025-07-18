// Code generated by protoc-gen-ts_proto. DO NOT EDIT.
// versions:
//   protoc-gen-ts_proto  v2.7.5
//   protoc               v5.29.3
// source: streams/management/service.proto
import {
    type CallOptions,
    type ChannelCredentials,
    Client,
    type ClientDuplexStream,
    type ClientOptions,
    type ClientUnaryCall,
    type Metadata,
    type ServiceError,
    type UntypedServiceImplementation,
    type handleBidiStreamingCall,
    type handleUnaryCall,
    makeGenericClientConstructor,
} from "@grpc/grpc-js";
import { BackchannelEventRequest, BackchannelEventResponse } from "./backchannel";
import {
    Event,
    GetEventRequest,
    GetLastEventByTypesRequest,
    GetLastEventRequest,
    GetLastHandledEventRequest,
} from "./event";
import {
    IntroduceGdprOnEventFieldRequest,
    IntroduceGdprOnEventFieldResponse,
    InvalidateGdprRequest,
    InvalidateGdprResponse,
} from "./gdpr";
import {
    PaginateEventsAfterEventIdRequest,
    PaginateEventsAfterEventIdResponse,
    PaginateEventsRequest,
    PaginateEventsResponse,
    PaginateStreamAfterEventIdRequest,
    PaginateStreamAfterEventIdResponse,
    PaginateStreamRequest,
    PaginateStreamResponse,
} from "./paginate";
import { PublishRequest, PublishResponse } from "./publish";
import { RenameEventTypeRequest, RenameEventTypeResponse } from "./rename";
import { CreateStreamSnapshotRequest, CreateStreamSnapshotResponse } from "./snapshot";
import { IsStreamEmptyRequest, IsStreamEmptyResponse } from "./stream";
import { SubscribeRequest, SubscribeResponse } from "./subscribe";

export type ServiceService = typeof ServiceService;
export const ServiceService = {
    publish: {
        path: "/freym.streams.management.Service/Publish",
        requestStream: false,
        responseStream: false,
        requestSerialize: (value: PublishRequest): Buffer =>
            Buffer.from(PublishRequest.encode(value).finish()),
        requestDeserialize: (value: Buffer): PublishRequest => PublishRequest.decode(value),
        responseSerialize: (value: PublishResponse): Buffer =>
            Buffer.from(PublishResponse.encode(value).finish()),
        responseDeserialize: (value: Buffer): PublishResponse => PublishResponse.decode(value),
    },
    subscribe: {
        path: "/freym.streams.management.Service/Subscribe",
        requestStream: true,
        responseStream: true,
        requestSerialize: (value: SubscribeRequest): Buffer =>
            Buffer.from(SubscribeRequest.encode(value).finish()),
        requestDeserialize: (value: Buffer): SubscribeRequest => SubscribeRequest.decode(value),
        responseSerialize: (value: SubscribeResponse): Buffer =>
            Buffer.from(SubscribeResponse.encode(value).finish()),
        responseDeserialize: (value: Buffer): SubscribeResponse => SubscribeResponse.decode(value),
    },
    getEvent: {
        path: "/freym.streams.management.Service/GetEvent",
        requestStream: false,
        responseStream: false,
        requestSerialize: (value: GetEventRequest): Buffer =>
            Buffer.from(GetEventRequest.encode(value).finish()),
        requestDeserialize: (value: Buffer): GetEventRequest => GetEventRequest.decode(value),
        responseSerialize: (value: Event): Buffer => Buffer.from(Event.encode(value).finish()),
        responseDeserialize: (value: Buffer): Event => Event.decode(value),
    },
    getLastEvent: {
        path: "/freym.streams.management.Service/GetLastEvent",
        requestStream: false,
        responseStream: false,
        requestSerialize: (value: GetLastEventRequest): Buffer =>
            Buffer.from(GetLastEventRequest.encode(value).finish()),
        requestDeserialize: (value: Buffer): GetLastEventRequest =>
            GetLastEventRequest.decode(value),
        responseSerialize: (value: Event): Buffer => Buffer.from(Event.encode(value).finish()),
        responseDeserialize: (value: Buffer): Event => Event.decode(value),
    },
    getLastHandledEvent: {
        path: "/freym.streams.management.Service/GetLastHandledEvent",
        requestStream: false,
        responseStream: false,
        requestSerialize: (value: GetLastHandledEventRequest): Buffer =>
            Buffer.from(GetLastHandledEventRequest.encode(value).finish()),
        requestDeserialize: (value: Buffer): GetLastHandledEventRequest =>
            GetLastHandledEventRequest.decode(value),
        responseSerialize: (value: Event): Buffer => Buffer.from(Event.encode(value).finish()),
        responseDeserialize: (value: Buffer): Event => Event.decode(value),
    },
    getLastEventByTypes: {
        path: "/freym.streams.management.Service/GetLastEventByTypes",
        requestStream: false,
        responseStream: false,
        requestSerialize: (value: GetLastEventByTypesRequest): Buffer =>
            Buffer.from(GetLastEventByTypesRequest.encode(value).finish()),
        requestDeserialize: (value: Buffer): GetLastEventByTypesRequest =>
            GetLastEventByTypesRequest.decode(value),
        responseSerialize: (value: Event): Buffer => Buffer.from(Event.encode(value).finish()),
        responseDeserialize: (value: Buffer): Event => Event.decode(value),
    },
    isStreamEmpty: {
        path: "/freym.streams.management.Service/IsStreamEmpty",
        requestStream: false,
        responseStream: false,
        requestSerialize: (value: IsStreamEmptyRequest): Buffer =>
            Buffer.from(IsStreamEmptyRequest.encode(value).finish()),
        requestDeserialize: (value: Buffer): IsStreamEmptyRequest =>
            IsStreamEmptyRequest.decode(value),
        responseSerialize: (value: IsStreamEmptyResponse): Buffer =>
            Buffer.from(IsStreamEmptyResponse.encode(value).finish()),
        responseDeserialize: (value: Buffer): IsStreamEmptyResponse =>
            IsStreamEmptyResponse.decode(value),
    },
    paginateStream: {
        path: "/freym.streams.management.Service/PaginateStream",
        requestStream: false,
        responseStream: false,
        requestSerialize: (value: PaginateStreamRequest): Buffer =>
            Buffer.from(PaginateStreamRequest.encode(value).finish()),
        requestDeserialize: (value: Buffer): PaginateStreamRequest =>
            PaginateStreamRequest.decode(value),
        responseSerialize: (value: PaginateStreamResponse): Buffer =>
            Buffer.from(PaginateStreamResponse.encode(value).finish()),
        responseDeserialize: (value: Buffer): PaginateStreamResponse =>
            PaginateStreamResponse.decode(value),
    },
    paginateStreamAfterEventId: {
        path: "/freym.streams.management.Service/PaginateStreamAfterEventId",
        requestStream: false,
        responseStream: false,
        requestSerialize: (value: PaginateStreamAfterEventIdRequest): Buffer =>
            Buffer.from(PaginateStreamAfterEventIdRequest.encode(value).finish()),
        requestDeserialize: (value: Buffer): PaginateStreamAfterEventIdRequest =>
            PaginateStreamAfterEventIdRequest.decode(value),
        responseSerialize: (value: PaginateStreamAfterEventIdResponse): Buffer =>
            Buffer.from(PaginateStreamAfterEventIdResponse.encode(value).finish()),
        responseDeserialize: (value: Buffer): PaginateStreamAfterEventIdResponse =>
            PaginateStreamAfterEventIdResponse.decode(value),
    },
    paginateEvents: {
        path: "/freym.streams.management.Service/PaginateEvents",
        requestStream: false,
        responseStream: false,
        requestSerialize: (value: PaginateEventsRequest): Buffer =>
            Buffer.from(PaginateEventsRequest.encode(value).finish()),
        requestDeserialize: (value: Buffer): PaginateEventsRequest =>
            PaginateEventsRequest.decode(value),
        responseSerialize: (value: PaginateEventsResponse): Buffer =>
            Buffer.from(PaginateEventsResponse.encode(value).finish()),
        responseDeserialize: (value: Buffer): PaginateEventsResponse =>
            PaginateEventsResponse.decode(value),
    },
    paginateEventsAfterEventId: {
        path: "/freym.streams.management.Service/PaginateEventsAfterEventId",
        requestStream: false,
        responseStream: false,
        requestSerialize: (value: PaginateEventsAfterEventIdRequest): Buffer =>
            Buffer.from(PaginateEventsAfterEventIdRequest.encode(value).finish()),
        requestDeserialize: (value: Buffer): PaginateEventsAfterEventIdRequest =>
            PaginateEventsAfterEventIdRequest.decode(value),
        responseSerialize: (value: PaginateEventsAfterEventIdResponse): Buffer =>
            Buffer.from(PaginateEventsAfterEventIdResponse.encode(value).finish()),
        responseDeserialize: (value: Buffer): PaginateEventsAfterEventIdResponse =>
            PaginateEventsAfterEventIdResponse.decode(value),
    },
    introduceGdprOnEventField: {
        path: "/freym.streams.management.Service/IntroduceGdprOnEventField",
        requestStream: false,
        responseStream: false,
        requestSerialize: (value: IntroduceGdprOnEventFieldRequest): Buffer =>
            Buffer.from(IntroduceGdprOnEventFieldRequest.encode(value).finish()),
        requestDeserialize: (value: Buffer): IntroduceGdprOnEventFieldRequest =>
            IntroduceGdprOnEventFieldRequest.decode(value),
        responseSerialize: (value: IntroduceGdprOnEventFieldResponse): Buffer =>
            Buffer.from(IntroduceGdprOnEventFieldResponse.encode(value).finish()),
        responseDeserialize: (value: Buffer): IntroduceGdprOnEventFieldResponse =>
            IntroduceGdprOnEventFieldResponse.decode(value),
    },
    invalidateGdpr: {
        path: "/freym.streams.management.Service/InvalidateGdpr",
        requestStream: false,
        responseStream: false,
        requestSerialize: (value: InvalidateGdprRequest): Buffer =>
            Buffer.from(InvalidateGdprRequest.encode(value).finish()),
        requestDeserialize: (value: Buffer): InvalidateGdprRequest =>
            InvalidateGdprRequest.decode(value),
        responseSerialize: (value: InvalidateGdprResponse): Buffer =>
            Buffer.from(InvalidateGdprResponse.encode(value).finish()),
        responseDeserialize: (value: Buffer): InvalidateGdprResponse =>
            InvalidateGdprResponse.decode(value),
    },
    backchannelEvent: {
        path: "/freym.streams.management.Service/BackchannelEvent",
        requestStream: false,
        responseStream: false,
        requestSerialize: (value: BackchannelEventRequest): Buffer =>
            Buffer.from(BackchannelEventRequest.encode(value).finish()),
        requestDeserialize: (value: Buffer): BackchannelEventRequest =>
            BackchannelEventRequest.decode(value),
        responseSerialize: (value: BackchannelEventResponse): Buffer =>
            Buffer.from(BackchannelEventResponse.encode(value).finish()),
        responseDeserialize: (value: Buffer): BackchannelEventResponse =>
            BackchannelEventResponse.decode(value),
    },
    createStreamSnapshot: {
        path: "/freym.streams.management.Service/CreateStreamSnapshot",
        requestStream: false,
        responseStream: false,
        requestSerialize: (value: CreateStreamSnapshotRequest): Buffer =>
            Buffer.from(CreateStreamSnapshotRequest.encode(value).finish()),
        requestDeserialize: (value: Buffer): CreateStreamSnapshotRequest =>
            CreateStreamSnapshotRequest.decode(value),
        responseSerialize: (value: CreateStreamSnapshotResponse): Buffer =>
            Buffer.from(CreateStreamSnapshotResponse.encode(value).finish()),
        responseDeserialize: (value: Buffer): CreateStreamSnapshotResponse =>
            CreateStreamSnapshotResponse.decode(value),
    },
    renameEventType: {
        path: "/freym.streams.management.Service/RenameEventType",
        requestStream: false,
        responseStream: false,
        requestSerialize: (value: RenameEventTypeRequest): Buffer =>
            Buffer.from(RenameEventTypeRequest.encode(value).finish()),
        requestDeserialize: (value: Buffer): RenameEventTypeRequest =>
            RenameEventTypeRequest.decode(value),
        responseSerialize: (value: RenameEventTypeResponse): Buffer =>
            Buffer.from(RenameEventTypeResponse.encode(value).finish()),
        responseDeserialize: (value: Buffer): RenameEventTypeResponse =>
            RenameEventTypeResponse.decode(value),
    },
} as const;

export interface ServiceServer extends UntypedServiceImplementation {
    publish: handleUnaryCall<PublishRequest, PublishResponse>;
    subscribe: handleBidiStreamingCall<SubscribeRequest, SubscribeResponse>;
    getEvent: handleUnaryCall<GetEventRequest, Event>;
    getLastEvent: handleUnaryCall<GetLastEventRequest, Event>;
    getLastHandledEvent: handleUnaryCall<GetLastHandledEventRequest, Event>;
    getLastEventByTypes: handleUnaryCall<GetLastEventByTypesRequest, Event>;
    isStreamEmpty: handleUnaryCall<IsStreamEmptyRequest, IsStreamEmptyResponse>;
    paginateStream: handleUnaryCall<PaginateStreamRequest, PaginateStreamResponse>;
    paginateStreamAfterEventId: handleUnaryCall<
        PaginateStreamAfterEventIdRequest,
        PaginateStreamAfterEventIdResponse
    >;
    paginateEvents: handleUnaryCall<PaginateEventsRequest, PaginateEventsResponse>;
    paginateEventsAfterEventId: handleUnaryCall<
        PaginateEventsAfterEventIdRequest,
        PaginateEventsAfterEventIdResponse
    >;
    introduceGdprOnEventField: handleUnaryCall<
        IntroduceGdprOnEventFieldRequest,
        IntroduceGdprOnEventFieldResponse
    >;
    invalidateGdpr: handleUnaryCall<InvalidateGdprRequest, InvalidateGdprResponse>;
    backchannelEvent: handleUnaryCall<BackchannelEventRequest, BackchannelEventResponse>;
    createStreamSnapshot: handleUnaryCall<
        CreateStreamSnapshotRequest,
        CreateStreamSnapshotResponse
    >;
    renameEventType: handleUnaryCall<RenameEventTypeRequest, RenameEventTypeResponse>;
}

export interface ServiceClient extends Client {
    publish(
        request: PublishRequest,
        callback: (error: ServiceError | null, response: PublishResponse) => void
    ): ClientUnaryCall;
    publish(
        request: PublishRequest,
        metadata: Metadata,
        callback: (error: ServiceError | null, response: PublishResponse) => void
    ): ClientUnaryCall;
    publish(
        request: PublishRequest,
        metadata: Metadata,
        options: Partial<CallOptions>,
        callback: (error: ServiceError | null, response: PublishResponse) => void
    ): ClientUnaryCall;
    subscribe(): ClientDuplexStream<SubscribeRequest, SubscribeResponse>;
    subscribe(
        options: Partial<CallOptions>
    ): ClientDuplexStream<SubscribeRequest, SubscribeResponse>;
    subscribe(
        metadata: Metadata,
        options?: Partial<CallOptions>
    ): ClientDuplexStream<SubscribeRequest, SubscribeResponse>;
    getEvent(
        request: GetEventRequest,
        callback: (error: ServiceError | null, response: Event) => void
    ): ClientUnaryCall;
    getEvent(
        request: GetEventRequest,
        metadata: Metadata,
        callback: (error: ServiceError | null, response: Event) => void
    ): ClientUnaryCall;
    getEvent(
        request: GetEventRequest,
        metadata: Metadata,
        options: Partial<CallOptions>,
        callback: (error: ServiceError | null, response: Event) => void
    ): ClientUnaryCall;
    getLastEvent(
        request: GetLastEventRequest,
        callback: (error: ServiceError | null, response: Event) => void
    ): ClientUnaryCall;
    getLastEvent(
        request: GetLastEventRequest,
        metadata: Metadata,
        callback: (error: ServiceError | null, response: Event) => void
    ): ClientUnaryCall;
    getLastEvent(
        request: GetLastEventRequest,
        metadata: Metadata,
        options: Partial<CallOptions>,
        callback: (error: ServiceError | null, response: Event) => void
    ): ClientUnaryCall;
    getLastHandledEvent(
        request: GetLastHandledEventRequest,
        callback: (error: ServiceError | null, response: Event) => void
    ): ClientUnaryCall;
    getLastHandledEvent(
        request: GetLastHandledEventRequest,
        metadata: Metadata,
        callback: (error: ServiceError | null, response: Event) => void
    ): ClientUnaryCall;
    getLastHandledEvent(
        request: GetLastHandledEventRequest,
        metadata: Metadata,
        options: Partial<CallOptions>,
        callback: (error: ServiceError | null, response: Event) => void
    ): ClientUnaryCall;
    getLastEventByTypes(
        request: GetLastEventByTypesRequest,
        callback: (error: ServiceError | null, response: Event) => void
    ): ClientUnaryCall;
    getLastEventByTypes(
        request: GetLastEventByTypesRequest,
        metadata: Metadata,
        callback: (error: ServiceError | null, response: Event) => void
    ): ClientUnaryCall;
    getLastEventByTypes(
        request: GetLastEventByTypesRequest,
        metadata: Metadata,
        options: Partial<CallOptions>,
        callback: (error: ServiceError | null, response: Event) => void
    ): ClientUnaryCall;
    isStreamEmpty(
        request: IsStreamEmptyRequest,
        callback: (error: ServiceError | null, response: IsStreamEmptyResponse) => void
    ): ClientUnaryCall;
    isStreamEmpty(
        request: IsStreamEmptyRequest,
        metadata: Metadata,
        callback: (error: ServiceError | null, response: IsStreamEmptyResponse) => void
    ): ClientUnaryCall;
    isStreamEmpty(
        request: IsStreamEmptyRequest,
        metadata: Metadata,
        options: Partial<CallOptions>,
        callback: (error: ServiceError | null, response: IsStreamEmptyResponse) => void
    ): ClientUnaryCall;
    paginateStream(
        request: PaginateStreamRequest,
        callback: (error: ServiceError | null, response: PaginateStreamResponse) => void
    ): ClientUnaryCall;
    paginateStream(
        request: PaginateStreamRequest,
        metadata: Metadata,
        callback: (error: ServiceError | null, response: PaginateStreamResponse) => void
    ): ClientUnaryCall;
    paginateStream(
        request: PaginateStreamRequest,
        metadata: Metadata,
        options: Partial<CallOptions>,
        callback: (error: ServiceError | null, response: PaginateStreamResponse) => void
    ): ClientUnaryCall;
    paginateStreamAfterEventId(
        request: PaginateStreamAfterEventIdRequest,
        callback: (error: ServiceError | null, response: PaginateStreamAfterEventIdResponse) => void
    ): ClientUnaryCall;
    paginateStreamAfterEventId(
        request: PaginateStreamAfterEventIdRequest,
        metadata: Metadata,
        callback: (error: ServiceError | null, response: PaginateStreamAfterEventIdResponse) => void
    ): ClientUnaryCall;
    paginateStreamAfterEventId(
        request: PaginateStreamAfterEventIdRequest,
        metadata: Metadata,
        options: Partial<CallOptions>,
        callback: (error: ServiceError | null, response: PaginateStreamAfterEventIdResponse) => void
    ): ClientUnaryCall;
    paginateEvents(
        request: PaginateEventsRequest,
        callback: (error: ServiceError | null, response: PaginateEventsResponse) => void
    ): ClientUnaryCall;
    paginateEvents(
        request: PaginateEventsRequest,
        metadata: Metadata,
        callback: (error: ServiceError | null, response: PaginateEventsResponse) => void
    ): ClientUnaryCall;
    paginateEvents(
        request: PaginateEventsRequest,
        metadata: Metadata,
        options: Partial<CallOptions>,
        callback: (error: ServiceError | null, response: PaginateEventsResponse) => void
    ): ClientUnaryCall;
    paginateEventsAfterEventId(
        request: PaginateEventsAfterEventIdRequest,
        callback: (error: ServiceError | null, response: PaginateEventsAfterEventIdResponse) => void
    ): ClientUnaryCall;
    paginateEventsAfterEventId(
        request: PaginateEventsAfterEventIdRequest,
        metadata: Metadata,
        callback: (error: ServiceError | null, response: PaginateEventsAfterEventIdResponse) => void
    ): ClientUnaryCall;
    paginateEventsAfterEventId(
        request: PaginateEventsAfterEventIdRequest,
        metadata: Metadata,
        options: Partial<CallOptions>,
        callback: (error: ServiceError | null, response: PaginateEventsAfterEventIdResponse) => void
    ): ClientUnaryCall;
    introduceGdprOnEventField(
        request: IntroduceGdprOnEventFieldRequest,
        callback: (error: ServiceError | null, response: IntroduceGdprOnEventFieldResponse) => void
    ): ClientUnaryCall;
    introduceGdprOnEventField(
        request: IntroduceGdprOnEventFieldRequest,
        metadata: Metadata,
        callback: (error: ServiceError | null, response: IntroduceGdprOnEventFieldResponse) => void
    ): ClientUnaryCall;
    introduceGdprOnEventField(
        request: IntroduceGdprOnEventFieldRequest,
        metadata: Metadata,
        options: Partial<CallOptions>,
        callback: (error: ServiceError | null, response: IntroduceGdprOnEventFieldResponse) => void
    ): ClientUnaryCall;
    invalidateGdpr(
        request: InvalidateGdprRequest,
        callback: (error: ServiceError | null, response: InvalidateGdprResponse) => void
    ): ClientUnaryCall;
    invalidateGdpr(
        request: InvalidateGdprRequest,
        metadata: Metadata,
        callback: (error: ServiceError | null, response: InvalidateGdprResponse) => void
    ): ClientUnaryCall;
    invalidateGdpr(
        request: InvalidateGdprRequest,
        metadata: Metadata,
        options: Partial<CallOptions>,
        callback: (error: ServiceError | null, response: InvalidateGdprResponse) => void
    ): ClientUnaryCall;
    backchannelEvent(
        request: BackchannelEventRequest,
        callback: (error: ServiceError | null, response: BackchannelEventResponse) => void
    ): ClientUnaryCall;
    backchannelEvent(
        request: BackchannelEventRequest,
        metadata: Metadata,
        callback: (error: ServiceError | null, response: BackchannelEventResponse) => void
    ): ClientUnaryCall;
    backchannelEvent(
        request: BackchannelEventRequest,
        metadata: Metadata,
        options: Partial<CallOptions>,
        callback: (error: ServiceError | null, response: BackchannelEventResponse) => void
    ): ClientUnaryCall;
    createStreamSnapshot(
        request: CreateStreamSnapshotRequest,
        callback: (error: ServiceError | null, response: CreateStreamSnapshotResponse) => void
    ): ClientUnaryCall;
    createStreamSnapshot(
        request: CreateStreamSnapshotRequest,
        metadata: Metadata,
        callback: (error: ServiceError | null, response: CreateStreamSnapshotResponse) => void
    ): ClientUnaryCall;
    createStreamSnapshot(
        request: CreateStreamSnapshotRequest,
        metadata: Metadata,
        options: Partial<CallOptions>,
        callback: (error: ServiceError | null, response: CreateStreamSnapshotResponse) => void
    ): ClientUnaryCall;
    renameEventType(
        request: RenameEventTypeRequest,
        callback: (error: ServiceError | null, response: RenameEventTypeResponse) => void
    ): ClientUnaryCall;
    renameEventType(
        request: RenameEventTypeRequest,
        metadata: Metadata,
        callback: (error: ServiceError | null, response: RenameEventTypeResponse) => void
    ): ClientUnaryCall;
    renameEventType(
        request: RenameEventTypeRequest,
        metadata: Metadata,
        options: Partial<CallOptions>,
        callback: (error: ServiceError | null, response: RenameEventTypeResponse) => void
    ): ClientUnaryCall;
}

export const ServiceClient = makeGenericClientConstructor(
    ServiceService,
    "freym.streams.management.Service"
) as unknown as {
    new (
        address: string,
        credentials: ChannelCredentials,
        options?: Partial<ClientOptions>
    ): ServiceClient;
    service: typeof ServiceService;
    serviceName: string;
};
