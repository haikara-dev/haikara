import React, { ReactElement } from "react";
import AdminLayout from "@/components/layouts/AdminLayout";
import { NextPageWithLayout } from "@/pages/_app";
import AddSiteForm from "@/components/site/AddSiteForm";
import Typography from "@mui/material/Typography";

const SiteAdd: NextPageWithLayout = () => {
  return (
    <div>
      <Typography variant="h3" component="h1">
        Add Site
      </Typography>

      <AddSiteForm />
    </div>
  );
};
SiteAdd.getLayout = (page: ReactElement) => <AdminLayout>{page}</AdminLayout>;
export default SiteAdd;
