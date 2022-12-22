import ChevronLeftIcon from "@heroicons/react/20/solid/ChevronLeftIcon";
import ChevronRightIcon from "@heroicons/react/20/solid/ChevronRightIcon";
import React, { FC, useEffect, useState } from "react";
import PaginationNavButton from "@/components/ui/PaginationNavButton";

type PaginationNavProps = {
  page: number;
  count: number;
  onChange: (e: React.ChangeEvent<unknown>, page: number) => void;
};

type Page = {
  label: string;
  current: boolean;
  data?: number;
};

const BUTTON_LENGTH = 7;

const PaginationNav: FC<PaginationNavProps> = ({ page, count, onChange }) => {
  const [pages, setPages] = useState<Page[]>([]);

  const onClickHandler = (e: React.ChangeEvent<unknown>, page: number) => {
    console.log("page", page);
    onChange(e, page);
  };

  useEffect(() => {
    const nextPages: Page[] = [];
    nextPages.push({
      label: "1",
      current: page === 1,
      data: 1,
    });

    if (count <= BUTTON_LENGTH) {
      for (let i = 2; i <= count; i++) {
        nextPages.push({
          label: i.toString(),
          current: page === i,
          data: i,
        });
      }
    } else if (page <= 4) {
      // 1 2 3 4 5 6 7
      for (let i = 2; i <= 5; i++) {
        nextPages.push({
          label: i.toString(),
          current: page === i,
          data: i,
        });
      }
      nextPages.push({
        label: "...",
        current: false,
      });
    } else if (page > count - 4) {
      // 1 ... 4 5 6 7 8
      nextPages.push({
        label: "...",
        current: false,
      });
      for (let i = count - 4; i < count; i++) {
        nextPages.push({
          label: i.toString(),
          current: page === i,
          data: i,
        });
      }
    } else {
      // 1 ... 4 5 6  ...8
      nextPages.push({
        label: "...",
        current: false,
      });
      nextPages.push({
        label: (page - 1).toString(),
        current: false,
        data: page - 1,
      });

      nextPages.push({
        label: page.toString(),
        current: true,
        data: page,
      });
      nextPages.push({
        label: (page + 1).toString(),
        current: false,
        data: page + 1,
      });

      nextPages.push({
        label: "...",
        current: false,
      });
    }

    nextPages.push({
      label: count.toString(),
      current: page === count,
      data: count,
    });
    setPages(nextPages);
  }, [page, count]);

  const prevPage = page - 1;
  const nextPage = page + 1;

  return (
    <div className="flex items-center justify-between border-t border-gray-200 bg-white px-4 py-3 sm:px-6">
      <div className="flex flex-1 items-center justify-between">
        <div>
          <nav
            className="isolate inline-flex -space-x-px rounded-md shadow-sm"
            aria-label="Pagination"
          >
            <PaginationNavButton
              current={false}
              round="left"
              data={Math.max(prevPage, 1)}
              onClick={onClickHandler}
            >
              <span className="sr-only">Previous</span>
              <ChevronLeftIcon className="w-5 h-5" />
            </PaginationNavButton>

            {pages.map((page, index) => {
              return (
                <PaginationNavButton
                  key={index}
                  current={page.current}
                  data={page.data}
                  onClick={onClickHandler}
                >
                  {page.label}
                </PaginationNavButton>
              );
            })}

            <PaginationNavButton
              current={false}
              round="right"
              data={Math.min(nextPage, count)}
              onClick={onClickHandler}
            >
              <span className="sr-only">Next</span>
              <ChevronRightIcon className="w-5 h-5" />
            </PaginationNavButton>
          </nav>
        </div>
      </div>
    </div>
  );
};
export default PaginationNav;
