// Code generated by protoc-gen-ts_proto. DO NOT EDIT.
// versions:
//   protoc-gen-ts_proto  v2.7.0
//   protoc               v5.29.3
// source: projections/delivery/shared.proto

/* eslint-disable */
import { BinaryReader, BinaryWriter } from "@bufbuild/protobuf/wire";

export const DeploymentTarget = {
    DEPLOYMENT_TARGET_BLUE: "DEPLOYMENT_TARGET_BLUE",
    DEPLOYMENT_TARGET_GREEN: "DEPLOYMENT_TARGET_GREEN",
} as const;

export type DeploymentTarget = (typeof DeploymentTarget)[keyof typeof DeploymentTarget];

export namespace DeploymentTarget {
    export type DEPLOYMENT_TARGET_BLUE = typeof DeploymentTarget.DEPLOYMENT_TARGET_BLUE;
    export type DEPLOYMENT_TARGET_GREEN = typeof DeploymentTarget.DEPLOYMENT_TARGET_GREEN;
}

export function deploymentTargetFromJSON(object: any): DeploymentTarget {
    switch (object) {
        case 0:
        case "DEPLOYMENT_TARGET_BLUE":
            return DeploymentTarget.DEPLOYMENT_TARGET_BLUE;
        case 1:
        case "DEPLOYMENT_TARGET_GREEN":
            return DeploymentTarget.DEPLOYMENT_TARGET_GREEN;
        default:
            throw new globalThis.Error(
                "Unrecognized enum value " + object + " for enum DeploymentTarget"
            );
    }
}

export function deploymentTargetToJSON(object: DeploymentTarget): string {
    switch (object) {
        case DeploymentTarget.DEPLOYMENT_TARGET_BLUE:
            return "DEPLOYMENT_TARGET_BLUE";
        case DeploymentTarget.DEPLOYMENT_TARGET_GREEN:
            return "DEPLOYMENT_TARGET_GREEN";
        default:
            throw new globalThis.Error(
                "Unrecognized enum value " + object + " for enum DeploymentTarget"
            );
    }
}

export function deploymentTargetToNumber(object: DeploymentTarget): number {
    switch (object) {
        case DeploymentTarget.DEPLOYMENT_TARGET_BLUE:
            return 0;
        case DeploymentTarget.DEPLOYMENT_TARGET_GREEN:
            return 1;
        default:
            throw new globalThis.Error(
                "Unrecognized enum value " + object + " for enum DeploymentTarget"
            );
    }
}

export interface AuthData {
    tenantId: string;
    userId: string;
    scopes: string[];
    data: { [key: string]: string };
}

export interface AuthData_DataEntry {
    key: string;
    value: string;
}

export interface DataFilter {
    fields: { [key: string]: DataFieldFilter };
    and: DataFilter[];
    or: DataFilter[];
}

export interface DataFilter_FieldsEntry {
    key: string;
    value: DataFieldFilter | undefined;
}

export interface DataFieldFilter {
    type: string;
    operation: string;
    value: string;
}

