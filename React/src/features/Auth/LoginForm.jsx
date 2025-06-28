import React, { useState } from "react";
import "../../Styles/FormLogin.css"; // Bạn có thể đổi tên file này thành FormRegister.css nếu muốn
import { LoginFormRequest, GoogleLoginRequest } from "../Service/Request";
import Toast from "../../Components/Toast";
import { useGoogleLogin } from '@react-oauth/google';
import { FaEye, FaEyeSlash ,FaUser} from "react-icons/fa";
import { useNavigate } from "react-router-dom";

const LoginForm = () => {
  const [showPassword, setShowPassword] = useState(false);
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
    setTimeout(() => setToast((prev) => ({ ...prev, visible: false })), 3000);
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
   
    if (!/[0-9]/.test(password)) {
      return "Mật khẩu phải có ít nhất một số.";
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
        setIsLoading(false);
        return;
      }
      const response = await LoginFormRequest(formData);
      if (response) {
        localStorage.setItem('access_token', response.access_token);
        showToast(response.message || "Đăng nhập thành công", "success");
        navigate("/dashboard");
      } else {
        showToast("Đăng nhập thất bại", "error");
      }
    } catch (error) {
      console.error("Login error:", error);
      showToast("Có lỗi xảy ra khi đăng nhập", "error");
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
    navigate("/forgotpassword");
  };

  const login = useGoogleLogin({
    onSuccess: async (response) => {
      try {
        await GoogleLoginRequest(response.access_token);
        showToast("Đăng nhập với Google thành công", "success");
        navigate("/dashboard");
      } catch (error) {
        console.error("Google login error:", error);
        showToast("Có lỗi xảy ra khi đăng nhập với Google", "error");
      }
    },
    onError: () => {
      showToast("Đăng nhập với Google thất bại", "error");
    },
  });

  return (
    <div className="register-container">
      <form onSubmit={handleSubmit} className="register-form">
        <h2>Đăng nhập</h2>
        <div className="password-wrapper">
          <input
            type="text"
            name="username"
            placeholder="Tên đăng nhập"
            value={formData.username}
            onChange={handleChange}
            disabled={isLoading}
            autoComplete="username"
          />
          <span className="toggle-password">
            <FaUser />
          </span>
        </div>

        <div className="password-wrapper">
          <input
            type={showPassword ? "text" : "password"}
            name="password"
            placeholder="Mật khẩu"
            value={formData.password}
            onChange={handleChange}
            disabled={isLoading}
            autoComplete="current-password"
          />
          <span
            className="toggle-password"
            onClick={() => setShowPassword((prev) => !prev)}
          >
            {showPassword ? <FaEyeSlash /> : <FaEye />}
          </span>
        </div>
        <div style={{ display: "flex", justifyContent: "space-between", marginBottom: "15px", marginTop: "15px" }}>
          <a
            href="#!"
            onClick={handleRegister}
            style={{ color: "#4a90e2", cursor: "pointer" }}
          >
            Bạn chưa có tài khoản?
          </a>
          <a
            href="#!"
            onClick={handleForgotPassword}
            style={{ color: "#4a90e2", cursor: "pointer" }}
          >
            Quên mật khẩu
          </a>
        </div>

        <button type="submit" disabled={isLoading}>
          {isLoading ? "Đang xử lý..." : "Đăng nhập"}
        </button>
       

        <button onClick={() => login()} type="button" style={{ display: 'flex', alignItems: 'center', justifyContent: 'center', gap: '10px',   }}>
          <img src="/src/Image/th.jpg" alt="Google" style={{ width: '25px', height: '25px' }} />
          Tiếp tục với Google
        </button>
      </form>

      {toast.visible && (
        <Toast
          messages={toast.message}
          type={toast.type}
          onClose={() => setToast((prev) => ({ ...prev, visible: false }))}
        />
      )}
    </div>
  );
};

export default LoginForm;
