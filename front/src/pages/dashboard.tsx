import React, { ReactElement } from "react";
import Typography from "@mui/material/Typography";
import AdminLayout from "@/components/layouts/AdminLayout";
import { NextPageWithLayout } from "@/pages/_app";

const Dashboard: NextPageWithLayout = () => {
  return (
    <div>
      <Typography variant="h3" component="h1">
        Dashboard
      </Typography>
    </div>
  );
};

Dashboard.getLayout = (page: ReactElement) => <AdminLayout>{page}</AdminLayout>;

export default Dashboard;
