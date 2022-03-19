import LoginRequiredLink from "@layouts/navbar/login-required.link";

import { Navbar as MantineNavbar } from "@mantine/core";
import Link from "next/link";

const Navbar = (): JSX.Element => {
  return (
    <MantineNavbar
      padding="xs"
      width={{ base: 100 }}
    >
      <MantineNavbar.Section>
        <Link href="/account">
          <a>account</a>
        </Link>
      </MantineNavbar.Section>
      <MantineNavbar.Section>
        <Link href="/encode">
          <a>encode</a>
        </Link>
      </MantineNavbar.Section>
      <MantineNavbar.Section>
        <Link href="/decode">
          <a>decode</a>
        </Link>
      </MantineNavbar.Section>
      <MantineNavbar.Section>
        <LoginRequiredLink to="/images">images</LoginRequiredLink>
      </MantineNavbar.Section>
    </MantineNavbar>
  );
};

export default Navbar;
