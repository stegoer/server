import ImagesFormComponent from "@features/images/components/images-form/images-form.component";
import { LSB_USED_MARK } from "@features/images/images.constants";
import { useDecodeImageMutation } from "@graphql/generated/codegen.generated";

import { Text } from "@mantine/core";
import { useCallback, useState } from "react";

import type { UseFormType } from "@features/images/images.types";
import type { ReactNode } from "react";

const DecodeImagesComponent = (): JSX.Element => {
  const [decodeImageResult, decodeImage] = useDecodeImageMutation();
  const [message, setMessage] = useState<string>();
  const [error, setError] = useState<ReactNode>();

  const onSubmit = useCallback(
    (values: UseFormType[`values`]) => {
      // eslint-disable-next-line unicorn/no-useless-undefined
      setError(undefined);

      if (values.file && values.channel) {
        void decodeImage({
          lsbUsed: values.lsbUsed / LSB_USED_MARK,
          channel: values.channel,
          file: values.file,
        }).then((result) => {
          if (result.error) {
            setError(result.error.message);
          } else if (result.data?.decodeImage.message) {
            setMessage(result.data.decodeImage.message);
          }
        });
      }
    },
    [decodeImage],
  );

  return (
    <ImagesFormComponent
      formType="decode"
      loading={decodeImageResult.fetching}
      onSubmit={onSubmit}
      error={error}
      setError={setError}
    >
      {decodeImageResult.data?.decodeImage && (
        <Text color="green">message: {message}</Text>
      )}
    </ImagesFormComponent>
  );
};

export default DecodeImagesComponent;
