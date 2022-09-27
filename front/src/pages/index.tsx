import React, { ReactElement } from "react";
import DefaultLayout from "@/components/layouts/DefaultLayout";
import { NextPageWithLayout } from "@/pages/_app";
import Typography from "@mui/material/Typography";

const Home: NextPageWithLayout = () => {
  return (
    <div>
      <Typography variant="h3" component="h1">
        haikara
      </Typography>
    </div>
  );
};

Home.getLayout = (page: ReactElement) => <DefaultLayout>{page}</DefaultLayout>;

export default Home;
