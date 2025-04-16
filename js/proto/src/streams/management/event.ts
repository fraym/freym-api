// Code generated by protoc-gen-ts_proto. DO NOT EDIT.
// versions:
//   protoc-gen-ts_proto  v2.7.0
//   protoc               v5.29.3
// source: streams/management/event.proto

/* eslint-disable */
import { BinaryReader, BinaryWriter } from "@bufbuild/protobuf/wire";

export interface GetEventRequest {
    tenantId: string;
    topic: string;
    id: string;
}

export interface GetLastEventRequest {
    tenantId: string;
    topic: string;
}

export interface GetLastEventByTypesRequest {
    tenantId: string;
    topic: string;
    types: string[];
}

/**
 * The Event is sent from server to client:
 * - Order is preserved per tenant and topic combination
 * - Persisted Events are send sequentially: the next event of a tenant and topic combination is sent to the client after a Handled message
 */
export interface Event {
    tenantId: string;
    topic: string;
    id: string;
    type: string;
    stream: string;
    reason: string;
    options: EventOptions | undefined;
    metadata: EventMetadata | undefined;
    payload: { [key: string]: EventPayload };
    raisedAt: string;
}

export interface Event_PayloadEntry {
    key: string;
    value: EventPayload | undefined;
}

export interface EventOptions {
    /**
     * Broadcast events are
     * - not persisted
     * - order is not guaranteed
     * - received by every subscriber of every group
     * - handling in clients is not guaranteed
     */
    broadcast: boolean;
}

export interface EventMetadata {
    correlationId: string;
    causationId: string;
    orderSerial: string;
    userId: string;
    /** when a deployment_id (!= 0) is provided the event will only be handled by subscriptions with a deployment_id that is equal or higher than the events deployment_id. */
    deploymentId: string;
}

export interface EventPayload {
    value: string;
    gdpr: EventGdprValue | undefined;
}

export interface EventGdprValue {
    id: string;
    default: string;
    isInvalidated: boolean;
}

function createBaseGetEventRequest(): GetEventRequest {
    return { tenantId: "", topic: "", id: "" };
}

export const GetEventRequest: MessageFns<GetEventRequest> = {
    encode(message: GetEventRequest, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
        if (message.tenantId !== "") {
            writer.uint32(10).string(message.tenantId);
        }
        if (message.topic !== "") {
            writer.uint32(18).string(message.topic);
        }
        if (message.id !== "") {
            writer.uint32(26).string(message.id);
        }
        return writer;
    },

    decode(input: BinaryReader | Uint8Array, length?: number): GetEventRequest {
        const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseGetEventRequest();
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

                    message.id = reader.string();
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

    fromJSON(object: any): GetEventRequest {
        return {
            tenantId: isSet(object.tenantId) ? globalThis.String(object.tenantId) : "",
            topic: isSet(object.topic) ? globalThis.String(object.topic) : "",
            id: isSet(object.id) ? globalThis.String(object.id) : "",
        };
    },

    toJSON(message: GetEventRequest): unknown {
        const obj: any = {};
        if (message.tenantId !== "") {
            obj.tenantId = message.tenantId;
        }
        if (message.topic !== "") {
            obj.topic = message.topic;
        }
        if (message.id !== "") {
            obj.id = message.id;
        }
        return obj;
    },

    create(base?: DeepPartial<GetEventRequest>): GetEventRequest {
        return GetEventRequest.fromPartial(base ?? {});
    },
    fromPartial(object: DeepPartial<GetEventRequest>): GetEventRequest {
        const message = createBaseGetEventRequest();
        message.tenantId = object.tenantId ?? "";
        message.topic = object.topic ?? "";
        message.id = object.id ?? "";
        return message;
    },
};

function createBaseGetLastEventRequest(): GetLastEventRequest {
    return { tenantId: "", topic: "" };
}

export const GetLastEventRequest: MessageFns<GetLastEventRequest> = {
    encode(message: GetLastEventRequest, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
        if (message.tenantId !== "") {
            writer.uint32(10).string(message.tenantId);
        }
        if (message.topic !== "") {
            writer.uint32(18).string(message.topic);
        }
        return writer;
    },

    decode(input: BinaryReader | Uint8Array, length?: number): GetLastEventRequest {
        const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseGetLastEventRequest();
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
            }
            if ((tag & 7) === 4 || tag === 0) {
                break;
            }
            reader.skip(tag & 7);
        }
        return message;
    },

    fromJSON(object: any): GetLastEventRequest {
        return {
            tenantId: isSet(object.tenantId) ? globalThis.String(object.tenantId) : "",
            topic: isSet(object.topic) ? globalThis.String(object.topic) : "",
        };
    },

    toJSON(message: GetLastEventRequest): unknown {
        const obj: any = {};
        if (message.tenantId !== "") {
            obj.tenantId = message.tenantId;
        }
        if (message.topic !== "") {
            obj.topic = message.topic;
        }
        return obj;
    },

    create(base?: DeepPartial<GetLastEventRequest>): GetLastEventRequest {
        return GetLastEventRequest.fromPartial(base ?? {});
    },
    fromPartial(object: DeepPartial<GetLastEventRequest>): GetLastEventRequest {
        const message = createBaseGetLastEventRequest();
        message.tenantId = object.tenantId ?? "";
        message.topic = object.topic ?? "";
        return message;
    },
};

