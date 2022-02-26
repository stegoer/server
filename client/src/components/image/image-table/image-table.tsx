import { IMAGE_TABLE_HEADERS } from "@constants/images.constants";

import { Table } from "@mantine/core";

import type { Image } from "@graphql/generated/codegen.generated";
import type { FC } from "react";

type Props = {
  data: Image[];
};

const ImageTable: FC<Props> = ({ data }) => {
  const rows = data.map((image, index) => (
    <tr key={index}>
      <td>{image.channel}</td>
      <td>{image.createdAt}</td>
    </tr>
  ));

  return (
    <Table striped highlightOnHover>
      <thead>
        <tr>
          {IMAGE_TABLE_HEADERS.map((header, index) => (
            <th key={index}>{header}</th>
          ))}
        </tr>
      </thead>
      <tbody>{rows}</tbody>
    </Table>
  );
};

export default ImageTable;
