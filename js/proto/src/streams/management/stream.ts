// Code generated by protoc-gen-ts_proto. DO NOT EDIT.
// versions:
//   protoc-gen-ts_proto  v2.7.0
//   protoc               v5.29.3
// source: streams/management/stream.proto

/* eslint-disable */
import { BinaryReader, BinaryWriter } from "@bufbuild/protobuf/wire";

export interface IsStreamEmptyRequest {
    tenantId: string;
    topic: string;
    stream: string;
}

export interface IsStreamEmptyResponse {
    isEmpty: boolean;
}

function createBaseIsStreamEmptyRequest(): IsStreamEmptyRequest {
    return { tenantId: "", topic: "", stream: "" };
}

export const IsStreamEmptyRequest: MessageFns<IsStreamEmptyRequest> = {
    encode(message: IsStreamEmptyRequest, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
        if (message.tenantId !== "") {
            writer.uint32(10).string(message.tenantId);
        }
        if (message.topic !== "") {
            writer.uint32(18).string(message.topic);
        }
        if (message.stream !== "") {
            writer.uint32(26).string(message.stream);
        }
        return writer;
    },

    decode(input: BinaryReader | Uint8Array, length?: number): IsStreamEmptyRequest {
        const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseIsStreamEmptyRequest();
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
            }
            if ((tag & 7) === 4 || tag === 0) {
                break;
            }
            reader.skip(tag & 7);
        }
        return message;
    },

    fromJSON(object: any): IsStreamEmptyRequest {
        return {
            tenantId: isSet(object.tenantId) ? globalThis.String(object.tenantId) : "",
            topic: isSet(object.topic) ? globalThis.String(object.topic) : "",
            stream: isSet(object.stream) ? globalThis.String(object.stream) : "",
        };
    },

    toJSON(message: IsStreamEmptyRequest): unknown {
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
        return obj;
    },

    create(base?: DeepPartial<IsStreamEmptyRequest>): IsStreamEmptyRequest {
        return IsStreamEmptyRequest.fromPartial(base ?? {});
    },
    fromPartial(object: DeepPartial<IsStreamEmptyRequest>): IsStreamEmptyRequest {
        const message = createBaseIsStreamEmptyRequest();
        message.tenantId = object.tenantId ?? "";
        message.topic = object.topic ?? "";
        message.stream = object.stream ?? "";
        return message;
    },
};

function createBaseIsStreamEmptyResponse(): IsStreamEmptyResponse {
    return { isEmpty: false };
}

export const IsStreamEmptyResponse: MessageFns<IsStreamEmptyResponse> = {
    encode(
        message: IsStreamEmptyResponse,
        writer: BinaryWriter = new BinaryWriter()
    ): BinaryWriter {
        if (message.isEmpty !== false) {
            writer.uint32(8).bool(message.isEmpty);
        }
        return writer;
    },

    decode(input: BinaryReader | Uint8Array, length?: number): IsStreamEmptyResponse {
        const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseIsStreamEmptyResponse();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1: {
                    if (tag !== 8) {
                        break;
                    }

                    message.isEmpty = reader.bool();
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

    fromJSON(object: any): IsStreamEmptyResponse {
        return { isEmpty: isSet(object.isEmpty) ? globalThis.Boolean(object.isEmpty) : false };
    },

    toJSON(message: IsStreamEmptyResponse): unknown {
        const obj: any = {};
        if (message.isEmpty !== false) {
            obj.isEmpty = message.isEmpty;
        }
        return obj;
    },

    create(base?: DeepPartial<IsStreamEmptyResponse>): IsStreamEmptyResponse {
        return IsStreamEmptyResponse.fromPartial(base ?? {});
    },
    fromPartial(object: DeepPartial<IsStreamEmptyResponse>): IsStreamEmptyResponse {
        const message = createBaseIsStreamEmptyResponse();
        message.isEmpty = object.isEmpty ?? false;
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
