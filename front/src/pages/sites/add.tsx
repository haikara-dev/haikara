import React, { ReactElement } from "react";
import AdminLayout from "@/components/layouts/AdminLayout";
import { NextPageWithLayout } from "@/pages/_app";
import AddSiteForm from "@/components/site/AddSiteForm";
import Typography from "@mui/material/Typography";
import { SiteWithSiteCrawlRule } from "@/features/Sites";
import { selectAuthUser } from "@/features/auth/authSlice";
import { useAppSelector } from "@/app/hooks";

const BACKEND_ADMIN_API_URL: string =
  process.env.NEXT_PUBLIC_BACKEND_ADMIN_API_URL!;

const SiteAdd: NextPageWithLayout = () => {
  const authUser = useAppSelector(selectAuthUser);
  const getRequestHeaders = async () => {
    const idToken = await authUser?.getIdToken();
    return {
      Authorization: `Bearer ${idToken}`,
    };
  };
  const addSite = async (site: SiteWithSiteCrawlRule) => {
    try {
      const headers = await getRequestHeaders();
      const res = await fetch(BACKEND_ADMIN_API_URL + "/sites", {
        method: "POST",
        headers: {
          ...headers,
          ...{
            "Content-Type": "application/json",
          },
        },
        body: JSON.stringify({
          ...site,
        }),
      });
      if (!res.ok) throw new Error(`${res.status}: ${res.statusText}`);

      // 一覧へ戻る
      // await loadData();
    } catch (err) {
      console.log(err);
    }
  };

  const getRssUrlByUrl = async (url: string): Promise<string> => {
    try {
      const headers = await getRequestHeaders();
      const queryParams = new URLSearchParams({ url: url });
      const res = await fetch(
        new URL(
          BACKEND_ADMIN_API_URL + "/sites/get-rss-url-by-url?" + queryParams
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
      console.log(json);
      return json.url;
    } catch (err) {
      console.log(err);
    }
    return "";
  };
  return (
    <div>
      <Typography variant="h3" component="h1">
        Add Site
      </Typography>

      <AddSiteForm addSite={addSite} getRssUrlByUrl={getRssUrlByUrl} />
    </div>
  );
};
SiteAdd.getLayout = (page: ReactElement) => <AdminLayout>{page}</AdminLayout>;
export default SiteAdd;
