import React, { ReactElement, useEffect, useState } from "react";
import DefaultLayout from "@/components/layouts/DefaultLayout";
import { NextPageWithLayout } from "@/pages/_app";

import PaginationHeader from "@/components/ui/PaginationHeader";

import { GetArticlesArg, useLazyGetArticlesQuery } from "@/services/userApi";
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

  const [
    getArticles,
    {
      data: articles = {
        totalCount: 0,
        totalPage: 1,
        pageSize: 10,
        data: [],
      },
      isLoading,
    },
  ] = useLazyGetArticlesQuery();

  const handleChangePagination = (
    e: React.ChangeEvent<unknown>,
    page: number
  ) => {
    router.push({ query: { ...query, page } });
  };

  useEffect(() => {
    setQuery(buildQuery());
  }, [router]);

  useEffect(() => {
    getArticles(query);
  }, [query]);

  return (
    <div>
      <h1>haikara</h1>

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
          <PaginationNav
            page={query.page!}
            count={articles.totalPage}
            onChange={handleChangePagination}
          />
          <div className="flex flex-wrap gap-3">
            {articles.data.map((article) => {
              return (
                <div key={article.id} className="max-w-sm">
                  {article.ogp_image_url && (
                    <div className="w-full h-auto relative">
                      <ImgproxyImage
                        src={article.ogp_image_url}
                        width={360 * 2}
                        height={189 * 2}
                        objectFit="contain"
                      />
                    </div>
                  )}

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
                        @ <StyledSiteName>{article.site.name}</StyledSiteName>
                      </a>
                    </div>
                  </div>

                  <div>
                    <a
                      href={article.url}
                      target="_blank"
                      rel="noopener noreferrer"
                    >
                      Read
                    </a>
                  </div>
                </div>
              );
            })}
          </div>
          <div
          // page={query.page}
          // count={articles.totalPage}
          // onChange={handleChangePagination}
          />
        </div>
      )}
    </div>
  );
};

Home.getLayout = (page: ReactElement) => <DefaultLayout>{page}</DefaultLayout>;

export default Home;
