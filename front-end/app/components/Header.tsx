import { Bars3Icon, MagnifyingGlassIcon, ShoppingCartIcon } from "@heroicons/react/16/solid";
import React from "react";


const Header: React.FC = () => {
  return (
    <header className="container mx-auto px-4 flex items-center justify-between p-4 bg-white shadow-md">
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
            <ShoppingCartIcon className="h-6 w-6 text-gray-500" />
            <span className="absolute -top-2 -right-2 inline-block w-4 h-4 bg-green-500 text-white text-xs rounded-full text-center">2</span>
        </div>
        <div className="relative">
          <MagnifyingGlassIcon className="h-6 w-6 text-gray-500 md:hidden" />
        </div>
        <Bars3Icon className="h-6 w-6 text-gray-500 md:hidden" />
        <button className="hidden md:flex items-center px-4 py-2 bg-[#ee4d2d] text-white rounded shadow">
          <span>Shopee</span>
          <svg className="ml-2 h-5 w-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
            <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M3 3h18M3 3a2 2 0 012-2h14a2 2 0 012 2M3 3l1.664 18.036A2 2 0 006.66 23h10.68a2 2 0 001.996-1.964L21 3M9 8h6M9 12h6M9 16h6" />
          </svg>
        </button>
      </div>
    </header>
  );
};

export default Header;
