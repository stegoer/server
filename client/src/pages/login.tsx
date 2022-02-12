import LoginButton from "@/components/login/login-button";
import LoginForm from "@/components/login/login-form";

import { Title } from "@mantine/core";

import type { NextPage } from "next";

const Login: NextPage = () => {
  return (
    <>
      <Title>Login</Title>
      <div>
        <LoginForm />
        <LoginButton />
      </div>
    </>
  );
};

export default Login;
