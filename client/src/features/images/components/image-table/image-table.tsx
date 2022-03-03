import { IMAGE_TABLE_HEADERS } from "@features/images/images.constants";

import { Table } from "@mantine/core";

import type { Image } from "@graphql/generated/codegen.generated";

type Props = {
  data: Image[];
};

const ImageTable = ({ data }: Props): JSX.Element => {
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
