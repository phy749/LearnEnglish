import React, { useState } from "react";
import "../../Styles/ChangePassword.css";
import { Navigate, useNavigate } from "react-router-dom";

const ResetPasswordForm = () => {
  const [formData, setFormData] = useState({
    username: "",
    email: "",
    password: "",
    confirmPassword: "",
  });
  const navigate =useNavigate();
  const [error, setError] = useState("");
  const handleLogin= (e) => {
    e.preventDefault();
    navigate("/login");
  };
  const handleChange = (e) => {
    const { name, value } = e.target;
    setFormData((prevData) => ({
      ...prevData,
      [name]: value,
    }));
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    if (formData.password !== formData.confirmPassword) {
      setError("Passwords do not match.");
      return;
    }

    setError("");
    console.log("Reset password data:", formData);
  };

  return (
    <div className="reset-container">
      <h2 className="reset-title">Reset Password</h2>
      <form onSubmit={handleSubmit} className="reset-form">
        <input
          type="text"
          name="username"
          placeholder="Username"
          value={formData.username}
          onChange={handleChange}
          required
          className="reset-input"
        />
        <input
          type="email"
          name="email"
          placeholder="Email"
          value={formData.email}
          onChange={handleChange}
          required
          className="reset-input"
        />
        <input
          type="password"
          name="password"
          placeholder="New Password"
          value={formData.password}
          onChange={handleChange}
          required
          className="reset-input"
        />
        <input
          type="password"
          name="confirmPassword"
          placeholder="Confirm Password"
          value={formData.confirmPassword}
          onChange={handleChange}
          required
          className="reset-input"
        />

        {error && <p className="reset-error">{error}</p>}

        <button type="submit" className="reset-button">
          Reset Password
        </button>
        <button onClick={handleLogin} >Đăng nhập</button>
      </form>
    </div>
  );
};

export default ResetPasswordForm;
