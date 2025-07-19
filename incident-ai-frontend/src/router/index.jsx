import { Routes, Route } from "react-router-dom";
import Home from "../pages/home";
import IncidentDetail from "../pages/incident/[id]";

export default function AppRouter() {
  return (
    <Routes>
      <Route path="/" element={<Home />} />
      <Route path="/incidents/:id" element={<IncidentDetail />} />
    </Routes>
  );
}
