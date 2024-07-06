import Link from "next/link";
import React from "react";

type ProductProps = {
  imageUrl: string;
  title: string;
  price: number;
  badgeText: string;
  category: string;
};

const Product: React.FC<ProductProps> = ({
  imageUrl,
  title,
  price = 0,
  badgeText,
  category,
}) => {
  return (
    <Link href="/product/1">
      <div className="mx-auto my-4 max-w-sm rounded-lg border border-gray-200 bg-white shadow-md">
        <div className="relative">
          <img
            className="rounded-[20px] p-2"
            src={imageUrl}
            alt="Product Image"
          />
          <span className="absolute left-4 top-4 rounded bg-blue-500 px-2 py-1 text-xs font-bold text-white">
            {badgeText}
          </span>
        </div>
        <div className="p-4 pt-0">
          <span className="mb-2 rounded text-[0.675rem] text-[#90908e]">
            {category}
          </span>
          <h5 className="line-clamp-2 pt-1 text-sm font-semibold hover:line-clamp-none md:text-base">
            {title}
          </h5>
          <div className="my-2 flex items-center">
            <div className="flex items-center text-yellow-400">
              <svg className="h-4 w-4 fill-current" viewBox="0 0 20 20">
                <path d="M10 15l-5.878 3.09L5.809 12 2 7.91 7.761 7 10 2l2.239 5 5.761.91L14.191 12l1.687 6.09z" />
              </svg>
              <svg className="h-4 w-4 fill-current" viewBox="0 0 20 20">
                <path d="M10 15l-5.878 3.09L5.809 12 2 7.91 7.761 7 10 2l2.239 5 5.761.91L14.191 12l1.687 6.09z" />
              </svg>
              <svg className="h-4 w-4 fill-current" viewBox="0 0 20 20">
                <path d="M10 15l-5.878 3.09L5.809 12 2 7.91 7.761 7 10 2l2.239 5 5.761.91L14.191 12l1.687 6.09z" />
              </svg>
              <svg className="h-4 w-4 fill-current" viewBox="0 0 20 20">
                <path d="M10 15l-5.878 3.09L5.809 12 2 7.91 7.761 7 10 2l2.239 5 5.761.91L14.191 12l1.687 6.09z" />
              </svg>
              <svg className="h-4 w-4 fill-current" viewBox="0 0 20 20">
                <path d="M10 15l-5.878 3.09L5.809 12 2 7.91 7.761 7 10 2l2.239 5 5.761.91L14.191 12l1.687 6.09z" />
              </svg>
            </div>
          </div>
          <div className="font-bold text-green-600 md:text-lg">
            {price.toLocaleString()} ₫
          </div>
        </div>
      </div>
    </Link>
  );
};

export default Product;
