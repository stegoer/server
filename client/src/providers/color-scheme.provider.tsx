import { ColorSchemeProvider as MantineColorSchemeProvider } from "@mantine/core";
import { useHotkeys } from "@mantine/hooks";

import type { ColorScheme } from "@mantine/core";
import type { FC } from "react";

type Props = {
  colorScheme: ColorScheme;
  toggleColorScheme(colorScheme?: ColorScheme): void;
};

const ColorSchemeProvider: FC<Props> = ({
  children,
  colorScheme,
  toggleColorScheme,
}) => {
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
