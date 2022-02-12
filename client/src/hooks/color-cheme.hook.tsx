import {
  useColorScheme as mantineUseColorScheme,
  useLocalStorageValue,
} from "@mantine/hooks";

import type { ColorScheme } from "@mantine/core";

const useColorScheme = (): [ColorScheme, (value?: ColorScheme) => void] => {
  const [colorScheme, setColorScheme] = useLocalStorageValue<ColorScheme>({
    key: `mantine-color-scheme`,
    defaultValue: mantineUseColorScheme(),
  });

  const toggleColorScheme = (value?: ColorScheme) =>
    setColorScheme(value || (colorScheme === `dark` ? `light` : `dark`));

  return [colorScheme, toggleColorScheme];
};

export default useColorScheme;
