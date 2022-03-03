import { PasswordInput as MantinePasswordInput } from "@mantine/core";
import { LockClosedIcon } from "@modulz/radix-icons";

import type { PasswordInputProps } from "@mantine/core/lib/components/PasswordInput/PasswordInput";
import type { UseForm } from "@mantine/hooks/lib/use-form/use-form";

type Props<T extends { confirmPassword: string }> = {
  form: UseForm<T>;
  props?: PasswordInputProps;
};

const ConfirmPasswordInput = <T extends { confirmPassword: string }>({
  form,
  props,
}: Props<T>) => {
  return (
    <MantinePasswordInput
      required
      label="Confirm Password"
      placeholder="Confirm Password"
      toggleTabIndex={0}
      icon={<LockClosedIcon />}
      {...form.getInputProps(`confirmPassword`)}
      {...props}
    />
  );
};

export default ConfirmPasswordInput;
