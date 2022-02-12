import ColorSchemeIcon from "./color-scheme-icon";

import { Center, useMantineColorScheme } from "@mantine/core";

import type { FC } from "react";

const ColorScheme: FC = () => {
  // eslint-disable-next-line @typescript-eslint/unbound-method
  const { colorScheme, toggleColorScheme } = useMantineColorScheme();
  const isDark = colorScheme === `dark`;

  return (
    <Center inline>
      <ColorSchemeIcon isDark={isDark} toggleColorScheme={toggleColorScheme} />
    </Center>
  );
};

export default ColorScheme;
