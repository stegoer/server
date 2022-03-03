import {
  GREEN_AVATAR,
  NotificationTitle,
} from "@constants/notifications.constants";

import type { User } from "@graphql/generated/codegen.generated";
import type { NotificationProps } from "@mantine/notifications";

const logoutNotification = (user: User): NotificationProps => {
  return {
    ...GREEN_AVATAR,
    title: NotificationTitle.ACCOUNT,
    message: `${user.username} successfully logged out`,
  };
};

export default logoutNotification;
