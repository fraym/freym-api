import { DeploymentTarget, ServiceClient } from "@fraym/proto/dist/index.freym.crud.delivery";
import { AuthData, getProtobufAuthData } from "./auth";
import { CrudData } from "./data";
import { Filter, getProtobufDataFilter } from "./filter";
import { GetSingleEntryOptions } from "./options";
import { Wait, getProtobufDataWait } from "./wait";

export const getCrudData = async <T extends CrudData>(
    type: string,
    authData: AuthData,
    id: string,
    filter: Filter,
    target: DeploymentTarget,
    serviceClient: ServiceClient,
    wait?: Wait,
    options?: GetSingleEntryOptions
): Promise<T | null> => {
    return new Promise<T | null>((resolve, reject) => {
        serviceClient.getData(
            {
                id,
                type,
                auth: getProtobufAuthData(authData),
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

                if (!result) {
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
