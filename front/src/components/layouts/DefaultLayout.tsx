import Head from "next/head";
import Header from "@/components/Header";
import { Box, Container } from "@mui/material";
import Footer from "@/components/Footer";
import React, { FC, ReactNode, useState } from "react";
import { styled } from "@mui/material/styles";

type DefaultLayoutProps = {
  children: ReactNode;
};

const drawerWidth = 0;

const Main = styled(Box, { shouldForwardProp: (prop) => prop !== "open" })<{
  open?: boolean;
}>(({ theme, open }) => ({
  minHeight: "100vh",
  flexGrow: 1,
  padding: theme.spacing(3),
  transition: theme.transitions.create("margin", {
    easing: theme.transitions.easing.sharp,
    duration: theme.transitions.duration.leavingScreen,
  }),
  marginLeft: `-${drawerWidth}px`,
  ...(open && {
    transition: theme.transitions.create("margin", {
      easing: theme.transitions.easing.easeOut,
      duration: theme.transitions.duration.enteringScreen,
    }),
    marginLeft: 0,
  }),
}));

const DrawerHeader = styled("div")(({ theme }) => ({
  display: "flex",
  alignItems: "center",
  padding: theme.spacing(0, 1),
  // necessary for content to be below app bar
  ...theme.mixins.toolbar,
  justifyContent: "flex-end",
}));

const DefaultLayout: FC<DefaultLayoutProps> = ({ children }) => {
  const [open, setOpen] = useState(false);
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

      <Header open={open} />

      <Main open={open} as="main">
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
      </Main>
    </Box>
  );
};
export default DefaultLayout;
