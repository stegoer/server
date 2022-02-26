import LogoutButton from "@components/buttons/logout.button";

import { Button, Group } from "@mantine/core";

import type { User } from "@graphql/generated/codegen.generated";
import type { FC } from "react";

type Props = {
  user: User;
  disabled: boolean;
  onUpdate(): void;
};

const AccountViewNavigation: FC<Props> = ({ user, disabled, onUpdate }) => {
  return (
    <Group>
      <Button onClick={onUpdate} disabled={disabled}>
        Update Account
      </Button>
      <LogoutButton user={user} disabled={disabled} />
    </Group>
  );
};

export default AccountViewNavigation;
