import useClient from "@hooks/client.hook";
import useUser from "@hooks/user.hook";

import { Provider } from "urql";

import type { PropsWithChildren } from "react";

type Props = PropsWithChildren<Record<never, never>>;

const GraphqlProvider = ({ children }: Props): JSX.Element => {
  const [user] = useUser();
  const client = useClient(!!user);

  return <Provider value={client}>{children}</Provider>;
};

export default GraphqlProvider;
