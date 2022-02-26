import ColorScheme from "@components/color-scheme/color-scheme";

import { Header as MantineHeader, Title } from "@mantine/core";
import Link from "next/link";

import type { FC } from "react";

const Header: FC = () => {
  return (
    <MantineHeader height={55} padding="xs">
      <div>
        <div style={{ float: `left` }}>
          <Link href="/" passHref>
            <a>
              <Title>stegoer</Title>
            </a>
          </Link>
        </div>
        <div style={{ float: `right` }}>
          <ColorScheme />
        </div>
      </div>
    </MantineHeader>
  );
};

export default Header;
