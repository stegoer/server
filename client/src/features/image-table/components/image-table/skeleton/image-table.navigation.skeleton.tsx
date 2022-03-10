import { ActionIcon, Group } from "@mantine/core";
import { ArrowLeftIcon, ArrowRightIcon } from "@modulz/radix-icons";

const ImageTableNavigationSkeleton = (): JSX.Element => {
  return (
    <Group>
      <ActionIcon disabled>
        <ArrowLeftIcon width={25} height={25} />
      </ActionIcon>
      <span>0</span>
      <ActionIcon disabled>
        <ArrowRightIcon width={25} height={25} />
      </ActionIcon>
    </Group>
  );
};

export default ImageTableNavigationSkeleton;
