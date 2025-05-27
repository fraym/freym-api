import { DeploymentTarget, ServiceClient } from "@fraym/proto/dist/index.freym.crud.delivery";
import { AuthData, getProtobufAuthData } from "./auth";
import { EventMetadata, fillMetadataWithDefaults } from "./eventMetadata";
import { Filter, getProtobufDataFilter } from "./filter";

export const deleteCrudData = async (
    type: string,
    authData: AuthData,
    id: string,
    filter: Filter,
    eventMetadata: Partial<EventMetadata> | null,
    target: DeploymentTarget,
    serviceClient: ServiceClient
): Promise<number> => {
    return new Promise<number>((resolve, reject) => {
        serviceClient.delete(
            {
                type,
                auth: getProtobufAuthData(authData),
                id,
                filter: getProtobufDataFilter(filter),
                eventMetadata: fillMetadataWithDefaults(eventMetadata, target),
            },
            (error, response) => {
                if (error) {
                    reject(error.message);
                    return;
                }

                resolve(parseInt(response.numberOfDeletedEntries, 10));
            }
        );
    });
};
