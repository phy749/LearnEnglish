import Dash from "../features/Dashboard/Dashboard";
import { Routes, Route } from "react-router-dom";
import Toast from '../Components/Toast';

function PageContent() {
  return (
    <div className="PageContent">
      <Routes>
        <Route path="/admin" element={<Dash />} />
      </Routes>
    </div>
  );
}
export default PageContent;
