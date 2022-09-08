import type { NextPage } from "next";
import Head from "next/head";
import { Container, Typography, Box, Stack, Card, Button } from "@mui/material";
import React, { useEffect, useState } from "react";
import Header from "@/components/Header";
import Footer from "@/components/Footer";
import { useAuthUserContext } from "@/lib/AuthUser";

const BACKEND_API_URL: string = process.env.NEXT_PUBLIC_BACKEND_API_URL!;

export type Article = {
  id: number;
  title: string;
  published_at: string;
};

const Articles: NextPage = () => {
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
      setData(json);
    } catch (err) {
      console.log(err);
    }
    setLoading(false);
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
          height: "100vh",
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
            <Stack gap={2} mt={2} pr={8}>
              {data.map((article) => {
                return (
                  <Card key={article.id}>
                    <Stack direction="row" gap={3} alignItems="center">
                      <div>
                        {new Date(article.published_at).toLocaleDateString()}
                      </div>
                      <div>{article.title}</div>
                    </Stack>
                  </Card>
                );
              })}
            </Stack>
          )}
        </Container>
      </Box>

      <Footer />
    </div>
  );
};

export default Articles;
