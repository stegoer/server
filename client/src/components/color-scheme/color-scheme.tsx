import ColorSchemeIcon from "./color-scheme-icon";

import {
  Center,
  Kbd,
  Text,
  Tooltip,
  useMantineColorScheme,
} from "@mantine/core";

import type { FC } from "react";

const ColorScheme: FC = () => {
  const scheme = useMantineColorScheme();
  const isDark = scheme.colorScheme === `dark`;

  const label = (
    <Text>
      Toggle color scheme:{` `}
      <Kbd>Ctrl</Kbd> + <Kbd>J</Kbd>
    </Text>
  );

  return (
    <Center inline>
      <Tooltip label={label}>
        <ColorSchemeIcon
          isDark={isDark}
          toggleColorScheme={() => scheme.toggleColorScheme()}
        />
      </Tooltip>
    </Center>
  );
};

export default ColorScheme;
