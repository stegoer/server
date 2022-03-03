import LocalStorageService from "@services/local-storage.service";

import { authExchange as urqlAuthExchange } from "@urql/exchange-auth";
import { makeOperation } from "urql";

import type { AuthState } from "@features/auth/auth.types";
import type { CombinedError, Operation } from "urql";

const addAuthToOperation = ({
  authState,
  operation,
}: {
  authState: AuthState;
  operation: Operation;
}) => {
  if (!authState || !authState.token) {
    return operation;
  }

  const fetchOptions =
    typeof operation.context.fetchOptions === `function`
      ? operation.context.fetchOptions()
      : operation.context.fetchOptions || {};

  return makeOperation(operation.kind, operation, {
    ...operation.context,
    fetchOptions: {
      ...fetchOptions,
      headers: {
        ...fetchOptions.headers,
        Authorization: authState.token,
      },
    },
  });
};

// eslint-disable-next-line @typescript-eslint/require-await
const getAuth = async ({ authState }: { authState: AuthState | null }) => {
  if (!authState) {
    const token = LocalStorageService.get(`token`);
    if (token) {
      return { token };
    }
  }

  LocalStorageService.remove(`token`);

  // eslint-disable-next-line unicorn/no-null
  return null;
};

const didAuthError = ({ error }: { error: CombinedError }) => {
  return error.graphQLErrors.some(
    (error) => error.extensions?.code === `AUTHORIZATION_ERROR`,
  );
};

const authExchange = urqlAuthExchange({
  addAuthToOperation,
  getAuth,
  didAuthError,
});

export default authExchange;
