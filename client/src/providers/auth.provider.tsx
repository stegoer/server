import AuthContext from "@context/auth.context";
import {
  useOverviewQuery,
  useRefreshTokenMutation,
} from "@graphql/generated/codegen.generated";
import useLocalStorageValue from "@hooks/local-storage.hook";
import useUser from "@hooks/user.hook";
import LocalStorageService from "@services/local-storage.service";

import { useCallback, useEffect } from "react";

import type { User } from "@graphql/generated/codegen.generated";
import type { PropsWithChildren } from "react";

export const REFRESH_INTERVAL = 600_000; // 10 minutes

export type AuthProviderProps = PropsWithChildren<Record<never, never>>;

const AuthProvider = ({ children }: AuthProviderProps): JSX.Element => {
  const [overviewQuery, fetchOverviewQuery] = useOverviewQuery();
  const [, refreshToken] = useRefreshTokenMutation();
  const [token, setToken] = useLocalStorageValue({ key: `token` });

  const [, setUser] = useUser();

  const updateToken = useCallback(() => {
    if (token) {
      void refreshToken({ token }, { requestPolicy: `network-only` }).then(
        (r) => {
          if (r.data?.refreshToken) {
            setToken(r.data.refreshToken.auth.token);
            setUser(r.data.refreshToken.user);
          }
        },
      );
    }
  }, [refreshToken, setToken, setUser, token]);

  // whenever token is changed/removed we want to fetch the latest data
  useEffect(() => fetchOverviewQuery(), [fetchOverviewQuery, token]);

  // whenever overview has new data we update user accordingly
  useEffect(() => {
    setUser(overviewQuery.data?.overview.user);
  }, [overviewQuery.data?.overview.user, setUser]);

  // every X seconds we want to refresh token
  useEffect(() => {
    const interval = setInterval(() => {
      updateToken();
    }, REFRESH_INTERVAL);

    return () => clearInterval(interval);
  }, [updateToken]);

  const afterLogin = useCallback(
    (token: string, user: User) => {
      setToken(token);
      setUser(user);
    },
    [setToken, setUser],
  );

  const logout = useCallback(() => {
    LocalStorageService.remove(`token`);
    // eslint-disable-next-line unicorn/no-useless-undefined
    setUser(undefined);
  }, [setUser]);

  return (
    <AuthContext.Provider
      value={{ fetching: overviewQuery.fetching, afterLogin, logout }}
    >
      {children}
    </AuthContext.Provider>
  );
};

export default AuthProvider;
