import Typography from "@mui/material/Typography";
import Box from "@mui/material/Box";
import Stack from "@mui/material/Stack";
import Card from "@mui/material/Card";
import Button from "@mui/material/Button";
import IconButton from "@mui/material/IconButton";
import Pagination from "@mui/material/Pagination";

import React, { ReactElement, useEffect, useState } from "react";
import DeleteIcon from "@mui/icons-material/Delete";
import { useRouter } from "next/router";
import { NextPageWithLayout } from "@/pages/_app";
import AdminLayout from "@/components/layouts/AdminLayout";

import {
  GetArticlesArg,
  useDeleteArticleMutation,
  useGetArticlesQuery,
} from "@/services/adminApi";
import PaginationHeader from "@/components/ui/PaginationHeader";

const Articles: NextPageWithLayout = () => {
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
    data: articles = {
      totalCount: 0,
      totalPage: 1,
      pageSize: 10,
      data: [],
    },
    isLoading,
  } = useGetArticlesQuery(query);

  const [deleteArticle, result] = useDeleteArticleMutation();

  const onClickDeleteHandler = async (
    id: number,
    e: React.MouseEvent<HTMLButtonElement>
  ) => {
    e.preventDefault();
    await deleteArticle(id);
  };

  const onClickGetOGPImageHandler = async (
    id: number,
    e: React.MouseEvent<HTMLButtonElement>
  ) => {
    e.preventDefault();
  };

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
      <Typography variant="h3" component="h1">
        Articles {query.site_id && <b> ( {query.site_id} ) </b>}
      </Typography>

      {isLoading ? (
        <div>Loading...</div>
      ) : (
        <Stack gap={3} alignItems="center">
          <PaginationHeader
            totalCount={articles.totalCount}
            page={query.page!}
            pageSize={articles.pageSize}
            dataSize={articles.data.length}
          />

          <Stack gap={2} mt={2} pr={2}>
            {articles.data.map((article) => {
              return (
                <Card key={article.id}>
                  <Stack direction="row" gap={3} alignItems="center">
                    <Button
                      onClick={onClickGetOGPImageHandler.bind(this, article.id)}
                    >
                      OGP
                    </Button>
                    <div>
                      {new Date(article.published_at).toLocaleDateString()}
                    </div>
                    <Box
                      sx={{
                        flexGrow: 1,
                      }}
                    >
                      {article.title}
                    </Box>
                    <Button
                      component="a"
                      href={article.url}
                      target="_blank"
                      rel="noopener noreferrer"
                    >
                      Read
                    </Button>
                    <IconButton
                      onClick={onClickDeleteHandler.bind(this, article.id)}
                      aria-label="remove"
                    >
                      <DeleteIcon />
                    </IconButton>
                  </Stack>
                </Card>
              );
            })}
          </Stack>
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

Articles.getLayout = (page: ReactElement) => <AdminLayout>{page}</AdminLayout>;

export default Articles;
