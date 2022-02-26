import ImageTableNavigationIcon from "@components/image/image-table/image-table-navigation-icon";

import { Group } from "@mantine/core";

import type { MoveDirection } from "@custom-types//images.types";
import type { FC } from "react";

type Props = {
  loading: boolean;
  isFirstPage: boolean;
  isLastPage: boolean;
  selectedPage: number;
  onMove(direction: MoveDirection): void;
};

const ImageTableNavigation: FC<Props> = ({
  loading,
  isFirstPage,
  isLastPage,
  selectedPage,
  onMove,
}) => {
  return (
    <Group>
      <ImageTableNavigationIcon
        disabled={loading || isFirstPage}
        direction="left"
        onMove={onMove}
      />
      <span>{selectedPage}</span>
      <ImageTableNavigationIcon
        disabled={loading || isLastPage}
        direction="right"
        onMove={onMove}
      />
    </Group>
  );
};

export default ImageTableNavigation;
