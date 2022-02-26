import { Box, Text } from "@mantine/core";
import { CheckIcon, Cross1Icon } from "@modulz/radix-icons";

import type { FC } from "react";

type Props = {
  meets: boolean;
  label: string;
};

const PasswordRequirement: FC<Props> = ({ meets, label }) => {
  return (
    <Text
      color={meets ? `teal` : `red`}
      sx={{ display: `flex`, alignItems: `center` }}
      mt={7}
      size="sm"
    >
      {meets ? <CheckIcon /> : <Cross1Icon />} <Box ml={10}>{label}</Box>
    </Text>
  );
};

export default PasswordRequirement;
