import { createContext } from "react";

import type { UserPayload } from "@custom-types//user.types";

const UserContext = createContext<UserPayload | undefined>(undefined);

export default UserContext;
