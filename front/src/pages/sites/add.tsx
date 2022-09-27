import React, { ReactElement } from "react";
import AdminLayout from "@/components/layouts/AdminLayout";
import { NextPageWithLayout } from "@/pages/_app";

const SiteAdd: NextPageWithLayout = () => {
  return <div>SiteAdd</div>;
};
SiteAdd.getLayout = (page: ReactElement) => <AdminLayout>{page}</AdminLayout>;
export default SiteAdd;
