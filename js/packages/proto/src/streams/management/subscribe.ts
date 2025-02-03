// Code generated by protoc-gen-ts_proto. DO NOT EDIT.
// versions:
//   protoc-gen-ts_proto  v2.6.1
//   protoc               v5.27.3
// source: streams/management/subscribe.proto

/* eslint-disable */
import { BinaryReader, BinaryWriter } from "@bufbuild/protobuf/wire";
import { Event } from "./event";

export interface SubscribeRequest {
    /**
     * Events have to be requeued in case of:
     * - connection is closed before Handled message is received
     * - NotHandled message is received
     */
    payload?:
        | //
        /** If not received within timeout after stream is opened, the stream is closed by the server. */
        { $case: "subscribe"; subscribe: Subscribe } //
        /**
         * The client sends a Handled message as soon as the current event of a tenants topic is successfully handled.
         * The client sends a Handled message with error when it failed to handle the current event of a tenants topic.
         * The server will panic if no subscriber is able to handle an event, this will increase awareness of implementation errors.
         */
        | { $case: "handled"; handled: Handled }
        | undefined;
}

export interface Subscribe {
    metadata: SubscribeMetadata | undefined;
    topics: string[];
}

export interface SubscribeMetadata {
    group: string;
    subscriberId: string;
    /** when a deployment_id (!= 0) is provided the event will only be handled by subscriptions with a deployment_id that is equal or higher than the events deployment_id. */
    deploymentId: string;
}

export interface Handled {
    tenantId: string;
    topic: string;
    error: string;
}

/** responses */
export interface SubscribeResponse {
    payload?:
        | { $case: "subscribed"; subscribed: Subscribed }
        | { $case: "panic"; panic: Panic }
        | {
              $case: "event";
              event: Event;
          }
        | undefined;
}

/**
 * The server will send a Subscribed in response to a new subscription.
 * In case of errors the error field is populated with a corresponding error message.
 * If the client does not receive this message within 5 sek after sending the Subscribe message, the stream is closed by the client.
 */
export interface Subscribed {
    /** The client will close the stream and retry the subscription non empty errors. */
    error: string;
}

/**
 * The server will send a Panic message when no other subscription was able to sucessfully handle an event.
 * Clients offer hooks that allow apps to observe panics and perform actions.
 */
export interface Panic {
    reason: string;
}

function createBaseSubscribeRequest(): SubscribeRequest {
    return { payload: undefined };
}

