import { Anchor } from "@mantine/core";

import type { FormType } from "@custom-types//account.types";
import type { FC } from "react";

type Props = {
  formType: FormType;
  toggleFormType(): void;
  disabled: boolean;
};

const AuthLink: FC<Props> = ({ formType, toggleFormType, disabled }) => {
  return (
    <Anchor
      component="button"
      type="button"
      color="gray"
      onClick={toggleFormType}
      size="sm"
      disabled={disabled}
    >
      {formType === `register`
        ? `Have an account? Login`
        : `Don't have an account? Register`}
    </Anchor>
  );
};

export default AuthLink;
