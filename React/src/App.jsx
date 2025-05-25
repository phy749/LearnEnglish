import LoginForm from "./features/Auth/LoginForm";
import RegisterForm from "./features/Auth/RegisterForm";
import ResetPassword from "./features/Auth/ResetPassword";

import { BrowserRouter, Routes, Route, Navigate } from "react-router-dom";
// import Dashboard from "./features/Dashboard/Dashboard";

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Navigate to="/login" replace />} />
        <Route path="/login" element={<LoginForm />} />
        <Route path="/register" element={<RegisterForm />} />
        <Route path="/forgotpassword" element={<ResetPassword />} />
        {/* <Route path="/dashboard" element={<Dashboard />} /> */}
        {/* <Route path="/forgot-password" element={<div>Forgot Password Page</div>} /> */}
      </Routes>
    </BrowserRouter>
  );
}

export default App;
