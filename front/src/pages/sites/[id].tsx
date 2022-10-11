import { NextPageWithLayout } from "@/pages/_app";
import React, { ReactElement, useEffect } from "react";
import AdminLayout from "@/components/layouts/AdminLayout";
import Typography from "@mui/material/Typography";
import { useRouter } from "next/router";
import { useLazyGetSiteWithSiteCrawlRuleQuery } from "@/services/adminApi";

const Site: NextPageWithLayout = () => {
  const router = useRouter();
  const { id } = router.query;
  const [getSiteWithSiteCrawlRule, { data: site }] =
    useLazyGetSiteWithSiteCrawlRuleQuery();
  useEffect(() => {
    if (id) {
      getSiteWithSiteCrawlRule(Number(id));
    }
  }, [id]);
  return (
    <div>
      <Typography variant="h3" component="h1">
        Site
      </Typography>
      {site && <div>site</div>}
    </div>
  );
};
Site.getLayout = (page: ReactElement) => <AdminLayout>{page}</AdminLayout>;
export default Site;
