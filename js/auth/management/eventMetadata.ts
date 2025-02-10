export interface EventMetadata {
    causationId: string;
    correlationId: string;
    deploymentId: string;
    userId: string;
}

export const fillMetadataWithDefaults = (
    metadata: Partial<EventMetadata> | null
): EventMetadata => {
    if (!metadata) {
        return {
            causationId: "",
            correlationId: "",
            deploymentId: "",
            userId: "",
        };
    }

    return {
        causationId: metadata.causationId || "",
        correlationId: metadata.correlationId || "",
        deploymentId: metadata.deploymentId || "",
        userId: metadata.userId || "",
    };
};
