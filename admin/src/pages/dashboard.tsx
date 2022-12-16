import React, { ReactElement } from "react";
import Typography from "@mui/material/Typography";
import AdminLayout from "@/components/layouts/AdminLayout";
import { NextPageWithLayout } from "@/pages/_app";
import DashBord from "@/components/dashboard/DashBord";
import AdminDashBord from "@/components/dashboard/AdminDashBord";
import { useAppSelector } from "@/app/hooks";
import { selectIsAdmin } from "@/features/auth/authSlice";

const DashboardPage: NextPageWithLayout = () => {
  const isAdmin = useAppSelector(selectIsAdmin);
  return (
    <div>
      <Typography variant="h3" component="h1">
        Dashboard
      </Typography>
      {isAdmin ? <AdminDashBord /> : <DashBord />}
    </div>
  );
};

DashboardPage.getLayout = (page: ReactElement) => (
  <AdminLayout>{page}</AdminLayout>
);

export default DashboardPage;
