import {
  useColorScheme as mantineUseColorScheme,
  useLocalStorageValue,
} from "@mantine/hooks";
import { useCallback } from "react";

import type { ColorScheme } from "@mantine/core";

const useColorScheme = (): [ColorScheme, (value?: ColorScheme) => void] => {
  const [colorScheme, setColorScheme] = useLocalStorageValue<ColorScheme>({
    key: `mantine-color-scheme`,
    defaultValue: mantineUseColorScheme(),
  });

  const toggleColorScheme = useCallback(
    (value?: ColorScheme) => {
      setColorScheme(value || (colorScheme === `dark` ? `light` : `dark`));
    },
    [colorScheme, setColorScheme],
  );

  return [colorScheme, toggleColorScheme];
};

export default useColorScheme;
