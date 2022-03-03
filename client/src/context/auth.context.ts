import { createContext } from "react";

import type { User } from "@graphql/generated/codegen.generated";

export type AuthPayload = {
  fetching: boolean;
  afterLogin(token: string, user: User): void;
  logout(): void;
};

const AuthContext = createContext<AuthPayload | undefined>(undefined);

export default AuthContext;
