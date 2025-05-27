import {
    EventMetadata as CrudEventMetadata,
    DeploymentTarget,
} from "@fraym/proto/dist/index.freym.crud.delivery";

export interface EventMetadata {
    correlationId: string;
    causationId: string;
}

export const fillMetadataWithDefaults = (
    metadata: Partial<EventMetadata> | null,
    target: DeploymentTarget
): CrudEventMetadata => {
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
