// Code generated by protoc-gen-ts_proto. DO NOT EDIT.
// versions:
//   protoc-gen-ts_proto  v2.7.0
//   protoc               v5.29.3
// source: deployments/management/rollback.proto

/* eslint-disable */
import { BinaryReader, BinaryWriter } from "@bufbuild/protobuf/wire";
import {
    DeploymentTarget,
    deploymentTargetFromJSON,
    deploymentTargetToJSON,
    deploymentTargetToNumber,
} from "./shared";

export interface RollbackDeploymentRequest {
    target: DeploymentTarget;
    namespace: string;
}

export interface RollbackDeploymentByIdRequest {
    deploymentId: string;
}

export interface RollbackDeploymentResponse {}

function createBaseRollbackDeploymentRequest(): RollbackDeploymentRequest {
    return { target: DeploymentTarget.DEPLOYMENT_TARGET_BLUE, namespace: "" };
}

export const RollbackDeploymentRequest: MessageFns<RollbackDeploymentRequest> = {
    encode(
        message: RollbackDeploymentRequest,
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

    decode(input: BinaryReader | Uint8Array, length?: number): RollbackDeploymentRequest {
        const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseRollbackDeploymentRequest();
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

    fromJSON(object: any): RollbackDeploymentRequest {
        return {
            target: isSet(object.target)
                ? deploymentTargetFromJSON(object.target)
                : DeploymentTarget.DEPLOYMENT_TARGET_BLUE,
            namespace: isSet(object.namespace) ? globalThis.String(object.namespace) : "",
        };
    },

    toJSON(message: RollbackDeploymentRequest): unknown {
        const obj: any = {};
        if (message.target !== DeploymentTarget.DEPLOYMENT_TARGET_BLUE) {
            obj.target = deploymentTargetToJSON(message.target);
        }
        if (message.namespace !== "") {
            obj.namespace = message.namespace;
        }
        return obj;
    },

    create(base?: DeepPartial<RollbackDeploymentRequest>): RollbackDeploymentRequest {
        return RollbackDeploymentRequest.fromPartial(base ?? {});
    },
    fromPartial(object: DeepPartial<RollbackDeploymentRequest>): RollbackDeploymentRequest {
        const message = createBaseRollbackDeploymentRequest();
        message.target = object.target ?? DeploymentTarget.DEPLOYMENT_TARGET_BLUE;
        message.namespace = object.namespace ?? "";
        return message;
    },
};

function createBaseRollbackDeploymentByIdRequest(): RollbackDeploymentByIdRequest {
    return { deploymentId: "0" };
}

export const RollbackDeploymentByIdRequest: MessageFns<RollbackDeploymentByIdRequest> = {
    encode(
        message: RollbackDeploymentByIdRequest,
        writer: BinaryWriter = new BinaryWriter()
    ): BinaryWriter {
        if (message.deploymentId !== "0") {
            writer.uint32(8).int64(message.deploymentId);
        }
        return writer;
    },

    decode(input: BinaryReader | Uint8Array, length?: number): RollbackDeploymentByIdRequest {
        const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseRollbackDeploymentByIdRequest();
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

    fromJSON(object: any): RollbackDeploymentByIdRequest {
        return {
            deploymentId: isSet(object.deploymentId) ? globalThis.String(object.deploymentId) : "0",
        };
    },

    toJSON(message: RollbackDeploymentByIdRequest): unknown {
        const obj: any = {};
        if (message.deploymentId !== "0") {
            obj.deploymentId = message.deploymentId;
        }
        return obj;
    },

    create(base?: DeepPartial<RollbackDeploymentByIdRequest>): RollbackDeploymentByIdRequest {
        return RollbackDeploymentByIdRequest.fromPartial(base ?? {});
    },
    fromPartial(object: DeepPartial<RollbackDeploymentByIdRequest>): RollbackDeploymentByIdRequest {
        const message = createBaseRollbackDeploymentByIdRequest();
        message.deploymentId = object.deploymentId ?? "0";
        return message;
    },
};

function createBaseRollbackDeploymentResponse(): RollbackDeploymentResponse {
    return {};
}

export const RollbackDeploymentResponse: MessageFns<RollbackDeploymentResponse> = {
    encode(_: RollbackDeploymentResponse, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
        return writer;
    },

    decode(input: BinaryReader | Uint8Array, length?: number): RollbackDeploymentResponse {
        const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseRollbackDeploymentResponse();
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

    fromJSON(_: any): RollbackDeploymentResponse {
        return {};
    },

    toJSON(_: RollbackDeploymentResponse): unknown {
        const obj: any = {};
        return obj;
    },

    create(base?: DeepPartial<RollbackDeploymentResponse>): RollbackDeploymentResponse {
        return RollbackDeploymentResponse.fromPartial(base ?? {});
    },
    fromPartial(_: DeepPartial<RollbackDeploymentResponse>): RollbackDeploymentResponse {
        const message = createBaseRollbackDeploymentResponse();
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
