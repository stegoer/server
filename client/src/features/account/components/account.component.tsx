import AccountNavigation from "@features/account/components/account.navigation";
import UpdateModal from "@features/account/components/modals/update.modal";
import UserData from "@features/account/components/user-data";

import { useState } from "react";

import type { User } from "@graphql/generated/codegen.generated";

type Props = {
  user: User;
};

const AccountComponent = ({ user }: Props): JSX.Element => {
  const [modelOpened, setModalOpened] = useState(false);

  return (
    <>
      <UserData user={user} />
      <UpdateModal
        user={user}
        opened={modelOpened}
        setOpened={setModalOpened}
      />
      <AccountNavigation
        user={user}
        disabled={modelOpened}
        onUpdate={() => setModalOpened(true)}
      />
    </>
  );
};

export default AccountComponent;
