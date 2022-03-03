import { calculateStrength } from "@components/input/password-strength/password-strength.constants";
import emailValidator from "@validators/email.validator";
import lengthValidator from "@validators/string.validator";

import { useForm } from "@mantine/hooks";

import type { FormType } from "@features/auth/auth.types";
import type { User } from "@graphql/generated/codegen.generated";

const useAuthForm = (
  formType: FormType,
  validatePassword: boolean,
  user?: User,
) => {
  const usernameValidator = lengthValidator(6);

  return useForm({
    initialValues: {
      username: user?.username ?? ``,
      email: user?.email ?? ``,
      password: ``,
      confirmPassword: ``,
    },

    validationRules: {
      username: (value) =>
        formType === `login` || usernameValidator(value.trim()),
      email: emailValidator,
      password: (value) =>
        !validatePassword ||
        formType === `login` ||
        calculateStrength(value) === 100,
      confirmPassword: (value, values) =>
        !validatePassword ||
        formType === `login` ||
        value === values?.password,
    },

    errorMessages: {
      username: `Username must contain at least 6 characters`,
      email: `Invalid email address`,
      password: `Invalid password`,
      confirmPassword: `Passwords don't match. Try again`,
    },
  });
};

export default useAuthForm;