function createBaseGetLastEventByTypesRequest(): GetLastEventByTypesRequest {
    return { tenantId: "", topic: "", types: [] };
}

export const GetLastEventByTypesRequest: MessageFns<GetLastEventByTypesRequest> = {
    encode(
        message: GetLastEventByTypesRequest,
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
        return writer;
    },

    decode(input: BinaryReader | Uint8Array, length?: number): GetLastEventByTypesRequest {
        const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseGetLastEventByTypesRequest();
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
            }
            if ((tag & 7) === 4 || tag === 0) {
                break;
            }
            reader.skip(tag & 7);
        }
        return message;
    },

    fromJSON(object: any): GetLastEventByTypesRequest {
        return {
            tenantId: isSet(object.tenantId) ? globalThis.String(object.tenantId) : "",
            topic: isSet(object.topic) ? globalThis.String(object.topic) : "",
            types: globalThis.Array.isArray(object?.types)
                ? object.types.map((e: any) => globalThis.String(e))
                : [],
        };
    },

    toJSON(message: GetLastEventByTypesRequest): unknown {
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
        return obj;
    },

    create(base?: DeepPartial<GetLastEventByTypesRequest>): GetLastEventByTypesRequest {
        return GetLastEventByTypesRequest.fromPartial(base ?? {});
    },
    fromPartial(object: DeepPartial<GetLastEventByTypesRequest>): GetLastEventByTypesRequest {
        const message = createBaseGetLastEventByTypesRequest();
        message.tenantId = object.tenantId ?? "";
        message.topic = object.topic ?? "";
        message.types = object.types?.map(e => e) || [];
        return message;
    },
};

function createBaseEvent(): Event {
    return {
        tenantId: "",
        topic: "",
        id: "",
        type: "",
        stream: "",
        reason: "",
        options: undefined,
        metadata: undefined,
        payload: {},
        raisedAt: "0",
    };
}

