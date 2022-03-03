import { SCHEMA } from "@graphql/graphql.constants";

import customScalarsExchange from "urql-custom-scalars-exchange";

const scalarsExchange = customScalarsExchange({
  schema: SCHEMA,
  scalars: {
    Time(value: string) {
      return new Date(Date.parse(value));
    },
  },
});

export default scalarsExchange;
