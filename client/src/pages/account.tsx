import AccountComponent from "@features/account/components/account.component";
import AuthComponent from "@features/auth/components/auth.component";
import useUser from "@hooks/user.hook";
import PageLayout from "@layouts/page.layout";

import { useState } from "react";

import type { NextPage } from "next";

const DEFAULT_TITLE = `account`;

const Account: NextPage = () => {
  const [user] = useUser();
  const [title, setTitle] = useState(DEFAULT_TITLE);

  return (
    <PageLayout title={title}>
      {user ? (
        <AccountComponent user={user} />
      ) : (
        <AuthComponent setTitle={setTitle} />
      )}
    </PageLayout>
  );
};

export default Account;
