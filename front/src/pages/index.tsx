import type { NextPage } from "next";
import Head from "next/head";
import Header from "@/components/Header";
import Footer from "@/components/Footer";
import { styled } from "@mui/material/styles";
import { Container } from "@mui/material";
const DrawerHeader = styled("div")(({ theme }) => ({
  display: "flex",
  alignItems: "center",
  padding: theme.spacing(0, 1),
  // necessary for content to be below app bar
  ...theme.mixins.toolbar,
  justifyContent: "flex-end",
}));

const Home: NextPage = () => {
  return (
    <div>
      <Head>
        <title>haikara</title>
        <meta name="description" content="haikara" />
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <Header />
      <main>
        <Container>
          <DrawerHeader />
          <h1>haikara</h1>
        </Container>
      </main>
      <Footer />
    </div>
  );
};

export default Home;
