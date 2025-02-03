// Code generated by protoc-gen-ts_proto. DO NOT EDIT.
// versions:
//   protoc-gen-ts_proto  v2.6.1
//   protoc               v5.27.3
// source: sync/management/lock.proto

/* eslint-disable */
import { BinaryReader, BinaryWriter } from "@bufbuild/protobuf/wire";

export interface LocalLockRequest {
    leaseId: string;
    tenantId: string;
    globalConcern: string;
    localConcern: string;
}

export interface LocalLockResponse {}

export interface LocalUnlockRequest {
    leaseId: string;
    tenantId: string;
    globalConcern: string;
    localConcern: string;
}

export interface LocalUnlockResponse {}

export interface GlobalLockRequest {
    leaseId: string;
    tenantId: string;
    globalConcern: string;
}

export interface GlobalLockResponse {}

export interface GlobalUnlockRequest {
    leaseId: string;
    tenantId: string;
    globalConcern: string;
}

export interface GlobalUnlockResponse {}

function createBaseLocalLockRequest(): LocalLockRequest {
    return { leaseId: "", tenantId: "", globalConcern: "", localConcern: "" };
}

export const LocalLockRequest: MessageFns<LocalLockRequest> = {
    encode(message: LocalLockRequest, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
        if (message.leaseId !== "") {
            writer.uint32(10).string(message.leaseId);
        }
        if (message.tenantId !== "") {
            writer.uint32(18).string(message.tenantId);
        }
        if (message.globalConcern !== "") {
            writer.uint32(26).string(message.globalConcern);
        }
        if (message.localConcern !== "") {
            writer.uint32(34).string(message.localConcern);
        }
        return writer;
    },

    decode(input: BinaryReader | Uint8Array, length?: number): LocalLockRequest {
        const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseLocalLockRequest();
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

                    message.globalConcern = reader.string();
                    continue;
                }
                case 4: {
                    if (tag !== 34) {
                        break;
                    }

                    message.localConcern = reader.string();
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

    fromJSON(object: any): LocalLockRequest {
        return {
            leaseId: isSet(object.leaseId) ? globalThis.String(object.leaseId) : "",
            tenantId: isSet(object.tenantId) ? globalThis.String(object.tenantId) : "",
            globalConcern: isSet(object.globalConcern)
                ? globalThis.String(object.globalConcern)
                : "",
            localConcern: isSet(object.localConcern) ? globalThis.String(object.localConcern) : "",
        };
    },

    toJSON(message: LocalLockRequest): unknown {
        const obj: any = {};
        if (message.leaseId !== "") {
            obj.leaseId = message.leaseId;
        }
        if (message.tenantId !== "") {
            obj.tenantId = message.tenantId;
        }
        if (message.globalConcern !== "") {
            obj.globalConcern = message.globalConcern;
        }
        if (message.localConcern !== "") {
            obj.localConcern = message.localConcern;
        }
        return obj;
    },

    create(base?: DeepPartial<LocalLockRequest>): LocalLockRequest {
        return LocalLockRequest.fromPartial(base ?? {});
    },
    fromPartial(object: DeepPartial<LocalLockRequest>): LocalLockRequest {
        const message = createBaseLocalLockRequest();
        message.leaseId = object.leaseId ?? "";
        message.tenantId = object.tenantId ?? "";
        message.globalConcern = object.globalConcern ?? "";
        message.localConcern = object.localConcern ?? "";
        return message;
    },
};

function createBaseLocalLockResponse(): LocalLockResponse {
    return {};
}

export const LocalLockResponse: MessageFns<LocalLockResponse> = {
    encode(_: LocalLockResponse, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
        return writer;
    },

    decode(input: BinaryReader | Uint8Array, length?: number): LocalLockResponse {
        const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseLocalLockResponse();
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

    fromJSON(_: any): LocalLockResponse {
        return {};
    },

    toJSON(_: LocalLockResponse): unknown {
        const obj: any = {};
        return obj;
    },

    create(base?: DeepPartial<LocalLockResponse>): LocalLockResponse {
        return LocalLockResponse.fromPartial(base ?? {});
    },
    fromPartial(_: DeepPartial<LocalLockResponse>): LocalLockResponse {
        const message = createBaseLocalLockResponse();
        return message;
    },
};

function createBaseLocalUnlockRequest(): LocalUnlockRequest {
    return { leaseId: "", tenantId: "", globalConcern: "", localConcern: "" };
}

export const LocalUnlockRequest: MessageFns<LocalUnlockRequest> = {
    encode(message: LocalUnlockRequest, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
        if (message.leaseId !== "") {
            writer.uint32(10).string(message.leaseId);
        }
        if (message.tenantId !== "") {
            writer.uint32(18).string(message.tenantId);
        }
        if (message.globalConcern !== "") {
            writer.uint32(26).string(message.globalConcern);
        }
        if (message.localConcern !== "") {
            writer.uint32(34).string(message.localConcern);
        }
        return writer;
    },

    decode(input: BinaryReader | Uint8Array, length?: number): LocalUnlockRequest {
        const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseLocalUnlockRequest();
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

                    message.globalConcern = reader.string();
                    continue;
                }
                case 4: {
                    if (tag !== 34) {
                        break;
                    }

                    message.localConcern = reader.string();
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

    fromJSON(object: any): LocalUnlockRequest {
        return {
            leaseId: isSet(object.leaseId) ? globalThis.String(object.leaseId) : "",
            tenantId: isSet(object.tenantId) ? globalThis.String(object.tenantId) : "",
            globalConcern: isSet(object.globalConcern)
                ? globalThis.String(object.globalConcern)
                : "",
            localConcern: isSet(object.localConcern) ? globalThis.String(object.localConcern) : "",
        };
    },

    toJSON(message: LocalUnlockRequest): unknown {
        const obj: any = {};
        if (message.leaseId !== "") {
            obj.leaseId = message.leaseId;
        }
        if (message.tenantId !== "") {
            obj.tenantId = message.tenantId;
        }
        if (message.globalConcern !== "") {
            obj.globalConcern = message.globalConcern;
        }
        if (message.localConcern !== "") {
            obj.localConcern = message.localConcern;
        }
        return obj;
    },

    create(base?: DeepPartial<LocalUnlockRequest>): LocalUnlockRequest {
        return LocalUnlockRequest.fromPartial(base ?? {});
    },
    fromPartial(object: DeepPartial<LocalUnlockRequest>): LocalUnlockRequest {
        const message = createBaseLocalUnlockRequest();
        message.leaseId = object.leaseId ?? "";
        message.tenantId = object.tenantId ?? "";
        message.globalConcern = object.globalConcern ?? "";
        message.localConcern = object.localConcern ?? "";
        return message;
    },
};

function createBaseLocalUnlockResponse(): LocalUnlockResponse {
    return {};
}

export const LocalUnlockResponse: MessageFns<LocalUnlockResponse> = {
    encode(_: LocalUnlockResponse, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
        return writer;
    },

    decode(input: BinaryReader | Uint8Array, length?: number): LocalUnlockResponse {
        const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseLocalUnlockResponse();
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

    fromJSON(_: any): LocalUnlockResponse {
        return {};
    },

    toJSON(_: LocalUnlockResponse): unknown {
        const obj: any = {};
        return obj;
    },

    create(base?: DeepPartial<LocalUnlockResponse>): LocalUnlockResponse {
        return LocalUnlockResponse.fromPartial(base ?? {});
    },
    fromPartial(_: DeepPartial<LocalUnlockResponse>): LocalUnlockResponse {
        const message = createBaseLocalUnlockResponse();
        return message;
    },
};

function createBaseGlobalLockRequest(): GlobalLockRequest {
    return { leaseId: "", tenantId: "", globalConcern: "" };
}

export const GlobalLockRequest: MessageFns<GlobalLockRequest> = {
    encode(message: GlobalLockRequest, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
        if (message.leaseId !== "") {
            writer.uint32(10).string(message.leaseId);
        }
        if (message.tenantId !== "") {
            writer.uint32(18).string(message.tenantId);
        }
        if (message.globalConcern !== "") {
            writer.uint32(26).string(message.globalConcern);
        }
        return writer;
    },

    decode(input: BinaryReader | Uint8Array, length?: number): GlobalLockRequest {
        const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseGlobalLockRequest();
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

                    message.globalConcern = reader.string();
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

    fromJSON(object: any): GlobalLockRequest {
        return {
            leaseId: isSet(object.leaseId) ? globalThis.String(object.leaseId) : "",
            tenantId: isSet(object.tenantId) ? globalThis.String(object.tenantId) : "",
            globalConcern: isSet(object.globalConcern)
                ? globalThis.String(object.globalConcern)
                : "",
        };
    },

    toJSON(message: GlobalLockRequest): unknown {
        const obj: any = {};
        if (message.leaseId !== "") {
            obj.leaseId = message.leaseId;
        }
        if (message.tenantId !== "") {
            obj.tenantId = message.tenantId;
        }
        if (message.globalConcern !== "") {
            obj.globalConcern = message.globalConcern;
        }
        return obj;
    },

    create(base?: DeepPartial<GlobalLockRequest>): GlobalLockRequest {
        return GlobalLockRequest.fromPartial(base ?? {});
    },
    fromPartial(object: DeepPartial<GlobalLockRequest>): GlobalLockRequest {
        const message = createBaseGlobalLockRequest();
        message.leaseId = object.leaseId ?? "";
        message.tenantId = object.tenantId ?? "";
        message.globalConcern = object.globalConcern ?? "";
        return message;
    },
};

function createBaseGlobalLockResponse(): GlobalLockResponse {
    return {};
}

export const GlobalLockResponse: MessageFns<GlobalLockResponse> = {
    encode(_: GlobalLockResponse, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
        return writer;
    },

    decode(input: BinaryReader | Uint8Array, length?: number): GlobalLockResponse {
        const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseGlobalLockResponse();
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

    fromJSON(_: any): GlobalLockResponse {
        return {};
    },

    toJSON(_: GlobalLockResponse): unknown {
        const obj: any = {};
        return obj;
    },

    create(base?: DeepPartial<GlobalLockResponse>): GlobalLockResponse {
        return GlobalLockResponse.fromPartial(base ?? {});
    },
    fromPartial(_: DeepPartial<GlobalLockResponse>): GlobalLockResponse {
        const message = createBaseGlobalLockResponse();
        return message;
    },
};

function createBaseGlobalUnlockRequest(): GlobalUnlockRequest {
    return { leaseId: "", tenantId: "", globalConcern: "" };
}

export const GlobalUnlockRequest: MessageFns<GlobalUnlockRequest> = {
    encode(message: GlobalUnlockRequest, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
        if (message.leaseId !== "") {
            writer.uint32(10).string(message.leaseId);
        }
        if (message.tenantId !== "") {
            writer.uint32(18).string(message.tenantId);
        }
        if (message.globalConcern !== "") {
            writer.uint32(26).string(message.globalConcern);
        }
        return writer;
    },

    decode(input: BinaryReader | Uint8Array, length?: number): GlobalUnlockRequest {
        const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseGlobalUnlockRequest();
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

                    message.globalConcern = reader.string();
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

    fromJSON(object: any): GlobalUnlockRequest {
        return {
            leaseId: isSet(object.leaseId) ? globalThis.String(object.leaseId) : "",
            tenantId: isSet(object.tenantId) ? globalThis.String(object.tenantId) : "",
            globalConcern: isSet(object.globalConcern)
                ? globalThis.String(object.globalConcern)
                : "",
        };
    },

    toJSON(message: GlobalUnlockRequest): unknown {
        const obj: any = {};
        if (message.leaseId !== "") {
            obj.leaseId = message.leaseId;
        }
        if (message.tenantId !== "") {
            obj.tenantId = message.tenantId;
        }
        if (message.globalConcern !== "") {
            obj.globalConcern = message.globalConcern;
        }
        return obj;
    },

    create(base?: DeepPartial<GlobalUnlockRequest>): GlobalUnlockRequest {
        return GlobalUnlockRequest.fromPartial(base ?? {});
    },
    fromPartial(object: DeepPartial<GlobalUnlockRequest>): GlobalUnlockRequest {
        const message = createBaseGlobalUnlockRequest();
        message.leaseId = object.leaseId ?? "";
        message.tenantId = object.tenantId ?? "";
        message.globalConcern = object.globalConcern ?? "";
        return message;
    },
};

function createBaseGlobalUnlockResponse(): GlobalUnlockResponse {
    return {};
}

export const GlobalUnlockResponse: MessageFns<GlobalUnlockResponse> = {
    encode(_: GlobalUnlockResponse, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
        return writer;
    },

    decode(input: BinaryReader | Uint8Array, length?: number): GlobalUnlockResponse {
        const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseGlobalUnlockResponse();
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

    fromJSON(_: any): GlobalUnlockResponse {
        return {};
    },

    toJSON(_: GlobalUnlockResponse): unknown {
        const obj: any = {};
        return obj;
    },

    create(base?: DeepPartial<GlobalUnlockResponse>): GlobalUnlockResponse {
        return GlobalUnlockResponse.fromPartial(base ?? {});
    },
    fromPartial(_: DeepPartial<GlobalUnlockResponse>): GlobalUnlockResponse {
        const message = createBaseGlobalUnlockResponse();
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
