import {
    DeploymentTarget,
    ServiceClient,
} from "@fraym/proto/dist/index.freym.projections.delivery";
import { AuthData, getProtobufAuthData } from "./auth";
import { ProjectionData } from "./data";
import { Filter, getProtobufDataFilter } from "./filter";
import { GetSingleEntryOptions } from "./options";
import { Wait, getProtobufDataWait } from "./wait";

export const getProjectionData = async <T extends ProjectionData>(
    projection: string,
    auth: AuthData,
    dataId: string,
    filter: Filter,
    target: DeploymentTarget,
    serviceClient: ServiceClient,
    wait?: Wait,
    options?: GetSingleEntryOptions
): Promise<T | null> => {
    return new Promise<T | null>((resolve, reject) => {
        serviceClient.getData(
            {
                dataId,
                projection,
                auth: getProtobufAuthData(auth),
                filter: getProtobufDataFilter(filter),
                wait: getProtobufDataWait(wait),
                target,
                returnEmptyDataIfNotFound: options?.returnEmptyDataIfNotFound ?? false,
                useStrongConsistency: options?.useStrongConsistency ?? false,
                useStrongConsistencyById: options?.useStrongConsistencyById ?? "",
            },
            (error, response) => {
                if (error) {
                    reject(error.message);
                    return;
                }

                const result = response.result;

                if (!result || !result.data || Object.keys(result.data).length === 0) {
                    resolve(null);
                    return;
                }

                // eslint-disable-next-line @typescript-eslint/no-explicit-any
                const data: any = {};

                for (const key in result.data) {
                    if (result.data[key] !== undefined) {
                        data[key] = JSON.parse(result.data[key]);
                    }
                }

                resolve(data);
            }
        );
    });
};
