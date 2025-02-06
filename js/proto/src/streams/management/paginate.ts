// Code generated by protoc-gen-ts_proto. DO NOT EDIT.
// versions:
//   protoc-gen-ts_proto  v2.6.1
//   protoc               v5.29.3
// source: streams/management/paginate.proto

/* eslint-disable */
import { BinaryReader, BinaryWriter } from "@bufbuild/protobuf/wire";
import { Event } from "./event";

export interface PaginateStreamRequest {
    tenantId: string;
    topic: string;
    stream: string;
    page: string;
    perPage: string;
    /** when a deployment_id (!= 0) is provided the snapshot will only be used by queries with a deployment_id that is equal or higher than the shapshot deployment_id. */
    deploymentId: string;
    snapshotEventId: string;
}

export interface PaginateStreamResponse {
    events: Event[];
    snapshotEventId: string;
}

export interface PaginateStreamAfterEventIdRequest {
    tenantId: string;
    topic: string;
    stream: string;
    eventId: string;
    page: string;
    perPage: string;
    /** when a deployment_id (!= 0) is provided the snapshot will only be used by queries with a deployment_id that is equal or higher than the shapshot deployment_id. */
    deploymentId: string;
    snapshotEventId: string;
}

export interface PaginateStreamAfterEventIdResponse {
    events: Event[];
    snapshotEventId: string;
}

export interface PaginateEventsRequest {
    tenantId: string;
    topic: string;
    types: string[];
    page: string;
    perPage: string;
}

export interface PaginateEventsResponse {
    events: Event[];
}

export interface PaginateEventsAfterEventIdRequest {
    tenantId: string;
    topic: string;
    types: string[];
    eventId: string;
    page: string;
    perPage: string;
}

export interface PaginateEventsAfterEventIdResponse {
    events: Event[];
}

function createBasePaginateStreamRequest(): PaginateStreamRequest {
    return {
        tenantId: "",
        topic: "",
        stream: "",
        page: "0",
        perPage: "0",
        deploymentId: "0",
        snapshotEventId: "",
    };
}

