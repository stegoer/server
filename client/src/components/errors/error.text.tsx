import { Text } from "@mantine/core";

import type { SharedTextProps } from "@mantine/core";
import type { ReactNode } from "react";

type Props = {
  error: ReactNode;
  textProps?: SharedTextProps;
};

const ErrorText = ({ error, textProps }: Props): JSX.Element => {
  return (
    <Text color="red" size="sm" mt="sm" {...textProps}>
      {error}
    </Text>
  );
};

export default ErrorText;