export const Event: MessageFns<Event> = {
    encode(message: Event, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
        if (message.tenantId !== "") {
            writer.uint32(10).string(message.tenantId);
        }
        if (message.topic !== "") {
            writer.uint32(18).string(message.topic);
        }
        if (message.id !== "") {
            writer.uint32(26).string(message.id);
        }
        if (message.type !== "") {
            writer.uint32(34).string(message.type);
        }
        if (message.stream !== "") {
            writer.uint32(42).string(message.stream);
        }
        if (message.reason !== "") {
            writer.uint32(50).string(message.reason);
        }
        if (message.options !== undefined) {
            EventOptions.encode(message.options, writer.uint32(58).fork()).join();
        }
        if (message.metadata !== undefined) {
            EventMetadata.encode(message.metadata, writer.uint32(66).fork()).join();
        }
        Object.entries(message.payload).forEach(([key, value]) => {
            Event_PayloadEntry.encode({ key: key as any, value }, writer.uint32(74).fork()).join();
        });
        if (message.raisedAt !== "0") {
            writer.uint32(80).int64(message.raisedAt);
        }
        return writer;
    },

    decode(input: BinaryReader | Uint8Array, length?: number): Event {
        const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseEvent();
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

                    message.id = reader.string();
                    continue;
                }
                case 4: {
                    if (tag !== 34) {
                        break;
                    }

                    message.type = reader.string();
                    continue;
                }
                case 5: {
                    if (tag !== 42) {
                        break;
                    }

                    message.stream = reader.string();
                    continue;
                }
                case 6: {
                    if (tag !== 50) {
                        break;
                    }

                    message.reason = reader.string();
                    continue;
                }
                case 7: {
                    if (tag !== 58) {
                        break;
                    }

                    message.options = EventOptions.decode(reader, reader.uint32());
                    continue;
                }
                case 8: {
                    if (tag !== 66) {
                        break;
                    }

                    message.metadata = EventMetadata.decode(reader, reader.uint32());
                    continue;
                }
                case 9: {
                    if (tag !== 74) {
                        break;
                    }

                    const entry9 = Event_PayloadEntry.decode(reader, reader.uint32());
                    if (entry9.value !== undefined) {
                        message.payload[entry9.key] = entry9.value;
                    }
                    continue;
                }
                case 10: {
                    if (tag !== 80) {
                        break;
                    }

                    message.raisedAt = reader.int64().toString();
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

    fromJSON(object: any): Event {
        return {
            tenantId: isSet(object.tenantId) ? globalThis.String(object.tenantId) : "",
            topic: isSet(object.topic) ? globalThis.String(object.topic) : "",
            id: isSet(object.id) ? globalThis.String(object.id) : "",
            type: isSet(object.type) ? globalThis.String(object.type) : "",
            stream: isSet(object.stream) ? globalThis.String(object.stream) : "",
            reason: isSet(object.reason) ? globalThis.String(object.reason) : "",
            options: isSet(object.options) ? EventOptions.fromJSON(object.options) : undefined,
            metadata: isSet(object.metadata) ? EventMetadata.fromJSON(object.metadata) : undefined,
            payload: isObject(object.payload)
                ? Object.entries(object.payload).reduce<{ [key: string]: EventPayload }>(
                      (acc, [key, value]) => {
                          acc[key] = EventPayload.fromJSON(value);
                          return acc;
                      },
                      {}
                  )
                : {},
            raisedAt: isSet(object.raisedAt) ? globalThis.String(object.raisedAt) : "0",
        };
    },

    toJSON(message: Event): unknown {
        const obj: any = {};
        if (message.tenantId !== "") {
            obj.tenantId = message.tenantId;
        }
        if (message.topic !== "") {
            obj.topic = message.topic;
        }
        if (message.id !== "") {
            obj.id = message.id;
        }
        if (message.type !== "") {
            obj.type = message.type;
        }
        if (message.stream !== "") {
            obj.stream = message.stream;
        }
        if (message.reason !== "") {
            obj.reason = message.reason;
        }
        if (message.options !== undefined) {
            obj.options = EventOptions.toJSON(message.options);
        }
        if (message.metadata !== undefined) {
            obj.metadata = EventMetadata.toJSON(message.metadata);
        }
        if (message.payload) {
            const entries = Object.entries(message.payload);
            if (entries.length > 0) {
                obj.payload = {};
                entries.forEach(([k, v]) => {
                    obj.payload[k] = EventPayload.toJSON(v);
                });
            }
        }
        if (message.raisedAt !== "0") {
            obj.raisedAt = message.raisedAt;
        }
        return obj;
    },

    create(base?: DeepPartial<Event>): Event {
        return Event.fromPartial(base ?? {});
    },
    fromPartial(object: DeepPartial<Event>): Event {
        const message = createBaseEvent();
        message.tenantId = object.tenantId ?? "";
        message.topic = object.topic ?? "";
        message.id = object.id ?? "";
        message.type = object.type ?? "";
        message.stream = object.stream ?? "";
        message.reason = object.reason ?? "";
        message.options =
            object.options !== undefined && object.options !== null
                ? EventOptions.fromPartial(object.options)
                : undefined;
        message.metadata =
            object.metadata !== undefined && object.metadata !== null
                ? EventMetadata.fromPartial(object.metadata)
                : undefined;
        message.payload = Object.entries(object.payload ?? {}).reduce<{
            [key: string]: EventPayload;
        }>((acc, [key, value]) => {
            if (value !== undefined) {
                acc[key] = EventPayload.fromPartial(value);
            }
            return acc;
        }, {});
        message.raisedAt = object.raisedAt ?? "0";
        return message;
    },
};

function createBaseEvent_PayloadEntry(): Event_PayloadEntry {
    return { key: "", value: undefined };
}

export const Event_PayloadEntry: MessageFns<Event_PayloadEntry> = {
    encode(message: Event_PayloadEntry, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
        if (message.key !== "") {
            writer.uint32(10).string(message.key);
        }
        if (message.value !== undefined) {
            EventPayload.encode(message.value, writer.uint32(18).fork()).join();
        }
        return writer;
    },

    decode(input: BinaryReader | Uint8Array, length?: number): Event_PayloadEntry {
        const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseEvent_PayloadEntry();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1: {
                    if (tag !== 10) {
                        break;
                    }

                    message.key = reader.string();
                    continue;
                }
                case 2: {
                    if (tag !== 18) {
                        break;
                    }

                    message.value = EventPayload.decode(reader, reader.uint32());
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

    fromJSON(object: any): Event_PayloadEntry {
        return {
            key: isSet(object.key) ? globalThis.String(object.key) : "",
            value: isSet(object.value) ? EventPayload.fromJSON(object.value) : undefined,
        };
    },

    toJSON(message: Event_PayloadEntry): unknown {
        const obj: any = {};
        if (message.key !== "") {
            obj.key = message.key;
        }
        if (message.value !== undefined) {
            obj.value = EventPayload.toJSON(message.value);
        }
        return obj;
    },

    create(base?: DeepPartial<Event_PayloadEntry>): Event_PayloadEntry {
        return Event_PayloadEntry.fromPartial(base ?? {});
    },
    fromPartial(object: DeepPartial<Event_PayloadEntry>): Event_PayloadEntry {
        const message = createBaseEvent_PayloadEntry();
        message.key = object.key ?? "";
        message.value =
            object.value !== undefined && object.value !== null
                ? EventPayload.fromPartial(object.value)
                : undefined;
        return message;
    },
};

function createBaseEventOptions(): EventOptions {
    return { broadcast: false };
}

export const EventOptions: MessageFns<EventOptions> = {
    encode(message: EventOptions, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
        if (message.broadcast !== false) {
            writer.uint32(8).bool(message.broadcast);
        }
        return writer;
    },

    decode(input: BinaryReader | Uint8Array, length?: number): EventOptions {
        const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseEventOptions();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1: {
                    if (tag !== 8) {
                        break;
                    }

                    message.broadcast = reader.bool();
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

    fromJSON(object: any): EventOptions {
        return {
            broadcast: isSet(object.broadcast) ? globalThis.Boolean(object.broadcast) : false,
        };
    },

    toJSON(message: EventOptions): unknown {
        const obj: any = {};
        if (message.broadcast !== false) {
            obj.broadcast = message.broadcast;
        }
        return obj;
    },

    create(base?: DeepPartial<EventOptions>): EventOptions {
        return EventOptions.fromPartial(base ?? {});
    },
    fromPartial(object: DeepPartial<EventOptions>): EventOptions {
        const message = createBaseEventOptions();
        message.broadcast = object.broadcast ?? false;
        return message;
    },
};

function createBaseEventMetadata(): EventMetadata {
    return { correlationId: "", causationId: "", orderSerial: "0", userId: "", deploymentId: "0" };
}

export const EventMetadata: MessageFns<EventMetadata> = {
    encode(message: EventMetadata, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
        if (message.correlationId !== "") {
            writer.uint32(10).string(message.correlationId);
        }
        if (message.causationId !== "") {
            writer.uint32(18).string(message.causationId);
        }
        if (message.orderSerial !== "0") {
            writer.uint32(24).int64(message.orderSerial);
        }
        if (message.userId !== "") {
            writer.uint32(34).string(message.userId);
        }
        if (message.deploymentId !== "0") {
            writer.uint32(40).int64(message.deploymentId);
        }
        return writer;
    },

    decode(input: BinaryReader | Uint8Array, length?: number): EventMetadata {
        const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseEventMetadata();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1: {
                    if (tag !== 10) {
                        break;
                    }

                    message.correlationId = reader.string();
                    continue;
                }
                case 2: {
                    if (tag !== 18) {
                        break;
                    }

                    message.causationId = reader.string();
                    continue;
                }
                case 3: {
                    if (tag !== 24) {
                        break;
                    }

                    message.orderSerial = reader.int64().toString();
                    continue;
                }
                case 4: {
                    if (tag !== 34) {
                        break;
                    }

                    message.userId = reader.string();
                    continue;
                }
                case 5: {
                    if (tag !== 40) {
                        break;
                    }

                    message.deploymentId = reader.int64().toString();
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

    fromJSON(object: any): EventMetadata {
        return {
            correlationId: isSet(object.correlationId)
                ? globalThis.String(object.correlationId)
                : "",
            causationId: isSet(object.causationId) ? globalThis.String(object.causationId) : "",
            orderSerial: isSet(object.orderSerial) ? globalThis.String(object.orderSerial) : "0",
            userId: isSet(object.userId) ? globalThis.String(object.userId) : "",
            deploymentId: isSet(object.deploymentId) ? globalThis.String(object.deploymentId) : "0",
        };
    },

    toJSON(message: EventMetadata): unknown {
        const obj: any = {};
        if (message.correlationId !== "") {
            obj.correlationId = message.correlationId;
        }
        if (message.causationId !== "") {
            obj.causationId = message.causationId;
        }
        if (message.orderSerial !== "0") {
            obj.orderSerial = message.orderSerial;
        }
        if (message.userId !== "") {
            obj.userId = message.userId;
        }
        if (message.deploymentId !== "0") {
            obj.deploymentId = message.deploymentId;
        }
        return obj;
    },

    create(base?: DeepPartial<EventMetadata>): EventMetadata {
        return EventMetadata.fromPartial(base ?? {});
    },
    fromPartial(object: DeepPartial<EventMetadata>): EventMetadata {
        const message = createBaseEventMetadata();
        message.correlationId = object.correlationId ?? "";
        message.causationId = object.causationId ?? "";
        message.orderSerial = object.orderSerial ?? "0";
        message.userId = object.userId ?? "";
        message.deploymentId = object.deploymentId ?? "0";
        return message;
    },
};

function createBaseEventPayload(): EventPayload {
    return { value: "", gdpr: undefined };
}

export const EventPayload: MessageFns<EventPayload> = {
    encode(message: EventPayload, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
        if (message.value !== "") {
            writer.uint32(10).string(message.value);
        }
        if (message.gdpr !== undefined) {
            EventGdprValue.encode(message.gdpr, writer.uint32(18).fork()).join();
        }
        return writer;
    },

    decode(input: BinaryReader | Uint8Array, length?: number): EventPayload {
        const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseEventPayload();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1: {
                    if (tag !== 10) {
                        break;
                    }

                    message.value = reader.string();
                    continue;
                }
                case 2: {
                    if (tag !== 18) {
                        break;
                    }

                    message.gdpr = EventGdprValue.decode(reader, reader.uint32());
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

    fromJSON(object: any): EventPayload {
        return {
            value: isSet(object.value) ? globalThis.String(object.value) : "",
            gdpr: isSet(object.gdpr) ? EventGdprValue.fromJSON(object.gdpr) : undefined,
        };
    },

    toJSON(message: EventPayload): unknown {
        const obj: any = {};
        if (message.value !== "") {
            obj.value = message.value;
        }
        if (message.gdpr !== undefined) {
            obj.gdpr = EventGdprValue.toJSON(message.gdpr);
        }
        return obj;
    },

    create(base?: DeepPartial<EventPayload>): EventPayload {
        return EventPayload.fromPartial(base ?? {});
    },
    fromPartial(object: DeepPartial<EventPayload>): EventPayload {
        const message = createBaseEventPayload();
        message.value = object.value ?? "";
        message.gdpr =
            object.gdpr !== undefined && object.gdpr !== null
                ? EventGdprValue.fromPartial(object.gdpr)
                : undefined;
        return message;
    },
};

function createBaseEventGdprValue(): EventGdprValue {
    return { id: "", default: "", isInvalidated: false };
}

export const EventGdprValue: MessageFns<EventGdprValue> = {
    encode(message: EventGdprValue, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
        if (message.id !== "") {
            writer.uint32(10).string(message.id);
        }
        if (message.default !== "") {
            writer.uint32(18).string(message.default);
        }
        if (message.isInvalidated !== false) {
            writer.uint32(24).bool(message.isInvalidated);
        }
        return writer;
    },

    decode(input: BinaryReader | Uint8Array, length?: number): EventGdprValue {
        const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseEventGdprValue();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1: {
                    if (tag !== 10) {
                        break;
                    }

                    message.id = reader.string();
                    continue;
                }
                case 2: {
                    if (tag !== 18) {
                        break;
                    }

                    message.default = reader.string();
                    continue;
                }
                case 3: {
                    if (tag !== 24) {
                        break;
                    }

                    message.isInvalidated = reader.bool();
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

    fromJSON(object: any): EventGdprValue {
        return {
            id: isSet(object.id) ? globalThis.String(object.id) : "",
            default: isSet(object.default) ? globalThis.String(object.default) : "",
            isInvalidated: isSet(object.isInvalidated)
                ? globalThis.Boolean(object.isInvalidated)
                : false,
        };
    },

    toJSON(message: EventGdprValue): unknown {
        const obj: any = {};
        if (message.id !== "") {
            obj.id = message.id;
        }
        if (message.default !== "") {
            obj.default = message.default;
        }
        if (message.isInvalidated !== false) {
            obj.isInvalidated = message.isInvalidated;
        }
        return obj;
    },

    create(base?: DeepPartial<EventGdprValue>): EventGdprValue {
        return EventGdprValue.fromPartial(base ?? {});
    },
    fromPartial(object: DeepPartial<EventGdprValue>): EventGdprValue {
        const message = createBaseEventGdprValue();
        message.id = object.id ?? "";
        message.default = object.default ?? "";
        message.isInvalidated = object.isInvalidated ?? false;
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

function isObject(value: any): boolean {
    return typeof value === "object" && value !== null;
}

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
