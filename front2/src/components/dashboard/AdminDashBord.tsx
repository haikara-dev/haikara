import React from "react";
import { useGetAdminDashboardQuery } from "@/services/adminApi";

const AdminDashBord = () => {
  const { data: dashboard, isLoading } = useGetAdminDashboardQuery();

  return (
    <div>
      {isLoading && <div>Loading...</div>}
      {dashboard && (
        <div>
          <div>Site: {dashboard.siteSize}</div>
          <div>Feed: {dashboard.feedSize}</div>
          <div>Article: {dashboard.articleSize}</div>
          <div>OGP Image: {dashboard.ogpImageSize}</div>
          <div>User: {dashboard.userSize}</div>
        </div>
      )}
    </div>
  );
};
export default AdminDashBord;
