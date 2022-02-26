import generatedSchema from "@graphql/generated/introspection-schema.generated.json";

import type { IntrospectionQuery } from "graphql";

const SCHEMA = generatedSchema as unknown as IntrospectionQuery;

export default SCHEMA;
