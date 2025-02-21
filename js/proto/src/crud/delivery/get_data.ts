// Code generated by protoc-gen-ts_proto. DO NOT EDIT.
// versions:
//   protoc-gen-ts_proto  v2.6.1
//   protoc               v5.29.3
// source: crud/delivery/get_data.proto

/* eslint-disable */
import { BinaryReader, BinaryWriter } from "@bufbuild/protobuf/wire";
import {
    AuthData,
    DataFilter,
    DeploymentTarget,
    deploymentTargetFromJSON,
    deploymentTargetToJSON,
    deploymentTargetToNumber,
} from "./shared";

export interface GetDataRequest {
    type: string;
    auth: AuthData | undefined;
    id: string;
    filter: DataFilter | undefined;
    returnEmptyDataIfNotFound: boolean;
    wait: DataWait | undefined;
    useStrongConsistency: boolean;
    target: DeploymentTarget;
}

export interface GetDataResponse {
    result: Data | undefined;
}

export interface GetDataListRequest {
    type: string;
    auth: AuthData | undefined;
    limit: string;
    page: string;
    filter: DataFilter | undefined;
    order: DataOrder[];
    useStrongConsistency: boolean;
    target: DeploymentTarget;
}

export interface GetDataListResponse {
    result: Data[];
    limit: string;
    page: string;
    total: string;
}

export interface DataWait {
    conditionFilter: DataFilter | undefined;
    timeout: string;
}

export interface DataOrder {
    field: string;
    descending: boolean;
}

export interface Data {
    data: { [key: string]: string };
}

export interface Data_DataEntry {
    key: string;
    value: string;
}

function createBaseGetDataRequest(): GetDataRequest {
    return {
        type: "",
        auth: undefined,
        id: "",
        filter: undefined,
        returnEmptyDataIfNotFound: false,
        wait: undefined,
        useStrongConsistency: false,
        target: DeploymentTarget.DEPLOYMENT_TARGET_BLUE,
    };
}

