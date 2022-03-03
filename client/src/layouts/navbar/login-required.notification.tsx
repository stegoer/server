import { RED_CROSS } from "@constants/notifications.constants";

import { Code } from "@mantine/core";

import type { NotificationProps } from "@mantine/notifications";
import type { ReactNode } from "react";

const loginRequiredNotification = (
  title: ReactNode,
  to: string,
): NotificationProps => {
  return {
    ...RED_CROSS,
    title: title,
    message: (
      <span>Login is required to access the {<Code>{to}</Code>} page</span>
    ),
  };
};

export default loginRequiredNotification;
