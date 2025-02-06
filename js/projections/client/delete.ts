import { ServiceClient } from "@fraym/proto/dist/index.freym.projections.delivery";
import { AuthData, getProtobufAuthData } from "./auth";
import { EventMetadata, fillMetadataWithDefaults } from "./eventMetadata";
import { Filter, getProtobufDataFilter } from "./filter";

export const deleteProjectionData = async (
    projection: string,
    auth: AuthData,
    dataId: string,
    filter: Filter,
    eventMetadata: Partial<EventMetadata> | null = null,
    serviceClient: ServiceClient
): Promise<number> => {
    return new Promise<number>((resolve, reject) => {
        serviceClient.delete(
            {
                projection,
                auth: getProtobufAuthData(auth),
                dataId,
                filter: getProtobufDataFilter(filter),
                eventMetadata: fillMetadataWithDefaults(eventMetadata),
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
