import useColorScheme from "@hooks/color-cheme.hook";
import Header from "@layouts/header/header";
import Navbar from "@layouts/navbar/navbar";
import AuthProvider from "@providers/auth.provider";
import ColorSchemeProvider from "@providers/color-scheme.provider";
import GraphqlProvider from "@providers/graphql.provider";
import UserProvider from "@providers/user.provider";

import { AppShell, MantineProvider } from "@mantine/core";
import { NotificationsProvider } from "@mantine/notifications";

import type { PropsWithChildren } from "react";

type Props = PropsWithChildren<Record<never, never>>;

const AppProvider = ({ children }: Props): JSX.Element => {
  const [colorScheme, toggleColorScheme] = useColorScheme();

  return (
    <>
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
                    {children}
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

export default AppProvider;
