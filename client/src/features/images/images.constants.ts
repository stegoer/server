import type { ChannelSwitchStyleType } from "@features/images/images.types";

export const LSB_USED_MIN = 1;
export const LSB_USED_MAX = 8;
export const LSB_USED_MARK = 12.5;

export const CHANNEL_SWITCH_STYLES: ChannelSwitchStyleType[] = [
  { label: `use red color channel`, color: `red` },
  { label: `use green color channel`, color: `green` },
  { label: `use blue color channel`, color: `blue` },
];

export const IMAGE_DATA_URI_PREFIX = `data:image/png;base64,`;
