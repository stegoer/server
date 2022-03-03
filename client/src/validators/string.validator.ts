const lengthValidator = (length: number) => (string: string) => {
  return string.length >= length;
};

export default lengthValidator;
