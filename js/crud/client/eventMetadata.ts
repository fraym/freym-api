import { DeploymentTarget } from "@fraym/proto/dist/index.freym.crud.delivery";

export interface EventMetadata {
    correlationId: string;
    causationId: string;
    userId: string;
    target: DeploymentTarget;
}

export const fillMetadataWithDefaults = (
    metadata: Partial<EventMetadata> | null
): EventMetadata => {
    if (!metadata) {
        return {
            causationId: "",
            correlationId: "",
            userId: "",
            target: "DEPLOYMENT_TARGET_BLUE",
        };
    }

    return {
        causationId: metadata.causationId || "",
        correlationId: metadata.correlationId || "",
        userId: metadata.userId || "",
        target: metadata.target || "DEPLOYMENT_TARGET_BLUE",
    };
};
