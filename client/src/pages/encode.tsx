import EncodeImagesComponent from "@features/images/components/encode-images.component";
import PageLayout from "@layouts/page.layout";

import type { NextPage } from "next";

const Encode: NextPage = () => {
  return (
    <PageLayout title="encode">
      <EncodeImagesComponent />
    </PageLayout>
  );
};

export default Encode;
