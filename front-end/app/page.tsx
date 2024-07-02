import Image from "next/image";
import Header from "./components/Header";
import Menu from "./components/Menu";
import Carousel from "./components/Carousel";
import  SectionPadding from "./components/SectionPadding";

export default function Home() {
  return (
    <>
      <Header />
      <main className="md:container md:mx-auto">
        <Menu />
        <Carousel />
        <SectionPadding />
      </main>
    </>
  );
}
