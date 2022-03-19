import { IMAGE_TABLE_HOTKEY_NAVIGATION } from "@features/image-table/image-table.constants";

import { ActionIcon } from "@mantine/core";
import { useHotkeys } from "@mantine/hooks";
import { ArrowLeftIcon, ArrowRightIcon } from "@modulz/radix-icons";
import { useCallback } from "react";

import type { MoveDirection } from "@features/image-table/image-table.types";

export type ImageTableNavigationIconProps = {
  disabled: boolean;
  direction: MoveDirection;
  onMove(direction: MoveDirection): void;
};

const ImageTableNavigationIcon = ({
  disabled,
  direction,
  onMove,
}: ImageTableNavigationIconProps): JSX.Element => {
  const onClick = useCallback(() => {
    if (!disabled) {
      onMove(direction);
    }
  }, [direction, disabled, onMove]);

  useHotkeys([[IMAGE_TABLE_HOTKEY_NAVIGATION[direction], onClick]]);

  return (
    <ActionIcon
      onClick={onClick}
      disabled={disabled}
    >
      {direction === `left` ? (
        <ArrowLeftIcon
          width={25}
          height={25}
        />
      ) : (
        <ArrowRightIcon
          width={25}
          height={25}
        />
      )}
    </ActionIcon>
  );
};

export default ImageTableNavigationIcon;
