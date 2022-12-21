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

  return <div>{text}</div>;
};

export default PaginationHeader;
