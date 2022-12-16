import React, { ReactElement, useEffect } from "react";
import AdminLayout from "@/components/layouts/AdminLayout";
import { NextPageWithLayout } from "@/pages/_app";
import Typography from "@mui/material/Typography";
import EditSiteForm from "@/components/site/EditSiteForm";
import { useRouter } from "next/router";
import { useLazyGetSiteWithCrawlRuleAndCategoryQuery } from "@/services/adminApi";

const SiteEdit: NextPageWithLayout = () => {
  const router = useRouter();
  const { id } = router.query;
  const [getSiteWithCrawlRuleAndCategory, { data: site }] =
    useLazyGetSiteWithCrawlRuleAndCategoryQuery();
  useEffect(() => {
    if (id) {
      getSiteWithCrawlRuleAndCategory(Number(id));
    }
  }, [id]);
  return (
    <div>
      <Typography variant="h3" component="h1">
        Edit Site
      </Typography>
      {site && <EditSiteForm site={site} />}
    </div>
  );
};
SiteEdit.getLayout = (page: ReactElement) => <AdminLayout>{page}</AdminLayout>;
export default SiteEdit;
