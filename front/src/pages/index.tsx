import type { NextPage } from "next";
import Head from "next/head";
import Header from "@/components/Header";
import Footer from "@/components/Footer";

const Home: NextPage = () => {
  return (
    <div>
      <Head>
        <title>dailyfj</title>
        <meta name="description" content="dailyfj" />
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <Header />
      <main>
        <h1>dailyfj</h1>
      </main>
      <Footer />
    </div>
  );
};

export default Home;
