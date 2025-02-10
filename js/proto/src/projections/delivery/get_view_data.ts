// Code generated by protoc-gen-ts_proto. DO NOT EDIT.
// versions:
//   protoc-gen-ts_proto  v2.6.1
//   protoc               v5.29.3
// source: projections/delivery/get_view_data.proto

/* eslint-disable */
import { BinaryReader, BinaryWriter } from "@bufbuild/protobuf/wire";
import { AuthData, Data, DataFilter, DataOrder } from "./shared";

export interface GetViewDataRequest {
    view: string;
    auth: AuthData | undefined;
    filter: DataFilter | undefined;
    useStrongConsistency: boolean;
    deploymentId: string;
}

export interface GetViewDataResponse {
    result: Data | undefined;
}

export interface GetViewDataListRequest {
    view: string;
    auth: AuthData | undefined;
    limit: string;
    page: string;
    filter: DataFilter | undefined;
    order: DataOrder[];
    useStrongConsistency: boolean;
    deploymentId: string;
}

export interface GetViewDataListResponse {
    result: Data[];
    limit: string;
    page: string;
    total: string;
}

function createBaseGetViewDataRequest(): GetViewDataRequest {
    return {
        view: "",
        auth: undefined,
        filter: undefined,
        useStrongConsistency: false,
        deploymentId: "0",
    };
}

