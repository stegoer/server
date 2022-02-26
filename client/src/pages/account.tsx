import AccountView from "@components/account/account-view/account-view";
import AuthView from "@components/account/auth-view/auth-view";
import useUser from "@hooks/user.hook";

import { Paper } from "@mantine/core";

import type { NextPage } from "next";

const Account: NextPage = () => {
  const [user] = useUser();

  return (
    <Paper style={{ width: 300, position: `relative` }}>
      {user ? <AccountView user={user} /> : <AuthView />}
    </Paper>
  );
};

export default Account;
