// Code generated by protoc-gen-ts_proto. DO NOT EDIT.
// versions:
//   protoc-gen-ts_proto  v2.6.1
//   protoc               v5.27.3
// source: streams/management/gdpr.proto

/* eslint-disable */
import { BinaryReader, BinaryWriter } from "@bufbuild/protobuf/wire";

export interface IntroduceGdprOnEventFieldRequest {
    tenantId: string;
    topic: string;
    eventId: string;
    fieldName: string;
    defaultValue: string;
}

export interface IntroduceGdprOnEventFieldResponse {}

export interface InvalidateGdprRequest {
    tenantId: string;
    topic: string;
    gdprId: string;
}

export interface InvalidateGdprResponse {}

function createBaseIntroduceGdprOnEventFieldRequest(): IntroduceGdprOnEventFieldRequest {
    return { tenantId: "", topic: "", eventId: "", fieldName: "", defaultValue: "" };
}

export const IntroduceGdprOnEventFieldRequest: MessageFns<IntroduceGdprOnEventFieldRequest> = {
    encode(
        message: IntroduceGdprOnEventFieldRequest,
        writer: BinaryWriter = new BinaryWriter()
    ): BinaryWriter {
        if (message.tenantId !== "") {
            writer.uint32(10).string(message.tenantId);
        }
        if (message.topic !== "") {
            writer.uint32(18).string(message.topic);
        }
        if (message.eventId !== "") {
            writer.uint32(26).string(message.eventId);
        }
        if (message.fieldName !== "") {
            writer.uint32(34).string(message.fieldName);
        }
        if (message.defaultValue !== "") {
            writer.uint32(42).string(message.defaultValue);
        }
        return writer;
    },

    decode(input: BinaryReader | Uint8Array, length?: number): IntroduceGdprOnEventFieldRequest {
        const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseIntroduceGdprOnEventFieldRequest();
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

                    message.eventId = reader.string();
                    continue;
                }
                case 4: {
                    if (tag !== 34) {
                        break;
                    }

                    message.fieldName = reader.string();
                    continue;
                }
                case 5: {
                    if (tag !== 42) {
                        break;
                    }

                    message.defaultValue = reader.string();
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

    fromJSON(object: any): IntroduceGdprOnEventFieldRequest {
        return {
            tenantId: isSet(object.tenantId) ? globalThis.String(object.tenantId) : "",
            topic: isSet(object.topic) ? globalThis.String(object.topic) : "",
            eventId: isSet(object.eventId) ? globalThis.String(object.eventId) : "",
            fieldName: isSet(object.fieldName) ? globalThis.String(object.fieldName) : "",
            defaultValue: isSet(object.defaultValue) ? globalThis.String(object.defaultValue) : "",
        };
    },

    toJSON(message: IntroduceGdprOnEventFieldRequest): unknown {
        const obj: any = {};
        if (message.tenantId !== "") {
            obj.tenantId = message.tenantId;
        }
        if (message.topic !== "") {
            obj.topic = message.topic;
        }
        if (message.eventId !== "") {
            obj.eventId = message.eventId;
        }
        if (message.fieldName !== "") {
            obj.fieldName = message.fieldName;
        }
        if (message.defaultValue !== "") {
            obj.defaultValue = message.defaultValue;
        }
        return obj;
    },

    create(base?: DeepPartial<IntroduceGdprOnEventFieldRequest>): IntroduceGdprOnEventFieldRequest {
        return IntroduceGdprOnEventFieldRequest.fromPartial(base ?? {});
    },
    fromPartial(
        object: DeepPartial<IntroduceGdprOnEventFieldRequest>
    ): IntroduceGdprOnEventFieldRequest {
        const message = createBaseIntroduceGdprOnEventFieldRequest();
        message.tenantId = object.tenantId ?? "";
        message.topic = object.topic ?? "";
        message.eventId = object.eventId ?? "";
        message.fieldName = object.fieldName ?? "";
        message.defaultValue = object.defaultValue ?? "";
        return message;
    },
};

function createBaseIntroduceGdprOnEventFieldResponse(): IntroduceGdprOnEventFieldResponse {
    return {};
}

export const IntroduceGdprOnEventFieldResponse: MessageFns<IntroduceGdprOnEventFieldResponse> = {
    encode(
        _: IntroduceGdprOnEventFieldResponse,
        writer: BinaryWriter = new BinaryWriter()
    ): BinaryWriter {
        return writer;
    },

    decode(input: BinaryReader | Uint8Array, length?: number): IntroduceGdprOnEventFieldResponse {
        const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseIntroduceGdprOnEventFieldResponse();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
            }
            if ((tag & 7) === 4 || tag === 0) {
                break;
            }
            reader.skip(tag & 7);
        }
        return message;
    },

    fromJSON(_: any): IntroduceGdprOnEventFieldResponse {
        return {};
    },

    toJSON(_: IntroduceGdprOnEventFieldResponse): unknown {
        const obj: any = {};
        return obj;
    },

    create(
        base?: DeepPartial<IntroduceGdprOnEventFieldResponse>
    ): IntroduceGdprOnEventFieldResponse {
        return IntroduceGdprOnEventFieldResponse.fromPartial(base ?? {});
    },
    fromPartial(
        _: DeepPartial<IntroduceGdprOnEventFieldResponse>
    ): IntroduceGdprOnEventFieldResponse {
        const message = createBaseIntroduceGdprOnEventFieldResponse();
        return message;
    },
};

