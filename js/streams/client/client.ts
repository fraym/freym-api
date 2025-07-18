import { ServiceClient } from "@fraym/proto/dist/index.freym.streams.management";
import { credentials } from "@grpc/grpc-js";
import { getAllEvents, getAllEventsAfterEvent } from "./allEvents";
import { ClientConfig, useConfigDefaults } from "./config";
import { HandlerFunc, PublishEvent, SubscriptionEvent } from "./event";
import { getEvent } from "./getEvent";
import { getLastEvent } from "./getLastEvent";
import { getLastEventByTypes } from "./getLastEventByTypes";
import { getLastHandledEvent } from "./getLastHandledEvent";
import { introduceGdprOnEventField } from "./introduceGdpr";
import { sendInvalidateGdpr } from "./invalidateGdpr";
import { sendPublish } from "./publish";
import { renameEventType } from "./rename";
import { createStreamSnapshot, getStream, getStreamAfterEvent, isStreamEmpty } from "./stream";
import { Subscription, newSubscription } from "./subscribe";

export interface StreamIterator {
    forEach: (callback: (event: SubscriptionEvent) => void) => Promise<void>;
    forEachAfterEvent: (
        eventId: string,
        callback: (event: SubscriptionEvent) => void
    ) => Promise<void>;
    isEmpty: () => Promise<boolean>;
}

type LastEventCheck = (lastEvent: SubscriptionEvent | null) => boolean;

export interface Client {
    close: () => void;
    publish: (topic: string, events: PublishEvent[]) => Promise<void>;
    subscribe: (topics?: string[], ignoreUnhandledEvents?: boolean) => Subscription;
    getEvent: (tenantId: string, topic: string, eventId: string) => Promise<SubscriptionEvent>;
    getLastEvent: (tenantId: string, topic: string) => Promise<SubscriptionEvent | null>;
    getLastHandledEvent: (tenantId: string, topic: string) => Promise<SubscriptionEvent | null>;
    getLastEventByTypes: (
        tenantId: string,
        topic: string,
        types: string[]
    ) => Promise<SubscriptionEvent | null>;
    iterateAllEvents: (
        tenantId: string,
        topic: string,
        includedEventTypes: string[],
        perPage: number,
        handler: HandlerFunc
    ) => Promise<void>;
    iterateAllEventsAfterEvent: (
        tenantId: string,
        topic: string,
        includedEventTypes: string[],
        eventId: string,
        perPage: number,
        handler: HandlerFunc
    ) => Promise<void>;
    getStreamIterator: (
        topic: string,
        tenantId: string,
        stream: string,
        perPage: number
    ) => StreamIterator;
    createStreamSnapshot: (
        tenantId: string,
        topic: string,
        stream: string,
        lastSnapshottedEventId: string,
        snapshotEvent: PublishEvent
    ) => Promise<void>;
    invalidateGdprData: (tenantId: string, topic: string, gdprId: string) => Promise<void>;
    introduceGdprOnEventField: (
        tenantId: string,
        defaultValue: string,
        topic: string,
        eventId: string,
        fieldName: string
    ) => Promise<void>;
    renameEventType: (topic: string, oldEventType: string, newEventType: string) => Promise<void>;
}

