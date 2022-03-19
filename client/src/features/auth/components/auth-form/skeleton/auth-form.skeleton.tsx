import { Skeleton } from "@mantine/core";

const AuthFormSkeleton = (): JSX.Element => {
  const keys = [0, 1, 2, 3];

  return (
    <>
      {keys.map((key) => (
        <Skeleton
          key={key}
          height={30}
          mb={10}
          visible
        />
      ))}
    </>
  );
};

export default AuthFormSkeleton;
