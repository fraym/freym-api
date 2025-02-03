// Code generated by protoc-gen-ts_proto. DO NOT EDIT.
// versions:
//   protoc-gen-ts_proto  v2.6.1
//   protoc               v5.27.3
// source: auth/management/role_upsert.proto

/* eslint-disable */
import { BinaryReader, BinaryWriter } from "@bufbuild/protobuf/wire";
import { EventMetadata, RoleScope } from "./shared";

export interface UpsertRoleRequest {
    tenantId: string;
    id: string;
    allowedScopes: RoleScope[];
    eventMetadata: EventMetadata | undefined;
}

export interface UpsertRoleResponse {
    id: string;
}

function createBaseUpsertRoleRequest(): UpsertRoleRequest {
    return { tenantId: "", id: "", allowedScopes: [], eventMetadata: undefined };
}

export const UpsertRoleRequest: MessageFns<UpsertRoleRequest> = {
    encode(message: UpsertRoleRequest, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
        if (message.tenantId !== "") {
            writer.uint32(10).string(message.tenantId);
        }
        if (message.id !== "") {
            writer.uint32(18).string(message.id);
        }
        for (const v of message.allowedScopes) {
            RoleScope.encode(v!, writer.uint32(26).fork()).join();
        }
        if (message.eventMetadata !== undefined) {
            EventMetadata.encode(message.eventMetadata, writer.uint32(34).fork()).join();
        }
        return writer;
    },

    decode(input: BinaryReader | Uint8Array, length?: number): UpsertRoleRequest {
        const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseUpsertRoleRequest();
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

                    message.allowedScopes.push(RoleScope.decode(reader, reader.uint32()));
                    continue;
                }
                case 4: {
                    if (tag !== 34) {
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

    fromJSON(object: any): UpsertRoleRequest {
        return {
            tenantId: isSet(object.tenantId) ? globalThis.String(object.tenantId) : "",
            id: isSet(object.id) ? globalThis.String(object.id) : "",
            allowedScopes: globalThis.Array.isArray(object?.allowedScopes)
                ? object.allowedScopes.map((e: any) => RoleScope.fromJSON(e))
                : [],
            eventMetadata: isSet(object.eventMetadata)
                ? EventMetadata.fromJSON(object.eventMetadata)
                : undefined,
        };
    },

    toJSON(message: UpsertRoleRequest): unknown {
        const obj: any = {};
        if (message.tenantId !== "") {
            obj.tenantId = message.tenantId;
        }
        if (message.id !== "") {
            obj.id = message.id;
        }
        if (message.allowedScopes?.length) {
            obj.allowedScopes = message.allowedScopes.map(e => RoleScope.toJSON(e));
        }
        if (message.eventMetadata !== undefined) {
            obj.eventMetadata = EventMetadata.toJSON(message.eventMetadata);
        }
        return obj;
    },

    create(base?: DeepPartial<UpsertRoleRequest>): UpsertRoleRequest {
        return UpsertRoleRequest.fromPartial(base ?? {});
    },
    fromPartial(object: DeepPartial<UpsertRoleRequest>): UpsertRoleRequest {
        const message = createBaseUpsertRoleRequest();
        message.tenantId = object.tenantId ?? "";
        message.id = object.id ?? "";
        message.allowedScopes = object.allowedScopes?.map(e => RoleScope.fromPartial(e)) || [];
        message.eventMetadata =
            object.eventMetadata !== undefined && object.eventMetadata !== null
                ? EventMetadata.fromPartial(object.eventMetadata)
                : undefined;
        return message;
    },
};

function createBaseUpsertRoleResponse(): UpsertRoleResponse {
    return { id: "" };
}

export const UpsertRoleResponse: MessageFns<UpsertRoleResponse> = {
    encode(message: UpsertRoleResponse, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
        if (message.id !== "") {
            writer.uint32(10).string(message.id);
        }
        return writer;
    },

    decode(input: BinaryReader | Uint8Array, length?: number): UpsertRoleResponse {
        const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseUpsertRoleResponse();
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
            }
            if ((tag & 7) === 4 || tag === 0) {
                break;
            }
            reader.skip(tag & 7);
        }
        return message;
    },

    fromJSON(object: any): UpsertRoleResponse {
        return { id: isSet(object.id) ? globalThis.String(object.id) : "" };
    },

    toJSON(message: UpsertRoleResponse): unknown {
        const obj: any = {};
        if (message.id !== "") {
            obj.id = message.id;
        }
        return obj;
    },

    create(base?: DeepPartial<UpsertRoleResponse>): UpsertRoleResponse {
        return UpsertRoleResponse.fromPartial(base ?? {});
    },
    fromPartial(object: DeepPartial<UpsertRoleResponse>): UpsertRoleResponse {
        const message = createBaseUpsertRoleResponse();
        message.id = object.id ?? "";
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