export const SubscribeRequest: MessageFns<SubscribeRequest> = {
    encode(message: SubscribeRequest, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
        switch (message.payload?.$case) {
            case "subscribe":
                Subscribe.encode(message.payload.subscribe, writer.uint32(10).fork()).join();
                break;
            case "handled":
                Handled.encode(message.payload.handled, writer.uint32(18).fork()).join();
                break;
        }
        return writer;
    },

    decode(input: BinaryReader | Uint8Array, length?: number): SubscribeRequest {
        const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseSubscribeRequest();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1: {
                    if (tag !== 10) {
                        break;
                    }

                    message.payload = {
                        $case: "subscribe",
                        subscribe: Subscribe.decode(reader, reader.uint32()),
                    };
                    continue;
                }
                case 2: {
                    if (tag !== 18) {
                        break;
                    }

                    message.payload = {
                        $case: "handled",
                        handled: Handled.decode(reader, reader.uint32()),
                    };
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

    fromJSON(object: any): SubscribeRequest {
        return {
            payload: isSet(object.subscribe)
                ? { $case: "subscribe", subscribe: Subscribe.fromJSON(object.subscribe) }
                : isSet(object.handled)
                  ? { $case: "handled", handled: Handled.fromJSON(object.handled) }
                  : undefined,
        };
    },

    toJSON(message: SubscribeRequest): unknown {
        const obj: any = {};
        if (message.payload?.$case === "subscribe") {
            obj.subscribe = Subscribe.toJSON(message.payload.subscribe);
        } else if (message.payload?.$case === "handled") {
            obj.handled = Handled.toJSON(message.payload.handled);
        }
        return obj;
    },

    create(base?: DeepPartial<SubscribeRequest>): SubscribeRequest {
        return SubscribeRequest.fromPartial(base ?? {});
    },
    fromPartial(object: DeepPartial<SubscribeRequest>): SubscribeRequest {
        const message = createBaseSubscribeRequest();
        switch (object.payload?.$case) {
            case "subscribe": {
                if (object.payload?.subscribe !== undefined && object.payload?.subscribe !== null) {
                    message.payload = {
                        $case: "subscribe",
                        subscribe: Subscribe.fromPartial(object.payload.subscribe),
                    };
                }
                break;
            }
            case "handled": {
                if (object.payload?.handled !== undefined && object.payload?.handled !== null) {
                    message.payload = {
                        $case: "handled",
                        handled: Handled.fromPartial(object.payload.handled),
                    };
                }
                break;
            }
        }
        return message;
    },
};

function createBaseSubscribe(): Subscribe {
    return { metadata: undefined, topics: [] };
}

export const Subscribe: MessageFns<Subscribe> = {
    encode(message: Subscribe, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
        if (message.metadata !== undefined) {
            SubscribeMetadata.encode(message.metadata, writer.uint32(10).fork()).join();
        }
        for (const v of message.topics) {
            writer.uint32(18).string(v!);
        }
        return writer;
    },

    decode(input: BinaryReader | Uint8Array, length?: number): Subscribe {
        const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseSubscribe();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1: {
                    if (tag !== 10) {
                        break;
                    }

                    message.metadata = SubscribeMetadata.decode(reader, reader.uint32());
                    continue;
                }
                case 2: {
                    if (tag !== 18) {
                        break;
                    }

                    message.topics.push(reader.string());
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

    fromJSON(object: any): Subscribe {
        return {
            metadata: isSet(object.metadata)
                ? SubscribeMetadata.fromJSON(object.metadata)
                : undefined,
            topics: globalThis.Array.isArray(object?.topics)
                ? object.topics.map((e: any) => globalThis.String(e))
                : [],
        };
    },

    toJSON(message: Subscribe): unknown {
        const obj: any = {};
        if (message.metadata !== undefined) {
            obj.metadata = SubscribeMetadata.toJSON(message.metadata);
        }
        if (message.topics?.length) {
            obj.topics = message.topics;
        }
        return obj;
    },

    create(base?: DeepPartial<Subscribe>): Subscribe {
        return Subscribe.fromPartial(base ?? {});
    },
    fromPartial(object: DeepPartial<Subscribe>): Subscribe {
        const message = createBaseSubscribe();
        message.metadata =
            object.metadata !== undefined && object.metadata !== null
                ? SubscribeMetadata.fromPartial(object.metadata)
                : undefined;
        message.topics = object.topics?.map(e => e) || [];
        return message;
    },
};

function createBaseSubscribeMetadata(): SubscribeMetadata {
    return { group: "", subscriberId: "", deploymentId: "0" };
}

export const SubscribeMetadata: MessageFns<SubscribeMetadata> = {
    encode(message: SubscribeMetadata, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
        if (message.group !== "") {
            writer.uint32(10).string(message.group);
        }
        if (message.subscriberId !== "") {
            writer.uint32(18).string(message.subscriberId);
        }
        if (message.deploymentId !== "0") {
            writer.uint32(24).int64(message.deploymentId);
        }
        return writer;
    },

    decode(input: BinaryReader | Uint8Array, length?: number): SubscribeMetadata {
        const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseSubscribeMetadata();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1: {
                    if (tag !== 10) {
                        break;
                    }

                    message.group = reader.string();
                    continue;
                }
                case 2: {
                    if (tag !== 18) {
                        break;
                    }

                    message.subscriberId = reader.string();
                    continue;
                }
                case 3: {
                    if (tag !== 24) {
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

    fromJSON(object: any): SubscribeMetadata {
        return {
            group: isSet(object.group) ? globalThis.String(object.group) : "",
            subscriberId: isSet(object.subscriberId) ? globalThis.String(object.subscriberId) : "",
            deploymentId: isSet(object.deploymentId) ? globalThis.String(object.deploymentId) : "0",
        };
    },

    toJSON(message: SubscribeMetadata): unknown {
        const obj: any = {};
        if (message.group !== "") {
            obj.group = message.group;
        }
        if (message.subscriberId !== "") {
            obj.subscriberId = message.subscriberId;
        }
        if (message.deploymentId !== "0") {
            obj.deploymentId = message.deploymentId;
        }
        return obj;
    },

    create(base?: DeepPartial<SubscribeMetadata>): SubscribeMetadata {
        return SubscribeMetadata.fromPartial(base ?? {});
    },
    fromPartial(object: DeepPartial<SubscribeMetadata>): SubscribeMetadata {
        const message = createBaseSubscribeMetadata();
        message.group = object.group ?? "";
        message.subscriberId = object.subscriberId ?? "";
        message.deploymentId = object.deploymentId ?? "0";
        return message;
    },
};

function createBaseHandled(): Handled {
    return { tenantId: "", topic: "", error: "" };
}

export const Handled: MessageFns<Handled> = {
    encode(message: Handled, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
        if (message.tenantId !== "") {
            writer.uint32(10).string(message.tenantId);
        }
        if (message.topic !== "") {
            writer.uint32(18).string(message.topic);
        }
        if (message.error !== "") {
            writer.uint32(26).string(message.error);
        }
        return writer;
    },

    decode(input: BinaryReader | Uint8Array, length?: number): Handled {
        const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseHandled();
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

                    message.error = reader.string();
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

    fromJSON(object: any): Handled {
        return {
            tenantId: isSet(object.tenantId) ? globalThis.String(object.tenantId) : "",
            topic: isSet(object.topic) ? globalThis.String(object.topic) : "",
            error: isSet(object.error) ? globalThis.String(object.error) : "",
        };
    },

    toJSON(message: Handled): unknown {
        const obj: any = {};
        if (message.tenantId !== "") {
            obj.tenantId = message.tenantId;
        }
        if (message.topic !== "") {
            obj.topic = message.topic;
        }
        if (message.error !== "") {
            obj.error = message.error;
        }
        return obj;
    },

    create(base?: DeepPartial<Handled>): Handled {
        return Handled.fromPartial(base ?? {});
    },
    fromPartial(object: DeepPartial<Handled>): Handled {
        const message = createBaseHandled();
        message.tenantId = object.tenantId ?? "";
        message.topic = object.topic ?? "";
        message.error = object.error ?? "";
        return message;
    },
};

function createBaseSubscribeResponse(): SubscribeResponse {
    return { payload: undefined };
}

export const SubscribeResponse: MessageFns<SubscribeResponse> = {
    encode(message: SubscribeResponse, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
        switch (message.payload?.$case) {
            case "subscribed":
                Subscribed.encode(message.payload.subscribed, writer.uint32(10).fork()).join();
                break;
            case "panic":
                Panic.encode(message.payload.panic, writer.uint32(18).fork()).join();
                break;
            case "event":
                Event.encode(message.payload.event, writer.uint32(26).fork()).join();
                break;
        }
        return writer;
    },

    decode(input: BinaryReader | Uint8Array, length?: number): SubscribeResponse {
        const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseSubscribeResponse();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1: {
                    if (tag !== 10) {
                        break;
                    }

                    message.payload = {
                        $case: "subscribed",
                        subscribed: Subscribed.decode(reader, reader.uint32()),
                    };
                    continue;
                }
                case 2: {
                    if (tag !== 18) {
                        break;
                    }

                    message.payload = {
                        $case: "panic",
                        panic: Panic.decode(reader, reader.uint32()),
                    };
                    continue;
                }
                case 3: {
                    if (tag !== 26) {
                        break;
                    }

                    message.payload = {
                        $case: "event",
                        event: Event.decode(reader, reader.uint32()),
                    };
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

    fromJSON(object: any): SubscribeResponse {
        return {
            payload: isSet(object.subscribed)
                ? { $case: "subscribed", subscribed: Subscribed.fromJSON(object.subscribed) }
                : isSet(object.panic)
                  ? { $case: "panic", panic: Panic.fromJSON(object.panic) }
                  : isSet(object.event)
                    ? { $case: "event", event: Event.fromJSON(object.event) }
                    : undefined,
        };
    },

    toJSON(message: SubscribeResponse): unknown {
        const obj: any = {};
        if (message.payload?.$case === "subscribed") {
            obj.subscribed = Subscribed.toJSON(message.payload.subscribed);
        } else if (message.payload?.$case === "panic") {
            obj.panic = Panic.toJSON(message.payload.panic);
        } else if (message.payload?.$case === "event") {
            obj.event = Event.toJSON(message.payload.event);
        }
        return obj;
    },

    create(base?: DeepPartial<SubscribeResponse>): SubscribeResponse {
        return SubscribeResponse.fromPartial(base ?? {});
    },
    fromPartial(object: DeepPartial<SubscribeResponse>): SubscribeResponse {
        const message = createBaseSubscribeResponse();
        switch (object.payload?.$case) {
            case "subscribed": {
                if (
                    object.payload?.subscribed !== undefined &&
                    object.payload?.subscribed !== null
                ) {
                    message.payload = {
                        $case: "subscribed",
                        subscribed: Subscribed.fromPartial(object.payload.subscribed),
                    };
                }
                break;
            }
            case "panic": {
                if (object.payload?.panic !== undefined && object.payload?.panic !== null) {
                    message.payload = {
                        $case: "panic",
                        panic: Panic.fromPartial(object.payload.panic),
                    };
                }
                break;
            }
            case "event": {
                if (object.payload?.event !== undefined && object.payload?.event !== null) {
                    message.payload = {
                        $case: "event",
                        event: Event.fromPartial(object.payload.event),
                    };
                }
                break;
            }
        }
        return message;
    },
};

function createBaseSubscribed(): Subscribed {
    return { error: "" };
}

export const Subscribed: MessageFns<Subscribed> = {
    encode(message: Subscribed, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
        if (message.error !== "") {
            writer.uint32(10).string(message.error);
        }
        return writer;
    },

    decode(input: BinaryReader | Uint8Array, length?: number): Subscribed {
        const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseSubscribed();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1: {
                    if (tag !== 10) {
                        break;
                    }

                    message.error = reader.string();
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

    fromJSON(object: any): Subscribed {
        return { error: isSet(object.error) ? globalThis.String(object.error) : "" };
    },

    toJSON(message: Subscribed): unknown {
        const obj: any = {};
        if (message.error !== "") {
            obj.error = message.error;
        }
        return obj;
    },

    create(base?: DeepPartial<Subscribed>): Subscribed {
        return Subscribed.fromPartial(base ?? {});
    },
    fromPartial(object: DeepPartial<Subscribed>): Subscribed {
        const message = createBaseSubscribed();
        message.error = object.error ?? "";
        return message;
    },
};

function createBasePanic(): Panic {
    return { reason: "" };
}

export const Panic: MessageFns<Panic> = {
    encode(message: Panic, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
        if (message.reason !== "") {
            writer.uint32(10).string(message.reason);
        }
        return writer;
    },

    decode(input: BinaryReader | Uint8Array, length?: number): Panic {
        const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBasePanic();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1: {
                    if (tag !== 10) {
                        break;
                    }

                    message.reason = reader.string();
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

    fromJSON(object: any): Panic {
        return { reason: isSet(object.reason) ? globalThis.String(object.reason) : "" };
    },

    toJSON(message: Panic): unknown {
        const obj: any = {};
        if (message.reason !== "") {
            obj.reason = message.reason;
        }
        return obj;
    },

    create(base?: DeepPartial<Panic>): Panic {
        return Panic.fromPartial(base ?? {});
    },
    fromPartial(object: DeepPartial<Panic>): Panic {
        const message = createBasePanic();
        message.reason = object.reason ?? "";
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
