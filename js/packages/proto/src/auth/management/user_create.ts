// Code generated by protoc-gen-ts_proto. DO NOT EDIT.
// versions:
//   protoc-gen-ts_proto  v2.6.1
//   protoc               v5.27.3
// source: auth/management/user_create.proto

/* eslint-disable */
import { BinaryReader, BinaryWriter } from "@bufbuild/protobuf/wire";
import { EventMetadata } from "./shared";

export interface CreateUserRequest {
    tenantId: string;
    login: string;
    email: string;
    displayName: string;
    password: string;
    assignedRoleIds: string[];
    active: boolean;
    blockedUntil: string;
    eventMetadata: EventMetadata | undefined;
}

export interface CreateUserResponse {
    id: string;
    setInitialPasswordToken: string;
}

function createBaseCreateUserRequest(): CreateUserRequest {
    return {
        tenantId: "",
        login: "",
        email: "",
        displayName: "",
        password: "",
        assignedRoleIds: [],
        active: false,
        blockedUntil: "0",
        eventMetadata: undefined,
    };
}

export const CreateUserRequest: MessageFns<CreateUserRequest> = {
    encode(message: CreateUserRequest, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
        if (message.tenantId !== "") {
            writer.uint32(10).string(message.tenantId);
        }
        if (message.login !== "") {
            writer.uint32(18).string(message.login);
        }
        if (message.email !== "") {
            writer.uint32(26).string(message.email);
        }
        if (message.displayName !== "") {
            writer.uint32(34).string(message.displayName);
        }
        if (message.password !== "") {
            writer.uint32(42).string(message.password);
        }
        for (const v of message.assignedRoleIds) {
            writer.uint32(50).string(v!);
        }
        if (message.active !== false) {
            writer.uint32(56).bool(message.active);
        }
        if (message.blockedUntil !== "0") {
            writer.uint32(64).int64(message.blockedUntil);
        }
        if (message.eventMetadata !== undefined) {
            EventMetadata.encode(message.eventMetadata, writer.uint32(74).fork()).join();
        }
        return writer;
    },

    decode(input: BinaryReader | Uint8Array, length?: number): CreateUserRequest {
        const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseCreateUserRequest();
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

                    message.login = reader.string();
                    continue;
                }
                case 3: {
                    if (tag !== 26) {
                        break;
                    }

                    message.email = reader.string();
                    continue;
                }
                case 4: {
                    if (tag !== 34) {
                        break;
                    }

                    message.displayName = reader.string();
                    continue;
                }
                case 5: {
                    if (tag !== 42) {
                        break;
                    }

                    message.password = reader.string();
                    continue;
                }
                case 6: {
                    if (tag !== 50) {
                        break;
                    }

                    message.assignedRoleIds.push(reader.string());
                    continue;
                }
                case 7: {
                    if (tag !== 56) {
                        break;
                    }

                    message.active = reader.bool();
                    continue;
                }
                case 8: {
                    if (tag !== 64) {
                        break;
                    }

                    message.blockedUntil = reader.int64().toString();
                    continue;
                }
                case 9: {
                    if (tag !== 74) {
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

    fromJSON(object: any): CreateUserRequest {
        return {
            tenantId: isSet(object.tenantId) ? globalThis.String(object.tenantId) : "",
            login: isSet(object.login) ? globalThis.String(object.login) : "",
            email: isSet(object.email) ? globalThis.String(object.email) : "",
            displayName: isSet(object.displayName) ? globalThis.String(object.displayName) : "",
            password: isSet(object.password) ? globalThis.String(object.password) : "",
            assignedRoleIds: globalThis.Array.isArray(object?.assignedRoleIds)
                ? object.assignedRoleIds.map((e: any) => globalThis.String(e))
                : [],
            active: isSet(object.active) ? globalThis.Boolean(object.active) : false,
            blockedUntil: isSet(object.blockedUntil) ? globalThis.String(object.blockedUntil) : "0",
            eventMetadata: isSet(object.eventMetadata)
                ? EventMetadata.fromJSON(object.eventMetadata)
                : undefined,
        };
    },

    toJSON(message: CreateUserRequest): unknown {
        const obj: any = {};
        if (message.tenantId !== "") {
            obj.tenantId = message.tenantId;
        }
        if (message.login !== "") {
            obj.login = message.login;
        }
        if (message.email !== "") {
            obj.email = message.email;
        }
        if (message.displayName !== "") {
            obj.displayName = message.displayName;
        }
        if (message.password !== "") {
            obj.password = message.password;
        }
        if (message.assignedRoleIds?.length) {
            obj.assignedRoleIds = message.assignedRoleIds;
        }
        if (message.active !== false) {
            obj.active = message.active;
        }
        if (message.blockedUntil !== "0") {
            obj.blockedUntil = message.blockedUntil;
        }
        if (message.eventMetadata !== undefined) {
            obj.eventMetadata = EventMetadata.toJSON(message.eventMetadata);
        }
        return obj;
    },

    create(base?: DeepPartial<CreateUserRequest>): CreateUserRequest {
        return CreateUserRequest.fromPartial(base ?? {});
    },
    fromPartial(object: DeepPartial<CreateUserRequest>): CreateUserRequest {
        const message = createBaseCreateUserRequest();
        message.tenantId = object.tenantId ?? "";
        message.login = object.login ?? "";
        message.email = object.email ?? "";
        message.displayName = object.displayName ?? "";
        message.password = object.password ?? "";
        message.assignedRoleIds = object.assignedRoleIds?.map(e => e) || [];
        message.active = object.active ?? false;
        message.blockedUntil = object.blockedUntil ?? "0";
        message.eventMetadata =
            object.eventMetadata !== undefined && object.eventMetadata !== null
                ? EventMetadata.fromPartial(object.eventMetadata)
                : undefined;
        return message;
    },
};

function createBaseCreateUserResponse(): CreateUserResponse {
    return { id: "", setInitialPasswordToken: "" };
}

export const CreateUserResponse: MessageFns<CreateUserResponse> = {
    encode(message: CreateUserResponse, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
        if (message.id !== "") {
            writer.uint32(10).string(message.id);
        }
        if (message.setInitialPasswordToken !== "") {
            writer.uint32(18).string(message.setInitialPasswordToken);
        }
        return writer;
    },

    decode(input: BinaryReader | Uint8Array, length?: number): CreateUserResponse {
        const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseCreateUserResponse();
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

                    message.setInitialPasswordToken = reader.string();
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

    fromJSON(object: any): CreateUserResponse {
        return {
            id: isSet(object.id) ? globalThis.String(object.id) : "",
            setInitialPasswordToken: isSet(object.setInitialPasswordToken)
                ? globalThis.String(object.setInitialPasswordToken)
                : "",
        };
    },

    toJSON(message: CreateUserResponse): unknown {
        const obj: any = {};
        if (message.id !== "") {
            obj.id = message.id;
        }
        if (message.setInitialPasswordToken !== "") {
            obj.setInitialPasswordToken = message.setInitialPasswordToken;
        }
        return obj;
    },

    create(base?: DeepPartial<CreateUserResponse>): CreateUserResponse {
        return CreateUserResponse.fromPartial(base ?? {});
    },
    fromPartial(object: DeepPartial<CreateUserResponse>): CreateUserResponse {
        const message = createBaseCreateUserResponse();
        message.id = object.id ?? "";
        message.setInitialPasswordToken = object.setInitialPasswordToken ?? "";
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
