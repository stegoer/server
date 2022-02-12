import Errors from "@/components/errors/errors";
import ImageTable from "@/components/image/image-table";
import { useImagesQuery } from "@/graphql/generated/codegen.generated";

import { Loader, Title } from "@mantine/core";

import type { NextPage } from "next";

const Images: NextPage = () => {
  const [imagesQuery] = useImagesQuery({ variables: { first: 10 } });

  let data;
  if (imagesQuery.fetching) {
    data = <Loader />;
  } else if (imagesQuery.error) {
    data = <Errors data={imagesQuery.error} />;
  } else if (imagesQuery.data?.images.edges.length) {
    data = (
      <ImageTable
        data={imagesQuery.data.images.edges.map((image) => image.node)}
      />
    );
  }

  return (
    <>
      <Title>Images</Title>
      {data}
    </>
  );
};

export default Images;
