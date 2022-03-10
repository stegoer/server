import { TextInput } from "@mantine/core";
import { EnvelopeClosedIcon } from "@modulz/radix-icons";

import type { UseForm } from "@mantine/hooks/lib/use-form/use-form";

type Props<T extends { email: string }> = {
  form: UseForm<T>;
  disabled: boolean;
};

const EmailInput = <T extends { email: string }>({
  form,
  disabled,
}: Props<T>) => {
  return (
    <TextInput
      label="Email"
      placeholder="Your email"
      icon={<EnvelopeClosedIcon />}
      disabled={disabled}
      required
      {...form.getInputProps(`email`)}
    />
  );
};

export default EmailInput;
