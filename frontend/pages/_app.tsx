// pages/_app.tsx
import "bootstrap/dist/css/bootstrap.min.css";
import "../public/css/globals.css";
import { AppProps } from "next/app";
import { useEffect } from "react";
import Script from "next/script";

function MyApp({ Component, pageProps }: AppProps) {
  useEffect(() => {
    require("bootstrap/dist/js/bootstrap.bundle.min.js");
  }, []);

  return (
    <>
      {/* <!-- Bootstrap JS and dependencies --> */}
      <Script src="https://code.jquery.com/jquery-3.5.1.slim.min.js" />
      {/* <Script src="https://stackpath.bootstrapcdn.com/bootstrap/4.5.2/js/bootstrap.min.js" defer/> */}
      <Component {...pageProps} />
    </>
  );
}

export default MyApp;
