import { createContext } from "react";

import type { User } from "@graphql/generated/codegen.generated";
import type { Dispatch, SetStateAction } from "react";

export type UserPayload = readonly [
  User | undefined,
  Dispatch<SetStateAction<User | undefined>>,
];

const UserContext = createContext<UserPayload | undefined>(undefined);

export default UserContext;
