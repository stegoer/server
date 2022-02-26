import AuthContext from "@context/auth.context";

import { useContext } from "react";

import type { AuthPayload } from "@custom-types/auth.types";

const useAuth = (): AuthPayload => {
  const context = useContext(AuthContext);
  if (!context) {
    throw new Error(`useAuth must be used within a AuthProvider`);
  }
  return context;
};

export default useAuth;
