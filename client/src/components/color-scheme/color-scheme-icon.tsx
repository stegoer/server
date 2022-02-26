import { ActionIcon } from "@mantine/core";
import { MoonIcon, SunIcon } from "@modulz/radix-icons";

import type { ColorScheme } from "@mantine/styles/lib/theme/types";
import type { FC } from "react";

type Props = {
  isDark: boolean;
  toggleColorScheme(colorScheme?: ColorScheme): void;
};

const ColorSchemeIcon: FC<Props> = ({ isDark, toggleColorScheme }) => {
  const [width, height] = [25, 25];

  return (
    <ActionIcon
      variant="light"
      color={isDark ? `yellow` : `blue`}
      onClick={() => toggleColorScheme()}
      size="lg"
    >
      {isDark ? (
        <SunIcon style={{ width, height }} />
      ) : (
        <MoonIcon style={{ width, height }} />
      )}
    </ActionIcon>
  );
};

export default ColorSchemeIcon;
