import React, { ReactElement, useEffect, useState } from "react";
import DefaultLayout from "@/components/layouts/DefaultLayout";
import { NextPageWithLayout } from "@/pages/_app";

import PaginationHeader from "@/components/ui/PaginationHeader";

import { GetArticlesArg, useGetArticlesQuery } from "@/services/userApi";
import { useRouter } from "next/router";
import StyledSiteName from "@/components/site/StyledSiteName";
import ImgproxyImage from "@/components/ImgproxyImage";
import PaginationNav from "@/components/ui/PaginationNav";

const Home: NextPageWithLayout = () => {
  const buildQuery = (): GetArticlesArg => {
    const state: GetArticlesArg = {
      page: router.query.page ? parseInt(router.query.page.toString()) : 1,
    };
    if (router.query.site_id) {
      state.site_id = parseInt(router.query.site_id.toString());
    }
    return state;
  };
  const router = useRouter();
  const [query, setQuery] = useState<GetArticlesArg>(buildQuery());

  const {
    isLoading,
    data: articles = {
      totalCount: 0,
      totalPage: 1,
      pageSize: 10,
      data: [],
    },
  } = useGetArticlesQuery(query);

  const handleChangePagination = (
    e: React.ChangeEvent<unknown>,
    page: number
  ) => {
    router.push({ query: { ...query, page } });
  };

  useEffect(() => {
    setQuery(buildQuery());
  }, [router]);
  return (
    <div>
      <h1 className="text-center p-12 text-6xl">haikara</h1>

      {isLoading ? (
        <div>loading...</div>
      ) : (
        <div className={"gap-3 flex flex-col items-center"}>
          <PaginationHeader
            totalCount={articles.totalCount}
            page={query.page!}
            pageSize={articles.pageSize}
            dataSize={articles.data.length}
          />

          <div className="flex flex-wrap gap-6 justify-center">
            {articles.data.map((article) => {
              return (
                <div
                  key={article.id}
                  className="max-w-sm bg-gray-800 border border-gray-700 rounded-lg overflow-hidden shadow-lg"
                >
                  <div className="flex flex-col justify-between h-full pb-2">
                    {article.ogp_image_url && (
                      <div className="w-full h-auto relative">
                        <ImgproxyImage
                          src={article.ogp_image_url}
                          width={360 * 2}
                          height={189 * 2}
                          objectFit="contain"
                          className="bg-white"
                        />
                      </div>
                    )}
                    <div className="p-4 grow">
                      <div>
                        <h2>{article.title}</h2>
                        <div>
                          {new Date(article.published_at).toLocaleDateString()}
                        </div>
                        <div>
                          <a
                            href={article.site.url}
                            target="_blank"
                            rel="noopener noreferrer"
                          >
                            <StyledSiteName>{article.site.name}</StyledSiteName>
                          </a>
                        </div>
                      </div>
                    </div>

                    <div className="text-center">
                      <a
                        href={article.url}
                        target="_blank"
                        rel="noopener noreferrer"
                      >
                        Read
                      </a>
                    </div>
                  </div>
                </div>
              );
            })}
          </div>
          <PaginationNav
            page={query.page!}
            count={articles.totalPage}
            onChange={handleChangePagination}
          />
        </div>
      )}
    </div>
  );
};

Home.getLayout = (page: ReactElement) => <DefaultLayout>{page}</DefaultLayout>;

export default Home;
