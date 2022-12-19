import Head from "next/head";
import Box from "@mui/material/Box";
import Container from "@mui/material/Container";
import Footer from "@/components/Footer";
import React, { FC, ReactNode } from "react";
import styled from "@mui/material/styles/styled";
import Header from "@/components/Header";

type DefaultLayoutProps = {
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

const DefaultLayout: FC<DefaultLayoutProps> = ({ children }) => {
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

      <Header />

      <Box
        component="main"
        sx={{
          minHeight: "100vh",
          flexGrow: 1,
          padding: 3,
        }}
      >
        <DrawerHeader />
        <Container
          sx={{
            p: 2,
            minHeight: "70vh",
          }}
        >
          {children}
        </Container>
        <Footer />
      </Box>
    </Box>
  );
};
export default DefaultLayout;
