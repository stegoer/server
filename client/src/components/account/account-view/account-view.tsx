import AccountViewNavigation from "@components/account/account-view/account-view-navigation";
import UpdateModal from "@components/account/account-view/modals/update.modal";
import UserData from "@components/account/account-view/user-data";

import { Title } from "@mantine/core";
import { useState } from "react";

import type { User } from "@graphql/generated/codegen.generated";
import type { FC } from "react";

type Props = {
  user: User;
};

const AccountView: FC<Props> = ({ user }) => {
  const [modelOpened, setModalOpened] = useState(false);

  return (
    <>
      <Title>Account</Title>
      <UserData user={user} />
      <UpdateModal
        user={user}
        opened={modelOpened}
        setOpened={setModalOpened}
      />
      <AccountViewNavigation
        user={user}
        disabled={modelOpened}
        onUpdate={() => setModalOpened(true)}
      />
    </>
  );
};

export default AccountView;