export interface EventMetadata {
    causationId: string;
    correlationId: string;
    target: DeploymentTarget;
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

function createBaseAuthData(): AuthData {
    return { tenantId: "", userId: "", scopes: [], data: {} };
}

export const AuthData: MessageFns<AuthData> = {
    encode(message: AuthData, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
        if (message.tenantId !== "") {
            writer.uint32(10).string(message.tenantId);
        }
        if (message.userId !== "") {
            writer.uint32(18).string(message.userId);
        }
        for (const v of message.scopes) {
            writer.uint32(26).string(v!);
        }
        Object.entries(message.data).forEach(([key, value]) => {
            AuthData_DataEntry.encode({ key: key as any, value }, writer.uint32(34).fork()).join();
        });
        return writer;
    },

    decode(input: BinaryReader | Uint8Array, length?: number): AuthData {
        const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseAuthData();
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

                    message.userId = reader.string();
                    continue;
                }
                case 3: {
                    if (tag !== 26) {
                        break;
                    }

                    message.scopes.push(reader.string());
                    continue;
                }
                case 4: {
                    if (tag !== 34) {
                        break;
                    }

                    const entry4 = AuthData_DataEntry.decode(reader, reader.uint32());
                    if (entry4.value !== undefined) {
                        message.data[entry4.key] = entry4.value;
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

    fromJSON(object: any): AuthData {
        return {
            tenantId: isSet(object.tenantId) ? globalThis.String(object.tenantId) : "",
            userId: isSet(object.userId) ? globalThis.String(object.userId) : "",
            scopes: globalThis.Array.isArray(object?.scopes)
                ? object.scopes.map((e: any) => globalThis.String(e))
                : [],
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

    toJSON(message: AuthData): unknown {
        const obj: any = {};
        if (message.tenantId !== "") {
            obj.tenantId = message.tenantId;
        }
        if (message.userId !== "") {
            obj.userId = message.userId;
        }
        if (message.scopes?.length) {
            obj.scopes = message.scopes;
        }
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

    create(base?: DeepPartial<AuthData>): AuthData {
        return AuthData.fromPartial(base ?? {});
    },
    fromPartial(object: DeepPartial<AuthData>): AuthData {
        const message = createBaseAuthData();
        message.tenantId = object.tenantId ?? "";
        message.userId = object.userId ?? "";
        message.scopes = object.scopes?.map(e => e) || [];
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

function createBaseAuthData_DataEntry(): AuthData_DataEntry {
    return { key: "", value: "" };
}

export const AuthData_DataEntry: MessageFns<AuthData_DataEntry> = {
    encode(message: AuthData_DataEntry, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
        if (message.key !== "") {
            writer.uint32(10).string(message.key);
        }
        if (message.value !== "") {
            writer.uint32(18).string(message.value);
        }
        return writer;
    },

    decode(input: BinaryReader | Uint8Array, length?: number): AuthData_DataEntry {
        const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseAuthData_DataEntry();
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

    fromJSON(object: any): AuthData_DataEntry {
        return {
            key: isSet(object.key) ? globalThis.String(object.key) : "",
            value: isSet(object.value) ? globalThis.String(object.value) : "",
        };
    },

    toJSON(message: AuthData_DataEntry): unknown {
        const obj: any = {};
        if (message.key !== "") {
            obj.key = message.key;
        }
        if (message.value !== "") {
            obj.value = message.value;
        }
        return obj;
    },

    create(base?: DeepPartial<AuthData_DataEntry>): AuthData_DataEntry {
        return AuthData_DataEntry.fromPartial(base ?? {});
    },
    fromPartial(object: DeepPartial<AuthData_DataEntry>): AuthData_DataEntry {
        const message = createBaseAuthData_DataEntry();
        message.key = object.key ?? "";
        message.value = object.value ?? "";
        return message;
    },
};

function createBaseDataFilter(): DataFilter {
    return { fields: {}, and: [], or: [] };
}

export const DataFilter: MessageFns<DataFilter> = {
    encode(message: DataFilter, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
        Object.entries(message.fields).forEach(([key, value]) => {
            DataFilter_FieldsEntry.encode(
                { key: key as any, value },
                writer.uint32(10).fork()
            ).join();
        });
        for (const v of message.and) {
            DataFilter.encode(v!, writer.uint32(18).fork()).join();
        }
        for (const v of message.or) {
            DataFilter.encode(v!, writer.uint32(26).fork()).join();
        }
        return writer;
    },

    decode(input: BinaryReader | Uint8Array, length?: number): DataFilter {
        const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseDataFilter();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1: {
                    if (tag !== 10) {
                        break;
                    }

                    const entry1 = DataFilter_FieldsEntry.decode(reader, reader.uint32());
                    if (entry1.value !== undefined) {
                        message.fields[entry1.key] = entry1.value;
                    }
                    continue;
                }
                case 2: {
                    if (tag !== 18) {
                        break;
                    }

                    message.and.push(DataFilter.decode(reader, reader.uint32()));
                    continue;
                }
                case 3: {
                    if (tag !== 26) {
                        break;
                    }

                    message.or.push(DataFilter.decode(reader, reader.uint32()));
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

    fromJSON(object: any): DataFilter {
        return {
            fields: isObject(object.fields)
                ? Object.entries(object.fields).reduce<{ [key: string]: DataFieldFilter }>(
                      (acc, [key, value]) => {
                          acc[key] = DataFieldFilter.fromJSON(value);
                          return acc;
                      },
                      {}
                  )
                : {},
            and: globalThis.Array.isArray(object?.and)
                ? object.and.map((e: any) => DataFilter.fromJSON(e))
                : [],
            or: globalThis.Array.isArray(object?.or)
                ? object.or.map((e: any) => DataFilter.fromJSON(e))
                : [],
        };
    },

    toJSON(message: DataFilter): unknown {
        const obj: any = {};
        if (message.fields) {
            const entries = Object.entries(message.fields);
            if (entries.length > 0) {
                obj.fields = {};
                entries.forEach(([k, v]) => {
                    obj.fields[k] = DataFieldFilter.toJSON(v);
                });
            }
        }
        if (message.and?.length) {
            obj.and = message.and.map(e => DataFilter.toJSON(e));
        }
        if (message.or?.length) {
            obj.or = message.or.map(e => DataFilter.toJSON(e));
        }
        return obj;
    },

    create(base?: DeepPartial<DataFilter>): DataFilter {
        return DataFilter.fromPartial(base ?? {});
    },
    fromPartial(object: DeepPartial<DataFilter>): DataFilter {
        const message = createBaseDataFilter();
        message.fields = Object.entries(object.fields ?? {}).reduce<{
            [key: string]: DataFieldFilter;
        }>((acc, [key, value]) => {
            if (value !== undefined) {
                acc[key] = DataFieldFilter.fromPartial(value);
            }
            return acc;
        }, {});
        message.and = object.and?.map(e => DataFilter.fromPartial(e)) || [];
        message.or = object.or?.map(e => DataFilter.fromPartial(e)) || [];
        return message;
    },
};

function createBaseDataFilter_FieldsEntry(): DataFilter_FieldsEntry {
    return { key: "", value: undefined };
}

export const DataFilter_FieldsEntry: MessageFns<DataFilter_FieldsEntry> = {
    encode(
        message: DataFilter_FieldsEntry,
        writer: BinaryWriter = new BinaryWriter()
    ): BinaryWriter {
        if (message.key !== "") {
            writer.uint32(10).string(message.key);
        }
        if (message.value !== undefined) {
            DataFieldFilter.encode(message.value, writer.uint32(18).fork()).join();
        }
        return writer;
    },

    decode(input: BinaryReader | Uint8Array, length?: number): DataFilter_FieldsEntry {
        const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseDataFilter_FieldsEntry();
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

                    message.value = DataFieldFilter.decode(reader, reader.uint32());
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

    fromJSON(object: any): DataFilter_FieldsEntry {
        return {
            key: isSet(object.key) ? globalThis.String(object.key) : "",
            value: isSet(object.value) ? DataFieldFilter.fromJSON(object.value) : undefined,
        };
    },

    toJSON(message: DataFilter_FieldsEntry): unknown {
        const obj: any = {};
        if (message.key !== "") {
            obj.key = message.key;
        }
        if (message.value !== undefined) {
            obj.value = DataFieldFilter.toJSON(message.value);
        }
        return obj;
    },

    create(base?: DeepPartial<DataFilter_FieldsEntry>): DataFilter_FieldsEntry {
        return DataFilter_FieldsEntry.fromPartial(base ?? {});
    },
    fromPartial(object: DeepPartial<DataFilter_FieldsEntry>): DataFilter_FieldsEntry {
        const message = createBaseDataFilter_FieldsEntry();
        message.key = object.key ?? "";
        message.value =
            object.value !== undefined && object.value !== null
                ? DataFieldFilter.fromPartial(object.value)
                : undefined;
        return message;
    },
};

function createBaseDataFieldFilter(): DataFieldFilter {
    return { type: "", operation: "", value: "" };
}

export const DataFieldFilter: MessageFns<DataFieldFilter> = {
    encode(message: DataFieldFilter, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
        if (message.type !== "") {
            writer.uint32(10).string(message.type);
        }
        if (message.operation !== "") {
            writer.uint32(18).string(message.operation);
        }
        if (message.value !== "") {
            writer.uint32(26).string(message.value);
        }
        return writer;
    },

    decode(input: BinaryReader | Uint8Array, length?: number): DataFieldFilter {
        const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseDataFieldFilter();
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

                    message.operation = reader.string();
                    continue;
                }
                case 3: {
                    if (tag !== 26) {
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

    fromJSON(object: any): DataFieldFilter {
        return {
            type: isSet(object.type) ? globalThis.String(object.type) : "",
            operation: isSet(object.operation) ? globalThis.String(object.operation) : "",
            value: isSet(object.value) ? globalThis.String(object.value) : "",
        };
    },

    toJSON(message: DataFieldFilter): unknown {
        const obj: any = {};
        if (message.type !== "") {
            obj.type = message.type;
        }
        if (message.operation !== "") {
            obj.operation = message.operation;
        }
        if (message.value !== "") {
            obj.value = message.value;
        }
        return obj;
    },

    create(base?: DeepPartial<DataFieldFilter>): DataFieldFilter {
        return DataFieldFilter.fromPartial(base ?? {});
    },
    fromPartial(object: DeepPartial<DataFieldFilter>): DataFieldFilter {
        const message = createBaseDataFieldFilter();
        message.type = object.type ?? "";
        message.operation = object.operation ?? "";
        message.value = object.value ?? "";
        return message;
    },
};

function createBaseEventMetadata(): EventMetadata {
    return { causationId: "", correlationId: "", target: DeploymentTarget.DEPLOYMENT_TARGET_BLUE };
}

export const EventMetadata: MessageFns<EventMetadata> = {
    encode(message: EventMetadata, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
        if (message.causationId !== "") {
            writer.uint32(10).string(message.causationId);
        }
        if (message.correlationId !== "") {
            writer.uint32(18).string(message.correlationId);
        }
        if (message.target !== DeploymentTarget.DEPLOYMENT_TARGET_BLUE) {
            writer.uint32(24).int32(deploymentTargetToNumber(message.target));
        }
        return writer;
    },

    decode(input: BinaryReader | Uint8Array, length?: number): EventMetadata {
        const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseEventMetadata();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1: {
                    if (tag !== 10) {
                        break;
                    }

                    message.causationId = reader.string();
                    continue;
                }
                case 2: {
                    if (tag !== 18) {
                        break;
                    }

                    message.correlationId = reader.string();
                    continue;
                }
                case 3: {
                    if (tag !== 24) {
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

    fromJSON(object: any): EventMetadata {
        return {
            causationId: isSet(object.causationId) ? globalThis.String(object.causationId) : "",
            correlationId: isSet(object.correlationId)
                ? globalThis.String(object.correlationId)
                : "",
            target: isSet(object.target)
                ? deploymentTargetFromJSON(object.target)
                : DeploymentTarget.DEPLOYMENT_TARGET_BLUE,
        };
    },

    toJSON(message: EventMetadata): unknown {
        const obj: any = {};
        if (message.causationId !== "") {
            obj.causationId = message.causationId;
        }
        if (message.correlationId !== "") {
            obj.correlationId = message.correlationId;
        }
        if (message.target !== DeploymentTarget.DEPLOYMENT_TARGET_BLUE) {
            obj.target = deploymentTargetToJSON(message.target);
        }
        return obj;
    },

    create(base?: DeepPartial<EventMetadata>): EventMetadata {
        return EventMetadata.fromPartial(base ?? {});
    },
    fromPartial(object: DeepPartial<EventMetadata>): EventMetadata {
        const message = createBaseEventMetadata();
        message.causationId = object.causationId ?? "";
        message.correlationId = object.correlationId ?? "";
        message.target = object.target ?? DeploymentTarget.DEPLOYMENT_TARGET_BLUE;
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
