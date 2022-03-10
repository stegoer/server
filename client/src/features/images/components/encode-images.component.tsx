import ImageData from "@features/images/components/image-data";
import ImagesFormComponent from "@features/images/components/images-form/images-form.component";
import {
  IMAGE_DATA_URI_PREFIX,
  LSB_USED_MARK,
} from "@features/images/images.constants";
import { useEncodeImageMutation } from "@graphql/generated/codegen.generated";

import NextImage from "next/image";
import { useCallback, useState } from "react";

import type { UseFormType } from "@features/images/images.types";
import type { Image } from "@graphql/generated/codegen.generated";
import type { ReactNode } from "react";

const EncodeImagesComponent = (): JSX.Element => {
  const [encodeImageResult, encodeImage] = useEncodeImageMutation();
  const [image, setImage] = useState<Image>();
  const [error, setError] = useState<ReactNode>();

  const onSubmit = useCallback(
    (values: UseFormType[`values`]) => {
      // eslint-disable-next-line unicorn/no-useless-undefined
      setError(undefined);

      if (values.file && values.channel) {
        void encodeImage({
          message: values.message,
          lsbUsed: values.lsbUsed / LSB_USED_MARK,
          channel: values.channel,
          file: values.file,
        }).then((result) => {
          if (result.error) {
            setError(result.error.message);
          } else if (result.data?.encodeImage)
            setImage(result.data.encodeImage.image);
        });
      }
    },
    [encodeImage],
  );

  return (
    <ImagesFormComponent
      formType="encode"
      loading={encodeImageResult.fetching}
      onSubmit={onSubmit}
      error={error}
      setError={setError}
    >
      {image && <ImageData image={image} />}
      {encodeImageResult.data?.encodeImage && (
        <NextImage
          src={`${IMAGE_DATA_URI_PREFIX}${encodeImageResult.data.encodeImage.file.content}`}
          layout="fill"
        />
      )}
    </ImagesFormComponent>
  );
};

export default EncodeImagesComponent;
