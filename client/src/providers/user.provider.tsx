import UserContext from "@context/user.context";

import { useState } from "react";

import type { User } from "@graphql/generated/codegen.generated";
import type { FC } from "react";

const UserProvider: FC = ({ children }) => {
  const [user, setUser] = useState<User | undefined>();

  return (
    <UserContext.Provider value={[user, setUser]}>
      {children}
    </UserContext.Provider>
  );
};

export default UserProvider;
