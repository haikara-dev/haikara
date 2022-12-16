import { useGetDashboardQuery } from "@/services/userApi";
import React from "react";

const DashBord = () => {
  const { data: dashboard, isLoading } = useGetDashboardQuery();

  return (
    <div>
      {isLoading && <div>Loading...</div>}
      {dashboard && (
        <div>
          <div>Site: {dashboard.siteSize}</div>
          <div>Article: {dashboard.articleSize}</div>
        </div>
      )}
    </div>
  );
};
export default DashBord;
