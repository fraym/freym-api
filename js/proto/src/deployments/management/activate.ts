// Code generated by protoc-gen-ts_proto. DO NOT EDIT.
// versions:
//   protoc-gen-ts_proto  v2.6.1
//   protoc               v5.29.3
// source: deployments/management/activate.proto

/* eslint-disable */
import { BinaryReader, BinaryWriter } from "@bufbuild/protobuf/wire";

export interface ActivateDeploymentRequest {
    deploymentId: string;
}

export interface ActivateDeploymentResponse {}

function createBaseActivateDeploymentRequest(): ActivateDeploymentRequest {
    return { deploymentId: "0" };
}

export const ActivateDeploymentRequest: MessageFns<ActivateDeploymentRequest> = {
    encode(
        message: ActivateDeploymentRequest,
        writer: BinaryWriter = new BinaryWriter()
    ): BinaryWriter {
        if (message.deploymentId !== "0") {
            writer.uint32(8).int64(message.deploymentId);
        }
        return writer;
    },

    decode(input: BinaryReader | Uint8Array, length?: number): ActivateDeploymentRequest {
        const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseActivateDeploymentRequest();
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

    fromJSON(object: any): ActivateDeploymentRequest {
        return {
            deploymentId: isSet(object.deploymentId) ? globalThis.String(object.deploymentId) : "0",
        };
    },

    toJSON(message: ActivateDeploymentRequest): unknown {
        const obj: any = {};
        if (message.deploymentId !== "0") {
            obj.deploymentId = message.deploymentId;
        }
        return obj;
    },

    create(base?: DeepPartial<ActivateDeploymentRequest>): ActivateDeploymentRequest {
        return ActivateDeploymentRequest.fromPartial(base ?? {});
    },
    fromPartial(object: DeepPartial<ActivateDeploymentRequest>): ActivateDeploymentRequest {
        const message = createBaseActivateDeploymentRequest();
        message.deploymentId = object.deploymentId ?? "0";
        return message;
    },
};

function createBaseActivateDeploymentResponse(): ActivateDeploymentResponse {
    return {};
}

export const ActivateDeploymentResponse: MessageFns<ActivateDeploymentResponse> = {
    encode(_: ActivateDeploymentResponse, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
        return writer;
    },

    decode(input: BinaryReader | Uint8Array, length?: number): ActivateDeploymentResponse {
        const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseActivateDeploymentResponse();
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

    fromJSON(_: any): ActivateDeploymentResponse {
        return {};
    },

    toJSON(_: ActivateDeploymentResponse): unknown {
        const obj: any = {};
        return obj;
    },

    create(base?: DeepPartial<ActivateDeploymentResponse>): ActivateDeploymentResponse {
        return ActivateDeploymentResponse.fromPartial(base ?? {});
    },
    fromPartial(_: DeepPartial<ActivateDeploymentResponse>): ActivateDeploymentResponse {
        const message = createBaseActivateDeploymentResponse();
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
