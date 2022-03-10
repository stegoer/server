import ImageFileInput from "@components/input/image-file.input";
import MessageInput from "@components/input/message.input";
import ChannelSwitches from "@features/images/components/images-form/channel.switch";
import LSBUsedSlider from "@features/images/components/images-form/lsb-used.slider";

import type { FormType } from "@features/images/images.types";
import type { Channel } from "@graphql/generated/codegen.generated";
import type { UseForm } from "@mantine/hooks/lib/use-form/use-form";

type Props = {
  form: UseForm<{
    message: string;
    lsbUsed: number;
    channel?: Channel;
    file?: File;
  }>;
  formType: FormType;
  disabled: boolean;
};

const ImagesFormInput = ({ form, formType, disabled }: Props): JSX.Element => {
  return (
    <>
      {formType === `encode` && (
        <MessageInput
          form={form}
          placeholder={`Message to ${formType} into your image`}
          disabled={disabled}
        />
      )}
      <LSBUsedSlider form={form} />
      <ChannelSwitches form={form} disabled={disabled} />
      <ImageFileInput form={form} disabled={disabled} />
    </>
  );
};

export default ImagesFormInput;
