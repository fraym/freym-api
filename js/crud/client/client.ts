import { ServiceClient } from "@fraym/proto/dist/index.freym.crud.delivery";
import { credentials } from "@grpc/grpc-js";
import { AuthData } from "./auth";
import { CloneResponse, cloneCrudData } from "./clone";
import { DeliveryClientConfig, useDeliveryConfigDefaults } from "./config";
import { CreateResponse, createCrudData } from "./create";
import { CrudData } from "./data";
import { deleteCrudData } from "./delete";
import { EventMetadata, fillMetadataWithDefaults } from "./eventMetadata";
import { Filter } from "./filter";
import { getCrudData } from "./getData";
import { GetCrudDataList, getCrudDataList } from "./getDataList";
import { Order } from "./order";
import { UpdateResponse, updateCrudData } from "./update";
import { Wait } from "./wait";

export interface DeliveryClient {
    getData: <T extends CrudData>(
        type: string,
        authData: AuthData,
        id: string,
        filter?: Filter,
        returnEmptyDataIfNotFound?: boolean,
        useStrongConsistency?: boolean,
        deploymentId?: number,
        wait?: Wait
    ) => Promise<T | null>;
    getDataList: <T extends CrudData>(
        type: string,
        authData: AuthData,
        limit?: number,
        page?: number,
        filter?: Filter,
        order?: Order[],
        useStrongConsistency?: boolean,
        deploymentId?: number
    ) => Promise<GetCrudDataList<T>>;
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

export const newDeliveryClient = async (config?: DeliveryClientConfig): Promise<DeliveryClient> => {
    config = useDeliveryConfigDefaults(config);
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
        deploymentId: number | null = null,
        wait?: Wait
    ): Promise<T | null> => {
        return await getCrudData<T>(
            type,
            authData,
            id,
            filter,
            returnEmptyDataIfNotFound,
            !!useStrongConsistency,
            deploymentId,
            serviceClient,
            wait
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
        deploymentId: number | null = null
    ): Promise<GetCrudDataList<T>> => {
        return await getCrudDataList<T>(
            type,
            authData,
            limit,
            page,
            filter,
            order,
            !!useStrongConsistency,
            deploymentId,
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
            fillMetadataWithDefaults(eventMetadata),
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
            fillMetadataWithDefaults(eventMetadata),
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
            fillMetadataWithDefaults(eventMetadata),
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
            fillMetadataWithDefaults(eventMetadata),
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
            fillMetadataWithDefaults(eventMetadata),
            serviceClient
        );
    };

    const close = async () => {
        serviceClient.close();
    };

    return {
        getData,
        getDataList,
        create,
        update,
        clone,
        deleteDataById,
        deleteDataByFilter,
        close,
    };
};
