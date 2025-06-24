import { DeploymentTarget, ServiceClient } from "@fraym/proto/dist/index.freym.crud.delivery";
import { credentials } from "@grpc/grpc-js";
import { AuthData } from "./auth";
import { CloneResponse, cloneCrudData } from "./clone";
import { DeliveryClientConfig, useDeliveryConfigDefaults } from "./config";
import { CreateResponse, createCrudData } from "./create";
import { CrudData } from "./data";
import { deleteCrudData } from "./delete";
import { EventMetadata } from "./eventMetadata";
import { Filter } from "./filter";
import { getCrudData } from "./getData";
import { GetCrudDataList, getCrudDataList } from "./getDataList";
import { getViewData as getDataFromView } from "./getViewData";
import { GetViewDataList, getViewDataList as getDataListFromView } from "./getViewDataList";
import { Order } from "./order";
import { UpdateResponse, updateCrudData } from "./update";
import { UpdateByFilterResponse, updateCrudDataByFilter } from "./updateByFilter";
import { Wait } from "./wait";

export interface DeliveryClient {
    getData: <T extends CrudData>(
        type: string,
        authData: AuthData,
        id: string,
        filter?: Filter,
        returnEmptyDataIfNotFound?: boolean,
        useStrongConsistency?: boolean,
        target?: DeploymentTarget,
        wait?: Wait
    ) => Promise<T | null>;
    getViewData: <T extends CrudData>(
        view: string,
        authData: AuthData,
        filter?: Filter,
        useStrongConsistency?: boolean
    ) => Promise<T | null>;
    getDataList: <T extends CrudData>(
        type: string,
        authData: AuthData,
        limit?: number,
        page?: number,
        filter?: Filter,
        order?: Order[],
        useStrongConsistency?: boolean,
        target?: DeploymentTarget
    ) => Promise<GetCrudDataList<T>>;
    getViewDataList: <T extends CrudData>(
        view: string,
        authData: AuthData,
        limit?: number,
        page?: number,
        filter?: Filter,
        order?: Order[],
        useStrongConsistency?: boolean
    ) => Promise<GetCrudDataList<T> | null>;
    create: <T extends CrudData>(
        type: string,
        authData: AuthData,
        data: T,
        id?: string,
        eventMetadata?: Partial<EventMetadata>
    ) => Promise<CreateResponse<T>>;
    update: <T extends CrudData>(
        type: string,
        authData: AuthData,
        id: string,
        data: Partial<T>,
        eventMetadata?: Partial<EventMetadata>
    ) => Promise<UpdateResponse<T>>;
    updateByFilter: <T extends CrudData>(
        type: string,
        authData: AuthData,
        filter: Filter,
        data: Partial<T>,
        eventMetadata?: Partial<EventMetadata>
    ) => Promise<UpdateByFilterResponse>;
    clone: <T extends CrudData>(
        type: string,
        authData: AuthData,
        id: string,
        newId: string,
        data?: Partial<T>,
        eventMetadata?: Partial<EventMetadata>
    ) => Promise<CloneResponse<T>>;
    deleteDataById: (
        type: string,
        authData: AuthData,
        id: string,
        eventMetadata?: Partial<EventMetadata>
    ) => Promise<number>;
    deleteDataByFilter: (
        type: string,
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

    const getData = async <T extends CrudData>(
        type: string,
        authData: AuthData,
        id: string,
        filter: Filter = { fields: {}, and: [], or: [] },
        returnEmptyDataIfNotFound: boolean = false,
        useStrongConsistency: boolean = false,
        target: DeploymentTarget = "DEPLOYMENT_TARGET_BLUE",
        wait?: Wait
    ): Promise<T | null> => {
        return await getCrudData<T>(
            type,
            authData,
            id,
            filter,
            returnEmptyDataIfNotFound,
            !!useStrongConsistency,
            target,
            serviceClient,
            wait
        );
    };

    const getViewData = async <T extends CrudData>(
        view: string,
        auth: AuthData,
        filter: Filter = { fields: {}, and: [], or: [] },
        useStrongConsistency: boolean = false
    ): Promise<T | null> => {
        return await getDataFromView<T>(
            view,
            auth,
            filter,
            !!useStrongConsistency,
            config.deploymentTarget,
            serviceClient
        );
    };

    const getDataList = async <T extends CrudData>(
        type: string,
        authData: AuthData,
        limit: number = 0,
        page: number = 1,
        filter: Filter = { fields: {}, and: [], or: [] },
        order: Order[] = [],
        useStrongConsistency: boolean = false,
        target: DeploymentTarget = "DEPLOYMENT_TARGET_BLUE"
    ): Promise<GetCrudDataList<T>> => {
        return await getCrudDataList<T>(
            type,
            authData,
            limit,
            page,
            filter,
            order,
            !!useStrongConsistency,
            target,
            serviceClient
        );
    };

    const getViewDataList = async <T extends CrudData>(
        view: string,
        auth: AuthData,
        limit: number = 0,
        page: number = 1,
        filter: Filter = { fields: {}, and: [], or: [] },
        order: Order[] = [],
        useStrongConsistency: boolean = false
    ): Promise<GetViewDataList<T> | null> => {
        return await getDataListFromView<T>(
            view,
            auth,
            limit,
            page,
            filter,
            order,
            !!useStrongConsistency,
            config.deploymentTarget,
            serviceClient
        );
    };

    const create = async <T extends CrudData>(
        type: string,
        authData: AuthData,
        data: T,
        id: string = "",
        eventMetadata: Partial<EventMetadata> | null = null
    ) => {
        return await createCrudData<T>(
            type,
            authData,
            data,
            id,
            eventMetadata,
            config.deploymentTarget,
            serviceClient
        );
    };

    const update = async <T extends CrudData>(
        type: string,
        authData: AuthData,
        id: string,
        data: Partial<T>,
        eventMetadata: Partial<EventMetadata> | null = null
    ) => {
        return await updateCrudData<T>(
            type,
            authData,
            id,
            data,
            eventMetadata,
            config.deploymentTarget,
            serviceClient
        );
    };

    const updateByFilter = async <T extends CrudData>(
        type: string,
        authData: AuthData,
        filter: Filter,
        data: Partial<T>,
        eventMetadata: Partial<EventMetadata> | null = null
    ) => {
        return await updateCrudDataByFilter<T>(
            type,
            authData,
            filter,
            data,
            eventMetadata,
            config.deploymentTarget,
            serviceClient
        );
    };

    const clone = async <T extends CrudData>(
        type: string,
        authData: AuthData,
        id: string,
        newId: string,
        data: Partial<T> = {},
        eventMetadata: Partial<EventMetadata> | null = null
    ) => {
        return await cloneCrudData<T>(
            type,
            authData,
            id,
            newId,
            data,
            eventMetadata,
            config.deploymentTarget,
            serviceClient
        );
    };

    const deleteDataById = async (
        type: string,
        authData: AuthData,
        id: string,
        eventMetadata: Partial<EventMetadata> | null = null
    ) => {
        return await deleteCrudData(
            type,
            authData,
            id,
            { fields: {}, and: [], or: [] },
            eventMetadata,
            config.deploymentTarget,
            serviceClient
        );
    };

    const deleteDataByFilter = async (
        type: string,
        authData: AuthData,
        filter: Filter = { fields: {}, and: [], or: [] },
        eventMetadata: Partial<EventMetadata> | null = null
    ) => {
        return await deleteCrudData(
            type,
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
        create,
        update,
        updateByFilter,
        clone,
        deleteDataById,
        deleteDataByFilter,
        close,
    };
};
