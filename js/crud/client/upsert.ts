import { DeploymentTarget, ServiceClient } from "@fraym/proto/dist/index.freym.crud.delivery";
import { AuthData, getProtobufAuthData } from "./auth";
import { CrudData } from "./data";
import { EventMetadata, fillMetadataWithDefaults } from "./eventMetadata";

export type UpsertResponse<T extends CrudData> =
    | UpsertSuccessResponse<T>
    | UpsertValidationResponse;

export interface UpsertSuccessResponse<T extends CrudData> {
    data: T;
}

export interface UpsertValidationResponse {
    validationErrors: string[];
    fieldValidationErrors: Record<string, string>;
}

export const isUpsertSuccessResponse = <T extends CrudData>(
    response: UpsertResponse<T>
): response is UpsertSuccessResponse<T> => {
    return Object.prototype.hasOwnProperty.call(response, "data");
};

export const isCreateValidationResponse = <T extends CrudData>(
    response: UpsertResponse<T>
): response is UpsertValidationResponse => {
    return !Object.prototype.hasOwnProperty.call(response, "data");
};

export const upsertCrudData = async <T extends CrudData>(
    type: string,
    authData: AuthData,
    data: T,
    id: string,
    eventMetadata: Partial<EventMetadata> | null,
    target: DeploymentTarget,
    serviceClient: ServiceClient
): Promise<UpsertResponse<T>> => {
    const requestData: Record<string, string> = {};

    for (const key in data) {
        requestData[key] = JSON.stringify(data[key]);
    }

    return new Promise<UpsertResponse<T>>((resolve, reject) => {
        serviceClient.upsert(
            {
                type,
                auth: getProtobufAuthData(authData),
                data: requestData,
                id,
                eventMetadata: fillMetadataWithDefaults(eventMetadata, target),
            },
            (error, response) => {
                if (error) {
                    reject(error.message);
                    return;
                }

                if (
                    response.validationErrors.length > 0 ||
                    Object.keys(response.fieldValidationErrors).length > 0
                ) {
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
                });
            }
        );
    });
};
