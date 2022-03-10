import SubmitButton from "@components/buttons/submit.button";
import ErrorText from "@components/errors/error.text";
import ConfirmPasswordInput from "@components/input/confirm-password.input";
import EmailInput from "@components/input/email.input";
import PasswordStrength from "@components/input/password-strength/password-strength.input";
import UsernameInput from "@components/input/username.input";
import userNotUpdatedNotification from "@features/account/notifications/user-not-updated.notification";
import userUpdatedNotification from "@features/account/notifications/user-updated.notification";
import { useUpdateUserMutation } from "@graphql/generated/codegen.generated";
import useAuthForm from "@hooks/auth-form.hook";

import { Anchor, Collapse, Group, LoadingOverlay } from "@mantine/core";
import { useNotifications } from "@mantine/notifications";
import { useCallback, useState } from "react";

import type { FormType } from "@features/auth/auth.types";
import type { User } from "@graphql/generated/codegen.generated";

type Props = {
  user: User;
};

const DEFAULT_FORM_TYPE: FormType = `register`;

const getUpdatedValue = (user: User, key: keyof User, value?: string) =>
  value && value !== user[key] ? value : undefined;

const UserForm = ({ user }: Props): JSX.Element => {
  const [passwordOpen, setPasswordOpen] = useState(false);
  const form = useAuthForm(DEFAULT_FORM_TYPE, passwordOpen, user);
  const [updateResult, updateUser] = useUpdateUserMutation();
  const [error, setError] = useState<string>();
  const notifications = useNotifications();
  const loading = updateResult.fetching;

  const onSubmit = useCallback(
    (values: typeof form[`values`]) => {
      // eslint-disable-next-line unicorn/no-useless-undefined
      setError(undefined);

      const username = getUpdatedValue(
        user,
        `username`,
        values.username.trim(),
      );
      const email = getUpdatedValue(user, `email`, values.email.trim());
      const password = passwordOpen ? values.password : undefined;

      if (username || email || password) {
        void updateUser({ username, email, password }).then((result) => {
          if (result.error) {
            setError(result.error.message);
          } else {
            notifications.showNotification(userUpdatedNotification(user));
          }
        });
      } else {
        setError(`No values updated`);
        notifications.showNotification(userNotUpdatedNotification(user));
      }
    },
    [passwordOpen, notifications, updateUser, user],
  );

  const errorContent = <ErrorText error={error} />;

  return (
    <form onSubmit={form.onSubmit(onSubmit)}>
      <LoadingOverlay visible={loading} />

      <UsernameInput form={form} disabled={loading} />
      <EmailInput form={form} disabled={loading} />

      {error && !passwordOpen && errorContent}

      <Group position="apart" mt="xs">
        <Anchor
          size="sm"
          onClick={() => setPasswordOpen((current) => !current)}
        >
          Set new password?
        </Anchor>
        <Collapse in={passwordOpen}>
          <PasswordStrength form={form} disabled={loading} />
          <ConfirmPasswordInput form={form} disabled={loading} />
        </Collapse>

        {error && passwordOpen && errorContent}
        <SubmitButton disabled={loading}>Update</SubmitButton>
      </Group>
    </form>
  );
};

export default UserForm;
