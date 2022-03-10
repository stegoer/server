import DecodeImagesComponent from "@features/images/components/decode-images.component";
import PageLayout from "@layouts/page.layout";

import type { NextPage } from "next";

const Decode: NextPage = () => {
  return (
    <PageLayout title="decode">
      <DecodeImagesComponent />
    </PageLayout>
  );
};

export default Decode;
