import useColorScheme from "@/hooks/color-cheme.hook";
import Head from "@/layouts/head/head";
import Header from "@/layouts/header/header";
import Navbar from "@/layouts/navbar/navbar";
import ColorSchemeProvider from "@/providers/color-scheme.provider";
import GraphqlProvider from "@/providers/graphql.provider";
import "@/styles/globals.style.css";

import { AppShell, MantineProvider } from "@mantine/core";
import { useHotkeys } from "@mantine/hooks";
import React from "react";

import type { NextComponentType } from "next";
import type { AppContext, AppInitialProps, AppProps } from "next/app";

const App: NextComponentType<AppContext, AppInitialProps, AppProps> = ({
  Component,
  pageProps,
}: AppProps) => {
  const [colorScheme, toggleColorScheme] = useColorScheme();

  useHotkeys([[`mod+J`, () => toggleColorScheme()]]);

  return (
    <>
      <Head />
      <GraphqlProvider>
        <ColorSchemeProvider
          colorScheme={colorScheme}
          toggleColorScheme={toggleColorScheme}
        >
          <MantineProvider
            withGlobalStyles
            withNormalizeCSS
            theme={{ colorScheme }}
          >
            <AppShell padding="xl" navbar={<Navbar />} header={<Header />}>
              <Component {...pageProps} />
            </AppShell>
          </MantineProvider>
        </ColorSchemeProvider>
      </GraphqlProvider>
    </>
  );
};

export default App;
