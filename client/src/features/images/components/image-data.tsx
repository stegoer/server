import { List, Text } from "@mantine/core";

import type { Image } from "@graphql/generated/codegen.generated";

export type ImageDataProps = {
  image: Image;
};

const ImageData = ({ image }: ImageDataProps): JSX.Element => {
  return (
    <Text>
      Image {image.id}
      <List>
        <List.Item>Created: {image.createdAt.toLocaleString()}</List.Item>
        <List.Item>Message: {image.message}</List.Item>
        <List.Item>Least significant bits used: {image.lsbUsed}</List.Item>
        <List.Item>Channel: {image.channel}</List.Item>
      </List>
    </Text>
  );
};

export default ImageData;
