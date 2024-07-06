"use client";

import { faBagShopping, faCartShopping } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { Bars3Icon, MagnifyingGlassIcon } from "@heroicons/react/16/solid";
import React, { useState } from "react";
import Sidebar from "./Sidebar";

const Header: React.FC = () => {
  const [isOpen, setIsOpen] = useState(false);
  return (
    <>
      <header className="mx-auto flex items-center justify-between bg-white p-4 px-4 md:container">
        <div className="col-span-1 flex items-center">
          <img
            src="https://nathstore.site/images/nathstore-logo.png"
            alt="Avatar"
            className="w-36 rounded-full"
          />
        </div>

        {/* Search Bar - Hidden on mobile */}
        <div className="mx-4 hidden w-full max-w-md items-center rounded-lg border p-2 shadow-sm md:flex">
          <MagnifyingGlassIcon className="h-5 w-5 text-gray-400" />
          <input
            type="text"
            placeholder="Search for items..."
            className="ml-2 flex-grow border-none text-gray-600 focus:outline-none focus:ring-0"
          />
        </div>
        {/* Icons */}
        <div className="flex items-center space-x-4">
          <div className="relative">
            <FontAwesomeIcon
              icon={faBagShopping}
              className="h-6 w-6 cursor-pointer"
            />
            <span className="absolute -right-2 -top-2 inline-block h-4 w-4 rounded-full bg-green-500 text-center text-xs text-white">
              2
            </span>
          </div>
         
          <Bars3Icon
            className="h-6 w-6 text-gray-500 md:hidden cursor-pointer"
            onClick={() => setIsOpen(true)}
          />
          <button className="hidden items-center rounded border border-solid border-shopee bg-shopee px-4 py-2 font-bold text-white duration-300 ease-in hover:bg-white hover:text-shopee md:flex">
            <span className="text-sm">SHOPEE</span>
            <FontAwesomeIcon icon={faCartShopping} className="ml-2 h-4" />
          </button>
        </div>
      </header>
      <Sidebar isOpen={isOpen} setIsOpen={setIsOpen} />
    </>
  );
};

export default Header;
