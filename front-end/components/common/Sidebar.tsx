// components/Sidebar.tsx
"use client";

import { useEffect, useRef } from 'react';
import Link from 'next/link';

interface SidebarProps {
  isOpen: boolean;
  setIsOpen: (isOpen: boolean) => void;
}

const Sidebar: React.FC<SidebarProps> = ({ isOpen, setIsOpen }) => {
  const sidebarRef = useRef<HTMLDivElement>(null);

  useEffect(() => {
    const handleClickOutside = (event: MouseEvent) => {
      if (sidebarRef.current && !sidebarRef.current.contains(event.target as Node)) {
        setIsOpen(false);
      }
    };

    if (isOpen) {
      document.addEventListener("mousedown", handleClickOutside);
    } else {
      document.removeEventListener("mousedown", handleClickOutside);
    }

    return () => {
      document.removeEventListener("mousedown", handleClickOutside);
    };
  }, [isOpen]);

  return (
    <div className="relative" >
      <div
        className={`fixed inset-0 bg-gray-800 bg-opacity-50 z-40 transition-opacity duration-300 ${isOpen ? 'opacity-100 visible' : 'opacity-0 invisible'}`}
      >
        <div
          ref={sidebarRef}
          className={`fixed right-0 top-0 h-full w-64 pt-2 bg-white z-50 transform transition-transform duration-300 ${isOpen ? 'translate-x-0' : 'translate-x-full'}`}
        >
        <div className="flex">
            <img className="h-10" src="https://nathstore.site/images/nathstore-logo.png" />
          <button
            className="absolute top-3 right-4"
            onClick={() => setIsOpen(false)}
          >
            <span className="material-icons">close</span>
          </button>
          </div>
          <div className="flex flex-col space-y-4 p-4">
            <input
              type="text"
              placeholder="Search for items..."
              className="p-2 border border-gray-300 rounded"
            />

            <a href="#" className="text-teal-600">Browse Categories</a>

            <nav className="flex flex-col space-y-2">
              <Link href="/" className="text-lg font-semibold">Home</Link>
              <Link href="/shop" className="text-lg font-semibold">Shop</Link>
              <Link href="/mega-menu" className="text-lg font-semibold">Mega Menu</Link>
              <Link href="/blog" className="text-lg font-semibold">Blog</Link>
              <Link href="/pages" className="text-lg font-semibold">Pages</Link>
              <Link href="/language" className="text-lg font-semibold">Language</Link>
            </nav>

            <div className="mt-4">
              <h3 className="text-gray-600">Follow Us</h3>
              <div className="flex space-x-4 mt-2">
                <a href="#" className="text-gray-600">Facebook</a>
                <a href="#" className="text-gray-600">Instagram</a>
                <a href="#" className="text-gray-600">TikTok</a>
                <a href="#" className="text-gray-600">Shopify</a>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Sidebar;
