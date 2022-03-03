import LogoutButton from "@features/account/components/logout.button";

import { Button, Group } from "@mantine/core";

import type { User } from "@graphql/generated/codegen.generated";

type Props = {
  user: User;
  disabled: boolean;
  onUpdate(): void;
};

const AccountNavigation = ({
  user,
  disabled,
  onUpdate,
}: Props): JSX.Element => {
  return (
    <Group>
      <Button onClick={onUpdate} disabled={disabled}>
        Update Account
      </Button>
      <LogoutButton user={user} disabled={disabled} />
    </Group>
  );
};

export default AccountNavigation;
