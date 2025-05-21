import { useEffect, useState } from "react";
import { getOptionalClient } from "@/client/client";
import { JWT } from "@/client/client";
import { ClientOptions } from "@/client/options";
import { Client } from "@urql/core";

export const useClient = (url: string, jwt: JWT, options?: ClientOptions) => {
    const [client, setClient] = useState<Client | null>(null);

    useEffect(() => {
        setClient(getOptionalClient(url, jwt, options));
    }, [url, jwt, options]);

    return client;
};
