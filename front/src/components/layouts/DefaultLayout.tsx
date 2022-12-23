import Head from "next/head";
import Footer from "@/components/Footer";
import React, { FC, ReactNode } from "react";

type DefaultLayoutProps = {
  children: ReactNode;
};

const DefaultLayout: FC<DefaultLayoutProps> = ({ children }) => {
  return (
    <div className="bg-gray-900 text-white">
      <Head>
        <title>haikara</title>
        <meta name="description" content="haikara" />
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <main>
        <div className="">{children}</div>
        <Footer />
      </main>
    </div>
  );
};
export default DefaultLayout;
