import type { AppProps } from "next/app";
import { CssBaseline } from "@mui/material";
import AuthUserProvider from "@/lib/AuthUser";
import "../firebaseConfig";

function MyApp({ Component, pageProps }: AppProps) {
  return (
    <>
      <CssBaseline />
      <AuthUserProvider>
        <Component {...pageProps} />
      </AuthUserProvider>
    </>
  );
}

export default MyApp;
