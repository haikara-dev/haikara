import Head from "next/head";
import Box from "@mui/material/Box";
import Card from "@mui/material/Card";
import Footer from "@/components/Footer";
import React, { FC, ReactNode } from "react";
import styled from "@mui/material/styles/styled";
import DefaultHeader from "@/components/DefaultHeader";

type AuthLayoutProps = {
  children: ReactNode;
};

const DrawerHeader = styled("div")(({ theme }) => ({
  display: "flex",
  alignItems: "center",
  padding: theme.spacing(0, 1),
  // necessary for content to be below app bar
  ...theme.mixins.toolbar,
  justifyContent: "flex-end",
}));

const AuthLayout: FC<AuthLayoutProps> = ({ children }) => {
  return (
    <Box
      sx={{
        display: "flex",
      }}
    >
      <Head>
        <title>haikara</title>
        <meta name="description" content="haikara" />
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <DefaultHeader />
      <Box
        component="main"
        sx={{
          flexGrow: 1,
          padding: 3,
        }}
      >
        <DrawerHeader />
        <Box
          sx={{
            minHeight: "80vh",
            display: "flex",
            justifyContent: "center",
            alignItems: "center",
          }}
        >
          <Card
            sx={{
              width: "min(90vw, 400px)",
              p: 2,
            }}
          >
            {children}
          </Card>
        </Box>

        <Footer />
      </Box>
    </Box>
  );
};
export default AuthLayout;
