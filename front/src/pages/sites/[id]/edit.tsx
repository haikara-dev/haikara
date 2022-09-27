import React, { ReactElement } from "react";
import AdminLayout from "@/components/layouts/AdminLayout";
import { NextPageWithLayout } from "@/pages/_app";

const SiteEdit: NextPageWithLayout = () => {
  return <div>SiteEdit</div>;
};
SiteEdit.getLayout = (page: ReactElement) => <AdminLayout>{page}</AdminLayout>;
export default SiteEdit;
