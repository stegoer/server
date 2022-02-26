import useAuth from "@hooks/auth.hook";
import logoutNotification from "@notifications/logout.notification";

import { Button } from "@mantine/core";
import { useNotifications } from "@mantine/notifications";
import { useCallback } from "react";

import type { User } from "@graphql/generated/codegen.generated";
import type { FC } from "react";

type Props = {
  user: User;
  disabled: boolean;
};

const LogoutButton: FC<Props> = ({ user, disabled }) => {
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
