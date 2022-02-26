import { createContext } from "react";

import type { AuthPayload } from "@custom-types//auth.types";

const AuthContext = createContext<AuthPayload | undefined>(undefined);

export default AuthContext;
