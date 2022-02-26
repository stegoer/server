import { REFRESH_INTERVAL } from "@constants/user.constants";
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
import type { FC } from "react";

const AuthProvider: FC = ({ children }) => {
  const [overviewQuery, fetchOverviewQuery] = useOverviewQuery();
  const [, refreshToken] = useRefreshTokenMutation();
  const [token, setToken] = useLocalStorageValue({ key: `token` });

  const [, setUser] = useUser();

  // whenever overview has new data we update user accordingly
  useEffect(() => {
    setUser(overviewQuery.data?.overview.user);
  }, [overviewQuery.data?.overview.user, setUser]);

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

  // every X seconds we want to refresh token
  useEffect(() => {
    const interval = setInterval(() => {
      updateToken();
    }, REFRESH_INTERVAL);

    return () => {
      updateToken();
      clearInterval(interval);
    };
  }, [updateToken]);

  // whenever token is changed/removed we want to fetch the latest data
  useEffect(() => fetchOverviewQuery(), [fetchOverviewQuery, token]);

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
    <AuthContext.Provider value={{ afterLogin, logout }}>
      {children}
    </AuthContext.Provider>
  );
};

export default AuthProvider;
