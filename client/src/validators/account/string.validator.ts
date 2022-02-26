const stringValidator = (length: number) => (string: string) => {
  return string.trim().length >= length;
};

export default stringValidator;
