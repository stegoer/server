import {
  IMAGE_TABLE_HEADERS,
  IMAGE_TABLE_PER_PAGE,
} from "@features/images/images.constants";

import { Skeleton, Table } from "@mantine/core";

const ImageTableSkeleton = (): JSX.Element => {
  const rows = Array.from({ length: IMAGE_TABLE_PER_PAGE })
    .fill(0)
    .map((_, index) => {
      return (
        <tr key={index}>
          <td key={index} colSpan={IMAGE_TABLE_HEADERS.length}>
            <Skeleton key={index} height={10} m={5} animate={false} visible />
          </td>
        </tr>
      );
    });

  return (
    <Table striped>
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

export default ImageTableSkeleton;
