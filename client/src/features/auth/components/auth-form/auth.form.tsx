import ErrorText from "@components/errors/error.text";
import AuthFormInput from "@features/auth/components/auth-form/auth-form.input";
import AuthFormNavigation from "@features/auth/components/auth-form/auth-form.navigation";
import {
  useCreateUserMutation,
  useLoginMutation,
} from "@graphql/generated/codegen.generated";
import useAuthForm from "@hooks/auth-form.hook";
import useAuth from "@hooks/auth.hook";
import { capitalize } from "@utils/format.utils";

import { LoadingOverlay } from "@mantine/core";
import { useCallback, useState } from "react";

import type { FormType } from "@features/auth/auth.types";
import type { SetStateAction } from "react";

export type AuthFormProps = {
  formType: FormType;
  toggleFormType(value?: SetStateAction<FormType> | undefined): void;
};

const AuthForm = ({ formType, toggleFormType }: AuthFormProps): JSX.Element => {
  const form = useAuthForm(formType, true);
  const auth = useAuth();
  const [loginResult, login] = useLoginMutation();
  const [createUserResult, createUser] = useCreateUserMutation();
  const [error, setError] = useState<string>();

  const title = capitalize(formType);
  const loading = createUserResult.fetching || loginResult.fetching;

  const resetError = useCallback(() => {
    // eslint-disable-next-line unicorn/no-useless-undefined
    setError(undefined);
  }, []);

  const onToggle = useCallback(() => {
    form.reset();
    toggleFormType();
    resetError();
  }, [form, resetError, toggleFormType]);

  const onLogin = useCallback(
    (values: { email: string; password: string }) => {
      void login({
        email: values.email.trim(),
        password: values.password,
      }).then((result) => {
        if (result.error) {
          setError(result.error.message);
        } else if (result.data?.login) {
          auth.afterLogin(result.data.login.auth.token, result.data.login.user);
        }
      });
    },
    [auth, login],
  );

  const onRegister = useCallback(
    (values: { username: string; email: string; password: string }) => {
      void createUser({
        username: values.username.trim(),
        email: values.email.trim(),
        password: values.password,
      }).then((result) => {
        if (result.error) {
          setError(result.error.message);
        } else if (result.data?.createUser) {
          auth.afterLogin(
            result.data.createUser.auth.token,
            result.data.createUser.user,
          );
        }
      });
    },
    [auth, createUser],
  );

  const onSubmit = useCallback(
    (values: typeof form[`values`]) => {
      resetError();
      if (formType === `login`) {
        onLogin(values);
      } else {
        onRegister(values);
      }
    },
    [formType, onLogin, onRegister, resetError],
  );

  return (
    <form onSubmit={form.onSubmit(onSubmit)}>
      <LoadingOverlay visible={loading} />

      <AuthFormInput
        form={form}
        formType={formType}
        disabled={loading}
      />

      {error && <ErrorText error={error} />}

      <AuthFormNavigation
        formType={formType}
        loading={loading}
        title={title}
        onToggle={onToggle}
      />
    </form>
  );
};

export default AuthForm;
