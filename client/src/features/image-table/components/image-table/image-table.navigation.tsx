import ImageTableNavigationIcon from "@features/image-table/components/image-table/image-table.navigation.icon";

import { Group } from "@mantine/core";

import type { MoveDirection } from "@features/image-table/image-table.types";

export type ImageTableNavigationProps = {
  loading: boolean;
  isFirstPage: boolean;
  isLastPage: boolean;
  selectedPage: number;
  onMove(direction: MoveDirection): void;
};

const ImageTableNavigation = ({
  loading,
  isFirstPage,
  isLastPage,
  selectedPage,
  onMove,
}: ImageTableNavigationProps): JSX.Element => {
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
