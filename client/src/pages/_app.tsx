import AppProvider from "@providers/app.provider";
import "@styles/globals.style.css";

import type { NextComponentType } from "next";
import type { AppContext, AppInitialProps, AppProps } from "next/app";

const App: NextComponentType<AppContext, AppInitialProps, AppProps> = ({
  Component,
  pageProps,
}: AppProps) => {
  return (
    <AppProvider>
      <Component {...pageProps} />
    </AppProvider>
  );
};

export default App;
