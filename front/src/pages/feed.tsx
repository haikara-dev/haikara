import Typography from "@mui/material/Typography";
import Box from "@mui/material/Box";
import Stack from "@mui/material/Stack";
import Card from "@mui/material/Card";
import Button from "@mui/material/Button";
import IconButton from "@mui/material/IconButton";
import Pagination from "@mui/material/Pagination";

import React, { ReactElement, useEffect, useState } from "react";
import DeleteIcon from "@mui/icons-material/Delete";
import { NextPageWithLayout } from "@/pages/_app";
import AdminLayout from "@/components/layouts/AdminLayout";
import { useRouter } from "next/router";
import {
  useDeleteFeedMutation,
  useGetFeedsQuery,
  useRunParseFeedMutation,
} from "@/services/adminApi";
import PaginationHeader from "@/components/PaginationHeader";

const Feeds: NextPageWithLayout = () => {
  const router = useRouter();
  const [page, setPage] = useState<number>(
    router.query.page ? parseInt(router.query.page.toString()) : 1
  );

  const {
    data: feeds = {
      totalCount: 0,
      totalPage: 1,
      pageSize: 10,
      data: [],
    },
    isLoading,
    refetch,
  } = useGetFeedsQuery(page);

  const [deleteFeed, deleteFeedResulte] = useDeleteFeedMutation();

  const [runParseFeed, runParseFeedResulte] = useRunParseFeedMutation();

  const onClickDeleteHandler = async (
    id: number,
    e: React.MouseEvent<HTMLButtonElement>
  ) => {
    e.preventDefault();
    await deleteFeed(id);
    refetch();
  };

  const onClickRunHandler = async (
    id: number,
    e: React.MouseEvent<HTMLButtonElement>
  ) => {
    e.preventDefault();
    await runParseFeed(id);
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
        Feeds
      </Typography>

      {isLoading ? (
        <div>Loading...</div>
      ) : (
        <Stack gap={3} alignItems="center">
          <PaginationHeader
            totalCount={feeds.totalCount}
            page={page}
            pageSize={feeds.pageSize}
            dataSize={feeds.data.length}
          />
          <Stack gap={2} mt={2} pr={2}>
            {feeds.data.map((feed) => {
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
          <Pagination
            page={page}
            count={feeds.totalPage}
            onChange={handleChangePagination}
          />
        </Stack>
      )}
    </div>
  );
};

Feeds.getLayout = (page: ReactElement) => <AdminLayout>{page}</AdminLayout>;

export default Feeds;
