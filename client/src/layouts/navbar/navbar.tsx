import LoginRequiredLink from "@components/links/login-required.link";

import { Navbar as MantineNavbar } from "@mantine/core";
import Link from "next/link";

import type { FC } from "react";

const Navbar: FC = () => {
  return (
    <MantineNavbar padding="xs" width={{ base: 100 }}>
      <MantineNavbar.Section>
        <Link href="/account">
          <a>Account</a>
        </Link>
      </MantineNavbar.Section>
      <MantineNavbar.Section>
        <Link href="/encode">
          <a>Encode</a>
        </Link>
      </MantineNavbar.Section>
      <MantineNavbar.Section>
        <Link href="/decode">
          <a>Decode</a>
        </Link>
      </MantineNavbar.Section>
      <MantineNavbar.Section>
        <LoginRequiredLink to="/images">Images</LoginRequiredLink>
      </MantineNavbar.Section>
    </MantineNavbar>
  );
};

export default Navbar;
