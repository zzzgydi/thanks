import type { MetaFunction } from "@remix-run/node";
import {
  isRouteErrorResponse,
  Links,
  Meta,
  Outlet,
  Scripts,
  ScrollRestoration,
  useRouteError,
} from "@remix-run/react";
import "~/assets/global.css";

export const meta: MetaFunction = () => {
  return [
    { title: "OpenSource Thanks" },
    {
      name: "description",
      content: "Thank You for Your OpenSource Contributions",
    },
    { property: "og:title", content: "OpenSource Thanks" },
    {
      property: "og:description",
      content: "Thank You for Your OpenSource Contributions",
    },
    { property: "og:type", content: "website" },
    { property: "og:url", content: "https://opensourcethanks.com/" },
    { property: "og:site_name", content: "OpenSource Thanks" },
    { property: "twitter:title", content: "OpenSource Thanks" },
    {
      property: "twitter:description",
      content: "Thank You for Your OpenSource Contributions",
    },
    { property: "twitter:card", content: "summary_large_image" },
    { property: "twitter:url", content: "https://opensourcethanks.com/" },
  ];
};

export const links = () => [
  { rel: "icon", type: "image/png", href: "/logo.png" },
];

function Layout({ children }: { children: React.ReactNode }) {
  return (
    <html lang="en">
      <head>
        <meta charSet="utf-8" />
        <meta name="viewport" content="width=device-width, initial-scale=1" />
        <Meta />
        <Links />
      </head>
      <body>
        {children}
        <ScrollRestoration />
        <Scripts />
      </body>
    </html>
  );
}

export default function App() {
  return (
    <Layout>
      <Outlet />
    </Layout>
  );
}

export function ErrorBoundary() {
  const error = useRouteError();

  return (
    <Layout>
      <div className="relative flex min-h-screen min-h-dvh flex-col bg-background">
        {isRouteErrorResponse(error) ? (
          <div className="container">
            <h1>
              {error.status} {error.statusText}
            </h1>
            <p>{error.data}</p>
          </div>
        ) : error instanceof Error ? (
          <div className="container">
            <h1>Error</h1>
            <p>{error.message}</p>
            <p>The stack trace is:</p>
            <pre>{error.stack}</pre>
          </div>
        ) : (
          <div className="container">
            <h1>Unknown Error</h1>
          </div>
        )}
      </div>
    </Layout>
  );
}
