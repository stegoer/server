export type FormType = `login` | `register`;

export const isFormType = (s: string): s is FormType => {
  return [`login`, `register`].includes(s);
};

export type AuthState = {
  token?: string;
};
