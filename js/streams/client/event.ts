import {
    ErroneousEvent as ErroneousEventPb,
    Event,
} from "@fraym/proto/dist/index.freym.streams.management";

export interface SubscriptionEvent extends BaseEvent {
    topic: string;
    raisedAt: Date;
    orderSerial?: number;
    deploymentId?: number;
    retryNumber?: number;
}

export interface PublishEvent extends BaseEvent {
    broadcast?: boolean;
}

export interface BaseEvent {
    id: string;
    payload: Record<string, EventData>;
    tenantId: string;
    type?: string;
    stream?: string;
    correlationId?: string;
    causationId?: string;
    reason?: string;
    userId?: string;
}

/* eslint-disable-next-line @typescript-eslint/no-explicit-any */
export type EventData = any | GdprEventData;

export interface GdprEventData {
    id?: string;
    /* eslint-disable-next-line @typescript-eslint/no-explicit-any */
    value: any;
    /* eslint-disable-next-line @typescript-eslint/no-explicit-any */
    gdprDefault: any;
    isInvalidated: boolean;
}

export const isGdprEventData = (value: EventData): value is GdprEventData => {
    return (
        value &&
        typeof value === "object" &&
        Object.keys(value).length == 2 &&
        Object.prototype.hasOwnProperty.call(value, "value") &&
        Object.prototype.hasOwnProperty.call(value, "gdprDefault")
    );
};

export type HandlerFunc = (event: SubscriptionEvent) => Promise<void>;

export const getSubscriptionEvent = (event: Event): SubscriptionEvent => {
    const payload: Record<string, EventData> = {};

    for (const key in event.payload) {
        const data = event.payload[key];

        if (data === undefined) {
            continue;
        }

        if (data.gdpr) {
            payload[key] = {
                id: data.gdpr.id,
                value: JSON.parse(data.value),
                gdprDefault: data.gdpr.default ? JSON.parse(data.gdpr.default) : "",
                isInvalidated: data.gdpr.isInvalidated,
            };
        } else {
            payload[key] = JSON.parse(data.value);
        }
    }

    return {
        id: event.id,
        topic: event.topic,
        tenantId: event.tenantId,
        payload,
        raisedAt: new Date(parseInt(event.raisedAt.slice(0, -6))),
        stream: event.stream || undefined,
        type: event.type || undefined,
        causationId: event.metadata ? event.metadata.causationId : undefined,
        correlationId: event.metadata ? event.metadata.correlationId : undefined,
        reason: event.reason || undefined,
        orderSerial: event.metadata ? parseInt(event.metadata.orderSerial) : undefined,
        deploymentId:
            event.metadata && event.metadata.deploymentId
                ? parseInt(event.metadata.deploymentId)
                : undefined,
        userId: event.metadata ? event.metadata.userId || undefined : undefined,
        retryNumber: event.metadata ? parseInt(event.metadata.retryNumber) : undefined,
    };
};

export interface ErroneousEvent {
    event: SubscriptionEvent;
    consumerGroup: string;
    error: string;
}

export const getErroneousEvent = (event: ErroneousEventPb): ErroneousEvent | null => {
    if (!event.event) {
        return null;
    }

    return {
        event: getSubscriptionEvent(event.event),
        consumerGroup: event.consumerGroup,
        error: event.error,
    };
};
