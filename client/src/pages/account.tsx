import AccountComponent from "@features/account/components/account.component";
import AuthComponent from "@features/auth/components/auth.component";
import useUser from "@hooks/user.hook";
import PageLayout from "@layouts/page.layout";

import { Paper } from "@mantine/core";

import type { NextPage } from "next";

const Account: NextPage = () => {
  const [user] = useUser();

  return (
    <PageLayout title="account">
      <Paper style={{ width: 300, position: `relative` }}>
        {user ? <AccountComponent user={user} /> : <AuthComponent />}
      </Paper>
    </PageLayout>
  );
};

export default Account;
