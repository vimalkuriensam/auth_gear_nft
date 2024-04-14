import React from "react";
import { Navigate, Route, Routes } from "react-router-dom";

import { Home } from "../pages";
import customHistory from "../utils/history/history";
import CustomRouter from "./CustomRouter";

const AppRoutes = () => {
  return (
    <CustomRouter history={customHistory}>
      <Routes>
        <Route path="*" element={<Navigate to={"/home"} replace />} />
        <Route path="/home" element={<Home />} />
      </Routes>
    </CustomRouter>
  );
};

export default AppRoutes;
