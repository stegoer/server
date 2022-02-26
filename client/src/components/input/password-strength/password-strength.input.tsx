import { calculateStrength, Requirements } from "./constants";

import PasswordRequirement from "@components/input/password-strength/password-requirement";
import PasswordInput from "@components/input/password.input";

import { Popover, Progress } from "@mantine/core";
import { useState } from "react";

import type { PasswordInputProps } from "@mantine/core/lib/components/PasswordInput/PasswordInput";
import type { UseForm } from "@mantine/hooks/lib/use-form/use-form";

type Props<T> = {
  form: UseForm<{ password: string } & T>;
  inputProps?: PasswordInputProps;
};

const PasswordStrength = <T,>({ form, inputProps }: Props<T>) => {
  const [popoverOpened, setPopoverOpened] = useState(false);
  const [password, setPassword] = useState(``);

  const requirements = Requirements.map((requirement, index) => (
    <PasswordRequirement
      key={index}
      label={requirement.label}
      meets={requirement.re.test(password)}
    />
  ));

  const strength = calculateStrength(password);
  const color = strength === 100 ? `teal` : strength > 50 ? `yellow` : `red`;

  return (
    <Popover
      opened={popoverOpened}
      position="bottom"
      placement="start"
      withArrow
      styles={{ popover: { width: `100%` } }}
      noFocusTrap
      transition="pop-top-left"
      onFocusCapture={() => setPopoverOpened(true)}
      onBlurCapture={() => setPopoverOpened(false)}
      target={
        <PasswordInput
          form={form}
          props={{
            description: `Password should include at least 1 lowercase letter, 1 uppercase letter, 1 number and 1 special symbol`,
            value: password,
            onChange: (event) => {
              setPassword(event.currentTarget.value.trim());
              // eslint-disable-next-line @typescript-eslint/ban-ts-comment
              // @ts-ignore
              form.setFieldValue(`password`, event.currentTarget.value.trim());
            },
            ...inputProps,
          }}
        />
      }
    >
      <Progress
        color={color}
        value={strength}
        size={5}
        style={{ marginBottom: 10 }}
      />
      <PasswordRequirement
        label="Includes at least 6 characters"
        meets={password.length > 5}
      />
      {requirements}
    </Popover>
  );
};

export default PasswordStrength;