export const PaginateStreamRequest: MessageFns<PaginateStreamRequest> = {
    encode(
        message: PaginateStreamRequest,
        writer: BinaryWriter = new BinaryWriter()
    ): BinaryWriter {
        if (message.tenantId !== "") {
            writer.uint32(10).string(message.tenantId);
        }
        if (message.topic !== "") {
            writer.uint32(18).string(message.topic);
        }
        if (message.stream !== "") {
            writer.uint32(26).string(message.stream);
        }
        if (message.page !== "0") {
            writer.uint32(32).int64(message.page);
        }
        if (message.perPage !== "0") {
            writer.uint32(40).int64(message.perPage);
        }
        if (message.deploymentId !== "0") {
            writer.uint32(48).int64(message.deploymentId);
        }
        if (message.snapshotEventId !== "") {
            writer.uint32(58).string(message.snapshotEventId);
        }
        return writer;
    },

    decode(input: BinaryReader | Uint8Array, length?: number): PaginateStreamRequest {
        const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBasePaginateStreamRequest();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1: {
                    if (tag !== 10) {
                        break;
                    }

                    message.tenantId = reader.string();
                    continue;
                }
                case 2: {
                    if (tag !== 18) {
                        break;
                    }

                    message.topic = reader.string();
                    continue;
                }
                case 3: {
                    if (tag !== 26) {
                        break;
                    }

                    message.stream = reader.string();
                    continue;
                }
                case 4: {
                    if (tag !== 32) {
                        break;
                    }

                    message.page = reader.int64().toString();
                    continue;
                }
                case 5: {
                    if (tag !== 40) {
                        break;
                    }

                    message.perPage = reader.int64().toString();
                    continue;
                }
                case 6: {
                    if (tag !== 48) {
                        break;
                    }

                    message.deploymentId = reader.int64().toString();
                    continue;
                }
                case 7: {
                    if (tag !== 58) {
                        break;
                    }

                    message.snapshotEventId = reader.string();
                    continue;
                }
            }
            if ((tag & 7) === 4 || tag === 0) {
                break;
            }
            reader.skip(tag & 7);
        }
        return message;
    },

    fromJSON(object: any): PaginateStreamRequest {
        return {
            tenantId: isSet(object.tenantId) ? globalThis.String(object.tenantId) : "",
            topic: isSet(object.topic) ? globalThis.String(object.topic) : "",
            stream: isSet(object.stream) ? globalThis.String(object.stream) : "",
            page: isSet(object.page) ? globalThis.String(object.page) : "0",
            perPage: isSet(object.perPage) ? globalThis.String(object.perPage) : "0",
            deploymentId: isSet(object.deploymentId) ? globalThis.String(object.deploymentId) : "0",
            snapshotEventId: isSet(object.snapshotEventId)
                ? globalThis.String(object.snapshotEventId)
                : "",
        };
    },

    toJSON(message: PaginateStreamRequest): unknown {
        const obj: any = {};
        if (message.tenantId !== "") {
            obj.tenantId = message.tenantId;
        }
        if (message.topic !== "") {
            obj.topic = message.topic;
        }
        if (message.stream !== "") {
            obj.stream = message.stream;
        }
        if (message.page !== "0") {
            obj.page = message.page;
        }
        if (message.perPage !== "0") {
            obj.perPage = message.perPage;
        }
        if (message.deploymentId !== "0") {
            obj.deploymentId = message.deploymentId;
        }
        if (message.snapshotEventId !== "") {
            obj.snapshotEventId = message.snapshotEventId;
        }
        return obj;
    },

    create(base?: DeepPartial<PaginateStreamRequest>): PaginateStreamRequest {
        return PaginateStreamRequest.fromPartial(base ?? {});
    },
    fromPartial(object: DeepPartial<PaginateStreamRequest>): PaginateStreamRequest {
        const message = createBasePaginateStreamRequest();
        message.tenantId = object.tenantId ?? "";
        message.topic = object.topic ?? "";
        message.stream = object.stream ?? "";
        message.page = object.page ?? "0";
        message.perPage = object.perPage ?? "0";
        message.deploymentId = object.deploymentId ?? "0";
        message.snapshotEventId = object.snapshotEventId ?? "";
        return message;
    },
};

function createBasePaginateStreamResponse(): PaginateStreamResponse {
    return { events: [], snapshotEventId: "" };
}

export const PaginateStreamResponse: MessageFns<PaginateStreamResponse> = {
    encode(
        message: PaginateStreamResponse,
        writer: BinaryWriter = new BinaryWriter()
    ): BinaryWriter {
        for (const v of message.events) {
            Event.encode(v!, writer.uint32(10).fork()).join();
        }
        if (message.snapshotEventId !== "") {
            writer.uint32(18).string(message.snapshotEventId);
        }
        return writer;
    },

    decode(input: BinaryReader | Uint8Array, length?: number): PaginateStreamResponse {
        const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBasePaginateStreamResponse();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1: {
                    if (tag !== 10) {
                        break;
                    }

                    message.events.push(Event.decode(reader, reader.uint32()));
                    continue;
                }
                case 2: {
                    if (tag !== 18) {
                        break;
                    }

                    message.snapshotEventId = reader.string();
                    continue;
                }
            }
            if ((tag & 7) === 4 || tag === 0) {
                break;
            }
            reader.skip(tag & 7);
        }
        return message;
    },

    fromJSON(object: any): PaginateStreamResponse {
        return {
            events: globalThis.Array.isArray(object?.events)
                ? object.events.map((e: any) => Event.fromJSON(e))
                : [],
            snapshotEventId: isSet(object.snapshotEventId)
                ? globalThis.String(object.snapshotEventId)
                : "",
        };
    },

    toJSON(message: PaginateStreamResponse): unknown {
        const obj: any = {};
        if (message.events?.length) {
            obj.events = message.events.map(e => Event.toJSON(e));
        }
        if (message.snapshotEventId !== "") {
            obj.snapshotEventId = message.snapshotEventId;
        }
        return obj;
    },

    create(base?: DeepPartial<PaginateStreamResponse>): PaginateStreamResponse {
        return PaginateStreamResponse.fromPartial(base ?? {});
    },
    fromPartial(object: DeepPartial<PaginateStreamResponse>): PaginateStreamResponse {
        const message = createBasePaginateStreamResponse();
        message.events = object.events?.map(e => Event.fromPartial(e)) || [];
        message.snapshotEventId = object.snapshotEventId ?? "";
        return message;
    },
};

