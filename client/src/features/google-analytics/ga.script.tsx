import { GA_MEASUREMENT_ID, IS_PRODUCTION } from "@config/environment";
import * as gtag from "@features/google-analytics/gtag";

import { useRouter } from "next/router";
import Script from "next/script";
import { useEffect } from "react";

import type { PropsWithChildren } from "react";

export type GAScriptProps = PropsWithChildren<Record<never, never>>;

const GAScript = ({ children }: GAScriptProps): JSX.Element => {
  const router = useRouter();

  useEffect(() => {
    const handleRouteChange = (url: URL) => {
      /* invoke analytics function only for production */
      if (IS_PRODUCTION) gtag.pageview(url);
    };
    router.events.on(`routeChangeComplete`, handleRouteChange);

    return () => {
      router.events.off(`routeChangeComplete`, handleRouteChange);
    };
  }, [router.events]);

  return (
    <>
      {/* Global Site Tag (gtag.js) - Google Analytics */}
      {IS_PRODUCTION && (
        <>
          <Script
            strategy="afterInteractive"
            src={`https://www.googletagmanager.com/gtag/js?id=${GA_MEASUREMENT_ID}`}
          />
          <Script
            id="gtag-init"
            strategy="afterInteractive"
            dangerouslySetInnerHTML={{
              __html: `
            window.dataLayer = window.dataLayer || [];
            function gtag(){dataLayer.push(arguments);}
            gtag('js', new Date());
            gtag('config', '${GA_MEASUREMENT_ID}', {
              page_path: window.location.pathname,
            });
          `,
            }}
          />
        </>
      )}
      {children}
    </>
  );
};

export default GAScript;
