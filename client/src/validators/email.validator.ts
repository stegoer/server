const emailValidator = (email: string): boolean => {
  return /^\S+@\S+\.\S+$/.test(email.trim());
};

export default emailValidator;