export const newClient = async (inputConfig: ClientConfig): Promise<Client> => {
    const config = useConfigDefaults(inputConfig);
    const serviceClient = new ServiceClient(config.serverAddress, credentials.createInsecure(), {
        "grpc.keepalive_time_ms": config.keepaliveInterval,
        "grpc.keepalive_timeout_ms": config.keepaliveTimeout,
        "grpc.keepalive_permit_without_calls": 1,
        "grpc.max_receive_message_length": 2147483647,
    });

    const closeFunctions: (() => void)[] = [];

    const getLastEventCheck = async (
        tenantId: string,
        topic: string
    ): Promise<LastEventCheck | null> => {
        const now = new Date(new Date().getTime() + 3000);
        const lastEvent = await getLastEvent(tenantId, topic, serviceClient);

        if (!lastEvent) {
            return null;
        }

        const lastOrderSerial = lastEvent.orderSerial;

        return (lastEvent: SubscriptionEvent | null) => {
            if (!lastEvent) {
                return true;
            }

            if (lastOrderSerial == undefined) {
                return lastEvent.raisedAt > now;
            }

            const orderSerial = lastEvent.orderSerial ? lastEvent.orderSerial : 0;

            return orderSerial > lastOrderSerial;
        };
    };

    return {
        getEvent: async (tenantId, topic, eventId) => {
            return await getEvent(tenantId, topic, eventId, serviceClient);
        },
        getLastEvent: async (tenantId, topic) => {
            return await getLastEvent(tenantId, topic, serviceClient);
        },
        getLastHandledEvent: async (tenantId, topic) => {
            return await getLastHandledEvent(tenantId, topic, config.groupId, serviceClient);
        },
        getLastEventByTypes: async (tenantId, topic, types) => {
            return await getLastEventByTypes(tenantId, topic, types, serviceClient);
        },
        iterateAllEvents: async (tenantId, topic, includedEventTypes, perPage, handler) => {
            const lastEventCheck = await getLastEventCheck(tenantId, topic);

            if (!lastEventCheck) {
                return;
            }

            await getAllEvents(
                tenantId,
                topic,
                includedEventTypes,
                perPage,
                handler,
                lastEventCheck,
                serviceClient
            );
        },
        iterateAllEventsAfterEvent: async (
            tenantId,
            topic,
            includedEventTypes,
            eventId,
            perPage,
            handler
        ) => {
            const lastEventCheck = await getLastEventCheck(tenantId, topic);

            if (!lastEventCheck) {
                return;
            }

            await getAllEventsAfterEvent(
                tenantId,
                topic,
                includedEventTypes,
                eventId,
                perPage,
                handler,
                lastEventCheck,
                serviceClient
            );
        },
        publish: async (topic, events) => {
            return await sendPublish(topic, events, config.deploymentId, serviceClient);
        },
        getStreamIterator: (topic, tenantId, stream, perPage) => {
            return {
                forEach: async callback => {
                    const lastEventCheck = await getLastEventCheck(tenantId, topic);

                    if (!lastEventCheck) {
                        return;
                    }

                    return await getStream(
                        topic,
                        tenantId,
                        stream,
                        perPage,
                        async (event: SubscriptionEvent) => {
                            callback(event);
                        },
                        lastEventCheck,
                        config.deploymentId,
                        serviceClient
                    );
                },
                forEachAfterEvent: async (eventId, callback) => {
                    const lastEventCheck = await getLastEventCheck(tenantId, topic);

                    if (!lastEventCheck) {
                        return;
                    }

                    return await getStreamAfterEvent(
                        topic,
                        tenantId,
                        stream,
                        eventId,
                        perPage,
                        async (event: SubscriptionEvent) => {
                            callback(event);
                        },
                        lastEventCheck,
                        config.deploymentId,
                        serviceClient
                    );
                },
                isEmpty: async () => {
                    return isStreamEmpty(topic, tenantId, stream, serviceClient);
                },
            };
        },
        subscribe: (topics: string[] = [], ignoreUnhandledEvents: boolean = false) => {
            const subscription = newSubscription(
                topics,
                ignoreUnhandledEvents,
                config,
                serviceClient
            );

            closeFunctions.push(subscription.stop);

            return subscription;
        },
        introduceGdprOnEventField: async (
            tenantId: string,
            defaultValue: string,
            topic: string,
            eventId: string,
            fieldName: string
        ) => {
            return await introduceGdprOnEventField(
                tenantId,
                defaultValue,
                topic,
                eventId,
                fieldName,
                serviceClient
            );
        },
        invalidateGdprData: async (tenantId, topic, gdprId) => {
            return await sendInvalidateGdpr(tenantId, topic, gdprId, serviceClient);
        },
        createStreamSnapshot: async (
            tenantId: string,
            topic: string,
            stream: string,
            idOfLastEventThatGotSnapshotted: string,
            snapshotEvent: PublishEvent
        ) => {
            return await createStreamSnapshot(
                tenantId,
                topic,
                stream,
                idOfLastEventThatGotSnapshotted,
                snapshotEvent,
                config.deploymentId,
                serviceClient
            );
        },
        renameEventType: async (topic, oldEventType, newEventType) => {
            return await renameEventType(topic, oldEventType, newEventType, serviceClient);
        },
        close: () => {
            closeFunctions.forEach(close => close());
        },
    };
};
