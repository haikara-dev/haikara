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
  useDeleteArticleMutation,
  useGetArticlesQuery,
} from "@/services/adminApi";

const Articles: NextPageWithLayout = () => {
  const router = useRouter();
  const [page, setPage] = useState<number>(
    router.query.page ? parseInt(router.query.page.toString()) : 1
  );

  const {
    data: articles = {
      totalCount: 0,
      totalPage: 1,
      pageSize: 10,
      data: [],
    },
    isLoading,
    refetch,
  } = useGetArticlesQuery(page);

  const [deleteArticle, result] = useDeleteArticleMutation();

  const onClickDeleteHandler = async (
    id: number,
    e: React.MouseEvent<HTMLButtonElement>
  ) => {
    e.preventDefault();
    await deleteArticle(id);
    refetch();
  };

  const handleChangePagination = (
    e: React.ChangeEvent<unknown>,
    page: number
  ) => {
    router.push({ query: { page: page } });
  };

  useEffect(() => {
    setPage(router.query.page ? parseInt(router.query.page.toString()) : 1);
  }, [router]);

  return (
    <div>
      <Typography variant="h3" component="h1">
        Articles
      </Typography>

      {isLoading ? (
        <div>Loading...</div>
      ) : (
        <Stack gap={3} alignItems="center">
          <Stack>
            {articles.totalCount}件中　{(page - 1) * articles.pageSize + 1} -{" "}
            {(page - 1) * articles.pageSize + articles.data.length}件
          </Stack>
          <Stack gap={2} mt={2} pr={8}>
            {articles.data.map((article) => {
              return (
                <Card key={article.id}>
                  <Stack direction="row" gap={3} alignItems="center">
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
            page={page}
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
