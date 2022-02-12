import { ActionIcon } from "@mantine/core";
import { MoonIcon, SunIcon } from "@modulz/radix-icons";

import type { ColorScheme } from "@mantine/styles/lib/theme/types";
import type { FC } from "react";

type Properties = {
  isDark: boolean;
  toggleColorScheme(colorScheme?: ColorScheme): void;
};

const ColorSchemeIcon: FC<Properties> = ({ isDark, toggleColorScheme }) => {
  const [width, height] = [25, 25];

  return (
    <ActionIcon
      variant="light"
      color={isDark ? `yellow` : `blue`}
      onClick={() => toggleColorScheme()}
      title="Toggle color scheme"
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
