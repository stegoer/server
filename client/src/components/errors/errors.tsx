import { Alert, List, Text } from "@mantine/core";

import type { FC } from "react";
import type { CombinedError } from "urql";

type Props = {
  data: CombinedError;
};

const Errors: FC<Props> = ({ data }) => {
  if (data.networkError) {
    return <Text>Network error: {data.networkError.message}</Text>;
  } else if (data.graphQLErrors.length > 0) {
    return (
      <Alert title="Errors" color="red" variant="outline">
        <List>
          {data.graphQLErrors.map((error, index) => (
            <List.Item key={index}>{error.message}</List.Item>
          ))}
        </List>
      </Alert>
    );
  }
  return <>{data.message}</>;
};

export default Errors;
