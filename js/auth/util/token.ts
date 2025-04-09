import { SignJWT, jwtVerify } from "jose";

const alg = "HS256";

export const generateJwt = async (
    appSecret: string,
    tenantId: string,
    scopes: string[] = [],
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    data: Record<string, any> = {},
    expirationTime: string | number | Date = "5m"
) => {
    const secret = new TextEncoder().encode(appSecret);

    return await new SignJWT({
        type: "access_token",
        tenantId,
        scopes,
        data,
    })
        .setProtectedHeader({
            alg,
            typ: "JWT",
        })
        .setIssuedAt()
        .setNotBefore("0s")
        .setIssuer("auth")
        .setAudience(["fraym"])
        .setExpirationTime(expirationTime)
        .sign(secret);
};

// eslint-disable-next-line @typescript-eslint/no-explicit-any
export const addDataToJwt = async (appSecret: string, token: string, data: Record<string, any>) => {
    const secret = new TextEncoder().encode(appSecret);
    const { payload, protectedHeader } = await jwtVerify(token, secret, {
        clockTolerance: "10 seconds",
    });

    if (!payload.exp) {
        throw Error("expiration time is missing in JWT");
    }

    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    const newData: Record<string, any> = payload.data ?? {};

    for (const key in data) {
        newData[key] = data[key];
    }

    return new SignJWT({ ...payload, data: newData })
        .setProtectedHeader(protectedHeader)
        .sign(secret);
};

export interface TokenData {
    tenantId: string;
    userId: string;
    scopes: string[];
    exp: number;
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    data: Record<string, any> | null;
}

export const getTokenData = async (
    appSecret: string,
    token: string,
    requireUserId: boolean = true
): Promise<TokenData> => {
    const secret = new TextEncoder().encode(appSecret);
    const { payload } = await jwtVerify(token, secret, {
        clockTolerance: "10 seconds",
    });

    if (!payload.exp) {
        throw Error("expiration time is missing in JWT");
    }

    if (requireUserId && !payload.sub) {
        throw Error("user id (subject) is missing in JWT");
    }

    return {
        tenantId: (payload.tenantId as string) ?? "",
        scopes: (payload.scopes as string[]) ?? [],
        userId: (payload.sub as string) ?? "",
        exp: payload.exp,
        data: payload.data ?? null,
    };
};