export const GetViewDataRequest: MessageFns<GetViewDataRequest> = {
    encode(message: GetViewDataRequest, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
        if (message.view !== "") {
            writer.uint32(10).string(message.view);
        }
        if (message.auth !== undefined) {
            AuthData.encode(message.auth, writer.uint32(18).fork()).join();
        }
        if (message.filter !== undefined) {
            DataFilter.encode(message.filter, writer.uint32(26).fork()).join();
        }
        if (message.useStrongConsistency !== false) {
            writer.uint32(32).bool(message.useStrongConsistency);
        }
        if (message.deploymentId !== "0") {
            writer.uint32(40).int64(message.deploymentId);
        }
        return writer;
    },

    decode(input: BinaryReader | Uint8Array, length?: number): GetViewDataRequest {
        const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseGetViewDataRequest();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1: {
                    if (tag !== 10) {
                        break;
                    }

                    message.view = reader.string();
                    continue;
                }
                case 2: {
                    if (tag !== 18) {
                        break;
                    }

                    message.auth = AuthData.decode(reader, reader.uint32());
                    continue;
                }
                case 3: {
                    if (tag !== 26) {
                        break;
                    }

                    message.filter = DataFilter.decode(reader, reader.uint32());
                    continue;
                }
                case 4: {
                    if (tag !== 32) {
                        break;
                    }

                    message.useStrongConsistency = reader.bool();
                    continue;
                }
                case 5: {
                    if (tag !== 40) {
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

    fromJSON(object: any): GetViewDataRequest {
        return {
            view: isSet(object.view) ? globalThis.String(object.view) : "",
            auth: isSet(object.auth) ? AuthData.fromJSON(object.auth) : undefined,
            filter: isSet(object.filter) ? DataFilter.fromJSON(object.filter) : undefined,
            useStrongConsistency: isSet(object.useStrongConsistency)
                ? globalThis.Boolean(object.useStrongConsistency)
                : false,
            deploymentId: isSet(object.deploymentId) ? globalThis.String(object.deploymentId) : "0",
        };
    },

    toJSON(message: GetViewDataRequest): unknown {
        const obj: any = {};
        if (message.view !== "") {
            obj.view = message.view;
        }
        if (message.auth !== undefined) {
            obj.auth = AuthData.toJSON(message.auth);
        }
        if (message.filter !== undefined) {
            obj.filter = DataFilter.toJSON(message.filter);
        }
        if (message.useStrongConsistency !== false) {
            obj.useStrongConsistency = message.useStrongConsistency;
        }
        if (message.deploymentId !== "0") {
            obj.deploymentId = message.deploymentId;
        }
        return obj;
    },

    create(base?: DeepPartial<GetViewDataRequest>): GetViewDataRequest {
        return GetViewDataRequest.fromPartial(base ?? {});
    },
    fromPartial(object: DeepPartial<GetViewDataRequest>): GetViewDataRequest {
        const message = createBaseGetViewDataRequest();
        message.view = object.view ?? "";
        message.auth =
            object.auth !== undefined && object.auth !== null
                ? AuthData.fromPartial(object.auth)
                : undefined;
        message.filter =
            object.filter !== undefined && object.filter !== null
                ? DataFilter.fromPartial(object.filter)
                : undefined;
        message.useStrongConsistency = object.useStrongConsistency ?? false;
        message.deploymentId = object.deploymentId ?? "0";
        return message;
    },
};

function createBaseGetViewDataResponse(): GetViewDataResponse {
    return { result: undefined };
}

export const GetViewDataResponse: MessageFns<GetViewDataResponse> = {
    encode(message: GetViewDataResponse, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
        if (message.result !== undefined) {
            Data.encode(message.result, writer.uint32(10).fork()).join();
        }
        return writer;
    },

    decode(input: BinaryReader | Uint8Array, length?: number): GetViewDataResponse {
        const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseGetViewDataResponse();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1: {
                    if (tag !== 10) {
                        break;
                    }

                    message.result = Data.decode(reader, reader.uint32());
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

    fromJSON(object: any): GetViewDataResponse {
        return { result: isSet(object.result) ? Data.fromJSON(object.result) : undefined };
    },

    toJSON(message: GetViewDataResponse): unknown {
        const obj: any = {};
        if (message.result !== undefined) {
            obj.result = Data.toJSON(message.result);
        }
        return obj;
    },

    create(base?: DeepPartial<GetViewDataResponse>): GetViewDataResponse {
        return GetViewDataResponse.fromPartial(base ?? {});
    },
    fromPartial(object: DeepPartial<GetViewDataResponse>): GetViewDataResponse {
        const message = createBaseGetViewDataResponse();
        message.result =
            object.result !== undefined && object.result !== null
                ? Data.fromPartial(object.result)
                : undefined;
        return message;
    },
};

function createBaseGetViewDataListRequest(): GetViewDataListRequest {
    return {
        view: "",
        auth: undefined,
        limit: "0",
        page: "0",
        filter: undefined,
        order: [],
        useStrongConsistency: false,
        deploymentId: "0",
    };
}

export const GetViewDataListRequest: MessageFns<GetViewDataListRequest> = {
    encode(
        message: GetViewDataListRequest,
        writer: BinaryWriter = new BinaryWriter()
    ): BinaryWriter {
        if (message.view !== "") {
            writer.uint32(10).string(message.view);
        }
        if (message.auth !== undefined) {
            AuthData.encode(message.auth, writer.uint32(18).fork()).join();
        }
        if (message.limit !== "0") {
            writer.uint32(24).int64(message.limit);
        }
        if (message.page !== "0") {
            writer.uint32(32).int64(message.page);
        }
        if (message.filter !== undefined) {
            DataFilter.encode(message.filter, writer.uint32(42).fork()).join();
        }
        for (const v of message.order) {
            DataOrder.encode(v!, writer.uint32(50).fork()).join();
        }
        if (message.useStrongConsistency !== false) {
            writer.uint32(56).bool(message.useStrongConsistency);
        }
        if (message.deploymentId !== "0") {
            writer.uint32(64).int64(message.deploymentId);
        }
        return writer;
    },

    decode(input: BinaryReader | Uint8Array, length?: number): GetViewDataListRequest {
        const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseGetViewDataListRequest();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1: {
                    if (tag !== 10) {
                        break;
                    }

                    message.view = reader.string();
                    continue;
                }
                case 2: {
                    if (tag !== 18) {
                        break;
                    }

                    message.auth = AuthData.decode(reader, reader.uint32());
                    continue;
                }
                case 3: {
                    if (tag !== 24) {
                        break;
                    }

                    message.limit = reader.int64().toString();
                    continue;
                }
                case 4: {
                    if (tag !== 32) {
                        break;
                    }

                    message.page = reader.int64().toString();
                    continue;
                }
                case 5: {
                    if (tag !== 42) {
                        break;
                    }

                    message.filter = DataFilter.decode(reader, reader.uint32());
                    continue;
                }
                case 6: {
                    if (tag !== 50) {
                        break;
                    }

                    message.order.push(DataOrder.decode(reader, reader.uint32()));
                    continue;
                }
                case 7: {
                    if (tag !== 56) {
                        break;
                    }

                    message.useStrongConsistency = reader.bool();
                    continue;
                }
                case 8: {
                    if (tag !== 64) {
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

    fromJSON(object: any): GetViewDataListRequest {
        return {
            view: isSet(object.view) ? globalThis.String(object.view) : "",
            auth: isSet(object.auth) ? AuthData.fromJSON(object.auth) : undefined,
            limit: isSet(object.limit) ? globalThis.String(object.limit) : "0",
            page: isSet(object.page) ? globalThis.String(object.page) : "0",
            filter: isSet(object.filter) ? DataFilter.fromJSON(object.filter) : undefined,
            order: globalThis.Array.isArray(object?.order)
                ? object.order.map((e: any) => DataOrder.fromJSON(e))
                : [],
            useStrongConsistency: isSet(object.useStrongConsistency)
                ? globalThis.Boolean(object.useStrongConsistency)
                : false,
            deploymentId: isSet(object.deploymentId) ? globalThis.String(object.deploymentId) : "0",
        };
    },

    toJSON(message: GetViewDataListRequest): unknown {
        const obj: any = {};
        if (message.view !== "") {
            obj.view = message.view;
        }
        if (message.auth !== undefined) {
            obj.auth = AuthData.toJSON(message.auth);
        }
        if (message.limit !== "0") {
            obj.limit = message.limit;
        }
        if (message.page !== "0") {
            obj.page = message.page;
        }
        if (message.filter !== undefined) {
            obj.filter = DataFilter.toJSON(message.filter);
        }
        if (message.order?.length) {
            obj.order = message.order.map(e => DataOrder.toJSON(e));
        }
        if (message.useStrongConsistency !== false) {
            obj.useStrongConsistency = message.useStrongConsistency;
        }
        if (message.deploymentId !== "0") {
            obj.deploymentId = message.deploymentId;
        }
        return obj;
    },

    create(base?: DeepPartial<GetViewDataListRequest>): GetViewDataListRequest {
        return GetViewDataListRequest.fromPartial(base ?? {});
    },
    fromPartial(object: DeepPartial<GetViewDataListRequest>): GetViewDataListRequest {
        const message = createBaseGetViewDataListRequest();
        message.view = object.view ?? "";
        message.auth =
            object.auth !== undefined && object.auth !== null
                ? AuthData.fromPartial(object.auth)
                : undefined;
        message.limit = object.limit ?? "0";
        message.page = object.page ?? "0";
        message.filter =
            object.filter !== undefined && object.filter !== null
                ? DataFilter.fromPartial(object.filter)
                : undefined;
        message.order = object.order?.map(e => DataOrder.fromPartial(e)) || [];
        message.useStrongConsistency = object.useStrongConsistency ?? false;
        message.deploymentId = object.deploymentId ?? "0";
        return message;
    },
};

function createBaseGetViewDataListResponse(): GetViewDataListResponse {
    return { result: [], limit: "0", page: "0", total: "0" };
}

export const GetViewDataListResponse: MessageFns<GetViewDataListResponse> = {
    encode(
        message: GetViewDataListResponse,
        writer: BinaryWriter = new BinaryWriter()
    ): BinaryWriter {
        for (const v of message.result) {
            Data.encode(v!, writer.uint32(10).fork()).join();
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

    decode(input: BinaryReader | Uint8Array, length?: number): GetViewDataListResponse {
        const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseGetViewDataListResponse();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1: {
                    if (tag !== 10) {
                        break;
                    }

                    message.result.push(Data.decode(reader, reader.uint32()));
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

    fromJSON(object: any): GetViewDataListResponse {
        return {
            result: globalThis.Array.isArray(object?.result)
                ? object.result.map((e: any) => Data.fromJSON(e))
                : [],
            limit: isSet(object.limit) ? globalThis.String(object.limit) : "0",
            page: isSet(object.page) ? globalThis.String(object.page) : "0",
            total: isSet(object.total) ? globalThis.String(object.total) : "0",
        };
    },

    toJSON(message: GetViewDataListResponse): unknown {
        const obj: any = {};
        if (message.result?.length) {
            obj.result = message.result.map(e => Data.toJSON(e));
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

    create(base?: DeepPartial<GetViewDataListResponse>): GetViewDataListResponse {
        return GetViewDataListResponse.fromPartial(base ?? {});
    },
    fromPartial(object: DeepPartial<GetViewDataListResponse>): GetViewDataListResponse {
        const message = createBaseGetViewDataListResponse();
        message.result = object.result?.map(e => Data.fromPartial(e)) || [];
        message.limit = object.limit ?? "0";
        message.page = object.page ?? "0";
        message.total = object.total ?? "0";
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
