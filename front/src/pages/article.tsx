import Typography from "@mui/material/Typography";
import Box from "@mui/material/Box";
import Stack from "@mui/material/Stack";
import Card from "@mui/material/Card";
import Button from "@mui/material/Button";
import Table from "@mui/material/Table";
import TableBody from "@mui/material/TableBody";
import TableCell from "@mui/material/TableCell";
import TableContainer from "@mui/material/TableContainer";
import TableHead from "@mui/material/TableHead";
import TableRow from "@mui/material/TableRow";
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
  useRunGetOGPImageOfArticleMutation,
} from "@/services/adminApi";
import PaginationHeader from "@/components/ui/PaginationHeader";
import ImgproxyImage from "@/components/ImgproxyImage";

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

  const [deleteArticle] = useDeleteArticleMutation();

  const [runGetOGPImageOfArticle] = useRunGetOGPImageOfArticleMutation();

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
    runGetOGPImageOfArticle(id);
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

          <TableContainer>
            <Table>
              <TableHead>
                <TableRow>
                  <TableCell>OGP</TableCell>
                  <TableCell>OGP Image</TableCell>
                  <TableCell>Published</TableCell>
                  <TableCell>Article Title</TableCell>
                  <TableCell></TableCell>
                  <TableCell></TableCell>
                </TableRow>
              </TableHead>
              <TableBody>
                {articles.data.map((article) => {
                  return (
                    <TableRow key={article.id}>
                      <TableCell>
                        <Button
                          onClick={onClickGetOGPImageHandler.bind(
                            this,
                            article.id
                          )}
                        >
                          OGP
                        </Button>
                      </TableCell>
                      <TableCell>
                        {article.ogp_image_url && (
                          <div
                            style={{
                              display: "inline-block",
                              width: "200px",
                              height: "105px",
                              position: "relative",
                            }}
                          >
                            <ImgproxyImage
                              src={article.ogp_image_url}
                              width={200 * 2}
                              height={105 * 2}
                              objectFit="contain"
                            />
                          </div>
                        )}
                      </TableCell>
                      <TableCell>
                        {new Date(article.published_at).toLocaleDateString()}
                      </TableCell>
                      <TableCell>{article.title}</TableCell>
                      <TableCell>
                        <Button
                          component="a"
                          href={article.url}
                          target="_blank"
                          rel="noopener noreferrer"
                        >
                          Read
                        </Button>
                      </TableCell>
                      <TableCell>
                        <IconButton
                          onClick={onClickDeleteHandler.bind(this, article.id)}
                          aria-label="remove"
                        >
                          <DeleteIcon />
                        </IconButton>
                      </TableCell>
                    </TableRow>
                  );
                })}
              </TableBody>
            </Table>
          </TableContainer>

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
