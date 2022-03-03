import LocalStorageService from "@services/local-storage.service";

import { errorExchange as urqlErrorExchange } from "urql";

const errorExchange = urqlErrorExchange({
  onError: (error) => {
    const isAuthError = error.graphQLErrors.some(
      (error_) => error_.extensions?.code === `AUTHORIZATION_ERROR`,
    );

    if (isAuthError) {
      LocalStorageService.remove(`token`);
    }
  },
});

export default errorExchange;
