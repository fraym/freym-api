import { FieldDefinitionNode, GraphQLObjectType, Kind, NamedTypeNode, TypeNode } from "graphql";
import { TypeField } from "./data";
import { getFieldDirectives } from "./directive";
import { ensureValidName } from "./util";

export const getObjectFields = (t: GraphQLObjectType, namespace: string): TypeField[] => {
    const fields: TypeField[] = [];

    t.astNode?.fields?.forEach(f => {
        fields.push(getObjectField(f, namespace));
    });

    return fields;
};

const getObjectField = (f: FieldDefinitionNode, namespace: string): TypeField => {
    return {
        name: f.name.value,
        directives: getFieldDirectives(f),
        type: getFieldType(f.type, namespace),
    };
};

const getFieldType = (t: TypeNode, namespace: string): string[] => {
    switch (t.kind) {
        case Kind.NAMED_TYPE:
            return getFieldTypeFromNamedTypeNode(t, namespace);
        case Kind.LIST_TYPE:
            return ["List", ...getFieldType(t.type, namespace)];
        case Kind.NON_NULL_TYPE:
            return ["NonNull", ...getFieldType(t.type, namespace)];
    }
};

const getFieldTypeFromNamedTypeNode = (t: NamedTypeNode, namespace: string): string[] => {
    const name = t.name.value;

    if (
        name === "String" ||
        name === "Float" ||
        name === "ID" ||
        name === "Boolean" ||
        name === "Int" ||
        name === "DateTime" ||
        name === "EventEnvelope" ||
        name === "File"
    ) {
        return ["Named", name];
    }

    ensureValidName(`${namespace}${name}`);
    return ["Named", `${namespace}${name}`];
};
