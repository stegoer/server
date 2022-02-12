import Errors from "@/components/errors/errors";
import DisplayImage from "@/components/image/display-image";
import ImageFileInput from "@/components/image/image-file-input";
import {
  Channel,
  useCreateImageMutation,
} from "@/graphql/generated/codegen.generated";

import { Title } from "@mantine/core";
import { useEffect, useState } from "react";

import type { Image } from "@/graphql/generated/codegen.generated";
import type { NextPage } from "next";

const Encode: NextPage = () => {
  const [file, setFile] = useState<File | undefined>();
  const [image, setImage] = useState<Image | null>();

  const [createImageResult, createImage] = useCreateImageMutation();

  useEffect(() => {
    if (file) {
      void createImage({ channel: Channel.RedGreenBlue, file }).then((r) =>
        setImage(r.data?.createImage.image),
      );
    }
  }, [file, createImage, setImage]);

  let data;
  if (createImageResult.fetching) {
    data = <div>Loading...</div>;
  } else if (createImageResult.error) {
    data = <Errors data={createImageResult.error} />;
  } else if (!file) {
    data = <ImageFileInput setSelectedFile={setFile} />;
  } else {
    data = <h2>Result: {image && <DisplayImage data={image} />}</h2>;
  }

  return (
    <>
      <Title>Encode</Title>
      {data}
    </>
  );
};

export default Encode;
