import { TextInput } from "@mantine/core";
import { AvatarIcon } from "@modulz/radix-icons";

import type { UseForm } from "@mantine/hooks/lib/use-form/use-form";

type Props<T extends { username: string }> = {
  form: UseForm<T>;
};

const UsernameInput = <T extends { username: string }>({ form }: Props<T>) => {
  return (
    <TextInput
      label="Username"
      placeholder="Your username"
      icon={<AvatarIcon />}
      required
      {...form.getInputProps(`username`)}
    />
  );
};

export default UsernameInput;
