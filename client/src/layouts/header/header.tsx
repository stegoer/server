import ColorScheme from "@features/color-scheme/components/color-scheme";

import { Header as MantineHeader, Title } from "@mantine/core";
import Link from "next/link";

const Header = (): JSX.Element => {
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
