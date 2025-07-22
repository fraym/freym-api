import { DataListWait, DataWait } from "@fraym/proto/dist/index.freym.projections.delivery";
import { Filter, getProtobufDataFilter } from "./filter";

export interface Wait {
    timeout?: number;
    conditionFilter: Filter;
}

export const getProtobufDataWait = (wait?: Wait): DataWait | undefined => {
    if (!wait) {
        return undefined;
    }

    return {
        conditionFilter: getProtobufDataFilter(wait.conditionFilter),
        timeout: (wait.timeout ?? 0).toString(),
    };
};

export interface ListWait {
    timeout?: number;
    condition: DataListWaitCondition[];
}

export interface DataListWaitCondition {
    operation: string;
    value: number;
}

export const getProtobufDataListWait = (wait?: ListWait): DataListWait | undefined => {
    if (!wait) {
        return undefined;
    }

    return {
        condition: wait.condition.map(condition => ({
            operation: condition.operation,
            value: condition.value.toString(),
        })),
        timeout: (wait.timeout ?? 0).toString(),
    };
};
