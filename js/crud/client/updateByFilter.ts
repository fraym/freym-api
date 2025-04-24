import { ServiceClient } from "@fraym/proto/dist/index.freym.crud.delivery";
import { AuthData, getProtobufAuthData } from "./auth";
import { CrudData } from "./data";
import { EventMetadata } from "./eventMetadata";
import { Filter, getProtobufDataFilter } from "./filter";

export type UpdateByFilterResponse =
    | UpdateByFilterSuccessResponse
    | UpdateByFilterValidationResponse;

export interface UpdateByFilterSuccessResponse {
    numberOfUpdatedEntries: number;
}

export interface UpdateByFilterValidationResponse {
    validationErrors: Record<string, UpdateByFilterDataValidationResponse>;
}

export interface UpdateByFilterDataValidationResponse {
    validationErrors: string[];
    fieldValidationErrors: Record<string, string>;
}

export const isUpdateByFilterSuccessResponse = (
    response: UpdateByFilterResponse
): response is UpdateByFilterSuccessResponse => {
    return Object.prototype.hasOwnProperty.call(response, "numberOfUpdatedEntries");
};

export const isUpdateByFilterValidationResponse = (
    response: UpdateByFilterResponse
): response is UpdateByFilterValidationResponse => {
    return !Object.prototype.hasOwnProperty.call(response, "numberOfUpdatedEntries");
};

export const updateCrudDataByFilter = async <T extends CrudData>(
    type: string,
    authData: AuthData,
    filter: Filter,
    data: Partial<T>,
    eventMetadata: EventMetadata,
    serviceClient: ServiceClient
): Promise<UpdateByFilterResponse> => {
    const requestData: Record<string, string> = {};

    for (const key in data) {
        requestData[key] = JSON.stringify(data[key]);
    }

    return new Promise<UpdateByFilterResponse>((resolve, reject) => {
        serviceClient.updateByFilter(
            {
                type,
                auth: getProtobufAuthData(authData),
                data: requestData,
                filter: getProtobufDataFilter(filter),
                eventMetadata,
            },
            (error, response) => {
                if (error) {
                    reject(error.message);
                    return;
                }

                const numberOfUpdatedEntries = parseInt(response.numberOfUpdatedEntries, 10) || 0;

                if (Object.keys(response.validationErrors).length > 0) {
                    const validationErrors: Record<string, UpdateByFilterDataValidationResponse> =
                        {};

                    for (const key in response.validationErrors) {
                        const validationError = response.validationErrors[key];

                        if (!validationError) {
                            continue;
                        }

                        validationErrors[key] = {
                            validationErrors: validationError.validationErrors,
                            fieldValidationErrors: validationError?.fieldValidationErrors,
                        };
                    }

                    if (Object.keys(validationErrors).length === 0) {
                        resolve({ numberOfUpdatedEntries });
                        return;
                    }

                    resolve({ validationErrors });
                    return;
                }

                resolve({ numberOfUpdatedEntries });
            }
        );
    });
};
