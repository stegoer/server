import type { FC } from "react";

const LoginButton: FC = () => {
  return <button onClick={(...data) => console.log(...data)}>Login</button>;
};

export default LoginButton;
