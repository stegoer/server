import Head from "@layouts/head/head";

import type { PropsWithChildren } from "react";

type Props = PropsWithChildren<{
  title?: string;
}>;

const PageLayout = ({ children, title }: Props): JSX.Element => {
  return (
    <>
      <Head title={title} />
      {children}
    </>
  );
};

export default PageLayout;
