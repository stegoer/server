import ImageTableNavigationIcon from "@features/images/components/image-table/image-table-navigation-icon";

import { Group } from "@mantine/core";

import type { MoveDirection } from "@features/images/images.types";

type Props = {
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
}: Props): JSX.Element => {
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
