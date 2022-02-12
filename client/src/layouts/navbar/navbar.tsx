import { Navbar as MantineNavbar } from "@mantine/core";
import Link from "next/link";
import React from "react";

import type { FC } from "react";

const Navbar: FC = () => {
  return (
    <MantineNavbar padding="xs" width={{ base: 100 }}>
      <MantineNavbar.Section>
        <Link href="/login">
          <a>Login</a>
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
        <Link href="/images">
          <a>Images</a>
        </Link>
      </MantineNavbar.Section>
    </MantineNavbar>
  );
};

export default Navbar;
