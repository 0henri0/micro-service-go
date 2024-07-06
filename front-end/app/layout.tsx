// app/layout.tsx
import { Suspense } from "react";
import "./globals.css";
import Loading from "@/components/common/Loading";

export const metadata = {
  title: "Your Site Title",
  description: "Your site description",
  keywords: "keyword1, keyword2",
};

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <html lang="en">
      <head>{/* Any additional head elements */}</head>
      <body>
        <Suspense fallback={<Loading />}>{children}</Suspense>
      </body>
    </html>
  );
}
