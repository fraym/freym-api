import { ConstListValueNode, ConstObjectValueNode, ConstValueNode, Kind } from "graphql";

// eslint-disable-next-line @typescript-eslint/no-explicit-any
export const getArgumentValue = (v: ConstValueNode): any => {
    switch (v.kind) {
        case Kind.LIST:
            return getValueFromListValueNode(v);
        case Kind.STRING:
            return v.value;
        case Kind.FLOAT:
            return parseFloat(v.value);
        case Kind.INT:
            return parseInt(v.value);
        case Kind.BOOLEAN:
            return v.value;
        case Kind.NULL:
            return null;
        case Kind.ENUM:
            return v.value;
        case Kind.OBJECT:
            return getValueFromObjectValueNode(v);
    }
};

// eslint-disable-next-line @typescript-eslint/no-explicit-any
const getValueFromListValueNode = (v: ConstListValueNode): any[] => {
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    const listValue: any[] = [];

    v.values.forEach(el => {
        listValue.push(getArgumentValue(el));
    });

    return listValue;
};

// eslint-disable-next-line @typescript-eslint/no-explicit-any
const getValueFromObjectValueNode = (v: ConstObjectValueNode): Record<string, any> => {
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    const objectValue: Record<string, any> = {};

    v.fields.forEach(f => {
        objectValue[f.name.value] = getArgumentValue(f.value);
    });

    return objectValue;
};
