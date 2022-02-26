import type { MoveDirection } from "@custom-types/images.types";

export const IMAGE_TABLE_PER_PAGE = 10;
export const IMAGE_TABLE_HEADERS = [`Channel`, `Created`];
export const IMAGE_TABLE_HOTKEY_NAVIGATION: Record<
  MoveDirection,
  `ArrowLeft` | `ArrowRight`
> = {
  left: `ArrowLeft`,
  right: `ArrowRight`,
};
