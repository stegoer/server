import { Anchor, Text } from "@mantine/core";
import { useRouter } from "next/router";
import { useCallback } from "react";

import type { FC } from "react";

type Props = {
  disabled: boolean;
  onClick(): void;
};

const LoginAnchor: FC<Props> = ({
  children,
  disabled,
  onClick: onClickProperty,
}) => {
  const router = useRouter();

  const onClick = useCallback(() => {
    onClickProperty();
    void router.push(`account`);
  }, [onClickProperty, router]);

  return (
    <Anchor
      component="button"
      type="button"
      onClick={onClick}
      disabled={disabled}
    >
      {children || <Text>login</Text>}
    </Anchor>
  );
};

export default LoginAnchor;
