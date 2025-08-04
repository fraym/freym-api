export interface GetEntryOptions {
    useStrongConsistency?: boolean;
    useStrongConsistencyById?: string;
}

export interface GetSingleEntryOptions extends GetEntryOptions {
    returnEmptyDataIfNotFound?: boolean;
}
