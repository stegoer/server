import {
  NotificationTitle,
  RED_CROSS,
} from "@constants/notifications.constants";

import type { User } from "@graphql/generated/codegen.generated";
import type { NotificationProps } from "@mantine/notifications";

const userNotUpdatedNotification = (user: User): NotificationProps => {
  return {
    ...RED_CROSS,
    title: NotificationTitle.ACCOUNT,
    message: `No updated values for ${user.username}`,
  };
};

export default userNotUpdatedNotification;
