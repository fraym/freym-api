import { DeploymentTarget } from "@fraym/proto/dist/index.freym.projections.delivery";

export interface EventMetadata {
    causationId: string;
    correlationId: string;
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
