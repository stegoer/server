import ConfirmPasswordInput from "@components/input/confirm-password.input";
import EmailInput from "@components/input/email.input";
import PasswordStrength from "@components/input/password-strength/password-strength.input";
import PasswordInput from "@components/input/password.input";
import UsernameInput from "@components/input/username.input";

import type { FormType } from "@features/auth/auth.types";
import type { UseForm } from "@mantine/hooks/lib/use-form/use-form";

type Props = {
  form: UseForm<{
    username: string;
    email: string;
    password: string;
    confirmPassword: string;
  }>;
  formType: FormType;
  disabled: boolean;
};

const AuthFormInput = ({ form, formType, disabled }: Props): JSX.Element => {
  return (
    <>
      {formType === `register` && (
        <UsernameInput form={form} disabled={disabled} />
      )}
      <EmailInput form={form} disabled={disabled} />
      {formType === `register` ? (
        <PasswordStrength form={form} disabled={disabled} />
      ) : (
        <PasswordInput form={form} disabled={disabled} />
      )}
      {formType === `register` && (
        <ConfirmPasswordInput form={form} disabled={disabled} />
      )}
    </>
  );
};

export default AuthFormInput;
