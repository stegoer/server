import { UnstyledButton } from "@mantine/core";
import { useNotifications } from "@mantine/notifications";

import type { NotificationProps } from "@mantine/notifications";
import type { FC } from "react";

type Props = {
  notificationProps: NotificationProps;
};

const NotificationButton: FC<Props> = ({ children, notificationProps }) => {
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
