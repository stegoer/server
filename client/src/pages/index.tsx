import PageLayout from "@layouts/page.layout";

import { Title } from "@mantine/core";

import type { NextPage } from "next";

const Home: NextPage = () => {
  return (
    <PageLayout title="home">
      <Title>Home</Title>
    </PageLayout>
  );
};

export default Home;
