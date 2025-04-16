import { Metadata as AuthMetadata } from "@fraym/proto/dist/index.freym.auth.management";

export interface Metadata {
    causationId: string;
    correlationId: string;
    userId: string;
}

export const fillMetadataWithDefaults = (metadata: Partial<Metadata> | null): AuthMetadata => {
    if (!metadata) {
        return {
            causationId: "",
            correlationId: "",
            userId: "",
        };
    }

    return {
        causationId: metadata.causationId || "",
        correlationId: metadata.correlationId || "",
        userId: metadata.userId || "",
    };
};
