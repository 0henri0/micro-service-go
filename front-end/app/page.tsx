import Carousel from "@/components/homepage/Carousel";
import Header from "@/components/common/Header";
import ListProduct from "@/components/common/ListProduct";
import Menu from "@/components/common/Menu";
import SectionPadding from "@/components/homepage/SectionPadding";
import Loading from "../components/common/Loading";
import { Suspense } from "react";

export default function Home() {
  return (
    <>
      <Suspense fallback={<Loading />}>
        <Header />
        <main className="container mx-auto p-4 pt-0">
          <Menu />
          <Carousel />
          <SectionPadding />
          <ListProduct />
        </main>
      </Suspense>
    </>
  );
}
