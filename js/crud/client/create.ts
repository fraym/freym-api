import { ServiceClient } from "@fraym/proto/dist/index.freym.crud.delivery";
import { AuthData, getProtobufAuthData } from "./auth";
import { CrudData } from "./data";
import { EventMetadata } from "./eventMetadata";

export type CreateResponse<T extends CrudData> =
    | CreateSuccessResponse<T>
    | CreateValidationResponse;

export interface CreateSuccessResponse<T extends CrudData> {
    data: T;
}

export interface CreateValidationResponse {
    validationErrors: string[];
    fieldValidationErrors: Record<string, string>;
}

export const isCreateSuccessResponse = <T extends CrudData>(
    response: CreateResponse<T>
): response is CreateSuccessResponse<T> => {
    return Object.prototype.hasOwnProperty.call(response, "data");
};

export const isCreateValidationResponse = <T extends CrudData>(
    response: CreateResponse<T>
): response is CreateValidationResponse => {
    return !Object.prototype.hasOwnProperty.call(response, "data");
};

export const createCrudData = async <T extends CrudData>(
    type: string,
    authData: AuthData,
    data: T,
    id: string,
    eventMetadata: EventMetadata,
    serviceClient: ServiceClient
): Promise<CreateResponse<T>> => {
    const requestData: Record<string, string> = {};

    for (const key in data) {
        requestData[key] = JSON.stringify(data[key]);
    }

    return new Promise<CreateResponse<T>>((resolve, reject) => {
        serviceClient.create(
            {
                type,
                auth: getProtobufAuthData(authData),
                data: requestData,
                id,
                eventMetadata,
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
