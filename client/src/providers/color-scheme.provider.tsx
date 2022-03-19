import { ColorSchemeProvider as MantineColorSchemeProvider } from "@mantine/core";
import { useHotkeys } from "@mantine/hooks";

import type { ColorScheme } from "@mantine/core";
import type { PropsWithChildren } from "react";

export type ColorSchemeProviderProps = PropsWithChildren<{
  colorScheme: ColorScheme;
  toggleColorScheme(colorScheme?: ColorScheme): void;
}>;

const ColorSchemeProvider = ({
  children,
  colorScheme,
  toggleColorScheme,
}: ColorSchemeProviderProps): JSX.Element => {
  useHotkeys([[`mod+J`, () => toggleColorScheme()]]);

  return (
    <MantineColorSchemeProvider
      colorScheme={colorScheme}
      toggleColorScheme={toggleColorScheme}
    >
      {children}
    </MantineColorSchemeProvider>
  );
};

export default ColorSchemeProvider;