export const GetDataRequest: MessageFns<GetDataRequest> = {
    encode(message: GetDataRequest, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
        if (message.type !== "") {
            writer.uint32(10).string(message.type);
        }
        if (message.auth !== undefined) {
            AuthData.encode(message.auth, writer.uint32(18).fork()).join();
        }
        if (message.id !== "") {
            writer.uint32(26).string(message.id);
        }
        if (message.filter !== undefined) {
            DataFilter.encode(message.filter, writer.uint32(34).fork()).join();
        }
        if (message.returnEmptyDataIfNotFound !== false) {
            writer.uint32(40).bool(message.returnEmptyDataIfNotFound);
        }
        if (message.wait !== undefined) {
            DataWait.encode(message.wait, writer.uint32(50).fork()).join();
        }
        if (message.useStrongConsistency !== false) {
            writer.uint32(56).bool(message.useStrongConsistency);
        }
        if (message.target !== DeploymentTarget.DEPLOYMENT_TARGET_BLUE) {
            writer.uint32(64).int32(deploymentTargetToNumber(message.target));
        }
        return writer;
    },

    decode(input: BinaryReader | Uint8Array, length?: number): GetDataRequest {
        const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseGetDataRequest();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1: {
                    if (tag !== 10) {
                        break;
                    }

                    message.type = reader.string();
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

                    message.id = reader.string();
                    continue;
                }
                case 4: {
                    if (tag !== 34) {
                        break;
                    }

                    message.filter = DataFilter.decode(reader, reader.uint32());
                    continue;
                }
                case 5: {
                    if (tag !== 40) {
                        break;
                    }

                    message.returnEmptyDataIfNotFound = reader.bool();
                    continue;
                }
                case 6: {
                    if (tag !== 50) {
                        break;
                    }

                    message.wait = DataWait.decode(reader, reader.uint32());
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

                    message.target = deploymentTargetFromJSON(reader.int32());
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

    fromJSON(object: any): GetDataRequest {
        return {
            type: isSet(object.type) ? globalThis.String(object.type) : "",
            auth: isSet(object.auth) ? AuthData.fromJSON(object.auth) : undefined,
            id: isSet(object.id) ? globalThis.String(object.id) : "",
            filter: isSet(object.filter) ? DataFilter.fromJSON(object.filter) : undefined,
            returnEmptyDataIfNotFound: isSet(object.returnEmptyDataIfNotFound)
                ? globalThis.Boolean(object.returnEmptyDataIfNotFound)
                : false,
            wait: isSet(object.wait) ? DataWait.fromJSON(object.wait) : undefined,
            useStrongConsistency: isSet(object.useStrongConsistency)
                ? globalThis.Boolean(object.useStrongConsistency)
                : false,
            target: isSet(object.target)
                ? deploymentTargetFromJSON(object.target)
                : DeploymentTarget.DEPLOYMENT_TARGET_BLUE,
        };
    },

    toJSON(message: GetDataRequest): unknown {
        const obj: any = {};
        if (message.type !== "") {
            obj.type = message.type;
        }
        if (message.auth !== undefined) {
            obj.auth = AuthData.toJSON(message.auth);
        }
        if (message.id !== "") {
            obj.id = message.id;
        }
        if (message.filter !== undefined) {
            obj.filter = DataFilter.toJSON(message.filter);
        }
        if (message.returnEmptyDataIfNotFound !== false) {
            obj.returnEmptyDataIfNotFound = message.returnEmptyDataIfNotFound;
        }
        if (message.wait !== undefined) {
            obj.wait = DataWait.toJSON(message.wait);
        }
        if (message.useStrongConsistency !== false) {
            obj.useStrongConsistency = message.useStrongConsistency;
        }
        if (message.target !== DeploymentTarget.DEPLOYMENT_TARGET_BLUE) {
            obj.target = deploymentTargetToJSON(message.target);
        }
        return obj;
    },

    create(base?: DeepPartial<GetDataRequest>): GetDataRequest {
        return GetDataRequest.fromPartial(base ?? {});
    },
    fromPartial(object: DeepPartial<GetDataRequest>): GetDataRequest {
        const message = createBaseGetDataRequest();
        message.type = object.type ?? "";
        message.auth =
            object.auth !== undefined && object.auth !== null
                ? AuthData.fromPartial(object.auth)
                : undefined;
        message.id = object.id ?? "";
        message.filter =
            object.filter !== undefined && object.filter !== null
                ? DataFilter.fromPartial(object.filter)
                : undefined;
        message.returnEmptyDataIfNotFound = object.returnEmptyDataIfNotFound ?? false;
        message.wait =
            object.wait !== undefined && object.wait !== null
                ? DataWait.fromPartial(object.wait)
                : undefined;
        message.useStrongConsistency = object.useStrongConsistency ?? false;
        message.target = object.target ?? DeploymentTarget.DEPLOYMENT_TARGET_BLUE;
        return message;
    },
};

function createBaseGetDataResponse(): GetDataResponse {
    return { result: undefined };
}

export const GetDataResponse: MessageFns<GetDataResponse> = {
    encode(message: GetDataResponse, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
        if (message.result !== undefined) {
            Data.encode(message.result, writer.uint32(10).fork()).join();
        }
        return writer;
    },

    decode(input: BinaryReader | Uint8Array, length?: number): GetDataResponse {
        const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseGetDataResponse();
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

    fromJSON(object: any): GetDataResponse {
        return { result: isSet(object.result) ? Data.fromJSON(object.result) : undefined };
    },

    toJSON(message: GetDataResponse): unknown {
        const obj: any = {};
        if (message.result !== undefined) {
            obj.result = Data.toJSON(message.result);
        }
        return obj;
    },

    create(base?: DeepPartial<GetDataResponse>): GetDataResponse {
        return GetDataResponse.fromPartial(base ?? {});
    },
    fromPartial(object: DeepPartial<GetDataResponse>): GetDataResponse {
        const message = createBaseGetDataResponse();
        message.result =
            object.result !== undefined && object.result !== null
                ? Data.fromPartial(object.result)
                : undefined;
        return message;
    },
};

function createBaseGetDataListRequest(): GetDataListRequest {
    return {
        type: "",
        auth: undefined,
        limit: "0",
        page: "0",
        filter: undefined,
        order: [],
        useStrongConsistency: false,
        target: DeploymentTarget.DEPLOYMENT_TARGET_BLUE,
    };
}

export const GetDataListRequest: MessageFns<GetDataListRequest> = {
    encode(message: GetDataListRequest, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
        if (message.type !== "") {
            writer.uint32(10).string(message.type);
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
            DataFilter.encode(message.filter, writer.uint32(58).fork()).join();
        }
        for (const v of message.order) {
            DataOrder.encode(v!, writer.uint32(66).fork()).join();
        }
        if (message.useStrongConsistency !== false) {
            writer.uint32(72).bool(message.useStrongConsistency);
        }
        if (message.target !== DeploymentTarget.DEPLOYMENT_TARGET_BLUE) {
            writer.uint32(80).int32(deploymentTargetToNumber(message.target));
        }
        return writer;
    },

    decode(input: BinaryReader | Uint8Array, length?: number): GetDataListRequest {
        const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseGetDataListRequest();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1: {
                    if (tag !== 10) {
                        break;
                    }

                    message.type = reader.string();
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
                case 7: {
                    if (tag !== 58) {
                        break;
                    }

                    message.filter = DataFilter.decode(reader, reader.uint32());
                    continue;
                }
                case 8: {
                    if (tag !== 66) {
                        break;
                    }

                    message.order.push(DataOrder.decode(reader, reader.uint32()));
                    continue;
                }
                case 9: {
                    if (tag !== 72) {
                        break;
                    }

                    message.useStrongConsistency = reader.bool();
                    continue;
                }
                case 10: {
                    if (tag !== 80) {
                        break;
                    }

                    message.target = deploymentTargetFromJSON(reader.int32());
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

    fromJSON(object: any): GetDataListRequest {
        return {
            type: isSet(object.type) ? globalThis.String(object.type) : "",
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
            target: isSet(object.target)
                ? deploymentTargetFromJSON(object.target)
                : DeploymentTarget.DEPLOYMENT_TARGET_BLUE,
        };
    },

    toJSON(message: GetDataListRequest): unknown {
        const obj: any = {};
        if (message.type !== "") {
            obj.type = message.type;
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
        if (message.target !== DeploymentTarget.DEPLOYMENT_TARGET_BLUE) {
            obj.target = deploymentTargetToJSON(message.target);
        }
        return obj;
    },

    create(base?: DeepPartial<GetDataListRequest>): GetDataListRequest {
        return GetDataListRequest.fromPartial(base ?? {});
    },
    fromPartial(object: DeepPartial<GetDataListRequest>): GetDataListRequest {
        const message = createBaseGetDataListRequest();
        message.type = object.type ?? "";
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
        message.target = object.target ?? DeploymentTarget.DEPLOYMENT_TARGET_BLUE;
        return message;
    },
};

function createBaseGetDataListResponse(): GetDataListResponse {
    return { result: [], limit: "0", page: "0", total: "0" };
}

export const GetDataListResponse: MessageFns<GetDataListResponse> = {
    encode(message: GetDataListResponse, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
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

    decode(input: BinaryReader | Uint8Array, length?: number): GetDataListResponse {
        const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseGetDataListResponse();
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

    fromJSON(object: any): GetDataListResponse {
        return {
            result: globalThis.Array.isArray(object?.result)
                ? object.result.map((e: any) => Data.fromJSON(e))
                : [],
            limit: isSet(object.limit) ? globalThis.String(object.limit) : "0",
            page: isSet(object.page) ? globalThis.String(object.page) : "0",
            total: isSet(object.total) ? globalThis.String(object.total) : "0",
        };
    },

    toJSON(message: GetDataListResponse): unknown {
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

    create(base?: DeepPartial<GetDataListResponse>): GetDataListResponse {
        return GetDataListResponse.fromPartial(base ?? {});
    },
    fromPartial(object: DeepPartial<GetDataListResponse>): GetDataListResponse {
        const message = createBaseGetDataListResponse();
        message.result = object.result?.map(e => Data.fromPartial(e)) || [];
        message.limit = object.limit ?? "0";
        message.page = object.page ?? "0";
        message.total = object.total ?? "0";
        return message;
    },
};

function createBaseDataWait(): DataWait {
    return { conditionFilter: undefined, timeout: "0" };
}

export const DataWait: MessageFns<DataWait> = {
    encode(message: DataWait, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
        if (message.conditionFilter !== undefined) {
            DataFilter.encode(message.conditionFilter, writer.uint32(10).fork()).join();
        }
        if (message.timeout !== "0") {
            writer.uint32(16).int64(message.timeout);
        }
        return writer;
    },

    decode(input: BinaryReader | Uint8Array, length?: number): DataWait {
        const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseDataWait();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1: {
                    if (tag !== 10) {
                        break;
                    }

                    message.conditionFilter = DataFilter.decode(reader, reader.uint32());
                    continue;
                }
                case 2: {
                    if (tag !== 16) {
                        break;
                    }

                    message.timeout = reader.int64().toString();
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

    fromJSON(object: any): DataWait {
        return {
            conditionFilter: isSet(object.conditionFilter)
                ? DataFilter.fromJSON(object.conditionFilter)
                : undefined,
            timeout: isSet(object.timeout) ? globalThis.String(object.timeout) : "0",
        };
    },

    toJSON(message: DataWait): unknown {
        const obj: any = {};
        if (message.conditionFilter !== undefined) {
            obj.conditionFilter = DataFilter.toJSON(message.conditionFilter);
        }
        if (message.timeout !== "0") {
            obj.timeout = message.timeout;
        }
        return obj;
    },

    create(base?: DeepPartial<DataWait>): DataWait {
        return DataWait.fromPartial(base ?? {});
    },
    fromPartial(object: DeepPartial<DataWait>): DataWait {
        const message = createBaseDataWait();
        message.conditionFilter =
            object.conditionFilter !== undefined && object.conditionFilter !== null
                ? DataFilter.fromPartial(object.conditionFilter)
                : undefined;
        message.timeout = object.timeout ?? "0";
        return message;
    },
};

function createBaseDataOrder(): DataOrder {
    return { field: "", descending: false };
}

export const DataOrder: MessageFns<DataOrder> = {
    encode(message: DataOrder, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
        if (message.field !== "") {
            writer.uint32(10).string(message.field);
        }
        if (message.descending !== false) {
            writer.uint32(16).bool(message.descending);
        }
        return writer;
    },

    decode(input: BinaryReader | Uint8Array, length?: number): DataOrder {
        const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseDataOrder();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1: {
                    if (tag !== 10) {
                        break;
                    }

                    message.field = reader.string();
                    continue;
                }
                case 2: {
                    if (tag !== 16) {
                        break;
                    }

                    message.descending = reader.bool();
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

    fromJSON(object: any): DataOrder {
        return {
            field: isSet(object.field) ? globalThis.String(object.field) : "",
            descending: isSet(object.descending) ? globalThis.Boolean(object.descending) : false,
        };
    },

    toJSON(message: DataOrder): unknown {
        const obj: any = {};
        if (message.field !== "") {
            obj.field = message.field;
        }
        if (message.descending !== false) {
            obj.descending = message.descending;
        }
        return obj;
    },

    create(base?: DeepPartial<DataOrder>): DataOrder {
        return DataOrder.fromPartial(base ?? {});
    },
    fromPartial(object: DeepPartial<DataOrder>): DataOrder {
        const message = createBaseDataOrder();
        message.field = object.field ?? "";
        message.descending = object.descending ?? false;
        return message;
    },
};

function createBaseData(): Data {
    return { data: {} };
}

export const Data: MessageFns<Data> = {
    encode(message: Data, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
        Object.entries(message.data).forEach(([key, value]) => {
            Data_DataEntry.encode({ key: key as any, value }, writer.uint32(10).fork()).join();
        });
        return writer;
    },

    decode(input: BinaryReader | Uint8Array, length?: number): Data {
        const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseData();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1: {
                    if (tag !== 10) {
                        break;
                    }

                    const entry1 = Data_DataEntry.decode(reader, reader.uint32());
                    if (entry1.value !== undefined) {
                        message.data[entry1.key] = entry1.value;
                    }
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

    fromJSON(object: any): Data {
        return {
            data: isObject(object.data)
                ? Object.entries(object.data).reduce<{ [key: string]: string }>(
                      (acc, [key, value]) => {
                          acc[key] = String(value);
                          return acc;
                      },
                      {}
                  )
                : {},
        };
    },

    toJSON(message: Data): unknown {
        const obj: any = {};
        if (message.data) {
            const entries = Object.entries(message.data);
            if (entries.length > 0) {
                obj.data = {};
                entries.forEach(([k, v]) => {
                    obj.data[k] = v;
                });
            }
        }
        return obj;
    },

    create(base?: DeepPartial<Data>): Data {
        return Data.fromPartial(base ?? {});
    },
    fromPartial(object: DeepPartial<Data>): Data {
        const message = createBaseData();
        message.data = Object.entries(object.data ?? {}).reduce<{ [key: string]: string }>(
            (acc, [key, value]) => {
                if (value !== undefined) {
                    acc[key] = globalThis.String(value);
                }
                return acc;
            },
            {}
        );
        return message;
    },
};

function createBaseData_DataEntry(): Data_DataEntry {
    return { key: "", value: "" };
}

export const Data_DataEntry: MessageFns<Data_DataEntry> = {
    encode(message: Data_DataEntry, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
        if (message.key !== "") {
            writer.uint32(10).string(message.key);
        }
        if (message.value !== "") {
            writer.uint32(18).string(message.value);
        }
        return writer;
    },

    decode(input: BinaryReader | Uint8Array, length?: number): Data_DataEntry {
        const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseData_DataEntry();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1: {
                    if (tag !== 10) {
                        break;
                    }

                    message.key = reader.string();
                    continue;
                }
                case 2: {
                    if (tag !== 18) {
                        break;
                    }

                    message.value = reader.string();
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

    fromJSON(object: any): Data_DataEntry {
        return {
            key: isSet(object.key) ? globalThis.String(object.key) : "",
            value: isSet(object.value) ? globalThis.String(object.value) : "",
        };
    },

    toJSON(message: Data_DataEntry): unknown {
        const obj: any = {};
        if (message.key !== "") {
            obj.key = message.key;
        }
        if (message.value !== "") {
            obj.value = message.value;
        }
        return obj;
    },

    create(base?: DeepPartial<Data_DataEntry>): Data_DataEntry {
        return Data_DataEntry.fromPartial(base ?? {});
    },
    fromPartial(object: DeepPartial<Data_DataEntry>): Data_DataEntry {
        const message = createBaseData_DataEntry();
        message.key = object.key ?? "";
        message.value = object.value ?? "";
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

function isObject(value: any): boolean {
    return typeof value === "object" && value !== null;
}

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
