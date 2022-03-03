import { ActionIcon } from "@mantine/core";
import { MoonIcon, SunIcon } from "@modulz/radix-icons";

import type { ColorScheme } from "@mantine/styles/lib/theme/types";

type Props = {
  isDark: boolean;
  toggleColorScheme(colorScheme?: ColorScheme): void;
};

const ColorSchemeIcon = ({
  isDark,
  toggleColorScheme,
}: Props): JSX.Element => {
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
