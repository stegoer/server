import type { User } from "@graphql/generated/codegen.generated";
import type { Dispatch, SetStateAction } from "react";

export type UserPayload = readonly [
  User | undefined,
  Dispatch<SetStateAction<User | undefined>>,
];
