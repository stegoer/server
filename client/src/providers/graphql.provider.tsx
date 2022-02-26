import useClient from "@hooks/graphql-client.hook";
import useUser from "@hooks/user.hook";

import { Provider } from "urql";

import type { FC } from "react";

const GraphqlProvider: FC = ({ children }) => {
  const [user] = useUser();
  const client = useClient(!!user);

  return <Provider value={client}>{children}</Provider>;
};

export default GraphqlProvider;
