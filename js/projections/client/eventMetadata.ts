import {
    DeploymentTarget,
    EventMetadata as ProjectionsEventMetadata,
} from "@fraym/proto/dist/index.freym.projections.delivery";

export interface EventMetadata {
    causationId: string;
    correlationId: string;
    target: DeploymentTarget;
}

export const fillMetadataWithDefaults = (
    metadata: Partial<EventMetadata> | null
): ProjectionsEventMetadata => {
    if (!metadata) {
        return {
            causationId: "",
            correlationId: "",
            target: "DEPLOYMENT_TARGET_BLUE",
        };
    }

    return {
        causationId: metadata.causationId || "",
        correlationId: metadata.correlationId || "",
        target: metadata.target || "DEPLOYMENT_TARGET_BLUE",
    };
};
