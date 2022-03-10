import type {
  MoveDirection,
  MoveHotkey,
} from "@features/image-table/image-table.types";

export const IMAGE_TABLE_PER_PAGE = 10;
export const IMAGE_TABLE_HEADERS = [`Channel`, `Created`];
export const IMAGE_TABLE_HOTKEY_NAVIGATION: Record<MoveDirection, MoveHotkey> =
  {
    left: `ArrowLeft`,
    right: `ArrowRight`,
  };
