import React, { useState } from "react";
import "../../Styles/FormLogin.css";
import { LoginFormRequest } from "../Service/LoginFormRequest";
import Toast from "../../Components/Toast";
import { useGoogleLogin } from '@react-oauth/google';
import { useNavigate } from "react-router-dom";

const LoginForm = () => {
  const [formData, setFormData] = useState({
    username: "",
    password: "",
  });
  const navigate = useNavigate();
  const [isLoading, setIsLoading] = useState(false);
  const [toast, setToast] = useState({
    message: "",
    type: "",
    visible: false,
  });

  const showToast = (message, type = "error") => {
    setToast({ message, type, visible: true });
    setTimeout(() => setToast({ ...toast, visible: false }), 3000);
  };

  const handleChange = (e) => {
    setFormData((prev) => ({  
      ...prev,
      [e.target.name]: e.target.value,
    }));
  };

  const validateForm = (username, password) => {
    if (!username || !password) {
      return "Vui lòng nhập đầy đủ thông tin";
    }
    if (username.length < 6 || password.length < 8) {
      return "Tên đăng nhập và mật khẩu phải có ít nhất 6 ký tự";
    }
    if (!/[a-z]/.test(password)) {
      return "Mật khẩu phải có ít nhất một chữ thường.";
    }
    if (!/[A-Z]/.test(password)) {
      return "Mật khẩu phải có ít nhất một chữ in hoa.";
    }
    if (!/[0-9]/.test(password)) {
      return "Mật khẩu phải có ít nhất một số.";
    }
    if (!/[^A-Za-z0-9]/.test(password)) {
      return "Mật khẩu phải có ít nhất một ký tự đặc biệt.";
    }
    return null;
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    setIsLoading(true);
    try {
      const { username, password } = formData;
      const validationError = validateForm(username, password);
      
      if (validationError) {
        showToast(validationError, "error");
        return;
      }

      const response = await LoginFormRequest(formData);
      if (response) {
        showToast("Đăng nhập thành công", "success");
        navigate("/dashboard"); 
      } else {
        showToast("Đăng nhập thất bại", "error");
      }
    } catch (error) {
      console.error("Login error:", error);
      showToast(error.message || "Có lỗi xảy ra khi đăng nhập", "error");
    } finally {
      setIsLoading(false);
    }
  };

  const handleRegister = (e) => {
    e.preventDefault();
    navigate("/register");
  };

  const handleForgotPassword = (e) => {
    e.preventDefault();
    navigate("/forgot-password");
  };

  const login = useGoogleLogin({
    onSuccess: async (response) => {
      try {
   
        const userInfo = await fetch('https://www.googleapis.com/oauth2/v3/userinfo', {
          headers: { Authorization: `Bearer ${response.access_token}` },
        }).then(res => res.json());

        // Lưu thông tin người dùng vào localStorage
        localStorage.setItem('user', JSON.stringify({
          email: userInfo.email,
          name: userInfo.name,
          picture: userInfo.picture,
          googleId: userInfo.sub,
          accessToken: response.access_token
        }));

        showToast("Đăng nhập với Google thành công", "success");
        navigate("/dashboard");
      } catch (error) {
        console.error("Google login error:", error);
        showToast("Có lỗi xảy ra khi đăng nhập với Google", "error");
      }
    },
    onError: () => {
      showToast("Đăng nhập với Google thất bại", "error");
    }
  });

  return (
    <div className="container">
      <form onSubmit={handleSubmit} className="form">
        <h2>Đăng nhập</h2>

        <input
          type="text"
          name="username"
          placeholder="Tên đăng nhập"
          value={formData.username}
          onChange={handleChange}
          className="input"
          disabled={isLoading}
          autoComplete="username"
        />

        <input
          type="password"
          name="password"
          placeholder="Mật khẩu"
          value={formData.password}
          onChange={handleChange}
          className="input"
          disabled={isLoading}
          autoComplete="current-password"
        />
        <a onClick={handleRegister}>Quên mật khẩu</a>

        <button type="submit" className="button" disabled={isLoading}>
          {isLoading ? "Đang xử lý..." : "Đăng nhập"}
        </button>
        <button onClick={handleRegister} type="button" className="button">
          Đăng ký
        </button>
        <button onClick={handleForgotPassword} type="button" className="button">
          Quên mật khẩu
        </button>
        <button onClick={() => login()} type="button" className="button">
          Đăng nhập với Google
        </button>
      </form>
      {toast.visible && (
        <Toast
          messages={toast.message}
          type={toast.type}
          onClose={() => setToast({ ...toast, visible: false })}
        />
      )}
    </div>
  );
};

export default LoginForm;
