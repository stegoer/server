import { ColorSchemeProvider as MantineColorSchemeProvider } from "@mantine/core";

import type { ColorScheme } from "@mantine/core";
import type { FC } from "react";

type Properties = {
  colorScheme: ColorScheme;
  toggleColorScheme(colorScheme?: ColorScheme): void;
};

const ColorSchemeProvider: FC<Properties> = ({
  children,
  colorScheme,
  toggleColorScheme,
}) => {
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
