import Head from "next/head";
import Header from "@/components/Header";
import { Box, Container } from "@mui/material";
import Footer from "@/components/Footer";
import React, { FC, ReactNode } from "react";

type AdminLayoutProps = {
  children: ReactNode;
};
const AdminLayout: FC<AdminLayoutProps> = ({ children }) => {
  return (
    <div>
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
        }}
      >
        <Container
          sx={{
            p: 2,
          }}
        >
          {children}
        </Container>
      </Box>
      <Footer />
    </div>
  );
};
export default AdminLayout;
