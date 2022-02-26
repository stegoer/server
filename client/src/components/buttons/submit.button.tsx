import { Button } from "@mantine/core";

import type { FC } from "react";

type Props = {
  disabled: boolean;
};

const SubmitButton: FC<Props> = ({ disabled, children }) => {
  return (
    <Button type="submit" disabled={disabled}>
      {children}
    </Button>
  );
};

export default SubmitButton;
