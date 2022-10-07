import type { AppProps } from "next/app";
import { CssBaseline } from "@mui/material";
import { store } from "@/app/store";
import { Provider } from "react-redux";
import "../firebaseConfig";
import { ReactElement, ReactNode } from "react";
import { NextPage } from "next";
import ProtectedRouterComponent from "@/components/ProtectedRouterComponent";
import Head from "next/head";

export type NextPageWithLayout<P = {}, IP = P> = NextPage<P, IP> & {
  getLayout?: (page: ReactElement) => ReactNode;
};

type AppPropsWithLayout = AppProps & {
  Component: NextPageWithLayout;
};

function MyApp({ Component, pageProps }: AppPropsWithLayout) {
  // Use the layout defined at the page level, if available
  const getLayout = Component.getLayout ?? ((page) => page);
  return (
    <>
      <Head>
        <meta name="viewport" content="initial-scale=1, width=device-width" />
      </Head>
      <CssBaseline />
      <Provider store={store}>
        <ProtectedRouterComponent>
          {getLayout(<Component {...pageProps} />)}
        </ProtectedRouterComponent>
      </Provider>
    </>
  );
}

export default MyApp;
