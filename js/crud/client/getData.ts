import { ServiceClient } from "@fraym/proto/dist/index.freym.crud.delivery";
import { AuthData, getProtobufAuthData } from "./auth";
import { CrudData } from "./data";
import { Filter, getProtobufDataFilter } from "./filter";
import { Wait, getProtobufDataWait } from "./wait";

export const getCrudData = async <T extends CrudData>(
    type: string,
    authData: AuthData,
    id: string,
    filter: Filter,
    returnEmptyDataIfNotFound: boolean,
    useStrongConsistency: boolean,
    serviceClient: ServiceClient,
    wait?: Wait
): Promise<T | null> => {
    return new Promise<T | null>((resolve, reject) => {
        serviceClient.getData(
            {
                type,
                auth: getProtobufAuthData(authData),
                filter: getProtobufDataFilter(filter),
                id,
                returnEmptyDataIfNotFound,
                wait: getProtobufDataWait(wait),
                useStrongConsistency,
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
