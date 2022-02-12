import useClient from "@/graphql/client";

import { Provider } from "urql";

import type { FC } from "react";

const GraphqlProvider: FC = ({ children }) => {
  const client = useClient();

  return <Provider value={client}>{children}</Provider>;
};

export default GraphqlProvider;