function createBasePaginateStreamAfterEventIdRequest(): PaginateStreamAfterEventIdRequest {
    return {
        tenantId: "",
        topic: "",
        stream: "",
        eventId: "",
        page: "0",
        perPage: "0",
        deploymentId: "0",
        snapshotEventId: "",
    };
}

export const PaginateStreamAfterEventIdRequest: MessageFns<PaginateStreamAfterEventIdRequest> = {
    encode(
        message: PaginateStreamAfterEventIdRequest,
        writer: BinaryWriter = new BinaryWriter()
    ): BinaryWriter {
        if (message.tenantId !== "") {
            writer.uint32(10).string(message.tenantId);
        }
        if (message.topic !== "") {
            writer.uint32(18).string(message.topic);
        }
        if (message.stream !== "") {
            writer.uint32(26).string(message.stream);
        }
        if (message.eventId !== "") {
            writer.uint32(34).string(message.eventId);
        }
        if (message.page !== "0") {
            writer.uint32(40).int64(message.page);
        }
        if (message.perPage !== "0") {
            writer.uint32(48).int64(message.perPage);
        }
        if (message.deploymentId !== "0") {
            writer.uint32(56).int64(message.deploymentId);
        }
        if (message.snapshotEventId !== "") {
            writer.uint32(66).string(message.snapshotEventId);
        }
        return writer;
    },

    decode(input: BinaryReader | Uint8Array, length?: number): PaginateStreamAfterEventIdRequest {
        const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBasePaginateStreamAfterEventIdRequest();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1: {
                    if (tag !== 10) {
                        break;
                    }

                    message.tenantId = reader.string();
                    continue;
                }
                case 2: {
                    if (tag !== 18) {
                        break;
                    }

                    message.topic = reader.string();
                    continue;
                }
                case 3: {
                    if (tag !== 26) {
                        break;
                    }

                    message.stream = reader.string();
                    continue;
                }
                case 4: {
                    if (tag !== 34) {
                        break;
                    }

                    message.eventId = reader.string();
                    continue;
                }
                case 5: {
                    if (tag !== 40) {
                        break;
                    }

                    message.page = reader.int64().toString();
                    continue;
                }
                case 6: {
                    if (tag !== 48) {
                        break;
                    }

                    message.perPage = reader.int64().toString();
                    continue;
                }
                case 7: {
                    if (tag !== 56) {
                        break;
                    }

                    message.deploymentId = reader.int64().toString();
                    continue;
                }
                case 8: {
                    if (tag !== 66) {
                        break;
                    }

                    message.snapshotEventId = reader.string();
                    continue;
                }
            }
            if ((tag & 7) === 4 || tag === 0) {
                break;
            }
            reader.skip(tag & 7);
        }
        return message;
    },

    fromJSON(object: any): PaginateStreamAfterEventIdRequest {
        return {
            tenantId: isSet(object.tenantId) ? globalThis.String(object.tenantId) : "",
            topic: isSet(object.topic) ? globalThis.String(object.topic) : "",
            stream: isSet(object.stream) ? globalThis.String(object.stream) : "",
            eventId: isSet(object.eventId) ? globalThis.String(object.eventId) : "",
            page: isSet(object.page) ? globalThis.String(object.page) : "0",
            perPage: isSet(object.perPage) ? globalThis.String(object.perPage) : "0",
            deploymentId: isSet(object.deploymentId) ? globalThis.String(object.deploymentId) : "0",
            snapshotEventId: isSet(object.snapshotEventId)
                ? globalThis.String(object.snapshotEventId)
                : "",
        };
    },

    toJSON(message: PaginateStreamAfterEventIdRequest): unknown {
        const obj: any = {};
        if (message.tenantId !== "") {
            obj.tenantId = message.tenantId;
        }
        if (message.topic !== "") {
            obj.topic = message.topic;
        }
        if (message.stream !== "") {
            obj.stream = message.stream;
        }
        if (message.eventId !== "") {
            obj.eventId = message.eventId;
        }
        if (message.page !== "0") {
            obj.page = message.page;
        }
        if (message.perPage !== "0") {
            obj.perPage = message.perPage;
        }
        if (message.deploymentId !== "0") {
            obj.deploymentId = message.deploymentId;
        }
        if (message.snapshotEventId !== "") {
            obj.snapshotEventId = message.snapshotEventId;
        }
        return obj;
    },

    create(
        base?: DeepPartial<PaginateStreamAfterEventIdRequest>
    ): PaginateStreamAfterEventIdRequest {
        return PaginateStreamAfterEventIdRequest.fromPartial(base ?? {});
    },
    fromPartial(
        object: DeepPartial<PaginateStreamAfterEventIdRequest>
    ): PaginateStreamAfterEventIdRequest {
        const message = createBasePaginateStreamAfterEventIdRequest();
        message.tenantId = object.tenantId ?? "";
        message.topic = object.topic ?? "";
        message.stream = object.stream ?? "";
        message.eventId = object.eventId ?? "";
        message.page = object.page ?? "0";
        message.perPage = object.perPage ?? "0";
        message.deploymentId = object.deploymentId ?? "0";
        message.snapshotEventId = object.snapshotEventId ?? "";
        return message;
    },
};

