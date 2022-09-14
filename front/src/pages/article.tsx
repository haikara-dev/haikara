import type { NextPage } from "next";
import Head from "next/head";
import {
  Container,
  Typography,
  Box,
  Stack,
  Card,
  Button,
  IconButton,
} from "@mui/material";
import React, { useEffect, useState } from "react";
import Header from "@/components/Header";
import Footer from "@/components/Footer";
import { useAuthUserContext } from "@/lib/AuthUser";
import DeleteIcon from "@mui/icons-material/Delete";
import { textAlign } from "@mui/system";

const BACKEND_API_URL: string = process.env.NEXT_PUBLIC_BACKEND_API_URL!;

export type Article = {
  id: number;
  title: string;
  published_at: string;
};

const Articles: NextPage = () => {
  const [totalCount, setTotalCount] = useState<number>(0);
  const [data, setData] = useState<Article[]>([]);
  const [isLoading, setLoading] = useState(false);
  const { authUser } = useAuthUserContext();

  const getRequestHeaders = async () => {
    const idToken = await authUser?.getIdToken();
    return {
      Authorization: `Bearer ${idToken}`,
    };
  };

  const loadData = async () => {
    try {
      const headers = await getRequestHeaders();
      const res = await fetch(BACKEND_API_URL + "/articles", {
        method: "GET",
        headers: headers,
      });
      if (!res.ok) throw new Error(`${res.status}: ${res.statusText}`);
      const json = await res.json();
      setTotalCount(json.totalCount);
      setData(json.data);
    } catch (err) {
      console.log(err);
    }
    setLoading(false);
  };

  const removeArticle = async (id: number) => {
    try {
      const headers = await getRequestHeaders();
      const res = await fetch(
        new URL(id.toString(), BACKEND_API_URL + "/articles/"),
        {
          method: "DELETE",
          headers: {
            ...headers,
            ...{
              "Content-Type": "application/json",
            },
          },
        }
      );
      if (!res.ok) throw new Error(`${res.status}: ${res.statusText}`);
      await loadData();
    } catch (err) {
      console.log(err);
    }
  };

  const onClickDeleteHandler = (
    id: number,
    e: React.MouseEvent<HTMLButtonElement>
  ) => {
    e.preventDefault();
    removeArticle(id);
  };

  useEffect(() => {
    setLoading(true);
    loadData();
  }, []);

  return (
    <div>
      <Head>
        <title>DailyFJ</title>
        <meta name="description" content="DailyFJ" />
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <Header />
      <Box
        component="main"
        sx={{
          minHeight: "100vh",
        }}
      >
        <Container
          sx={{
            p: 2,
          }}
        >
          <Typography variant="h3" component="h1">
            Articles
          </Typography>

          {isLoading ? (
            <div>Loading...</div>
          ) : (
            <Box>
              <Box
                sx={{
                  textAlign: "center",
                }}
              >
                {totalCount}件中　{1} - {100}
              </Box>
              <Stack gap={2} mt={2} pr={8}>
                {data.map((article) => {
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
            </Box>
          )}
        </Container>
      </Box>

      <Footer />
    </div>
  );
};

export default Articles;
