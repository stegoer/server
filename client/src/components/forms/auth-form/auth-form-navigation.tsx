import SubmitButton from "@components/buttons/submit.button";
import AuthLink from "@components/links/auth.link";

import { Group } from "@mantine/core";

import type { FormType } from "@custom-types//account.types";
import type { FC } from "react";

type Props = {
  formType: FormType;
  loading: boolean;
  title: string;
  onToggle(): void;
};

const AuthFormNavigation: FC<Props> = ({
  formType,
  loading,
  title,
  onToggle,
}) => {
  return (
    <Group position="apart" mt="xl">
      <AuthLink
        formType={formType}
        toggleFormType={onToggle}
        disabled={loading}
      />
      <SubmitButton disabled={loading}>{title}</SubmitButton>
    </Group>
  );
};

export default AuthFormNavigation;
