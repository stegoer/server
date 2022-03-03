import createClient from "@graphql/client";

import { useMemo } from "react";

import type { Client } from "urql";

const useClient = (
  isAuthenticated: boolean,
  options?: RequestInit,
): Client => {
  return useMemo(
    () => createClient(options),
    // eslint-disable-next-line react-hooks/exhaustive-deps
    [createClient, isAuthenticated, options],
  );
};

export default useClient;
