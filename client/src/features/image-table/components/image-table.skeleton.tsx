import ImageTableNavigationSkeleton from "@features/image-table/components/image-table/skeleton/image-table.navigation.skeleton";
import TableSkeleton from "@features/image-table/components/image-table/skeleton/image-table.skeleton";

const ImageTableSkeleton = (): JSX.Element => {
  return (
    <>
      <TableSkeleton />
      <ImageTableNavigationSkeleton />
    </>
  );
};

export default ImageTableSkeleton;
