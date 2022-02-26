import { AvatarIcon, CheckIcon, Cross2Icon } from "@modulz/radix-icons";

import type { NotificationProps } from "@mantine/notifications";

export enum NotificationTitle {
  ACCOUNT = `Account`,
  ENCODE = `Encode`,
  DECODE = `Decode`,
  IMAGES = `Images`,
}

export const GREEN_AVATAR: NotificationProps = {
  message: `GREEN_AVATAR`,
  icon: <AvatarIcon />,
  color: `green`,
};

export const GREEN_CHECK: NotificationProps = {
  message: `GREEN_CHECK`,
  icon: <CheckIcon />,
  color: `green`,
};

export const RED_CROSS: NotificationProps = {
  message: `RED_CROSS`,
  icon: <Cross2Icon />,
  color: `red`,
};
