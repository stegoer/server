import ImageTableNavigationSkeleton from "@components/image/image-table/skeleton/image-table-navigation-skeleton";
import ImageTableSkeleton from "@components/image/image-table/skeleton/image-table-skeleton";

import type { FC } from "react";

const ImageSkeletonView: FC = () => {
  return (
    <>
      <ImageTableSkeleton />
      <ImageTableNavigationSkeleton />
    </>
  );
};

export default ImageSkeletonView;
