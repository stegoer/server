import { TextInput } from "@mantine/core";
import { EnvelopeClosedIcon } from "@modulz/radix-icons";

import type { UseForm } from "@mantine/hooks/lib/use-form/use-form";

type Props<T> = {
  form: UseForm<{ email: string } & T>;
};

const EmailInput = <T,>({ form }: Props<T>) => {
  return (
    <TextInput
      label="Email"
      placeholder="Your email"
      icon={<EnvelopeClosedIcon />}
      required
      {...form.getInputProps(`email`)}
    />
  );
};

export default EmailInput;
