import { UnstyledButton } from "@mantine/core";
import { useNotifications } from "@mantine/notifications";

import type { NotificationProps } from "@mantine/notifications";
import type { PropsWithChildren } from "react";

type Props = PropsWithChildren<{
  notificationProps: NotificationProps;
}>;

const NotificationButton = ({
  children,
  notificationProps,
}: Props): JSX.Element => {
  const notifications = useNotifications();

  return (
    <UnstyledButton
      onClick={() => notifications.showNotification(notificationProps)}
    >
      {children}
    </UnstyledButton>
  );
};

export default NotificationButton;
