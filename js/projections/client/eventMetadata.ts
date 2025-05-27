import {
    DeploymentTarget,
    EventMetadata as ProjectionsEventMetadata,
} from "@fraym/proto/dist/index.freym.projections.delivery";

export interface EventMetadata {
    causationId: string;
    correlationId: string;
}

export const fillMetadataWithDefaults = (
    metadata: Partial<EventMetadata> | null,
    target: DeploymentTarget
): ProjectionsEventMetadata => {
    if (!metadata) {
        return {
            causationId: "",
            correlationId: "",
            target,
        };
    }

    return {
        causationId: metadata.causationId || "",
        correlationId: metadata.correlationId || "",
        target,
    };
};