function createBasePaginateStreamAfterEventIdResponse(): PaginateStreamAfterEventIdResponse {
    return { events: [], snapshotEventId: "" };
}

export const PaginateStreamAfterEventIdResponse: MessageFns<PaginateStreamAfterEventIdResponse> = {
    encode(
        message: PaginateStreamAfterEventIdResponse,
        writer: BinaryWriter = new BinaryWriter()
    ): BinaryWriter {
        for (const v of message.events) {
            Event.encode(v!, writer.uint32(10).fork()).join();
        }
        if (message.snapshotEventId !== "") {
            writer.uint32(18).string(message.snapshotEventId);
        }
        return writer;
    },

    decode(input: BinaryReader | Uint8Array, length?: number): PaginateStreamAfterEventIdResponse {
        const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBasePaginateStreamAfterEventIdResponse();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1: {
                    if (tag !== 10) {
                        break;
                    }

                    message.events.push(Event.decode(reader, reader.uint32()));
                    continue;
                }
                case 2: {
                    if (tag !== 18) {
                        break;
                    }

                    message.snapshotEventId = reader.string();
                    continue;
                }
            }
            if ((tag & 7) === 4 || tag === 0) {
                break;
            }
            reader.skip(tag & 7);
        }
        return message;
    },

    fromJSON(object: any): PaginateStreamAfterEventIdResponse {
        return {
            events: globalThis.Array.isArray(object?.events)
                ? object.events.map((e: any) => Event.fromJSON(e))
                : [],
            snapshotEventId: isSet(object.snapshotEventId)
                ? globalThis.String(object.snapshotEventId)
                : "",
        };
    },

    toJSON(message: PaginateStreamAfterEventIdResponse): unknown {
        const obj: any = {};
        if (message.events?.length) {
            obj.events = message.events.map(e => Event.toJSON(e));
        }
        if (message.snapshotEventId !== "") {
            obj.snapshotEventId = message.snapshotEventId;
        }
        return obj;
    },

    create(
        base?: DeepPartial<PaginateStreamAfterEventIdResponse>
    ): PaginateStreamAfterEventIdResponse {
        return PaginateStreamAfterEventIdResponse.fromPartial(base ?? {});
    },
    fromPartial(
        object: DeepPartial<PaginateStreamAfterEventIdResponse>
    ): PaginateStreamAfterEventIdResponse {
        const message = createBasePaginateStreamAfterEventIdResponse();
        message.events = object.events?.map(e => Event.fromPartial(e)) || [];
        message.snapshotEventId = object.snapshotEventId ?? "";
        return message;
    },
};

