import type { User } from "@graphql/generated/codegen.generated";

export type AuthState = {
  token?: string;
};

export type AuthPayload = {
  afterLogin(token: string, user: User): void;
  logout(): void;
};
