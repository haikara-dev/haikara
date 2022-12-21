import { FC } from "react";

export type PaginationNavButtonProps = {
  label: string;
  current: boolean;
  round?: "left" | "right" | "none";
};
const PaginationNavButton: FC<PaginationNavButtonProps> = ({
  label,
  current,
  round = "none",
}) => {
  //  rounded-l-md

  return (
    <>
      {current ? (
        <a
          href="#"
          className={`relative items-center border border-indigo-500 bg-indigo-50 px-4 py-2 text-sm font-medium text-gray-500 hover:bg-gray-50 focus:z-20`}
        >
          {label}
        </a>
      ) : (
        <a
          href="#"
          className={`relative items-center border border-gray-300 bg-white px-4 py-2 text-sm font-medium text-gray-500 hover:bg-gray-50 focus:z-20`}
        >
          {label}
        </a>
      )}
    </>
  );
};

export default PaginationNavButton;