function createBasePaginateEventsRequest(): PaginateEventsRequest {
    return { tenantId: "", topic: "", types: [], page: "0", perPage: "0" };
}

export const PaginateEventsRequest: MessageFns<PaginateEventsRequest> = {
    encode(
        message: PaginateEventsRequest,
        writer: BinaryWriter = new BinaryWriter()
    ): BinaryWriter {
        if (message.tenantId !== "") {
            writer.uint32(10).string(message.tenantId);
        }
        if (message.topic !== "") {
            writer.uint32(18).string(message.topic);
        }
        for (const v of message.types) {
            writer.uint32(26).string(v!);
        }
        if (message.page !== "0") {
            writer.uint32(32).int64(message.page);
        }
        if (message.perPage !== "0") {
            writer.uint32(40).int64(message.perPage);
        }
        return writer;
    },

    decode(input: BinaryReader | Uint8Array, length?: number): PaginateEventsRequest {
        const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBasePaginateEventsRequest();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1: {
                    if (tag !== 10) {
                        break;
                    }

                    message.tenantId = reader.string();
                    continue;
                }
                case 2: {
                    if (tag !== 18) {
                        break;
                    }

                    message.topic = reader.string();
                    continue;
                }
                case 3: {
                    if (tag !== 26) {
                        break;
                    }

                    message.types.push(reader.string());
                    continue;
                }
                case 4: {
                    if (tag !== 32) {
                        break;
                    }

                    message.page = reader.int64().toString();
                    continue;
                }
                case 5: {
                    if (tag !== 40) {
                        break;
                    }

                    message.perPage = reader.int64().toString();
                    continue;
                }
            }
            if ((tag & 7) === 4 || tag === 0) {
                break;
            }
            reader.skip(tag & 7);
        }
        return message;
    },

    fromJSON(object: any): PaginateEventsRequest {
        return {
            tenantId: isSet(object.tenantId) ? globalThis.String(object.tenantId) : "",
            topic: isSet(object.topic) ? globalThis.String(object.topic) : "",
            types: globalThis.Array.isArray(object?.types)
                ? object.types.map((e: any) => globalThis.String(e))
                : [],
            page: isSet(object.page) ? globalThis.String(object.page) : "0",
            perPage: isSet(object.perPage) ? globalThis.String(object.perPage) : "0",
        };
    },

    toJSON(message: PaginateEventsRequest): unknown {
        const obj: any = {};
        if (message.tenantId !== "") {
            obj.tenantId = message.tenantId;
        }
        if (message.topic !== "") {
            obj.topic = message.topic;
        }
        if (message.types?.length) {
            obj.types = message.types;
        }
        if (message.page !== "0") {
            obj.page = message.page;
        }
        if (message.perPage !== "0") {
            obj.perPage = message.perPage;
        }
        return obj;
    },

    create(base?: DeepPartial<PaginateEventsRequest>): PaginateEventsRequest {
        return PaginateEventsRequest.fromPartial(base ?? {});
    },
    fromPartial(object: DeepPartial<PaginateEventsRequest>): PaginateEventsRequest {
        const message = createBasePaginateEventsRequest();
        message.tenantId = object.tenantId ?? "";
        message.topic = object.topic ?? "";
        message.types = object.types?.map(e => e) || [];
        message.page = object.page ?? "0";
        message.perPage = object.perPage ?? "0";
        return message;
    },
};

function createBasePaginateEventsResponse(): PaginateEventsResponse {
    return { events: [] };
}

