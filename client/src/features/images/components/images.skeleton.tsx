import ImageTableNavigationSkeleton from "@features/images/components/image-table/skeleton/image-table-navigation.skeleton";
import ImageTableSkeleton from "@features/images/components/image-table/skeleton/image-table.skeleton";

const ImagesSkeleton = (): JSX.Element => {
  return (
    <>
      <ImageTableSkeleton />
      <ImageTableNavigationSkeleton />
    </>
  );
};

export default ImagesSkeleton;
