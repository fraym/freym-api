import { ServiceClient } from "@fraym/proto/dist/index.freym.projections.delivery";
import { AuthData, getProtobufAuthData } from "./auth";
import { ProjectionData } from "./data";
import { Filter, getProtobufDataFilter } from "./filter";
import { Wait, getProtobufDataWait } from "./wait";

export const getProjectionData = async <T extends ProjectionData>(
    projection: string,
    auth: AuthData,
    dataId: string,
    filter: Filter,
    returnEmptyDataIfNotFound: boolean,
    useStrongConsistency: boolean,
    deploymentId: number | null,
    serviceClient: ServiceClient,
    wait?: Wait
): Promise<T | null> => {
    return new Promise<T | null>((resolve, reject) => {
        serviceClient.getData(
            {
                projection,
                auth: getProtobufAuthData(auth),
                dataId,
                filter: getProtobufDataFilter(filter),
                returnEmptyDataIfNotFound,
                wait: getProtobufDataWait(wait),
                useStrongConsistency,
                deploymentId: deploymentId?.toString() ?? "",
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
