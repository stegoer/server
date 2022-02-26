import ImageSkeletonView from "@components/image/image-skeleton-view";
import ImageView from "@components/image/image-view";
import useUser from "@hooks/user.hook";

import { Title } from "@mantine/core";

import type { NextPage } from "next";

const Images: NextPage = () => {
  const [user] = useUser();

  return (
    <>
      <Title>Images</Title>
      {user ? <ImageView /> : <ImageSkeletonView />}
    </>
  );
};

export default Images;
