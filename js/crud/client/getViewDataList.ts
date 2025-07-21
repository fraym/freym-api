import { DeploymentTarget, ServiceClient } from "@fraym/proto/dist/index.freym.crud.delivery";
import { AuthData, getProtobufAuthData } from "./auth";
import { CrudData } from "./data";
import { Filter, getProtobufDataFilter } from "./filter";
import { Order, getProtobufDataOrder } from "./order";
import { ListWait, getProtobufDataListWait } from "./wait";

export interface GetViewDataList<T extends CrudData> {
    limit: number;
    page: number;
    total: number;
    data: T[];
}

export const getViewDataList = async <T extends CrudData>(
    view: string,
    auth: AuthData,
    limit: number,
    page: number,
    filter: Filter,
    order: Order[],
    useStrongConsistency: boolean,
    target: DeploymentTarget,
    serviceClient: ServiceClient,
    wait?: ListWait
): Promise<GetViewDataList<T> | null> => {
    return new Promise<GetViewDataList<T> | null>((resolve, reject) => {
        serviceClient.getViewDataList(
            {
                view,
                auth: getProtobufAuthData(auth),
                limit: limit.toString(),
                page: page.toString(),
                filter: getProtobufDataFilter(filter),
                order: getProtobufDataOrder(order),
                useStrongConsistency,
                target,
                wait: getProtobufDataListWait(wait),
            },
            (error, response) => {
                if (error) {
                    reject(error.message);
                    return;
                }

                // eslint-disable-next-line @typescript-eslint/no-explicit-any
                const data: any[] = [];

                for (const result of response.result) {
                    // eslint-disable-next-line @typescript-eslint/no-explicit-any
                    const dataRecord: Record<string, any> = {};
                    const resultData = result.data;

                    for (const key in resultData) {
                        if (resultData[key] !== undefined) {
                            dataRecord[key] = JSON.parse(resultData[key]);
                        }
                    }

                    data.push(dataRecord);
                }

                resolve({
                    limit: parseInt(response.limit, 10),
                    page: parseInt(response.page, 10),
                    total: parseInt(response.total, 10),
                    data,
                });
            }
        );
    });
};
