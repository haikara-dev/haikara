import Typography from "@mui/material/Typography";
import Stack from "@mui/material/Stack";
import Card from "@mui/material/Card";
import Button from "@mui/material/Button";
import Pagination from "@mui/material/Pagination";

import React, { ReactElement, useEffect, useState } from "react";
import SiteRow from "@/components/site/SiteRow";
import EditSiteFormDialog from "@/components/site/EditSiteFormDialog";
import DryRunDialog from "@/components/site/DryRunDialog";
import { NextPageWithLayout } from "@/pages/_app";
import AdminLayout from "@/components/layouts/AdminLayout";
import { useRouter } from "next/router";
import { DryRunResult, Site, SiteWithSiteCrawlRule } from "@/features/Sites";
import { selectAuthUser } from "@/features/auth/authSlice";
import { useAppSelector } from "@/app/hooks";

const BACKEND_API_URL: string = process.env.NEXT_PUBLIC_BACKEND_API_URL!;
const BACKEND_ADMIN_API_URL: string =
  process.env.NEXT_PUBLIC_BACKEND_ADMIN_API_URL!;

const Sites: NextPageWithLayout = () => {
  const router = useRouter();
  const [page, setPage] = useState<number>(
    router.query.page ? parseInt(router.query.page.toString()) : 1
  );
  const [totalCount, setTotalCount] = useState<number>(0);
  const [totalPage, setTotalPage] = useState<number>(1);
  const [pageSize, setPageSize] = useState<number>(10);

  const [data, setData] = useState<Site[]>([]);
  const [isLoading, setLoading] = useState(false);
  const authUser = useAppSelector(selectAuthUser);
  const [editOpen, setEditOpen] = useState(false);
  const [editTarget, setEditTarget] = useState<Site | null>(null);

  const [dryOpen, setDryOpen] = useState(false);
  const [dryRunResult, setDryRunResult] = useState<DryRunResult | null>(null);

  const handleAddOpen = () => {
    router.push("/sites/add");
  };

  const handleEditOpen = (site: Site) => {
    setEditOpen(true);
    setEditTarget(site);
  };

  const handleEditClose = () => {
    setEditOpen(false);
    setEditTarget(null);
  };

  const openDryDialog = (result: DryRunResult) => {
    setDryRunResult(result);
    setDryOpen(true);
  };

  const handleDryClose = () => {
    setDryOpen(false);
  };

  const getRequestHeaders = async () => {
    const idToken = await authUser?.getIdToken();
    return {
      Authorization: `Bearer ${idToken}`,
    };
  };

  const loadData = async () => {
    try {
      const headers = await getRequestHeaders();
      const queryParams = new URLSearchParams({ page: page.toString() });
      const res = await fetch(BACKEND_ADMIN_API_URL + "/sites?" + queryParams, {
        method: "GET",
        headers: headers,
      });
      if (!res.ok) throw new Error(`${res.status}: ${res.statusText}`);
      const json = await res.json();

      for (const site of json.data) {
        site.cannot_crawl = site.cannot_crawl_at ? true : false;
      }

      setTotalCount(json.totalCount);
      setTotalPage(json.totalPage);
      setPageSize(json.pageSize);
      setData(json.data);
    } catch (err) {
      console.log(err);
    }
    setLoading(false);
  };

  const loadSiteWithSiteCrawlRule = async (
    id: number
  ): Promise<SiteWithSiteCrawlRule | null> => {
    try {
      const headers = await getRequestHeaders();
      const res = await fetch(
        new URL(id.toString(), BACKEND_ADMIN_API_URL + "/sites/"),
        {
          method: "GET",
          headers: headers,
        }
      );
      if (!res.ok) throw new Error(`${res.status}: ${res.statusText}`);
      const json = await res.json();

      console.log("cannot_crawl", json.cannot_crawl_at ? true : false);
      if (json.edges.site_crawl_rule) {
        return {
          id: json.id,
          name: json.name,
          url: json.url,
          feed_url: json.feed_url,
          active: json.active,
          cannot_crawl_at: json.cannot_crawl_at,
          cannot_crawl: json.cannot_crawl_at ? true : false,
          site_crawl_rule: json.edges.site_crawl_rule,
        };
      } else {
        return {
          id: json.id,
          name: json.name,
          url: json.url,
          feed_url: json.feed_url,
          active: json.active,
          cannot_crawl_at: json.cannot_crawl_at,
          cannot_crawl: json.cannot_crawl_at ? true : false,
          site_crawl_rule: {
            article_selector: "",
            title_selector: "",
            link_selector: "",
            description_selector: "",
            has_data_to_list: true,
            date_selector: "",
            date_layout: "",
            is_time_humanize: false,
            is_spa: false,
          },
        };
      }
    } catch (err) {
      console.log(err);
      return null;
    }
  };

  const activeSite = async (id: number) => {
    try {
      const headers = await getRequestHeaders();
      const res = await fetch(
        new URL(id.toString(), BACKEND_ADMIN_API_URL + "/sites/active/"),
        {
          method: "PATCH",
          headers: {
            ...headers,
            ...{
              "Content-Type": "application/json",
            },
          },
          body: JSON.stringify({
            active: true,
          }),
        }
      );
      if (!res.ok) throw new Error(`${res.status}: ${res.statusText}`);
      await loadData();
    } catch (err) {
      console.log(err);
    }
  };

  const deActiveSite = async (id: number) => {
    try {
      const headers = await getRequestHeaders();
      const res = await fetch(
        new URL(id.toString(), BACKEND_ADMIN_API_URL + "/sites/deActive/"),
        {
          method: "PATCH",
          headers: {
            ...headers,
            ...{
              "Content-Type": "application/json",
            },
          },
          body: JSON.stringify({
            active: false,
          }),
        }
      );
      if (!res.ok) throw new Error(`${res.status}: ${res.statusText}`);
      await loadData();
    } catch (err) {
      console.log(err);
    }
  };

  const removeSite = async (id: number) => {
    try {
      const headers = await getRequestHeaders();
      const res = await fetch(
        new URL(id.toString(), BACKEND_ADMIN_API_URL + "/sites/"),
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

  const updateSite = async (site: SiteWithSiteCrawlRule) => {
    console.log("site", site);
    try {
      const headers = await getRequestHeaders();
      const res = await fetch(
        new URL(site.id.toString(), BACKEND_ADMIN_API_URL + "/sites/"),
        {
          method: "PUT",
          headers: {
            ...headers,
            ...{
              "Content-Type": "application/json",
            },
          },
          body: JSON.stringify({
            ...site,
          }),
        }
      );
      if (!res.ok) throw new Error(`${res.status}: ${res.statusText}`);
      await loadData();
    } catch (err) {
      console.log(err);
    }
  };

  const runCrawling = async (id: number) => {
    try {
      const headers = await getRequestHeaders();
      const res = await fetch(
        new URL(id.toString(), BACKEND_ADMIN_API_URL + "/sites/run-crawling/"),
        {
          method: "GET",
          headers: {
            ...headers,
            ...{
              "Content-Type": "application/json",
            },
          },
        }
      );
      if (!res.ok) throw new Error(`${res.status}: ${res.statusText}`);

      const json = await res.json();
      console.log(json);
    } catch (err) {
      console.log(err);
    }
  };

  const dryRunCrawling = async (id: number) => {
    setDryRunResult(null);
    try {
      const headers = await getRequestHeaders();
      const res = await fetch(
        new URL(
          id.toString(),
          BACKEND_ADMIN_API_URL + "/sites/dry-run-crawling/"
        ),
        {
          method: "GET",
          headers: {
            ...headers,
            ...{
              "Content-Type": "application/json",
            },
          },
        }
      );
      if (!res.ok) throw new Error(`${res.status}: ${res.statusText}`);

      const json = await res.json();
      openDryDialog(json);
    } catch (err) {
      console.log(err);
    }
  };

  const getRssUrl = async (id: number): Promise<string> => {
    try {
      const headers = await getRequestHeaders();
      const res = await fetch(
        new URL(id.toString(), BACKEND_ADMIN_API_URL + "/sites/get-rss-url/"),
        {
          method: "GET",
          headers: {
            ...headers,
            ...{
              "Content-Type": "application/json",
            },
          },
        }
      );
      if (!res.ok) throw new Error(`${res.status}: ${res.statusText}`);

      const json = await res.json();
      console.log(json);
      return json.url;
    } catch (err) {
      console.log(err);
    }
    return "";
  };

  const handleChangePagination = (
    e: React.ChangeEvent<unknown>,
    page: number
  ) => {
    router.push({ query: { page: page } });
  };

  useEffect(() => {
    setLoading(true);
  }, []);

  useEffect(() => {
    setPage(router.query.page ? parseInt(router.query.page.toString()) : 1);
  }, [router]);

  useEffect(() => {
    loadData();
  }, [page]);

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
          <Stack>
            {totalCount}件中　{(page - 1) * pageSize + 1} -{" "}
            {(page - 1) * pageSize + data.length}件
          </Stack>
          <Stack gap={2} mt={2} pr={8}>
            {data.map((site) => {
              return (
                <Card key={site.id}>
                  <SiteRow
                    key={site.id}
                    site={site}
                    activeSite={activeSite}
                    deActiveSite={deActiveSite}
                    removeSite={removeSite}
                    openDialog={handleEditOpen}
                    runCrawling={runCrawling}
                    dryRunCrawling={dryRunCrawling}
                  />
                </Card>
              );
            })}
          </Stack>
          <Pagination
            page={page}
            count={totalPage}
            onChange={handleChangePagination}
          />
        </Stack>
      )}

      {editTarget && (
        <EditSiteFormDialog
          open={editOpen}
          handleClose={handleEditClose}
          site={editTarget}
          updateSite={updateSite}
          onEndEdit={handleEditClose}
          getRssUrl={getRssUrl}
          loadSiteWithSiteCrawlRule={loadSiteWithSiteCrawlRule}
        />
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
