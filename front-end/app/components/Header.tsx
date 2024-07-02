/* eslint-disable @next/next/no-img-element */
import {
  faBagShopping,
  faCartShopping,
} from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import {
  Bars3Icon,
  MagnifyingGlassIcon,
  ShoppingCartIcon,
} from "@heroicons/react/16/solid";
import React from "react";

const Header: React.FC = () => {
  return (
    <header className="container px-4 mx-auto flex items-center justify-between p-4 bg-white">
      <div className="col-span-1 flex items-center">
        <img
          src="https://nathstore.site/images/nathstore-logo.png"
          alt="Avatar"
          className="w-36 rounded-full"
        />
      </div>

      {/* Search Bar - Hidden on mobile */}
      <div className="hidden md:flex items-center border rounded-lg shadow-sm p-2 w-full max-w-md mx-4">
        <MagnifyingGlassIcon className="h-5 w-5 text-gray-400" />
        <input
          type="text"
          placeholder="Search for items..."
          className="ml-2 flex-grow border-none focus:ring-0 focus:outline-none text-gray-600"
        />
      </div>

      {/* Icons */}
      <div className="flex items-center space-x-4">
        <div className="relative hidden md:block">
          <FontAwesomeIcon icon={faBagShopping} className="h-6 w-6 cursor-pointer" />
          <span className="absolute -top-2 -right-2 inline-block w-4 h-4 bg-green-500 text-white text-xs rounded-full text-center">
            2
          </span>
        </div>
        <div className="relative">
          <MagnifyingGlassIcon className="h-6 w-6 text-gray-500 md:hidden" />
        </div>
        <Bars3Icon className="h-6 w-6 text-gray-500 md:hidden" />
        <button className="hidden md:flex items-center px-4 py-2 bg-shopee text-white font-bold ease-in duration-300 rounded hover:bg-white hover:text-shopee border border-solid border-shopee">
          <span className="text-sm">SHOPEE</span>
          <FontAwesomeIcon icon={faCartShopping} className="ml-2 h-4" />
        </button>
      </div>
    </header>
  );
};

export default Header;
