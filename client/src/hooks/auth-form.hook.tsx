import { calculateStrength } from "@components/input/password-strength/constants";
import emailValidator from "@validators/account/email.validator";
import stringValidator from "@validators/account/string.validator";

import { useForm } from "@mantine/hooks";

import type { FormType } from "@custom-types//account.types";
import type { User } from "@graphql/generated/codegen.generated";

const useAuthForm = (
  formType: FormType,
  validatePassword: boolean,
  user?: User,
) => {
  const usernameValidator = stringValidator(6);

  return useForm({
    initialValues: {
      username: user ? user.username : ``,
      email: user ? user.email : ``,
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
        calculateStrength(value.trim()) === 100,
      confirmPassword: (value, values) =>
        !validatePassword ||
        formType === `login` ||
        value.trim() === values?.password.trim(),
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
