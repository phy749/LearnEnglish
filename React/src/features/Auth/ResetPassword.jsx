import React, { useState } from "react";
import "../../Styles/ResetPassword.css";
import { useNavigate } from "react-router-dom";

const ResetPassword = () => {
  const [step, setStep] = useState(1); // 1: nhập email, 2: đặt lại mật khẩu

  const [formData, setFormData] = useState({
    email: "",
    username: "",
    password: "",
    confirmPassword: "",
  });

  const [error, setError] = useState("");
  const navigate = useNavigate();

  const handleLogin = (e) => {
    e.preventDefault();
    navigate("/login");
  };

  // Hàm xử lý thay đổi các input
  const handleChange = (e) => {
    const { name, value } = e.target;
    setFormData((prev) => ({
      ...prev,
      [name]: value,
    }));
  };

  // Xử lý submit bước 1 (email + username)
  const handleEmailSubmit = (e) => {
    e.preventDefault();

    if (!formData.email) {
      setError("Email is required.");
      return;
    }
    if (!formData.username) {
      setError("Username is required.");
      return;
    }

    setError("");
    console.log("Gửi email reset:", formData.email, formData.username);
    // TODO: Gửi email và username lên server để nhận link reset mật khẩu
    setStep(2); // chuyển sang bước 2
  };

  // Xử lý submit bước 2 (đặt mật khẩu mới)
  const handlePasswordSubmit = (e) => {
    e.preventDefault();

    if (formData.password !== formData.confirmPassword) {
      setError("Passwords do not match.");
      return;
    }

    setError("");
    console.log("Reset with new password:", formData.password);
    // TODO: Gửi password mới lên server
  };

  return (
    <div className="reset-container">
      {step === 1 && (
        <>
          <h2 className="reset-title">Forgot Password</h2>
          <form onSubmit={handleEmailSubmit} className="reset-form">
            <input
              type="email"
              name="email"
              placeholder="Enter your email"
              value={formData.email}
              onChange={handleChange}
              className="reset-input"
              required
            />
            <input
              type="text"
              name="username"
              placeholder="Enter your username"
              value={formData.username}
              onChange={handleChange}
              className="reset-input"
              required
            />
            {error && <p className="reset-error">{error}</p>}
            <button type="submit" className="reset-button">
              Send Reset Link
            </button>
            <button onClick={handleLogin}>Đăng nhập</button>
          </form>
        </>
      )}

      {step === 2 && (
        <>
          <h2 className="reset-title">Set New Password</h2>
          <form onSubmit={handlePasswordSubmit} className="reset-form">
            <input
              type="password"
              name="password"
              placeholder="New Password"
              value={formData.password}
              onChange={handleChange}
              className="reset-input"
              required
            />
            <input
              type="password"
              name="confirmPassword"
              placeholder="Confirm Password"
              value={formData.confirmPassword}
              onChange={handleChange}
              className="reset-input"
              required
            />
            {error && <p className="reset-error">{error}</p>}
            <button type="submit" className="reset-button">
              Reset Password
            </button>
            <button onClick={handleLogin}>Đăng nhập</button>
          </form>
        </>
      )}
    </div>
  );
};

export default ResetPassword;
