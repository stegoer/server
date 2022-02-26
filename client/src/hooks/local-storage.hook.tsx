import { STORAGE_KEY_PREFIX } from "@constants/storage.constants";

import { useLocalStorageValue as mantineUseLocalStorageValue } from "@mantine/hooks";

const useLocalStorageValue = <T extends string>({
  key,
  defaultValue,
}: {
  key: string;
  defaultValue?: T;
}): readonly [T, (value: T | ((previousState: T) => T)) => void] => {
  return mantineUseLocalStorageValue({
    key: `${STORAGE_KEY_PREFIX}${key}`,
    defaultValue,
  });
};

export default useLocalStorageValue;
