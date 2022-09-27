import { NextPageWithLayout } from "@/pages/_app";
import React, { ReactElement } from "react";
import AdminLayout from "@/components/layouts/AdminLayout";

const Site: NextPageWithLayout = () => {
  return <div>Site</div>;
};
Site.getLayout = (page: ReactElement) => <AdminLayout>{page}</AdminLayout>;
export default Site;
