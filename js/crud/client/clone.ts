import { DeploymentTarget, ServiceClient } from "@fraym/proto/dist/index.freym.crud.delivery";
import { AuthData, getProtobufAuthData } from "./auth";
import { CrudData } from "./data";
import { EventMetadata, fillMetadataWithDefaults } from "./eventMetadata";

export type CloneResponse<T extends CrudData> = CloneSuccessResponse<T> | CloneValidationResponse;

export interface CloneSuccessResponse<T extends CrudData> {
    data: T;
}

export interface CloneValidationResponse {
    validationErrors: string[];
    fieldValidationErrors: Record<string, string>;
}

export const isCloneSuccessResponse = <T extends CrudData>(
    response: CloneResponse<T>
): response is CloneSuccessResponse<T> => {
    return Object.prototype.hasOwnProperty.call(response, "data");
};

export const isCloneValidationResponse = <T extends CrudData>(
    response: CloneResponse<T>
): response is CloneValidationResponse => {
    return !Object.prototype.hasOwnProperty.call(response, "data");
};

export const cloneCrudData = async <T extends CrudData>(
    type: string,
    authData: AuthData,
    id: string,
    newId: string,
    data: Partial<T>,
    eventMetadata: Partial<EventMetadata> | null,
    target: DeploymentTarget,
    serviceClient: ServiceClient
): Promise<CloneResponse<T>> => {
    const requestData: Record<string, string> = {};

    for (const key in data) {
        requestData[key] = JSON.stringify(data[key]);
    }

    return new Promise<CloneResponse<T>>((resolve, reject) => {
        serviceClient.clone(
            {
                type,
                auth: getProtobufAuthData(authData),
                id,
                newId,
                eventMetadata: fillMetadataWithDefaults(eventMetadata, target),
                data: requestData,
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

                resolve({ data });
            }
        );
    });
};
