import type { FC } from "react";

const LoginForm: FC = () => {
  return (
    <form onSubmit={(event) => console.log(event.target)}>
      <label htmlFor="uname">
        <b>Username</b>
      </label>
      <input
        id="name"
        type="text"
        placeholder="Enter Username"
        name="uname"
        required
      />
      <br />
      <label htmlFor="psw">
        <b>Password</b>
      </label>
      <input
        id="password"
        type="password"
        placeholder="Enter Password"
        name="password"
        required
      />
    </form>
  );
};

export default LoginForm;
