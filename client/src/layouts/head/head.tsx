import NextHead from "next/head";

export type HeadProps = {
  title?: string;
};

const BASE_TITLE = `stegoer`;

const Head = ({ title }: HeadProps): JSX.Element => {
  return (
    <NextHead>
      <title>{title ? `${BASE_TITLE} | ${title}` : BASE_TITLE}</title>
      <link
        rel="manifest"
        href="site.webmanifest.json"
      />
      <link
        rel="shortcut icon"
        href="/images/favicon.ico"
      />
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
