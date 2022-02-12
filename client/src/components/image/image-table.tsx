import { Table } from "@mantine/core";

import type { Image } from "@/graphql/generated/codegen.generated";
import type { FC } from "react";

type Properties = {
  data: Image[];
};

const ImageTable: FC<Properties> = ({ data }) => {
  const rows = data.map((image, index) => (
    <tr key={index}>
      <td>{image.channel}</td>
      <td>{image.createdAt.toLocaleString()}</td>
    </tr>
  ));

  return (
    <Table>
      <thead>
        <tr>
          <th>Channel</th>
          <th>Created</th>
        </tr>
      </thead>
      <tbody>{rows}</tbody>
    </Table>
  );
};

export default ImageTable;
