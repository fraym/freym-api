import { ServiceClient } from "@fraym/proto/dist/index.freym.projections.delivery";
import { credentials } from "@grpc/grpc-js";
import { AuthData } from "./auth";
import { DeliveryClientConfig, useDeliveryConfigDefaults } from "./config";
import { ProjectionData } from "./data";
import { deleteProjectionData } from "./delete";
import { EventMetadata } from "./eventMetadata";
import { Filter } from "./filter";
import { getProjectionData } from "./getData";
import { GetProjectionDataList, getProjectionDataList } from "./getDataList";
import { getViewData as getDataFromView } from "./getViewData";
import { GetViewDataList, getViewDataList as getDataListFromView } from "./getViewDataList";
import { GetEntryOptions, GetSingleEntryOptions } from "./options";
import { Order } from "./order";
import { UpsertResponse, upsertProjectionData } from "./upsert";
import { ListWait, Wait } from "./wait";

export interface DeliveryClient {
    getData: <T extends ProjectionData>(
        projection: string,
        authData: AuthData,
        id: string,
        filter?: Filter,
        wait?: Wait,
        options?: GetSingleEntryOptions
    ) => Promise<T | null>;
    getViewData: <T extends ProjectionData>(
        view: string,
        authData: AuthData,
        filter?: Filter,
        wait?: Wait,
        options?: GetSingleEntryOptions
    ) => Promise<T | null>;
    getDataList: <T extends ProjectionData>(
        projection: string,
        authData: AuthData,
        limit?: number,
        page?: number,
        filter?: Filter,
        order?: Order[],
        wait?: ListWait,
        options?: GetEntryOptions
    ) => Promise<GetProjectionDataList<T> | null>;
    getViewDataList: <T extends ProjectionData>(
        view: string,
        authData: AuthData,
        limit?: number,
        page?: number,
        filter?: Filter,
        order?: Order[],
        wait?: ListWait,
        options?: GetEntryOptions
    ) => Promise<GetProjectionDataList<T> | null>;
    upsertData: <T extends ProjectionData>(
        projection: string,
        authData: AuthData,
        dataId: string,
        payload: Partial<T>,
        eventMetadata?: Partial<EventMetadata>
    ) => Promise<UpsertResponse<T>>;
    deleteDataById: (
        projection: string,
        authData: AuthData,
        dataId: string,
        eventMetadata?: Partial<EventMetadata>
    ) => Promise<number>;
    deleteDataByFilter: (
        projection: string,
        authData: AuthData,
        filter?: Filter,
        eventMetadata?: Partial<EventMetadata>
    ) => Promise<number>;
    close: () => Promise<void>;
}

export const newDeliveryClient = async (
    inputConfig?: DeliveryClientConfig
): Promise<DeliveryClient> => {
    const config = useDeliveryConfigDefaults(inputConfig);
    const serviceClient = new ServiceClient(config.serverAddress, credentials.createInsecure(), {
        "grpc.keepalive_time_ms": config.keepaliveInterval,
        "grpc.keepalive_timeout_ms": config.keepaliveTimeout,
        "grpc.keepalive_permit_without_calls": 1,
    });

    const getData = async <T extends ProjectionData>(
        projection: string,
        auth: AuthData,
        id: string,
        filter: Filter = { fields: {}, and: [], or: [] },
        wait?: Wait,
        options?: GetSingleEntryOptions
    ): Promise<T | null> => {
        return await getProjectionData<T>(
            projection,
            auth,
            id,
            filter,
            config.deploymentTarget,
            serviceClient,
            wait,
            options
        );
    };

    const getViewData = async <T extends ProjectionData>(
        view: string,
        auth: AuthData,
        filter: Filter = { fields: {}, and: [], or: [] },
        wait?: Wait,
        options?: GetSingleEntryOptions
    ): Promise<T | null> => {
        return await getDataFromView<T>(
            view,
            auth,
            filter,
            config.deploymentTarget,
            serviceClient,
            wait,
            options
        );
    };

    const getDataList = async <T extends ProjectionData>(
        projection: string,
        auth: AuthData,
        limit: number = 0,
        page: number = 1,
        filter: Filter = { fields: {}, and: [], or: [] },
        order: Order[] = [],
        wait?: ListWait,
        options?: GetEntryOptions
    ): Promise<GetProjectionDataList<T> | null> => {
        return await getProjectionDataList<T>(
            projection,
            auth,
            limit,
            page,
            filter,
            order,
            config.deploymentTarget,
            serviceClient,
            wait,
            options
        );
    };

    const getViewDataList = async <T extends ProjectionData>(
        view: string,
        auth: AuthData,
        limit: number = 0,
        page: number = 1,
        filter: Filter = { fields: {}, and: [], or: [] },
        order: Order[] = [],
        wait?: ListWait,
        options?: GetEntryOptions
    ): Promise<GetViewDataList<T> | null> => {
        return await getDataListFromView<T>(
            view,
            auth,
            limit,
            page,
            filter,
            order,
            config.deploymentTarget,
            serviceClient,
            wait,
            options
        );
    };

    const upsertData = async <T extends ProjectionData>(
        projection: string,
        authData: AuthData,
        dataId: string,
        payload: Partial<T>,
        eventMetadata: Partial<EventMetadata> | null = null
    ): Promise<UpsertResponse<T>> => {
        return upsertProjectionData<T>(
            projection,
            authData,
            dataId,
            payload,
            eventMetadata,
            config.deploymentTarget,
            serviceClient
        );
    };

    const deleteDataById = async (
        projection: string,
        authData: AuthData,
        dataId: string,
        eventMetadata: Partial<EventMetadata> | null = null
    ): Promise<number> => {
        return deleteProjectionData(
            projection,
            authData,
            dataId,
            { fields: {}, and: [], or: [] },
            eventMetadata,
            config.deploymentTarget,
            serviceClient
        );
    };

    const deleteDataByFilter = async (
        projection: string,
        authData: AuthData,
        filter: Filter = { fields: {}, and: [], or: [] },
        eventMetadata: Partial<EventMetadata> | null = null
    ): Promise<number> => {
        return deleteProjectionData(
            projection,
            authData,
            "",
            filter,
            eventMetadata,
            config.deploymentTarget,
            serviceClient
        );
    };

    const close = async () => {
        serviceClient.close();
    };

    return {
        getData,
        getViewData,
        getDataList,
        getViewDataList,
        upsertData,
        deleteDataById,
        deleteDataByFilter,
        close,
    };
};
