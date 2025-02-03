// Code generated by protoc-gen-ts_proto. DO NOT EDIT.
// versions:
//   protoc-gen-ts_proto  v2.6.1
//   protoc               v5.27.3
// source: auth/management/user_delete.proto

/* eslint-disable */
import { BinaryReader, BinaryWriter } from "@bufbuild/protobuf/wire";
import { EventMetadata } from "./shared";

export interface DeleteUserRequest {
    tenantId: string;
    id: string;
    eventMetadata: EventMetadata | undefined;
}

export interface DeleteUserResponse {}

function createBaseDeleteUserRequest(): DeleteUserRequest {
    return { tenantId: "", id: "", eventMetadata: undefined };
}

export const DeleteUserRequest: MessageFns<DeleteUserRequest> = {
    encode(message: DeleteUserRequest, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
        if (message.tenantId !== "") {
            writer.uint32(10).string(message.tenantId);
        }
        if (message.id !== "") {
            writer.uint32(18).string(message.id);
        }
        if (message.eventMetadata !== undefined) {
            EventMetadata.encode(message.eventMetadata, writer.uint32(26).fork()).join();
        }
        return writer;
    },

    decode(input: BinaryReader | Uint8Array, length?: number): DeleteUserRequest {
        const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseDeleteUserRequest();
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

                    message.id = reader.string();
                    continue;
                }
                case 3: {
                    if (tag !== 26) {
                        break;
                    }

                    message.eventMetadata = EventMetadata.decode(reader, reader.uint32());
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

    fromJSON(object: any): DeleteUserRequest {
        return {
            tenantId: isSet(object.tenantId) ? globalThis.String(object.tenantId) : "",
            id: isSet(object.id) ? globalThis.String(object.id) : "",
            eventMetadata: isSet(object.eventMetadata)
                ? EventMetadata.fromJSON(object.eventMetadata)
                : undefined,
        };
    },

    toJSON(message: DeleteUserRequest): unknown {
        const obj: any = {};
        if (message.tenantId !== "") {
            obj.tenantId = message.tenantId;
        }
        if (message.id !== "") {
            obj.id = message.id;
        }
        if (message.eventMetadata !== undefined) {
            obj.eventMetadata = EventMetadata.toJSON(message.eventMetadata);
        }
        return obj;
    },

    create(base?: DeepPartial<DeleteUserRequest>): DeleteUserRequest {
        return DeleteUserRequest.fromPartial(base ?? {});
    },
    fromPartial(object: DeepPartial<DeleteUserRequest>): DeleteUserRequest {
        const message = createBaseDeleteUserRequest();
        message.tenantId = object.tenantId ?? "";
        message.id = object.id ?? "";
        message.eventMetadata =
            object.eventMetadata !== undefined && object.eventMetadata !== null
                ? EventMetadata.fromPartial(object.eventMetadata)
                : undefined;
        return message;
    },
};

function createBaseDeleteUserResponse(): DeleteUserResponse {
    return {};
}

export const DeleteUserResponse: MessageFns<DeleteUserResponse> = {
    encode(_: DeleteUserResponse, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
        return writer;
    },

    decode(input: BinaryReader | Uint8Array, length?: number): DeleteUserResponse {
        const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseDeleteUserResponse();
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

    fromJSON(_: any): DeleteUserResponse {
        return {};
    },

    toJSON(_: DeleteUserResponse): unknown {
        const obj: any = {};
        return obj;
    },

    create(base?: DeepPartial<DeleteUserResponse>): DeleteUserResponse {
        return DeleteUserResponse.fromPartial(base ?? {});
    },
    fromPartial(_: DeepPartial<DeleteUserResponse>): DeleteUserResponse {
        const message = createBaseDeleteUserResponse();
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
