import { ServiceClient } from "@fraym/proto/dist/index.freym.projections.delivery";
import { AuthData, getProtobufAuthData } from "./auth";
import { EventMetadata, fillMetadataWithDefaults } from "./eventMetadata";

export type UpsertResponse<T extends object> = UpsertSuccessResponse<T> | UpsertValidationResponse;

export interface UpsertSuccessResponse<T extends object> {
    data: T;
    id: string;
}

export interface UpsertValidationResponse {
    validationErrors: string[];
    fieldValidationErrors: Record<string, string>;
}

export const isUpsertSuccessResponse = <T extends object>(
    response: UpsertResponse<T>
): response is UpsertSuccessResponse<T> => {
    return Object.prototype.hasOwnProperty.call(response, "data");
};

export const isUpsertValidationResponse = <T extends object>(
    response: UpsertResponse<T>
): response is UpsertValidationResponse => {
    return !Object.prototype.hasOwnProperty.call(response, "data");
};

export const upsertProjectionData = async <T extends object>(
    projection: string,
    auth: AuthData,
    dataId: string,
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    payload: Record<string, any>,
    eventMetadata: Partial<EventMetadata> | null = null,
    serviceClient: ServiceClient
): Promise<UpsertResponse<T>> => {
    const usedPayload: Record<string, string> = {};

    for (const key in payload) {
        usedPayload[key] = JSON.stringify(payload[key]);
    }

    return new Promise<UpsertResponse<T>>((resolve, reject) => {
        serviceClient.upsert(
            {
                projection,
                auth: getProtobufAuthData(auth),
                dataId,
                payload: usedPayload,
                eventMetadata: fillMetadataWithDefaults(eventMetadata),
            },
            (error, response) => {
                if (error) {
                    reject(error.message);
                    return;
                }

                if (response.validationErrors || response.fieldValidationErrors) {
                    resolve({
                        validationErrors: response.validationErrors,
                        fieldValidationErrors: response.fieldValidationErrors,
                    });
                    return;
                }

                // eslint-disable-next-line @typescript-eslint/no-explicit-any
                const data: any = {};

                for (const key in response.newData) {
                    if (response.newData[key] !== undefined) {
                        data[key] = JSON.parse(response.newData[key]);
                    }
                }

                resolve({
                    data,
                    id: response.id,
                });
            }
        );
    });
};
