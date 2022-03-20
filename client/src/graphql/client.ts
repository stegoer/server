import { IS_PRODUCTION, SERVER_GRAPHQL_ENDPOINT } from "@config/environment";
import Exchanges from "@graphql/exchanges/base.exchange";
import devtoolsExchange from "@graphql/exchanges/devtools.exchange";

import { createClient as createURQLClient } from "urql";

const createClient = (options?: RequestInit) => {
  return createURQLClient({
    url: SERVER_GRAPHQL_ENDPOINT,
    exchanges: IS_PRODUCTION ? Exchanges : [devtoolsExchange, ...Exchanges],
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
