import type { Image } from "@graphql/generated/codegen.generated";
import type { FC } from "react";

type Props = {
  data: Image;
};

const DisplayImage: FC<Props> = ({ data }) => {
  return (
    <div>
      <h3>Image {data.id}</h3>
      <h4>Created at: {data.createdAt.toLocaleString()}</h4>
      <h4>Channel: {data.channel}</h4>
    </div>
  );
};

export default DisplayImage;
