import {
  LSB_USED_MARK,
  LSB_USED_MAX,
  LSB_USED_MIN,
} from "@features/images/images.constants";
import { Channel } from "@graphql/generated/codegen.generated";

import { useForm } from "@mantine/hooks";

import type { FormType } from "@features/images/images.types";

const useImagesForm = (formType: FormType) => {
  return useForm<{
    message: string;
    lsbUsed: number;
    channel?: Channel;
    file?: File;
  }>({
    initialValues: {
      message: ``,
      lsbUsed: LSB_USED_MARK,
      channel: Channel.RedGreenBlue,
      file: undefined,
    },

    validationRules: {
      message: (value) => formType === `decode` || !!value,
      lsbUsed: (value) => {
        value = (LSB_USED_MAX * LSB_USED_MARK) / value;
        return value >= LSB_USED_MIN && value <= LSB_USED_MAX;
      },
      channel: (value) => value !== undefined,
      file: (value) => value !== undefined,
    },

    errorMessages: {
      message: `Message can't be empty`,
      lsbUsed: `Least significant bits should be within the range [${LSB_USED_MIN}:${LSB_USED_MAX}]`,
      channel: `At least one color channel is required to encode the message`,
      file: `Please choose an image file to encode your message into`,
    },
  });
};

export default useImagesForm;
