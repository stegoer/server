const LocalStorageService = {
  get(key: string): string | undefined {
    if (typeof localStorage !== `undefined`) {
      return localStorage.getItem(key) || undefined;
    }
  },
  set(key: string, value: string) {
    if (typeof localStorage !== `undefined`) {
      localStorage.setItem(key, value);
    }
  },
  remove(key: string) {
    if (typeof localStorage !== `undefined`) {
      localStorage.removeItem(key);
    }
  },
  clear() {
    if (typeof localStorage !== `undefined`) {
      localStorage.clear();
    }
  },
};

export default LocalStorageService;
