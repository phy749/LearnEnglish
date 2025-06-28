import {
  AppstoreOutlined,
  ShopOutlined,
  ShoppingCartOutlined,
  UserOutlined,
} from "@ant-design/icons";

import {
  MdOutlineLogout,
  MdDiscount,
  MdOutlineWarehouse,
  MdOutlineShoppingCart,
} from "react-icons/md";

import { Menu } from "antd";
import { useEffect, useState } from "react";
import { useLocation, useNavigate } from "react-router-dom";

import "../Styles/SideMenu.css"

function SideMenu() {
  const location = useLocation();
  const [selectedKeys, setSelectedKeys] = useState("/");

  const navigate = useNavigate();

  useEffect(() => {
    setSelectedKeys(location.pathname);
  }, [location.pathname]);

  const handleLogout = () => {
    localStorage.clear();
    navigate("/");
  };

  return (
    <div className="side-menu">
      <div className="user-info">
        {/* <img
          src="https://i.pravatar.cc/100"
          alt="avatar"
          className="avatar"
        /> */}
        <h3>John Smith</h3>
        <p>johnsmith@gmail.com</p>
      </div>

      <Menu
        mode="vertical"
        onClick={(item) => {
          if (item.key === "/") {
            handleLogout();
          } else {
            navigate(item.key);
          }
        }}
        selectedKeys={[selectedKeys]}
        className="side-menu-items"
        items={[
          {
            label: "Dashboard",
            icon: <AppstoreOutlined />,
            key: "/admin",
          },
          {
            label: "Game",
            key: "/admin/admingame",
            icon: <ShopOutlined />,
          },
          {
            label: "Acount game",
            key: "/admin/adminaccountgame",
            icon: <ShoppingCartOutlined />,
          },
          {
            label: "Account",
            key: "/admin/adminaccount",
            icon: <UserOutlined />,
          },
          {
            label: "Discount",
            key: "/admin/admindiscount",
            icon: <MdDiscount />,
          },
          {
            label: "Publisher",
            key: "/admin/adminpublisher",
            icon: <MdOutlineWarehouse />,
          },
          {
            label: "Cart",
            key: "/admin/admincard",
            icon: <MdOutlineShoppingCart />,
          },
          {
            label: "Logout",
            key: "/",
            icon: <MdOutlineLogout />,
          },
        ]}
      />
    </div>
  );
}

export default SideMenu;
