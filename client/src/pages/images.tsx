import ImageTableComponent from "@features/image-table/components/image-table.component";
import ImageTableSkeleton from "@features/image-table/components/image-table.skeleton";
import useUser from "@hooks/user.hook";
import loginRequiredNotification from "@layouts/navbar/login-required.notification";
import PageLayout from "@layouts/page.layout";

import { useNotifications } from "@mantine/notifications";
import { useEffect } from "react";

import type { NextPage } from "next";

const NOTIFICATION_TIMER = 2000; // 2 seconds

const Images: NextPage = () => {
  const notifications = useNotifications();
  const [user] = useUser();

  // 2 seconds after mount show notification if user is not logged in
  useEffect(() => {
    const timer = setTimeout(() => {
      if (!user) {
        notifications.showNotification(
          loginRequiredNotification(`Images`, `/images`),
        );
      }
    }, NOTIFICATION_TIMER);
    return () => clearTimeout(timer);
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  return (
    <PageLayout title="images">
      {user ? <ImageTableComponent /> : <ImageTableSkeleton />}
    </PageLayout>
  );
};

export default Images;
