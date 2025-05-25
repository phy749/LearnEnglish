import React, { useState } from "react";
import "../../Styles/ResetPassword.css";
import { Navigate, useNavigate } from "react-router-dom";


const ResetPassword = () => {
  const [step, setStep] = useState(1); // 1: nhập email, 2: đặt lại mật khẩu
  const [email, setEmail] = useState("");
  const [formData, setFormData] = useState({
    password: "",
    confirmPassword: "",
  });
  const navigate =useNavigate();
  const [error, setError] = useState("");
  const handleLogin= (e) => {
    e.preventDefault();
    navigate("/login");
  };
  const handleEmailSubmit = (e) => {
    e.preventDefault();
    if (!email) {
      setError("Email is required.");
      return;
    }

    setError("");
    console.log("Gửi email reset:", email);
    // TODO: Gửi email lên server
    setStep(2);
  };

  const handlePasswordChange = (e) => {
    const { name, value } = e.target;
    setFormData((prev) => ({ ...prev, [name]: value }));
  };

  const handlePasswordSubmit = (e) => {
    e.preventDefault();
    if (formData.password !== formData.confirmPassword) {
      setError("Passwords do not match.");
      return;
    }

    setError("");
    console.log("Reset with new password:", formData);
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
              placeholder="Enter your email"
              value={email}
              onChange={(e) => setEmail(e.target.value)}
              className="reset-input"
              required
            />
            {error && <p className="reset-error">{error}</p>}
            <button type="submit" className="reset-button">
              Send Reset Link
            </button>
            <button onClick={handleLogin} >Đăng nhập</button>
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
              onChange={handlePasswordChange}
              className="reset-input"
              required
            />
            <input
              type="password"
              name="confirmPassword"
              placeholder="Confirm Password"
              value={formData.confirmPassword}
              onChange={handlePasswordChange}
              className="reset-input"
              required
            />
            {error && <p className="reset-error">{error}</p>}
            <button type="submit" className="reset-button">
              Reset Password
            </button>
            <button onClick={handleLogin} >Đăng nhập</button>
          </form>
        </>
      )}
    </div>
  );
};

export default ResetPassword;
