// Code generated by protoc-gen-ts_proto. DO NOT EDIT.
// versions:
//   protoc-gen-ts_proto  v2.7.0
//   protoc               v5.29.3
// source: auth/management/roles_get.proto

/* eslint-disable */
import { BinaryReader, BinaryWriter } from "@bufbuild/protobuf/wire";
import { RoleScope } from "./shared";

export interface GetRolesRequest {
    tenantId: string;
    limit: string;
    page: string;
}

export interface GetRolesResponse {
    roles: Role[];
    limit: string;
    page: string;
    total: string;
}

export interface Role {
    id: string;
    allowedScopes: RoleScope[];
}

function createBaseGetRolesRequest(): GetRolesRequest {
    return { tenantId: "", limit: "0", page: "0" };
}

export const GetRolesRequest: MessageFns<GetRolesRequest> = {
    encode(message: GetRolesRequest, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
        if (message.tenantId !== "") {
            writer.uint32(10).string(message.tenantId);
        }
        if (message.limit !== "0") {
            writer.uint32(16).int64(message.limit);
        }
        if (message.page !== "0") {
            writer.uint32(24).int64(message.page);
        }
        return writer;
    },

    decode(input: BinaryReader | Uint8Array, length?: number): GetRolesRequest {
        const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseGetRolesRequest();
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
                    if (tag !== 16) {
                        break;
                    }

                    message.limit = reader.int64().toString();
                    continue;
                }
                case 3: {
                    if (tag !== 24) {
                        break;
                    }

                    message.page = reader.int64().toString();
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

    fromJSON(object: any): GetRolesRequest {
        return {
            tenantId: isSet(object.tenantId) ? globalThis.String(object.tenantId) : "",
            limit: isSet(object.limit) ? globalThis.String(object.limit) : "0",
            page: isSet(object.page) ? globalThis.String(object.page) : "0",
        };
    },

    toJSON(message: GetRolesRequest): unknown {
        const obj: any = {};
        if (message.tenantId !== "") {
            obj.tenantId = message.tenantId;
        }
        if (message.limit !== "0") {
            obj.limit = message.limit;
        }
        if (message.page !== "0") {
            obj.page = message.page;
        }
        return obj;
    },

    create(base?: DeepPartial<GetRolesRequest>): GetRolesRequest {
        return GetRolesRequest.fromPartial(base ?? {});
    },
    fromPartial(object: DeepPartial<GetRolesRequest>): GetRolesRequest {
        const message = createBaseGetRolesRequest();
        message.tenantId = object.tenantId ?? "";
        message.limit = object.limit ?? "0";
        message.page = object.page ?? "0";
        return message;
    },
};

function createBaseGetRolesResponse(): GetRolesResponse {
    return { roles: [], limit: "0", page: "0", total: "0" };
}

export const GetRolesResponse: MessageFns<GetRolesResponse> = {
    encode(message: GetRolesResponse, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
        for (const v of message.roles) {
            Role.encode(v!, writer.uint32(10).fork()).join();
        }
        if (message.limit !== "0") {
            writer.uint32(16).int64(message.limit);
        }
        if (message.page !== "0") {
            writer.uint32(24).int64(message.page);
        }
        if (message.total !== "0") {
            writer.uint32(32).int64(message.total);
        }
        return writer;
    },

    decode(input: BinaryReader | Uint8Array, length?: number): GetRolesResponse {
        const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseGetRolesResponse();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1: {
                    if (tag !== 10) {
                        break;
                    }

                    message.roles.push(Role.decode(reader, reader.uint32()));
                    continue;
                }
                case 2: {
                    if (tag !== 16) {
                        break;
                    }

                    message.limit = reader.int64().toString();
                    continue;
                }
                case 3: {
                    if (tag !== 24) {
                        break;
                    }

                    message.page = reader.int64().toString();
                    continue;
                }
                case 4: {
                    if (tag !== 32) {
                        break;
                    }

                    message.total = reader.int64().toString();
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

    fromJSON(object: any): GetRolesResponse {
        return {
            roles: globalThis.Array.isArray(object?.roles)
                ? object.roles.map((e: any) => Role.fromJSON(e))
                : [],
            limit: isSet(object.limit) ? globalThis.String(object.limit) : "0",
            page: isSet(object.page) ? globalThis.String(object.page) : "0",
            total: isSet(object.total) ? globalThis.String(object.total) : "0",
        };
    },

    toJSON(message: GetRolesResponse): unknown {
        const obj: any = {};
        if (message.roles?.length) {
            obj.roles = message.roles.map(e => Role.toJSON(e));
        }
        if (message.limit !== "0") {
            obj.limit = message.limit;
        }
        if (message.page !== "0") {
            obj.page = message.page;
        }
        if (message.total !== "0") {
            obj.total = message.total;
        }
        return obj;
    },

    create(base?: DeepPartial<GetRolesResponse>): GetRolesResponse {
        return GetRolesResponse.fromPartial(base ?? {});
    },
    fromPartial(object: DeepPartial<GetRolesResponse>): GetRolesResponse {
        const message = createBaseGetRolesResponse();
        message.roles = object.roles?.map(e => Role.fromPartial(e)) || [];
        message.limit = object.limit ?? "0";
        message.page = object.page ?? "0";
        message.total = object.total ?? "0";
        return message;
    },
};

function createBaseRole(): Role {
    return { id: "", allowedScopes: [] };
}

export const Role: MessageFns<Role> = {
    encode(message: Role, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
        if (message.id !== "") {
            writer.uint32(10).string(message.id);
        }
        for (const v of message.allowedScopes) {
            RoleScope.encode(v!, writer.uint32(18).fork()).join();
        }
        return writer;
    },

    decode(input: BinaryReader | Uint8Array, length?: number): Role {
        const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseRole();
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

                    message.allowedScopes.push(RoleScope.decode(reader, reader.uint32()));
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

    fromJSON(object: any): Role {
        return {
            id: isSet(object.id) ? globalThis.String(object.id) : "",
            allowedScopes: globalThis.Array.isArray(object?.allowedScopes)
                ? object.allowedScopes.map((e: any) => RoleScope.fromJSON(e))
                : [],
        };
    },

    toJSON(message: Role): unknown {
        const obj: any = {};
        if (message.id !== "") {
            obj.id = message.id;
        }
        if (message.allowedScopes?.length) {
            obj.allowedScopes = message.allowedScopes.map(e => RoleScope.toJSON(e));
        }
        return obj;
    },

    create(base?: DeepPartial<Role>): Role {
        return Role.fromPartial(base ?? {});
    },
    fromPartial(object: DeepPartial<Role>): Role {
        const message = createBaseRole();
        message.id = object.id ?? "";
        message.allowedScopes = object.allowedScopes?.map(e => RoleScope.fromPartial(e)) || [];
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
