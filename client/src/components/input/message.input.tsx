import { Textarea } from "@mantine/core";

import type { UseForm } from "@mantine/hooks/lib/use-form/use-form";

export type MessageInputProps<T extends { message: string }> = {
  form: UseForm<T>;
  placeholder: string;
  disabled: boolean;
};

const MessageInput = <T extends { message: string }>({
  form,
  placeholder,
  disabled,
}: MessageInputProps<T>) => {
  return (
    <Textarea
      label="Message"
      placeholder={placeholder}
      required
      disabled={disabled}
      minRows={2}
      autosize
      {...form.getInputProps(`message`)}
    />
  );
};

export default MessageInput;
