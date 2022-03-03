import {
  GREEN_CHECK,
  NotificationTitle,
} from "@constants/notifications.constants";

import type { User } from "@graphql/generated/codegen.generated";
import type { NotificationProps } from "@mantine/notifications";

const userUpdatedNotification = (user: User): NotificationProps => {
  return {
    ...GREEN_CHECK,
    title: NotificationTitle.ACCOUNT,
    message: `${user.username} successfully updated`,
  };
};

export default userUpdatedNotification;
