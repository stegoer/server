type Requirement = {
  re: RegExp;
  label: string;
};

export const Requirements: Requirement[] = [
  { re: /\d/, label: `Includes number` },
  { re: /[a-z]/, label: `Includes lowercase letter` },
  { re: /[A-Z]/, label: `Includes uppercase letter` },
  { re: /[!#$%&'()*+,.:;<=>?@^|-]/, label: `Includes special symbol` },
];

export const calculateStrength = (password: string) => {
  let multiplier = password.length > 5 ? 0 : 1;

  for (const requirement of Requirements) {
    if (!requirement.re.test(password)) {
      multiplier += 1;
    }
  }

  return Math.max(100 - (100 / (Requirements.length + 1)) * multiplier, 10);
};
