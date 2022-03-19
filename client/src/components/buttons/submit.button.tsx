import { Button } from "@mantine/core";

import type { PropsWithChildren } from "react";

export type SubmitButtonProps = PropsWithChildren<{
  disabled: boolean;
}>;

const SubmitButton = ({
  children,
  disabled,
}: SubmitButtonProps): JSX.Element => {
  return (
    <Button
      type="submit"
      disabled={disabled}
    >
      {children}
    </Button>
  );
};

export default SubmitButton;
