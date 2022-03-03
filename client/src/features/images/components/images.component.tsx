import ImageTable from "@features/images/components/image-table/image-table";
import ImageTableNavigation from "@features/images/components/image-table/image-table-navigation";
import { IMAGE_TABLE_PER_PAGE } from "@features/images/images.constants";
import { useImagesQuery } from "@graphql/generated/codegen.generated";

import { Skeleton } from "@mantine/core";
import { useCallback, useEffect, useState } from "react";

import type { MoveDirection } from "@features/images/images.types";
import type {
  Image,
  ImagesConnection,
} from "@graphql/generated/codegen.generated";

const calculateEdgesIndexes = (page: number): readonly [number, number] => {
  return [(page - 1) * IMAGE_TABLE_PER_PAGE, page * IMAGE_TABLE_PER_PAGE];
};

const getImageNodes = (page: number, images: ImagesConnection): Image[] => {
  return images.edges
    .slice(...calculateEdgesIndexes(page))
    .map((image) => image.node);
};

const ImagesComponent = (): JSX.Element => {
  // table navigation/pagination
  const [page, setPage] = useState(1);
  const [imageRows, setImageRows] = useState<Image[]>([]);
  // relay pagination based query
  const [first, setFirst] = useState<number | undefined>(IMAGE_TABLE_PER_PAGE);
  const [last, setLast] = useState<number>();
  const [startCursor, setStartCursor] = useState<string>();
  const [endCursor, setEndCursor] = useState<string>();
  const [imagesQuery, fetchImages] = useImagesQuery({
    variables: {
      first,
      last,
      after: endCursor,
      before: startCursor,
    },
  });
  // UI constants
  const loading = !imagesQuery.data && imagesQuery.fetching;
  const isFirstPage = page === 1;
  const isLastPage = Boolean(
    imagesQuery.data?.images.totalCount === 0 ||
      (imagesQuery.data &&
        page ===
          Math.ceil(
            imagesQuery.data.images.totalCount / IMAGE_TABLE_PER_PAGE,
          )),
  );

  // fetch images after variables get updated
  useEffect(() => {
    void fetchImages();
  }, [fetchImages, page, first, last, startCursor, endCursor]);

  // set rows based on latest data and selected page
  useEffect(() => {
    if (imagesQuery.data) {
      setImageRows(getImageNodes(page, imagesQuery.data.images));
    }
  }, [imagesQuery.data, page]);

  const onMove = useCallback(
    (direction: MoveDirection) => {
      const isLeft = direction === `left`;

      // don't set new variables if we would move outside of bounds
      if ((isLeft && isFirstPage) || (!isLeft && isLastPage)) {
        return;
      }

      setPage((previousPage) =>
        isLeft ? previousPage - 1 : previousPage + 1,
      );
      setFirst(isLeft ? undefined : IMAGE_TABLE_PER_PAGE);
      setLast(isLeft ? IMAGE_TABLE_PER_PAGE : undefined);
      setStartCursor(
        isLeft
          ? imagesQuery.data?.images.pageInfo.startCursor ?? undefined
          : undefined,
      );
      setEndCursor(
        isLeft
          ? undefined
          : imagesQuery.data?.images.pageInfo.endCursor ?? undefined,
      );
    },
    [
      imagesQuery.data?.images.pageInfo.endCursor,
      imagesQuery.data?.images.pageInfo.startCursor,
      isFirstPage,
      isLastPage,
    ],
  );

  return (
    <Skeleton visible={loading}>
      {imagesQuery.data && <ImageTable data={imageRows} />}
      <ImageTableNavigation
        loading={loading}
        isFirstPage={isFirstPage}
        isLastPage={isLastPage}
        selectedPage={page}
        onMove={onMove}
      />
    </Skeleton>
  );
};

export default ImagesComponent;
