import generatedSchema from "@graphql/generated/introspection-schema.generated.json";

import type { IntrospectionQuery } from "graphql";

export const SCHEMA = generatedSchema as unknown as IntrospectionQuery;
