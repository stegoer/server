import { SERVER_GRAPHQL_ENDPOINT } from "@config/environment";
import Exchanges from "@exchanges/base.exchange";
import devtoolsExchange from "@exchanges/devtools.exchange";

import { createClient as createURQLClient } from "urql";

const createClient = (options?: RequestInit) => {
  return createURQLClient({
    url: SERVER_GRAPHQL_ENDPOINT,
    exchanges:
      process.env.NODE_ENV === `development`
        ? [devtoolsExchange, ...Exchanges]
        : Exchanges,
    requestPolicy: `cache-and-network`,
    fetchOptions: () => {
      return {
        headers: {
          ...(options?.headers ? options.headers : {}),
        },
      };
    },
  });
};

export default createClient;
