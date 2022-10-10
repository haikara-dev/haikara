import Typography from "@mui/material/Typography";
import Stack from "@mui/material/Stack";
import Card from "@mui/material/Card";
import Button from "@mui/material/Button";
import Pagination from "@mui/material/Pagination";

import React, { ReactElement, useEffect, useState } from "react";
import SiteRow from "@/components/site/SiteRow";
import DryRunDialog from "@/components/site/DryRunDialog";
import { NextPageWithLayout } from "@/pages/_app";
import AdminLayout from "@/components/layouts/AdminLayout";
import { useRouter } from "next/router";
import { DryRunResult, useGetSitesQuery } from "@/services/adminApi";
import PaginationHeader from "@/components/PaginationHeader";

const Sites: NextPageWithLayout = () => {
  const router = useRouter();
  const [page, setPage] = useState<number>(
    router.query.page ? parseInt(router.query.page.toString()) : 1
  );

  const {
    data: sites = {
      totalCount: 0,
      totalPage: 1,
      pageSize: 10,
      data: [],
    },
    isLoading,
  } = useGetSitesQuery(page);

  const [dryOpen, setDryOpen] = useState(false);
  const [dryRunResult, setDryRunResult] = useState<DryRunResult | null>(null);

  const handleAddOpen = () => {
    router.push("/sites/add");
  };

  const openDryDialog = (result: DryRunResult) => {
    setDryRunResult(result);
    setDryOpen(true);
  };

  const handleDryClose = () => {
    setDryOpen(false);
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
        Sites
      </Typography>

      <Button variant="outlined" onClick={handleAddOpen}>
        Add Site
      </Button>

      {isLoading ? (
        <div>Loading...</div>
      ) : (
        <Stack gap={3} alignItems="center">
          <PaginationHeader
            totalCount={sites.totalCount}
            page={page}
            pageSize={sites.pageSize}
            dataSize={sites.data.length}
          />
          <Stack gap={2} mt={2} pr={2}>
            {sites.data.map((site) => {
              return (
                <Card key={site.id}>
                  <SiteRow
                    key={site.id}
                    site={site}
                    openDryDialog={openDryDialog}
                  />
                </Card>
              );
            })}
          </Stack>
          <Pagination
            page={page}
            count={sites.totalPage}
            onChange={handleChangePagination}
          />
        </Stack>
      )}

      {dryRunResult && (
        <DryRunDialog
          open={dryOpen}
          handleClose={handleDryClose}
          dryRunResult={dryRunResult}
        />
      )}
    </div>
  );
};

Sites.getLayout = (page: ReactElement) => <AdminLayout>{page}</AdminLayout>;

export default Sites;
