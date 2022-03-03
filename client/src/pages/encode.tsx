import DisplayImage from "@components/image/display-image";
import ImageFileInput from "@components/input/image-file.input";
import {
  Channel,
  useCreateImageMutation,
} from "@graphql/generated/codegen.generated";
import PageLayout from "@layouts/page.layout";

import { Title } from "@mantine/core";
import { useEffect, useState } from "react";

import type { Image } from "@graphql/generated/codegen.generated";
import type { NextPage } from "next";

const Encode: NextPage = () => {
  const [file, setFile] = useState<File | undefined>();
  const [image, setImage] = useState<Image | null>();

  const [createImageResult, createImage] = useCreateImageMutation();

  useEffect(() => {
    if (file) {
      void createImage({ channel: Channel.RedGreen, file }).then((r) =>
        setImage(r.data?.createImage.image),
      );
    }
  }, [file, createImage, setImage]);

  let data;
  if (createImageResult.fetching) {
    data = <div>Loading...</div>;
  } else if (createImageResult.error) {
    data = <span>Error {createImageResult.error}</span>;
  } else if (!file) {
    data = <ImageFileInput setSelectedFile={setFile} />;
  } else {
    data = <h2>Result: {image && <DisplayImage data={image} />}</h2>;
  }

  return (
    <PageLayout title="encode">
      <Title>Encode</Title>
      {data}
    </PageLayout>
  );
};

export default Encode;
