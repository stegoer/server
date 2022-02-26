import { SERVER_GRAPHQL_ENDPOINT } from "@config/environment";
import Exchanges from "@exchanges/base.exchange";

import { createClient as createURQLClient } from "urql";

const createClient = (options?: RequestInit) => {
  return createURQLClient({
    url: SERVER_GRAPHQL_ENDPOINT,
    exchanges: Exchanges,
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
