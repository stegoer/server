import schema from "@/graphql/generated/introspection-schema.generated.json";

import customScalarsExchange from "urql-custom-scalars-exchange";

import type { IntrospectionQuery } from "graphql";

const scalarsExchange = customScalarsExchange({
  schema: schema as unknown as IntrospectionQuery,
  scalars: {
    Time(value: string) {
      return new Date(Date.parse(value));
    },
  },
});

export default scalarsExchange;
