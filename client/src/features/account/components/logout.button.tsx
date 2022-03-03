import logoutNotification from "@features/account/notifications/logout.notification";
import useAuth from "@hooks/auth.hook";

import { Button } from "@mantine/core";
import { useNotifications } from "@mantine/notifications";
import { useCallback } from "react";

import type { User } from "@graphql/generated/codegen.generated";

type Props = {
  user: User;
  disabled: boolean;
};

const LogoutButton = ({ user, disabled }: Props): JSX.Element => {
  const auth = useAuth();
  const notifications = useNotifications();

  const onClick = useCallback(() => {
    auth.logout();
    notifications.showNotification(logoutNotification(user));
  }, [auth, notifications, user]);

  return (
    <Button onClick={onClick} disabled={disabled}>
      Logout
    </Button>
  );
};

export default LogoutButton;
