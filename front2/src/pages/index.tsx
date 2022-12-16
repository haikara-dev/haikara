import React, { ReactElement, useEffect, useState } from "react";
import DefaultLayout from "@/components/layouts/DefaultLayout";
import { NextPageWithLayout } from "@/pages/_app";
import Typography from "@mui/material/Typography";
import { useAppSelector } from "@/app/hooks";
import { selectAuthUser } from "@/features/auth/authSlice";
import Stack from "@mui/material/Stack";
import PaginationHeader from "@/components/ui/PaginationHeader";
import Card from "@mui/material/Card";
import Button from "@mui/material/Button";
import Box from "@mui/material/Box";
import CardMedia from "@mui/material/CardMedia";
import CardContent from "@mui/material/CardContent";
import CardActions from "@mui/material/CardActions";
import ImgproxyImage from "@/components/ImgproxyImage";
import Pagination from "@mui/material/Pagination";
import { GetArticlesArg, useLazyGetArticlesQuery } from "@/services/adminApi";
import { useRouter } from "next/router";
import StyledSiteName from "@/components/site/StyledSiteName";

const Home: NextPageWithLayout = () => {
  const authUser = useAppSelector(selectAuthUser);
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
    if (authUser) {
      getArticles(query);
    }
  }, [query]);

  return (
    <div>
      <Typography variant="h3" component="h1">
        haikara
      </Typography>

      {authUser && isLoading ? (
        <div>loading...</div>
      ) : (
        <Stack gap={3} alignItems="center">
          <PaginationHeader
            totalCount={articles.totalCount}
            page={query.page!}
            pageSize={articles.pageSize}
            dataSize={articles.data.length}
          />

          <Box
            sx={{
              display: "flex",
              flexWrap: "wrap",
              gap: 3,
            }}
          >
            {articles.data.map((article) => {
              return (
                <Card key={article.id} sx={{ maxWidth: 345 }}>
                  {article.ogp_image_url && (
                    <CardMedia
                      style={{
                        width: "1OO%",
                        height: "auto",
                        position: "relative",
                      }}
                    >
                      <ImgproxyImage
                        src={article.ogp_image_url}
                        width={360 * 2}
                        height={189 * 2}
                        objectFit="contain"
                      />
                    </CardMedia>
                  )}

                  <CardContent>
                    <Typography variant="h6" component="h2">
                      {article.title}
                    </Typography>
                    <div>
                      {new Date(article.published_at).toLocaleDateString()}
                    </div>
                    <div>
                      <Button
                        component="a"
                        href={article.site.url}
                        target="_blank"
                        rel="noopener noreferrer"
                      >
                        @ <StyledSiteName>{article.site.name}</StyledSiteName>
                      </Button>
                    </div>
                  </CardContent>

                  <CardActions>
                    <Button
                      component="a"
                      href={article.url}
                      target="_blank"
                      rel="noopener noreferrer"
                    >
                      Read
                    </Button>
                  </CardActions>
                </Card>
              );
            })}
          </Box>
          <Pagination
            page={query.page}
            count={articles.totalPage}
            onChange={handleChangePagination}
          />
        </Stack>
      )}
    </div>
  );
};

Home.getLayout = (page: ReactElement) => <DefaultLayout>{page}</DefaultLayout>;

export default Home;
