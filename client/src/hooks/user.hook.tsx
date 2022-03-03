import UserContext from "@context/user.context";

import { useContext } from "react";

import type { UserPayload } from "@context/user.context";

const useUser = (): UserPayload => {
  const context = useContext(UserContext);
  if (!context) {
    throw new Error(`useUser must be used within a UserProvider`);
  }
  return context;
};

export default useUser;
