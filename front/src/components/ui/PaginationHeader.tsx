import Stack from "@mui/material/Stack";
import React from "react";

export type PaginationHeaderProps = {
  totalCount: number;
  page: number;
  pageSize: number;
  dataSize: number;
};

const PaginationHeader: React.FC<PaginationHeaderProps> = ({
  totalCount,
  page,
  pageSize,
  dataSize,
}) => {
  const text = `${totalCount}件中 ${(page - 1) * pageSize + 1} - ${
    (page - 1) * pageSize + dataSize
  }件`;

  return <Stack>{text}</Stack>;
};

export default PaginationHeader;
