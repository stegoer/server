import SubmitButton from "@components/buttons/submit.button";
import AuthLink from "@features/auth/components/auth.link";

import { Group } from "@mantine/core";

import type { FormType } from "@features/auth/auth.types";

type Props = {
  formType: FormType;
  loading: boolean;
  title: string;
  onToggle(): void;
};

const AuthFormNavigation = ({
  formType,
  loading,
  title,
  onToggle,
}: Props): JSX.Element => {
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
