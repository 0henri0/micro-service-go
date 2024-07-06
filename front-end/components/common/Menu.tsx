import React from "react";
import { DevicePhoneMobileIcon } from "@heroicons/react/16/solid";

const Menu: React.FC = () => {
  return (
    <nav className="border-t">
      <div className="flex items-center justify-between p-4">
        <div className="flex space-x-4 font-bold">
          <a href="#" className="text-gray-700 hover:text-gray-900">
            Tất cả sản phẩm
          </a>
          <a href="#" className="text-gray-700 hover:text-gray-900">
            Hàng sẵn
          </a>
          <a href="#" className="text-gray-700 hover:text-gray-900">
            Hàng order
          </a>
        </div>
        <div className="flex items-center space-x-2 font-bold">
          <DevicePhoneMobileIcon className="h-5 w-5 text-gray-500" />
          <span className="text-gray-700">Hotline</span>
          <span className="text-teal-600">0963494219</span>
        </div>
      </div>
    </nav>
  );
};

export default Menu;
