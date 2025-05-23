  import React, { useEffect } from "react";
  import "../Styles/Toast.css";

  const Toast = ({ messages, type, duration = 3000, onClose }) => {
    useEffect(() => {
      console.log("Toast message:", messages);
      console.log("Toast message:", type);

      const timer = setTimeout(() => {
        onClose();
      }, duration);

      return () => clearTimeout(timer);
    }, [duration, onClose]);

    const getIcon = () => {
      switch (type) {
        case "success":
          return "👍"; // icon thành công
        case "error":
          return "❌"; // icon lỗi
        case "warning":
          return "⚠️"; // icon cảnh báo
        case "info":
          return "ℹ️"; // icon thông tin
        default:
          return "";
      }
    };
    return (
      <div className={`toast ${type ? `toast-${type}` : ""}`}>
        {getIcon()}
        {messages}
      </div>
    );
  };

  export default Toast;