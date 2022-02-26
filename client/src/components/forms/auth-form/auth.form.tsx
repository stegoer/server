import { capitalize } from "../../../utils/string.utils";

import AuthFormNavigation from "@components/forms/auth-form/auth-form-navigation";
import AuthFormInput from "@components/input/auth-form.input";
import {
  useCreateUserMutation,
  useLoginMutation,
} from "@graphql/generated/codegen.generated";
import useAuthForm from "@hooks/auth-form.hook";
import useAuth from "@hooks/auth.hook";

import { LoadingOverlay, Text, Title } from "@mantine/core";
import { useToggle } from "@mantine/hooks";
import { useCallback, useEffect, useState } from "react";

import type { FormType } from "@custom-types//account.types";
import type { FC } from "react";

const AuthForm: FC = () => {
  const [formType, toggleFormType] = useToggle<FormType>(`login`, [
    `login`,
    `register`,
  ]);
  const form = useAuthForm(formType, true);

  const auth = useAuth();
  const [loginResult, login] = useLoginMutation();
  const [createUserResult, createUser] = useCreateUserMutation();
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string>();

  const title = capitalize(formType);

  useEffect(
    () => setLoading(createUserResult.fetching || loginResult.fetching),
    [createUserResult.fetching, loginResult.fetching],
  );

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
        password: values.password.trim(),
      }).then((result) => {
        if (result.error) {
          setError(result.error.message);
        } else if (result.data?.login) {
          auth.afterLogin(
            result.data.login.auth.token,
            result.data.login.user,
          );
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
        password: values.password.trim(),
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
    <>
      <Title>{title}</Title>
      <form onSubmit={form.onSubmit(onSubmit)}>
        <LoadingOverlay visible={loading} />

        <AuthFormInput form={form} formType={formType} />

        {error && (
          <Text color="red" size="sm" mt="sm">
            {error}
          </Text>
        )}

        <AuthFormNavigation
          formType={formType}
          loading={loading}
          title={title}
          onToggle={onToggle}
        />
      </form>
    </>
  );
};

export default AuthForm;
