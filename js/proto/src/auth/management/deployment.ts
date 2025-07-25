// Code generated by protoc-gen-ts_proto. DO NOT EDIT.
// versions:
//   protoc-gen-ts_proto  v2.7.5
//   protoc               v5.29.3
// source: auth/management/deployment.proto
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

export interface DeploySchemaRequest {
    deploymentId: string;
    namespace: string;
    permissions: string[];
    options: DeploymentOptions | undefined;
}

export interface DeploySchemaResponse {}

export interface ActivateSchemaRequest {
    deploymentId: string;
}

export interface ActivateSchemaResponse {}

export interface ConfirmSchemaRequest {
    deploymentId: string;
}

export interface ConfirmSchemaResponse {}

export interface RollbackSchemaRequest {
    target: DeploymentTarget;
    namespace: string;
}

export interface RollbackSchemaByDeploymentRequest {
    deploymentId: string;
}

export interface RollbackSchemaResponse {}

export interface GetSchemaDeploymentRequest {
    deploymentId: string;
}

export interface GetSchemaDeploymentResponse {
    progress: number;
}

export interface DeploymentOptions {
    target: DeploymentTarget;
    force: boolean;
}

function createBaseDeploySchemaRequest(): DeploySchemaRequest {
    return { deploymentId: "0", namespace: "", permissions: [], options: undefined };
}

