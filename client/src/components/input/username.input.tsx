import { TextInput } from "@mantine/core";
import { AvatarIcon } from "@modulz/radix-icons";

import type { UseForm } from "@mantine/hooks/lib/use-form/use-form";

export type UsernameInputProps<T extends { username: string }> = {
  form: UseForm<T>;
  disabled: boolean;
};

const UsernameInput = <T extends { username: string }>({
  form,
  disabled,
}: UsernameInputProps<T>) => {
  return (
    <TextInput
      label="Username"
      placeholder="Your username"
      icon={<AvatarIcon />}
      disabled={disabled}
      required
      {...form.getInputProps(`username`)}
    />
  );
};

export default UsernameInput;
