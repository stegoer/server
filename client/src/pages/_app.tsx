import useColorScheme from "@hooks/color-cheme.hook";
import Head from "@layouts/head/head";
import Header from "@layouts/header/header";
import Navbar from "@layouts/navbar/navbar";
import AuthProvider from "@providers/auth.provider";
import ColorSchemeProvider from "@providers/color-scheme.provider";
import GraphqlProvider from "@providers/graphql.provider";
import UserProvider from "@providers/user.provider";
import "@styles/base/globals.style.css";

import { AppShell, MantineProvider } from "@mantine/core";
import { NotificationsProvider } from "@mantine/notifications";

import type { NextComponentType } from "next";
import type { AppContext, AppInitialProps, AppProps } from "next/app";

const App: NextComponentType<AppContext, AppInitialProps, AppProps> = ({
  Component,
  pageProps,
}: AppProps) => {
  const [colorScheme, toggleColorScheme] = useColorScheme();

  return (
    <>
      <Head />
      <UserProvider>
        <GraphqlProvider>
          <AuthProvider>
            <ColorSchemeProvider
              colorScheme={colorScheme}
              toggleColorScheme={toggleColorScheme}
            >
              <MantineProvider
                withGlobalStyles
                withNormalizeCSS
                theme={{ colorScheme }}
              >
                <NotificationsProvider limit={3}>
                  <AppShell
                    padding="xl"
                    navbar={<Navbar />}
                    header={<Header />}
                  >
                    <Component {...pageProps} />
                  </AppShell>
                </NotificationsProvider>
              </MantineProvider>
            </ColorSchemeProvider>
          </AuthProvider>
        </GraphqlProvider>
      </UserProvider>
    </>
  );
};

export default App;