export const DeploySchemaRequest: MessageFns<DeploySchemaRequest> = {
    encode(message: DeploySchemaRequest, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
        if (message.deploymentId !== "0") {
            writer.uint32(8).int64(message.deploymentId);
        }
        if (message.namespace !== "") {
            writer.uint32(18).string(message.namespace);
        }
        for (const v of message.permissions) {
            writer.uint32(26).string(v!);
        }
        if (message.options !== undefined) {
            DeploymentOptions.encode(message.options, writer.uint32(34).fork()).join();
        }
        return writer;
    },

    decode(input: BinaryReader | Uint8Array, length?: number): DeploySchemaRequest {
        const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
        const end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseDeploySchemaRequest();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1: {
                    if (tag !== 8) {
                        break;
                    }

                    message.deploymentId = reader.int64().toString();
                    continue;
                }
                case 2: {
                    if (tag !== 18) {
                        break;
                    }

                    message.namespace = reader.string();
                    continue;
                }
                case 3: {
                    if (tag !== 26) {
                        break;
                    }

                    message.permissions.push(reader.string());
                    continue;
                }
                case 4: {
                    if (tag !== 34) {
                        break;
                    }

                    message.options = DeploymentOptions.decode(reader, reader.uint32());
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

    fromJSON(object: any): DeploySchemaRequest {
        return {
            deploymentId: isSet(object.deploymentId) ? globalThis.String(object.deploymentId) : "0",
            namespace: isSet(object.namespace) ? globalThis.String(object.namespace) : "",
            permissions: globalThis.Array.isArray(object?.permissions)
                ? object.permissions.map((e: any) => globalThis.String(e))
                : [],
            options: isSet(object.options) ? DeploymentOptions.fromJSON(object.options) : undefined,
        };
    },

    toJSON(message: DeploySchemaRequest): unknown {
        const obj: any = {};
        if (message.deploymentId !== "0") {
            obj.deploymentId = message.deploymentId;
        }
        if (message.namespace !== "") {
            obj.namespace = message.namespace;
        }
        if (message.permissions?.length) {
            obj.permissions = message.permissions;
        }
        if (message.options !== undefined) {
            obj.options = DeploymentOptions.toJSON(message.options);
        }
        return obj;
    },

    create(base?: DeepPartial<DeploySchemaRequest>): DeploySchemaRequest {
        return DeploySchemaRequest.fromPartial(base ?? {});
    },
    fromPartial(object: DeepPartial<DeploySchemaRequest>): DeploySchemaRequest {
        const message = createBaseDeploySchemaRequest();
        message.deploymentId = object.deploymentId ?? "0";
        message.namespace = object.namespace ?? "";
        message.permissions = object.permissions?.map(e => e) || [];
        message.options =
            object.options !== undefined && object.options !== null
                ? DeploymentOptions.fromPartial(object.options)
                : undefined;
        return message;
    },
};

function createBaseDeploySchemaResponse(): DeploySchemaResponse {
    return {};
}

export const DeploySchemaResponse: MessageFns<DeploySchemaResponse> = {
    encode(_: DeploySchemaResponse, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
        return writer;
    },

    decode(input: BinaryReader | Uint8Array, length?: number): DeploySchemaResponse {
        const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
        const end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseDeploySchemaResponse();
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

    fromJSON(_: any): DeploySchemaResponse {
        return {};
    },

    toJSON(_: DeploySchemaResponse): unknown {
        const obj: any = {};
        return obj;
    },

    create(base?: DeepPartial<DeploySchemaResponse>): DeploySchemaResponse {
        return DeploySchemaResponse.fromPartial(base ?? {});
    },
    fromPartial(_: DeepPartial<DeploySchemaResponse>): DeploySchemaResponse {
        const message = createBaseDeploySchemaResponse();
        return message;
    },
};

function createBaseActivateSchemaRequest(): ActivateSchemaRequest {
    return { deploymentId: "0" };
}

export const ActivateSchemaRequest: MessageFns<ActivateSchemaRequest> = {
    encode(
        message: ActivateSchemaRequest,
        writer: BinaryWriter = new BinaryWriter()
    ): BinaryWriter {
        if (message.deploymentId !== "0") {
            writer.uint32(8).int64(message.deploymentId);
        }
        return writer;
    },

    decode(input: BinaryReader | Uint8Array, length?: number): ActivateSchemaRequest {
        const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
        const end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseActivateSchemaRequest();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1: {
                    if (tag !== 8) {
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

    fromJSON(object: any): ActivateSchemaRequest {
        return {
            deploymentId: isSet(object.deploymentId) ? globalThis.String(object.deploymentId) : "0",
        };
    },

    toJSON(message: ActivateSchemaRequest): unknown {
        const obj: any = {};
        if (message.deploymentId !== "0") {
            obj.deploymentId = message.deploymentId;
        }
        return obj;
    },

    create(base?: DeepPartial<ActivateSchemaRequest>): ActivateSchemaRequest {
        return ActivateSchemaRequest.fromPartial(base ?? {});
    },
    fromPartial(object: DeepPartial<ActivateSchemaRequest>): ActivateSchemaRequest {
        const message = createBaseActivateSchemaRequest();
        message.deploymentId = object.deploymentId ?? "0";
        return message;
    },
};

function createBaseActivateSchemaResponse(): ActivateSchemaResponse {
    return {};
}

export const ActivateSchemaResponse: MessageFns<ActivateSchemaResponse> = {
    encode(_: ActivateSchemaResponse, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
        return writer;
    },

    decode(input: BinaryReader | Uint8Array, length?: number): ActivateSchemaResponse {
        const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
        const end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseActivateSchemaResponse();
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

    fromJSON(_: any): ActivateSchemaResponse {
        return {};
    },

    toJSON(_: ActivateSchemaResponse): unknown {
        const obj: any = {};
        return obj;
    },

    create(base?: DeepPartial<ActivateSchemaResponse>): ActivateSchemaResponse {
        return ActivateSchemaResponse.fromPartial(base ?? {});
    },
    fromPartial(_: DeepPartial<ActivateSchemaResponse>): ActivateSchemaResponse {
        const message = createBaseActivateSchemaResponse();
        return message;
    },
};

function createBaseConfirmSchemaRequest(): ConfirmSchemaRequest {
    return { deploymentId: "0" };
}

export const ConfirmSchemaRequest: MessageFns<ConfirmSchemaRequest> = {
    encode(message: ConfirmSchemaRequest, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
        if (message.deploymentId !== "0") {
            writer.uint32(8).int64(message.deploymentId);
        }
        return writer;
    },

    decode(input: BinaryReader | Uint8Array, length?: number): ConfirmSchemaRequest {
        const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
        const end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseConfirmSchemaRequest();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1: {
                    if (tag !== 8) {
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

    fromJSON(object: any): ConfirmSchemaRequest {
        return {
            deploymentId: isSet(object.deploymentId) ? globalThis.String(object.deploymentId) : "0",
        };
    },

    toJSON(message: ConfirmSchemaRequest): unknown {
        const obj: any = {};
        if (message.deploymentId !== "0") {
            obj.deploymentId = message.deploymentId;
        }
        return obj;
    },

    create(base?: DeepPartial<ConfirmSchemaRequest>): ConfirmSchemaRequest {
        return ConfirmSchemaRequest.fromPartial(base ?? {});
    },
    fromPartial(object: DeepPartial<ConfirmSchemaRequest>): ConfirmSchemaRequest {
        const message = createBaseConfirmSchemaRequest();
        message.deploymentId = object.deploymentId ?? "0";
        return message;
    },
};

function createBaseConfirmSchemaResponse(): ConfirmSchemaResponse {
    return {};
}

export const ConfirmSchemaResponse: MessageFns<ConfirmSchemaResponse> = {
    encode(_: ConfirmSchemaResponse, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
        return writer;
    },

    decode(input: BinaryReader | Uint8Array, length?: number): ConfirmSchemaResponse {
        const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
        const end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseConfirmSchemaResponse();
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

    fromJSON(_: any): ConfirmSchemaResponse {
        return {};
    },

    toJSON(_: ConfirmSchemaResponse): unknown {
        const obj: any = {};
        return obj;
    },

    create(base?: DeepPartial<ConfirmSchemaResponse>): ConfirmSchemaResponse {
        return ConfirmSchemaResponse.fromPartial(base ?? {});
    },
    fromPartial(_: DeepPartial<ConfirmSchemaResponse>): ConfirmSchemaResponse {
        const message = createBaseConfirmSchemaResponse();
        return message;
    },
};

function createBaseRollbackSchemaRequest(): RollbackSchemaRequest {
    return { target: DeploymentTarget.DEPLOYMENT_TARGET_BLUE, namespace: "" };
}

export const RollbackSchemaRequest: MessageFns<RollbackSchemaRequest> = {
    encode(
        message: RollbackSchemaRequest,
        writer: BinaryWriter = new BinaryWriter()
    ): BinaryWriter {
        if (message.target !== DeploymentTarget.DEPLOYMENT_TARGET_BLUE) {
            writer.uint32(8).int32(deploymentTargetToNumber(message.target));
        }
        if (message.namespace !== "") {
            writer.uint32(18).string(message.namespace);
        }
        return writer;
    },

    decode(input: BinaryReader | Uint8Array, length?: number): RollbackSchemaRequest {
        const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
        const end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseRollbackSchemaRequest();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1: {
                    if (tag !== 8) {
                        break;
                    }

                    message.target = deploymentTargetFromJSON(reader.int32());
                    continue;
                }
                case 2: {
                    if (tag !== 18) {
                        break;
                    }

                    message.namespace = reader.string();
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

    fromJSON(object: any): RollbackSchemaRequest {
        return {
            target: isSet(object.target)
                ? deploymentTargetFromJSON(object.target)
                : DeploymentTarget.DEPLOYMENT_TARGET_BLUE,
            namespace: isSet(object.namespace) ? globalThis.String(object.namespace) : "",
        };
    },

    toJSON(message: RollbackSchemaRequest): unknown {
        const obj: any = {};
        if (message.target !== DeploymentTarget.DEPLOYMENT_TARGET_BLUE) {
            obj.target = deploymentTargetToJSON(message.target);
        }
        if (message.namespace !== "") {
            obj.namespace = message.namespace;
        }
        return obj;
    },

    create(base?: DeepPartial<RollbackSchemaRequest>): RollbackSchemaRequest {
        return RollbackSchemaRequest.fromPartial(base ?? {});
    },
    fromPartial(object: DeepPartial<RollbackSchemaRequest>): RollbackSchemaRequest {
        const message = createBaseRollbackSchemaRequest();
        message.target = object.target ?? DeploymentTarget.DEPLOYMENT_TARGET_BLUE;
        message.namespace = object.namespace ?? "";
        return message;
    },
};

function createBaseRollbackSchemaByDeploymentRequest(): RollbackSchemaByDeploymentRequest {
    return { deploymentId: "0" };
}

export const RollbackSchemaByDeploymentRequest: MessageFns<RollbackSchemaByDeploymentRequest> = {
    encode(
        message: RollbackSchemaByDeploymentRequest,
        writer: BinaryWriter = new BinaryWriter()
    ): BinaryWriter {
        if (message.deploymentId !== "0") {
            writer.uint32(8).int64(message.deploymentId);
        }
        return writer;
    },

    decode(input: BinaryReader | Uint8Array, length?: number): RollbackSchemaByDeploymentRequest {
        const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
        const end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseRollbackSchemaByDeploymentRequest();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1: {
                    if (tag !== 8) {
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

    fromJSON(object: any): RollbackSchemaByDeploymentRequest {
        return {
            deploymentId: isSet(object.deploymentId) ? globalThis.String(object.deploymentId) : "0",
        };
    },

    toJSON(message: RollbackSchemaByDeploymentRequest): unknown {
        const obj: any = {};
        if (message.deploymentId !== "0") {
            obj.deploymentId = message.deploymentId;
        }
        return obj;
    },

    create(
        base?: DeepPartial<RollbackSchemaByDeploymentRequest>
    ): RollbackSchemaByDeploymentRequest {
        return RollbackSchemaByDeploymentRequest.fromPartial(base ?? {});
    },
    fromPartial(
        object: DeepPartial<RollbackSchemaByDeploymentRequest>
    ): RollbackSchemaByDeploymentRequest {
        const message = createBaseRollbackSchemaByDeploymentRequest();
        message.deploymentId = object.deploymentId ?? "0";
        return message;
    },
};

function createBaseRollbackSchemaResponse(): RollbackSchemaResponse {
    return {};
}

export const RollbackSchemaResponse: MessageFns<RollbackSchemaResponse> = {
    encode(_: RollbackSchemaResponse, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
        return writer;
    },

    decode(input: BinaryReader | Uint8Array, length?: number): RollbackSchemaResponse {
        const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
        const end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseRollbackSchemaResponse();
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

    fromJSON(_: any): RollbackSchemaResponse {
        return {};
    },

    toJSON(_: RollbackSchemaResponse): unknown {
        const obj: any = {};
        return obj;
    },

    create(base?: DeepPartial<RollbackSchemaResponse>): RollbackSchemaResponse {
        return RollbackSchemaResponse.fromPartial(base ?? {});
    },
    fromPartial(_: DeepPartial<RollbackSchemaResponse>): RollbackSchemaResponse {
        const message = createBaseRollbackSchemaResponse();
        return message;
    },
};

function createBaseGetSchemaDeploymentRequest(): GetSchemaDeploymentRequest {
    return { deploymentId: "0" };
}

export const GetSchemaDeploymentRequest: MessageFns<GetSchemaDeploymentRequest> = {
    encode(
        message: GetSchemaDeploymentRequest,
        writer: BinaryWriter = new BinaryWriter()
    ): BinaryWriter {
        if (message.deploymentId !== "0") {
            writer.uint32(8).int64(message.deploymentId);
        }
        return writer;
    },

    decode(input: BinaryReader | Uint8Array, length?: number): GetSchemaDeploymentRequest {
        const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
        const end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseGetSchemaDeploymentRequest();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1: {
                    if (tag !== 8) {
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

    fromJSON(object: any): GetSchemaDeploymentRequest {
        return {
            deploymentId: isSet(object.deploymentId) ? globalThis.String(object.deploymentId) : "0",
        };
    },

    toJSON(message: GetSchemaDeploymentRequest): unknown {
        const obj: any = {};
        if (message.deploymentId !== "0") {
            obj.deploymentId = message.deploymentId;
        }
        return obj;
    },

    create(base?: DeepPartial<GetSchemaDeploymentRequest>): GetSchemaDeploymentRequest {
        return GetSchemaDeploymentRequest.fromPartial(base ?? {});
    },
    fromPartial(object: DeepPartial<GetSchemaDeploymentRequest>): GetSchemaDeploymentRequest {
        const message = createBaseGetSchemaDeploymentRequest();
        message.deploymentId = object.deploymentId ?? "0";
        return message;
    },
};

function createBaseGetSchemaDeploymentResponse(): GetSchemaDeploymentResponse {
    return { progress: 0 };
}

export const GetSchemaDeploymentResponse: MessageFns<GetSchemaDeploymentResponse> = {
    encode(
        message: GetSchemaDeploymentResponse,
        writer: BinaryWriter = new BinaryWriter()
    ): BinaryWriter {
        if (message.progress !== 0) {
            writer.uint32(8).uint32(message.progress);
        }
        return writer;
    },

    decode(input: BinaryReader | Uint8Array, length?: number): GetSchemaDeploymentResponse {
        const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
        const end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseGetSchemaDeploymentResponse();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1: {
                    if (tag !== 8) {
                        break;
                    }

                    message.progress = reader.uint32();
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

    fromJSON(object: any): GetSchemaDeploymentResponse {
        return { progress: isSet(object.progress) ? globalThis.Number(object.progress) : 0 };
    },

    toJSON(message: GetSchemaDeploymentResponse): unknown {
        const obj: any = {};
        if (message.progress !== 0) {
            obj.progress = Math.round(message.progress);
        }
        return obj;
    },

    create(base?: DeepPartial<GetSchemaDeploymentResponse>): GetSchemaDeploymentResponse {
        return GetSchemaDeploymentResponse.fromPartial(base ?? {});
    },
    fromPartial(object: DeepPartial<GetSchemaDeploymentResponse>): GetSchemaDeploymentResponse {
        const message = createBaseGetSchemaDeploymentResponse();
        message.progress = object.progress ?? 0;
        return message;
    },
};

function createBaseDeploymentOptions(): DeploymentOptions {
    return { target: DeploymentTarget.DEPLOYMENT_TARGET_BLUE, force: false };
}

export const DeploymentOptions: MessageFns<DeploymentOptions> = {
    encode(message: DeploymentOptions, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
        if (message.target !== DeploymentTarget.DEPLOYMENT_TARGET_BLUE) {
            writer.uint32(8).int32(deploymentTargetToNumber(message.target));
        }
        if (message.force !== false) {
            writer.uint32(16).bool(message.force);
        }
        return writer;
    },

    decode(input: BinaryReader | Uint8Array, length?: number): DeploymentOptions {
        const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
        const end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseDeploymentOptions();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1: {
                    if (tag !== 8) {
                        break;
                    }

                    message.target = deploymentTargetFromJSON(reader.int32());
                    continue;
                }
                case 2: {
                    if (tag !== 16) {
                        break;
                    }

                    message.force = reader.bool();
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

    fromJSON(object: any): DeploymentOptions {
        return {
            target: isSet(object.target)
                ? deploymentTargetFromJSON(object.target)
                : DeploymentTarget.DEPLOYMENT_TARGET_BLUE,
            force: isSet(object.force) ? globalThis.Boolean(object.force) : false,
        };
    },

    toJSON(message: DeploymentOptions): unknown {
        const obj: any = {};
        if (message.target !== DeploymentTarget.DEPLOYMENT_TARGET_BLUE) {
            obj.target = deploymentTargetToJSON(message.target);
        }
        if (message.force !== false) {
            obj.force = message.force;
        }
        return obj;
    },

    create(base?: DeepPartial<DeploymentOptions>): DeploymentOptions {
        return DeploymentOptions.fromPartial(base ?? {});
    },
    fromPartial(object: DeepPartial<DeploymentOptions>): DeploymentOptions {
        const message = createBaseDeploymentOptions();
        message.target = object.target ?? DeploymentTarget.DEPLOYMENT_TARGET_BLUE;
        message.force = object.force ?? false;
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
