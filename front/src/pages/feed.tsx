import {
  Typography,
  Box,
  Stack,
  Card,
  Button,
  IconButton,
} from "@mui/material";
import React, { ReactElement, useEffect, useState } from "react";
import { useAuthUserContext } from "@/lib/AuthUser";
import DeleteIcon from "@mui/icons-material/Delete";
import { NextPageWithLayout } from "@/pages/_app";
import AdminLayout from "@/components/layouts/AdminLayout";

const BACKEND_API_URL: string = process.env.NEXT_PUBLIC_BACKEND_API_URL!;
const BACKEND_ADMIN_API_URL: string =
  process.env.NEXT_PUBLIC_BACKEND_ADMIN_API_URL!;

export type Feed = {
  id: number;
  count: number;
  created_at: string;
  indexed_at: string;
  site_id: number;
  site_name: string;
};

const Feeds: NextPageWithLayout = () => {
  const [data, setData] = useState<Feed[]>([]);
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
      const res = await fetch(BACKEND_ADMIN_API_URL + "/feeds/lite", {
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

  const removeFeed = async (id: number) => {
    try {
      const headers = await getRequestHeaders();
      const res = await fetch(
        new URL(id.toString(), BACKEND_ADMIN_API_URL + "/feeds/"),
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

  const runParse = async (id: number) => {
    try {
      const headers = await getRequestHeaders();
      const res = await fetch(
        new URL(id.toString(), BACKEND_ADMIN_API_URL + "/feeds/parse/"),
        {
          method: "GET",
          headers: headers,
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
    removeFeed(id);
  };

  const onClickRunHandler = (
    id: number,
    e: React.MouseEvent<HTMLButtonElement>
  ) => {
    e.preventDefault();
    runParse(id);
  };

  useEffect(() => {
    setLoading(true);
    loadData();
  }, []);

  return (
    <div>
      <Typography variant="h3" component="h1">
        Feeds
      </Typography>

      {isLoading ? (
        <div>Loading...</div>
      ) : (
        <Stack gap={2} mt={2} pr={8}>
          {data.map((feed) => {
            return (
              <Card key={feed.id}>
                <Stack direction="row" gap={3} alignItems="center">
                  <Button onClick={onClickRunHandler.bind(this, feed.id)}>
                    Run
                  </Button>
                  <div>{new Date(feed.created_at).toLocaleString()}</div>
                  <Box
                    sx={{
                      flexGrow: 1,
                    }}
                  >
                    {feed.site_name}
                  </Box>
                  {feed.indexed_at && (
                    <div>{new Date(feed.indexed_at).toLocaleString()}</div>
                  )}
                  <div>{feed.count}</div>
                  <IconButton
                    onClick={onClickDeleteHandler.bind(this, feed.id)}
                    aria-label="remove"
                  >
                    <DeleteIcon />
                  </IconButton>
                </Stack>
              </Card>
            );
          })}
        </Stack>
      )}
    </div>
  );
};

Feeds.getLayout = (page: ReactElement) => <AdminLayout>{page}</AdminLayout>;

export default Feeds;
