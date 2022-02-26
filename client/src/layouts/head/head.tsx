import NextHead from "next/head";

import type { FC } from "react";

const Head: FC = () => {
  return (
    <NextHead>
      <title>stegoer</title>
      <link rel="manifest" href="site.webmanifest.json" />
      <link rel="shortcut icon" href="/images/favicon.ico" />
      <link
        rel="apple-touch-icon"
        sizes="180x180"
        href="/images/apple-touch-icon.png"
      />
      <link
        rel="icon"
        type="image/png"
        sizes="32x32"
        href="/images/favicon-32x32.png"
      />
      <link
        rel="icon"
        type="image/png"
        sizes="16x16"
        href="/images/favicon-16x16.png"
      />
    </NextHead>
  );
};

export default Head;