export const PaginateEventsResponse: MessageFns<PaginateEventsResponse> = {
    encode(
        message: PaginateEventsResponse,
        writer: BinaryWriter = new BinaryWriter()
    ): BinaryWriter {
        for (const v of message.events) {
            Event.encode(v!, writer.uint32(10).fork()).join();
        }
        return writer;
    },

    decode(input: BinaryReader | Uint8Array, length?: number): PaginateEventsResponse {
        const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBasePaginateEventsResponse();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1: {
                    if (tag !== 10) {
                        break;
                    }

                    message.events.push(Event.decode(reader, reader.uint32()));
                    continue;
                }
            }
            if ((tag & 7) === 4 || tag === 0) {
                break;
            }
            reader.skip(tag & 7);
        }
        return message;
    },

    fromJSON(object: any): PaginateEventsResponse {
        return {
            events: globalThis.Array.isArray(object?.events)
                ? object.events.map((e: any) => Event.fromJSON(e))
                : [],
        };
    },

    toJSON(message: PaginateEventsResponse): unknown {
        const obj: any = {};
        if (message.events?.length) {
            obj.events = message.events.map(e => Event.toJSON(e));
        }
        return obj;
    },

    create(base?: DeepPartial<PaginateEventsResponse>): PaginateEventsResponse {
        return PaginateEventsResponse.fromPartial(base ?? {});
    },
    fromPartial(object: DeepPartial<PaginateEventsResponse>): PaginateEventsResponse {
        const message = createBasePaginateEventsResponse();
        message.events = object.events?.map(e => Event.fromPartial(e)) || [];
        return message;
    },
};

function createBasePaginateEventsAfterEventIdRequest(): PaginateEventsAfterEventIdRequest {
    return { tenantId: "", topic: "", types: [], eventId: "", page: "0", perPage: "0" };
}

export const PaginateEventsAfterEventIdRequest: MessageFns<PaginateEventsAfterEventIdRequest> = {
    encode(
        message: PaginateEventsAfterEventIdRequest,
        writer: BinaryWriter = new BinaryWriter()
    ): BinaryWriter {
        if (message.tenantId !== "") {
            writer.uint32(10).string(message.tenantId);
        }
        if (message.topic !== "") {
            writer.uint32(18).string(message.topic);
        }
        for (const v of message.types) {
            writer.uint32(26).string(v!);
        }
        if (message.eventId !== "") {
            writer.uint32(34).string(message.eventId);
        }
        if (message.page !== "0") {
            writer.uint32(40).int64(message.page);
        }
        if (message.perPage !== "0") {
            writer.uint32(48).int64(message.perPage);
        }
        return writer;
    },

    decode(input: BinaryReader | Uint8Array, length?: number): PaginateEventsAfterEventIdRequest {
        const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBasePaginateEventsAfterEventIdRequest();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1: {
                    if (tag !== 10) {
                        break;
                    }

                    message.tenantId = reader.string();
                    continue;
                }
                case 2: {
                    if (tag !== 18) {
                        break;
                    }

                    message.topic = reader.string();
                    continue;
                }
                case 3: {
                    if (tag !== 26) {
                        break;
                    }

                    message.types.push(reader.string());
                    continue;
                }
                case 4: {
                    if (tag !== 34) {
                        break;
                    }

                    message.eventId = reader.string();
                    continue;
                }
                case 5: {
                    if (tag !== 40) {
                        break;
                    }

                    message.page = reader.int64().toString();
                    continue;
                }
                case 6: {
                    if (tag !== 48) {
                        break;
                    }

                    message.perPage = reader.int64().toString();
                    continue;
                }
            }
            if ((tag & 7) === 4 || tag === 0) {
                break;
            }
            reader.skip(tag & 7);
        }
        return message;
    },

    fromJSON(object: any): PaginateEventsAfterEventIdRequest {
        return {
            tenantId: isSet(object.tenantId) ? globalThis.String(object.tenantId) : "",
            topic: isSet(object.topic) ? globalThis.String(object.topic) : "",
            types: globalThis.Array.isArray(object?.types)
                ? object.types.map((e: any) => globalThis.String(e))
                : [],
            eventId: isSet(object.eventId) ? globalThis.String(object.eventId) : "",
            page: isSet(object.page) ? globalThis.String(object.page) : "0",
            perPage: isSet(object.perPage) ? globalThis.String(object.perPage) : "0",
        };
    },

    toJSON(message: PaginateEventsAfterEventIdRequest): unknown {
        const obj: any = {};
        if (message.tenantId !== "") {
            obj.tenantId = message.tenantId;
        }
        if (message.topic !== "") {
            obj.topic = message.topic;
        }
        if (message.types?.length) {
            obj.types = message.types;
        }
        if (message.eventId !== "") {
            obj.eventId = message.eventId;
        }
        if (message.page !== "0") {
            obj.page = message.page;
        }
        if (message.perPage !== "0") {
            obj.perPage = message.perPage;
        }
        return obj;
    },

    create(
        base?: DeepPartial<PaginateEventsAfterEventIdRequest>
    ): PaginateEventsAfterEventIdRequest {
        return PaginateEventsAfterEventIdRequest.fromPartial(base ?? {});
    },
    fromPartial(
        object: DeepPartial<PaginateEventsAfterEventIdRequest>
    ): PaginateEventsAfterEventIdRequest {
        const message = createBasePaginateEventsAfterEventIdRequest();
        message.tenantId = object.tenantId ?? "";
        message.topic = object.topic ?? "";
        message.types = object.types?.map(e => e) || [];
        message.eventId = object.eventId ?? "";
        message.page = object.page ?? "0";
        message.perPage = object.perPage ?? "0";
        return message;
    },
};

