import { STORAGE_KEY_PREFIX } from "@constants/storage.constants";

const LocalStorageService = {
  get(key: string): string | undefined {
    if (typeof localStorage !== `undefined`) {
      return localStorage.getItem(`${STORAGE_KEY_PREFIX}${key}`) || undefined;
    }
  },
  set(key: string, value: string) {
    if (typeof localStorage !== `undefined`) {
      localStorage.setItem(`${STORAGE_KEY_PREFIX}${key}`, value);
    }
  },
  remove(key: string) {
    if (typeof localStorage !== `undefined`) {
      localStorage.removeItem(`${STORAGE_KEY_PREFIX}${key}`);
    }
  },
  clear() {
    if (typeof localStorage !== `undefined`) {
      localStorage.clear();
    }
  },
};

export default LocalStorageService;
