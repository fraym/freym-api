// Code generated by protoc-gen-ts_proto. DO NOT EDIT.
// versions:
//   protoc-gen-ts_proto  v2.7.0
//   protoc               v5.29.3
// source: auth/management/role_delete.proto

/* eslint-disable */
import { BinaryReader, BinaryWriter } from "@bufbuild/protobuf/wire";
import { EventMetadata } from "./shared";

export interface DeleteRoleRequest {
    id: string;
    tenantId: string;
    eventMetadata: EventMetadata | undefined;
}

export interface DeleteRoleResponse {}

function createBaseDeleteRoleRequest(): DeleteRoleRequest {
    return { id: "", tenantId: "", eventMetadata: undefined };
}

export const DeleteRoleRequest: MessageFns<DeleteRoleRequest> = {
    encode(message: DeleteRoleRequest, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
        if (message.id !== "") {
            writer.uint32(10).string(message.id);
        }
        if (message.tenantId !== "") {
            writer.uint32(18).string(message.tenantId);
        }
        if (message.eventMetadata !== undefined) {
            EventMetadata.encode(message.eventMetadata, writer.uint32(26).fork()).join();
        }
        return writer;
    },

    decode(input: BinaryReader | Uint8Array, length?: number): DeleteRoleRequest {
        const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseDeleteRoleRequest();
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

                    message.tenantId = reader.string();
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

    fromJSON(object: any): DeleteRoleRequest {
        return {
            id: isSet(object.id) ? globalThis.String(object.id) : "",
            tenantId: isSet(object.tenantId) ? globalThis.String(object.tenantId) : "",
            eventMetadata: isSet(object.eventMetadata)
                ? EventMetadata.fromJSON(object.eventMetadata)
                : undefined,
        };
    },

    toJSON(message: DeleteRoleRequest): unknown {
        const obj: any = {};
        if (message.id !== "") {
            obj.id = message.id;
        }
        if (message.tenantId !== "") {
            obj.tenantId = message.tenantId;
        }
        if (message.eventMetadata !== undefined) {
            obj.eventMetadata = EventMetadata.toJSON(message.eventMetadata);
        }
        return obj;
    },

    create(base?: DeepPartial<DeleteRoleRequest>): DeleteRoleRequest {
        return DeleteRoleRequest.fromPartial(base ?? {});
    },
    fromPartial(object: DeepPartial<DeleteRoleRequest>): DeleteRoleRequest {
        const message = createBaseDeleteRoleRequest();
        message.id = object.id ?? "";
        message.tenantId = object.tenantId ?? "";
        message.eventMetadata =
            object.eventMetadata !== undefined && object.eventMetadata !== null
                ? EventMetadata.fromPartial(object.eventMetadata)
                : undefined;
        return message;
    },
};

function createBaseDeleteRoleResponse(): DeleteRoleResponse {
    return {};
}

export const DeleteRoleResponse: MessageFns<DeleteRoleResponse> = {
    encode(_: DeleteRoleResponse, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
        return writer;
    },

    decode(input: BinaryReader | Uint8Array, length?: number): DeleteRoleResponse {
        const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseDeleteRoleResponse();
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

    fromJSON(_: any): DeleteRoleResponse {
        return {};
    },

    toJSON(_: DeleteRoleResponse): unknown {
        const obj: any = {};
        return obj;
    },

    create(base?: DeepPartial<DeleteRoleResponse>): DeleteRoleResponse {
        return DeleteRoleResponse.fromPartial(base ?? {});
    },
    fromPartial(_: DeepPartial<DeleteRoleResponse>): DeleteRoleResponse {
        const message = createBaseDeleteRoleResponse();
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
