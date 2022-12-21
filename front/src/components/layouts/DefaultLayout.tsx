import Head from "next/head";
import Footer from "@/components/Footer";
import React, { FC, ReactNode } from "react";

type DefaultLayoutProps = {
  children: ReactNode;
};

const DefaultLayout: FC<DefaultLayoutProps> = ({ children }) => {
  return (
    <div>
      <Head>
        <title>haikara</title>
        <meta name="description" content="haikara" />
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <main className="bg-blue-100">
        <div className="">{children}</div>
        <Footer />
      </main>
    </div>
  );
};
export default DefaultLayout;
