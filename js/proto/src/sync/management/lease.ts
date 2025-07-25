// Code generated by protoc-gen-ts_proto. DO NOT EDIT.
// versions:
//   protoc-gen-ts_proto  v2.7.5
//   protoc               v5.29.3
// source: sync/management/lease.proto
/* eslint-disable */
import { BinaryReader, BinaryWriter } from "@bufbuild/protobuf/wire";

export interface Lock {
    leaseId: string;
    tenantId: string;
    resource: string[];
}

export interface CreateLeaseRequest {
    ip: string;
    app: string;
    ttl: number;
    leaseId: string;
    alreadyRegisteredLocks: Lock[];
    alreadyRegisteredRlocks: Lock[];
}

export interface CreateLeaseResponse {
    leaseId: string;
}

export interface KeepLeaseRequest {
    leaseId: string;
    ttl: number;
}

export interface KeepLeaseResponse {}

export interface DropLeaseRequest {
    leaseId: string;
}

export interface DropLeaseResponse {}

function createBaseLock(): Lock {
    return { leaseId: "", tenantId: "", resource: [] };
}

export const Lock: MessageFns<Lock> = {
    encode(message: Lock, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
        if (message.leaseId !== "") {
            writer.uint32(10).string(message.leaseId);
        }
        if (message.tenantId !== "") {
            writer.uint32(18).string(message.tenantId);
        }
        for (const v of message.resource) {
            writer.uint32(26).string(v!);
        }
        return writer;
    },

    decode(input: BinaryReader | Uint8Array, length?: number): Lock {
        const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
        const end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseLock();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1: {
                    if (tag !== 10) {
                        break;
                    }

                    message.leaseId = reader.string();
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

                    message.resource.push(reader.string());
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

    fromJSON(object: any): Lock {
        return {
            leaseId: isSet(object.leaseId) ? globalThis.String(object.leaseId) : "",
            tenantId: isSet(object.tenantId) ? globalThis.String(object.tenantId) : "",
            resource: globalThis.Array.isArray(object?.resource)
                ? object.resource.map((e: any) => globalThis.String(e))
                : [],
        };
    },

    toJSON(message: Lock): unknown {
        const obj: any = {};
        if (message.leaseId !== "") {
            obj.leaseId = message.leaseId;
        }
        if (message.tenantId !== "") {
            obj.tenantId = message.tenantId;
        }
        if (message.resource?.length) {
            obj.resource = message.resource;
        }
        return obj;
    },

    create(base?: DeepPartial<Lock>): Lock {
        return Lock.fromPartial(base ?? {});
    },
    fromPartial(object: DeepPartial<Lock>): Lock {
        const message = createBaseLock();
        message.leaseId = object.leaseId ?? "";
        message.tenantId = object.tenantId ?? "";
        message.resource = object.resource?.map(e => e) || [];
        return message;
    },
};

function createBaseCreateLeaseRequest(): CreateLeaseRequest {
    return {
        ip: "",
        app: "",
        ttl: 0,
        leaseId: "",
        alreadyRegisteredLocks: [],
        alreadyRegisteredRlocks: [],
    };
}

export const CreateLeaseRequest: MessageFns<CreateLeaseRequest> = {
    encode(message: CreateLeaseRequest, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
        if (message.ip !== "") {
            writer.uint32(10).string(message.ip);
        }
        if (message.app !== "") {
            writer.uint32(18).string(message.app);
        }
        if (message.ttl !== 0) {
            writer.uint32(24).int32(message.ttl);
        }
        if (message.leaseId !== "") {
            writer.uint32(34).string(message.leaseId);
        }
        for (const v of message.alreadyRegisteredLocks) {
            Lock.encode(v!, writer.uint32(42).fork()).join();
        }
        for (const v of message.alreadyRegisteredRlocks) {
            Lock.encode(v!, writer.uint32(50).fork()).join();
        }
        return writer;
    },

    decode(input: BinaryReader | Uint8Array, length?: number): CreateLeaseRequest {
        const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
        const end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseCreateLeaseRequest();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1: {
                    if (tag !== 10) {
                        break;
                    }

                    message.ip = reader.string();
                    continue;
                }
                case 2: {
                    if (tag !== 18) {
                        break;
                    }

                    message.app = reader.string();
                    continue;
                }
                case 3: {
                    if (tag !== 24) {
                        break;
                    }

                    message.ttl = reader.int32();
                    continue;
                }
                case 4: {
                    if (tag !== 34) {
                        break;
                    }

                    message.leaseId = reader.string();
                    continue;
                }
                case 5: {
                    if (tag !== 42) {
                        break;
                    }

                    message.alreadyRegisteredLocks.push(Lock.decode(reader, reader.uint32()));
                    continue;
                }
                case 6: {
                    if (tag !== 50) {
                        break;
                    }

                    message.alreadyRegisteredRlocks.push(Lock.decode(reader, reader.uint32()));
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

    fromJSON(object: any): CreateLeaseRequest {
        return {
            ip: isSet(object.ip) ? globalThis.String(object.ip) : "",
            app: isSet(object.app) ? globalThis.String(object.app) : "",
            ttl: isSet(object.ttl) ? globalThis.Number(object.ttl) : 0,
            leaseId: isSet(object.leaseId) ? globalThis.String(object.leaseId) : "",
            alreadyRegisteredLocks: globalThis.Array.isArray(object?.alreadyRegisteredLocks)
                ? object.alreadyRegisteredLocks.map((e: any) => Lock.fromJSON(e))
                : [],
            alreadyRegisteredRlocks: globalThis.Array.isArray(object?.alreadyRegisteredRlocks)
                ? object.alreadyRegisteredRlocks.map((e: any) => Lock.fromJSON(e))
                : [],
        };
    },

    toJSON(message: CreateLeaseRequest): unknown {
        const obj: any = {};
        if (message.ip !== "") {
            obj.ip = message.ip;
        }
        if (message.app !== "") {
            obj.app = message.app;
        }
        if (message.ttl !== 0) {
            obj.ttl = Math.round(message.ttl);
        }
        if (message.leaseId !== "") {
            obj.leaseId = message.leaseId;
        }
        if (message.alreadyRegisteredLocks?.length) {
            obj.alreadyRegisteredLocks = message.alreadyRegisteredLocks.map(e => Lock.toJSON(e));
        }
        if (message.alreadyRegisteredRlocks?.length) {
            obj.alreadyRegisteredRlocks = message.alreadyRegisteredRlocks.map(e => Lock.toJSON(e));
        }
        return obj;
    },

    create(base?: DeepPartial<CreateLeaseRequest>): CreateLeaseRequest {
        return CreateLeaseRequest.fromPartial(base ?? {});
    },
    fromPartial(object: DeepPartial<CreateLeaseRequest>): CreateLeaseRequest {
        const message = createBaseCreateLeaseRequest();
        message.ip = object.ip ?? "";
        message.app = object.app ?? "";
        message.ttl = object.ttl ?? 0;
        message.leaseId = object.leaseId ?? "";
        message.alreadyRegisteredLocks =
            object.alreadyRegisteredLocks?.map(e => Lock.fromPartial(e)) || [];
        message.alreadyRegisteredRlocks =
            object.alreadyRegisteredRlocks?.map(e => Lock.fromPartial(e)) || [];
        return message;
    },
};

function createBaseCreateLeaseResponse(): CreateLeaseResponse {
    return { leaseId: "" };
}

export const CreateLeaseResponse: MessageFns<CreateLeaseResponse> = {
    encode(message: CreateLeaseResponse, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
        if (message.leaseId !== "") {
            writer.uint32(10).string(message.leaseId);
        }
        return writer;
    },

    decode(input: BinaryReader | Uint8Array, length?: number): CreateLeaseResponse {
        const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
        const end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseCreateLeaseResponse();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1: {
                    if (tag !== 10) {
                        break;
                    }

                    message.leaseId = reader.string();
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

    fromJSON(object: any): CreateLeaseResponse {
        return { leaseId: isSet(object.leaseId) ? globalThis.String(object.leaseId) : "" };
    },

    toJSON(message: CreateLeaseResponse): unknown {
        const obj: any = {};
        if (message.leaseId !== "") {
            obj.leaseId = message.leaseId;
        }
        return obj;
    },

    create(base?: DeepPartial<CreateLeaseResponse>): CreateLeaseResponse {
        return CreateLeaseResponse.fromPartial(base ?? {});
    },
    fromPartial(object: DeepPartial<CreateLeaseResponse>): CreateLeaseResponse {
        const message = createBaseCreateLeaseResponse();
        message.leaseId = object.leaseId ?? "";
        return message;
    },
};

function createBaseKeepLeaseRequest(): KeepLeaseRequest {
    return { leaseId: "", ttl: 0 };
}

export const KeepLeaseRequest: MessageFns<KeepLeaseRequest> = {
    encode(message: KeepLeaseRequest, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
        if (message.leaseId !== "") {
            writer.uint32(10).string(message.leaseId);
        }
        if (message.ttl !== 0) {
            writer.uint32(16).int32(message.ttl);
        }
        return writer;
    },

    decode(input: BinaryReader | Uint8Array, length?: number): KeepLeaseRequest {
        const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
        const end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseKeepLeaseRequest();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1: {
                    if (tag !== 10) {
                        break;
                    }

                    message.leaseId = reader.string();
                    continue;
                }
                case 2: {
                    if (tag !== 16) {
                        break;
                    }

                    message.ttl = reader.int32();
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

    fromJSON(object: any): KeepLeaseRequest {
        return {
            leaseId: isSet(object.leaseId) ? globalThis.String(object.leaseId) : "",
            ttl: isSet(object.ttl) ? globalThis.Number(object.ttl) : 0,
        };
    },

    toJSON(message: KeepLeaseRequest): unknown {
        const obj: any = {};
        if (message.leaseId !== "") {
            obj.leaseId = message.leaseId;
        }
        if (message.ttl !== 0) {
            obj.ttl = Math.round(message.ttl);
        }
        return obj;
    },

    create(base?: DeepPartial<KeepLeaseRequest>): KeepLeaseRequest {
        return KeepLeaseRequest.fromPartial(base ?? {});
    },
    fromPartial(object: DeepPartial<KeepLeaseRequest>): KeepLeaseRequest {
        const message = createBaseKeepLeaseRequest();
        message.leaseId = object.leaseId ?? "";
        message.ttl = object.ttl ?? 0;
        return message;
    },
};

function createBaseKeepLeaseResponse(): KeepLeaseResponse {
    return {};
}

export const KeepLeaseResponse: MessageFns<KeepLeaseResponse> = {
    encode(_: KeepLeaseResponse, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
        return writer;
    },

    decode(input: BinaryReader | Uint8Array, length?: number): KeepLeaseResponse {
        const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
        const end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseKeepLeaseResponse();
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

    fromJSON(_: any): KeepLeaseResponse {
        return {};
    },

    toJSON(_: KeepLeaseResponse): unknown {
        const obj: any = {};
        return obj;
    },

    create(base?: DeepPartial<KeepLeaseResponse>): KeepLeaseResponse {
        return KeepLeaseResponse.fromPartial(base ?? {});
    },
    fromPartial(_: DeepPartial<KeepLeaseResponse>): KeepLeaseResponse {
        const message = createBaseKeepLeaseResponse();
        return message;
    },
};

function createBaseDropLeaseRequest(): DropLeaseRequest {
    return { leaseId: "" };
}

export const DropLeaseRequest: MessageFns<DropLeaseRequest> = {
    encode(message: DropLeaseRequest, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
        if (message.leaseId !== "") {
            writer.uint32(10).string(message.leaseId);
        }
        return writer;
    },

    decode(input: BinaryReader | Uint8Array, length?: number): DropLeaseRequest {
        const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
        const end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseDropLeaseRequest();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1: {
                    if (tag !== 10) {
                        break;
                    }

                    message.leaseId = reader.string();
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

    fromJSON(object: any): DropLeaseRequest {
        return { leaseId: isSet(object.leaseId) ? globalThis.String(object.leaseId) : "" };
    },

    toJSON(message: DropLeaseRequest): unknown {
        const obj: any = {};
        if (message.leaseId !== "") {
            obj.leaseId = message.leaseId;
        }
        return obj;
    },

    create(base?: DeepPartial<DropLeaseRequest>): DropLeaseRequest {
        return DropLeaseRequest.fromPartial(base ?? {});
    },
    fromPartial(object: DeepPartial<DropLeaseRequest>): DropLeaseRequest {
        const message = createBaseDropLeaseRequest();
        message.leaseId = object.leaseId ?? "";
        return message;
    },
};

function createBaseDropLeaseResponse(): DropLeaseResponse {
    return {};
}

export const DropLeaseResponse: MessageFns<DropLeaseResponse> = {
    encode(_: DropLeaseResponse, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
        return writer;
    },

    decode(input: BinaryReader | Uint8Array, length?: number): DropLeaseResponse {
        const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
        const end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseDropLeaseResponse();
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

    fromJSON(_: any): DropLeaseResponse {
        return {};
    },

    toJSON(_: DropLeaseResponse): unknown {
        const obj: any = {};
        return obj;
    },

    create(base?: DeepPartial<DropLeaseResponse>): DropLeaseResponse {
        return DropLeaseResponse.fromPartial(base ?? {});
    },
    fromPartial(_: DeepPartial<DropLeaseResponse>): DropLeaseResponse {
        const message = createBaseDropLeaseResponse();
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
