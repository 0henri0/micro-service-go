import React from "react";
import Product from "./Product";

const products = [
  {
    imageUrl: "https://cf.shopee.vn/file/86eafd25757191758c0b6a153d8e497a",
    category: "Lego",
    title: "[Mã Skamlsc814 Giảm 10% Đơn 100K] Đồ Chơi Lắp Ráp...",
    price: 273000,
    badgeText: "Hàng Order",
  },
  {
    imageUrl: "https://cf.shopee.vn/file/79863f82ef4d9982ee5d6257a4d01135",
    category: "Lego",
    title: "[Mã Skamlsc721 Giảm 10% Đơn 100K] Đồ Chơi Xếp Hình...",
    price: 273000,
    badgeText: "Hàng Order",
  },
  {
    imageUrl: "https://cf.shopee.vn/file/86eafd25757191758c0b6a153d8e497a",
    category: "Lego",
    title: "[Mã Skamlsc814 Giảm 10% Đơn 100K] Đồ Chơi Lắp Ráp...",
    price: 273000,
    badgeText: "Hàng Order",
  },
  {
    imageUrl: "https://cf.shopee.vn/file/86eafd25757191758c0b6a153d8e497a",
    category: "Lego",
    title: "[Mã Skamlsc814 Giảm 10% Đơn 100K] Đồ Chơi Lắp Ráp...",
    price: 273000,
    badgeText: "Hàng Order",
  },
  {
    imageUrl: "https://cf.shopee.vn/file/86eafd25757191758c0b6a153d8e497a",
    category: "Lego",
    title: "[Mã Skamlsc814 Giảm 10% Đơn 100K] Đồ Chơi Lắp Ráp...",
    price: 273000,
    badgeText: "Hàng Order",
  },
];

const ListProduct: React.FC = () => {
  return (
    <div className="mt-2 grid grid-cols-2 gap-2 md:gap-4 md:grid-cols-3 xl:grid-cols-4 xl:gap-8">
      {products.map((product, index) => (
        <Product
          key={index}
          imageUrl={product.imageUrl}
          category="Lego"
          title="[Mã Skamlsc814 Giảm 10% Đơn 100K] Đồ Chơi Lắp Ráp..."
          price={273000}
          badgeText="Hàng Order"
        />
      ))}
    </div>
  );
};

export default ListProduct;
