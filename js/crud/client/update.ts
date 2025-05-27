import { DeploymentTarget, ServiceClient } from "@fraym/proto/dist/index.freym.crud.delivery";
import { AuthData, getProtobufAuthData } from "./auth";
import { CrudData } from "./data";
import { EventMetadata, fillMetadataWithDefaults } from "./eventMetadata";

export type UpdateResponse<T extends CrudData> =
    | UpdateSuccessResponse<T>
    | UpdateValidationResponse;

export interface UpdateSuccessResponse<T extends CrudData> {
    data: T;
}

export interface UpdateValidationResponse {
    validationErrors: string[];
    fieldValidationErrors: Record<string, string>;
}

export const isUpdateSuccessResponse = <T extends CrudData>(
    response: UpdateResponse<T>
): response is UpdateSuccessResponse<T> => {
    return Object.prototype.hasOwnProperty.call(response, "data");
};

export const isUpdateValidationResponse = <T extends CrudData>(
    response: UpdateResponse<T>
): response is UpdateValidationResponse => {
    return !Object.prototype.hasOwnProperty.call(response, "data");
};

export const updateCrudData = async <T extends CrudData>(
    type: string,
    authData: AuthData,
    id: string,
    data: Partial<T>,
    eventMetadata: Partial<EventMetadata> | null,
    target: DeploymentTarget,
    serviceClient: ServiceClient
): Promise<UpdateResponse<T>> => {
    const requestData: Record<string, string> = {};

    for (const key in data) {
        requestData[key] = JSON.stringify(data[key]);
    }

    return new Promise<UpdateResponse<T>>((resolve, reject) => {
        serviceClient.update(
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
