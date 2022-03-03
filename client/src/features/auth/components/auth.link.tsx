import { Anchor } from "@mantine/core";

import type { FormType } from "@features/auth/auth.types";

type Props = {
  formType: FormType;
  toggleFormType(): void;
  disabled: boolean;
};

const AuthLink = ({
  formType,
  toggleFormType,
  disabled,
}: Props): JSX.Element => {
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