function createBaseInvalidateGdprRequest(): InvalidateGdprRequest {
    return { tenantId: "", topic: "", gdprId: "" };
}

export const InvalidateGdprRequest: MessageFns<InvalidateGdprRequest> = {
    encode(
        message: InvalidateGdprRequest,
        writer: BinaryWriter = new BinaryWriter()
    ): BinaryWriter {
        if (message.tenantId !== "") {
            writer.uint32(10).string(message.tenantId);
        }
        if (message.topic !== "") {
            writer.uint32(18).string(message.topic);
        }
        if (message.gdprId !== "") {
            writer.uint32(26).string(message.gdprId);
        }
        return writer;
    },

    decode(input: BinaryReader | Uint8Array, length?: number): InvalidateGdprRequest {
        const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseInvalidateGdprRequest();
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

                    message.gdprId = reader.string();
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

    fromJSON(object: any): InvalidateGdprRequest {
        return {
            tenantId: isSet(object.tenantId) ? globalThis.String(object.tenantId) : "",
            topic: isSet(object.topic) ? globalThis.String(object.topic) : "",
            gdprId: isSet(object.gdprId) ? globalThis.String(object.gdprId) : "",
        };
    },

    toJSON(message: InvalidateGdprRequest): unknown {
        const obj: any = {};
        if (message.tenantId !== "") {
            obj.tenantId = message.tenantId;
        }
        if (message.topic !== "") {
            obj.topic = message.topic;
        }
        if (message.gdprId !== "") {
            obj.gdprId = message.gdprId;
        }
        return obj;
    },

    create(base?: DeepPartial<InvalidateGdprRequest>): InvalidateGdprRequest {
        return InvalidateGdprRequest.fromPartial(base ?? {});
    },
    fromPartial(object: DeepPartial<InvalidateGdprRequest>): InvalidateGdprRequest {
        const message = createBaseInvalidateGdprRequest();
        message.tenantId = object.tenantId ?? "";
        message.topic = object.topic ?? "";
        message.gdprId = object.gdprId ?? "";
        return message;
    },
};

function createBaseInvalidateGdprResponse(): InvalidateGdprResponse {
    return {};
}

export const InvalidateGdprResponse: MessageFns<InvalidateGdprResponse> = {
    encode(_: InvalidateGdprResponse, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
        return writer;
    },

    decode(input: BinaryReader | Uint8Array, length?: number): InvalidateGdprResponse {
        const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseInvalidateGdprResponse();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
            }
            if ((tag & 7) === 4 || tag === 0) {
                break;
            }
            reader.skip(tag & 7);
        }
        return message;
    },

    fromJSON(_: any): InvalidateGdprResponse {
        return {};
    },

    toJSON(_: InvalidateGdprResponse): unknown {
        const obj: any = {};
        return obj;
    },

    create(base?: DeepPartial<InvalidateGdprResponse>): InvalidateGdprResponse {
        return InvalidateGdprResponse.fromPartial(base ?? {});
    },
    fromPartial(_: DeepPartial<InvalidateGdprResponse>): InvalidateGdprResponse {
        const message = createBaseInvalidateGdprResponse();
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
