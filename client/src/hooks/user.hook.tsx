import UserContext from "@context/user.context";

import { useContext } from "react";

import type { UserPayload } from "@custom-types/user.types";

const useUser = (): UserPayload => {
  const context = useContext(UserContext);
  if (!context) {
    throw new Error(`useUser must be used within a UserProvider`);
  }
  return context;
};

export default useUser;
