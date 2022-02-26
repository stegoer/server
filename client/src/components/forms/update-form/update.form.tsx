import SubmitButton from "@components/buttons/submit.button";
import ConfirmPasswordInput from "@components/input/confirm-password.input";
import EmailInput from "@components/input/email.input";
import PasswordStrength from "@components/input/password-strength/password-strength.input";
import UsernameInput from "@components/input/username.input";
import { useUpdateUserMutation } from "@graphql/generated/codegen.generated";
import useAuthForm from "@hooks/auth-form.hook";
import userNotUpdatedNotification from "@notifications/user-not-updated.notification";
import userUpdatedNotification from "@notifications/user-updated.notification";

import { Anchor, Collapse, Group, LoadingOverlay, Text } from "@mantine/core";
import { useNotifications } from "@mantine/notifications";
import { useCallback, useState } from "react";

import type { User } from "@graphql/generated/codegen.generated";
import type { FC } from "react";

type Props = {
  user: User;
};

const getUpdatedValue = (user: User, key: keyof User, value?: string) =>
  value && value !== user[key] ? value : undefined;

const UserForm: FC<Props> = ({ user }) => {
  const [passwordOpen, setPasswordOpen] = useState(false);
  const form = useAuthForm(`register`, passwordOpen, user);
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
      const password = passwordOpen ? values.password.trim() : undefined;

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

  const errorContent = (
    <Text color="red" size="sm" mt="sm">
      {error}
    </Text>
  );

  return (
    <form onSubmit={form.onSubmit(onSubmit)}>
      <LoadingOverlay visible={loading} />
      <UsernameInput form={form} />
      <EmailInput form={form} />

      {error && !passwordOpen && errorContent}

      <Group position="apart" mt="xs">
        <Anchor
          size="sm"
          onClick={() => setPasswordOpen((current) => !current)}
        >
          Set new password?
        </Anchor>
        <Collapse in={passwordOpen}>
          <PasswordStrength form={form} />
          <ConfirmPasswordInput form={form} />
        </Collapse>

        {error && passwordOpen && errorContent}
        <SubmitButton disabled={loading}>Update</SubmitButton>
      </Group>
    </form>
  );
};

export default UserForm;