function createBasePaginateEventsAfterEventIdResponse(): PaginateEventsAfterEventIdResponse {
    return { events: [] };
}

export const PaginateEventsAfterEventIdResponse: MessageFns<PaginateEventsAfterEventIdResponse> = {
    encode(
        message: PaginateEventsAfterEventIdResponse,
        writer: BinaryWriter = new BinaryWriter()
    ): BinaryWriter {
        for (const v of message.events) {
            Event.encode(v!, writer.uint32(10).fork()).join();
        }
        return writer;
    },

    decode(input: BinaryReader | Uint8Array, length?: number): PaginateEventsAfterEventIdResponse {
        const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBasePaginateEventsAfterEventIdResponse();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1: {
                    if (tag !== 10) {
                        break;
                    }

                    message.events.push(Event.decode(reader, reader.uint32()));
                    continue;
                }
            }
            if ((tag & 7) === 4 || tag === 0) {
                break;
            }
            reader.skip(tag & 7);
        }
        return message;
    },

    fromJSON(object: any): PaginateEventsAfterEventIdResponse {
        return {
            events: globalThis.Array.isArray(object?.events)
                ? object.events.map((e: any) => Event.fromJSON(e))
                : [],
        };
    },

    toJSON(message: PaginateEventsAfterEventIdResponse): unknown {
        const obj: any = {};
        if (message.events?.length) {
            obj.events = message.events.map(e => Event.toJSON(e));
        }
        return obj;
    },

    create(
        base?: DeepPartial<PaginateEventsAfterEventIdResponse>
    ): PaginateEventsAfterEventIdResponse {
        return PaginateEventsAfterEventIdResponse.fromPartial(base ?? {});
    },
    fromPartial(
        object: DeepPartial<PaginateEventsAfterEventIdResponse>
    ): PaginateEventsAfterEventIdResponse {
        const message = createBasePaginateEventsAfterEventIdResponse();
        message.events = object.events?.map(e => Event.fromPartial(e)) || [];
        return message;
    },
};

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

type DeepPartial<T> = T extends Builtin
    ? T
    : T extends globalThis.Array<infer U>
      ? globalThis.Array<DeepPartial<U>>
      : T extends ReadonlyArray<infer U>
        ? ReadonlyArray<DeepPartial<U>>
        : T extends { $case: string }
          ? { [K in keyof Omit<T, "$case">]?: DeepPartial<T[K]> } & { $case: T["$case"] }
          : T extends {}
            ? { [K in keyof T]?: DeepPartial<T[K]> }
            : Partial<T>;

function isSet(value: any): boolean {
    return value !== null && value !== undefined;
}

interface MessageFns<T> {
    encode(message: T, writer?: BinaryWriter): BinaryWriter;
    decode(input: BinaryReader | Uint8Array, length?: number): T;
    fromJSON(object: any): T;
    toJSON(message: T): unknown;
    create(base?: DeepPartial<T>): T;
    fromPartial(object: DeepPartial<T>): T;
}
