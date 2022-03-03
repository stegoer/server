import { SCHEMA } from "@graphql/graphql.constants";

import { cacheExchange } from "@urql/exchange-graphcache";
import { relayPagination } from "@urql/exchange-graphcache/extras";

import type { OverviewPayload } from "@graphql/generated/codegen.generated";
import type { Data } from "@urql/exchange-graphcache";

const graphcacheExchange = cacheExchange({
  schema: SCHEMA,
  resolvers: { Query: { images: relayPagination() } },
  keys: {
    OverviewPayload: (data: Data & OverviewPayload) => data.user.id,
  },
});

export default graphcacheExchange;
