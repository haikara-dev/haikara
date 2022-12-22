import { FC } from "react";

export type PaginationNavButtonProps = {
  current: boolean;
  round?: "left" | "right" | "none";
  children: React.ReactNode;
  data?: number;
  onClick?: (e: React.ChangeEvent<unknown>, page: number) => void;
};
const PaginationNavButton: FC<PaginationNavButtonProps> = ({
  current,
  round = "none",
  children,
  data,
  onClick,
}) => {
  //  rounded-l-md

  const roundClass =
    round === "left" ? "rounded-l-md" : round === "right" ? "rounded-r-md" : "";

  const colorClass = current
    ? "border-indigo-500 bg-indigo-50"
    : "border-gray-300 bg-white";

  const onClickHandler = (e: React.ChangeEvent<unknown>) => {
    e.preventDefault();
    if (onClick) {
      onClick(e, data!);
    }
  };

  return (
    <>
      {current ? (
        <a
          href="#"
          onClick={onClickHandler}
          className={`relative items-center border ${colorClass} px-4 py-2 text-sm font-medium text-gray-500 hover:bg-gray-50 focus:z-20 ${roundClass}`}
        >
          {children}
        </a>
      ) : (
        <a
          href="#"
          onClick={onClickHandler}
          className={`relative items-center border ${colorClass} px-4 py-2 text-sm font-medium text-gray-500 hover:bg-gray-50 focus:z-20 ${roundClass}`}
        >
          {children}
        </a>
      )}
    </>
  );
};

export default PaginationNavButton;
