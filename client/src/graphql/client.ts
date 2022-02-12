import scalarsExchange from "@/graphql/exchange/scalars.exchange";
import LocalStorageService from "@/services/local-storage.service";

import { multipartFetchExchange } from "@urql/exchange-multipart-fetch";
import { useMemo } from "react";
import { createClient } from "urql";

const useClient = (options?: RequestInit) => {
  const token = LocalStorageService.get(`token`) ?? ``;

  return useMemo(() => {
    return createClient({
      url: `${process.env.NEXT_PUBLIC_SERVER_URI as string}/graphql`,
      exchanges: [scalarsExchange, multipartFetchExchange],
      fetchOptions: () => {
        return {
          headers: {
            Authorization: token,
            ...(options?.headers ? options.headers : {}),
          },
        };
      },
    });
  }, [token, options]);
};

export default useClient;
