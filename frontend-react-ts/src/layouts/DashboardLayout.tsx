import type React from "react";
import SideBar from "../components/dashboard/SideBar";
import NavBar from "../components/dashboard/NavBar";

type DashboardLayoutProps = {
  children: React.ReactNode;
};

const DashboardLayout: React.FC<DashboardLayoutProps> = ({ children }) => {
  return (
    <div>
      <NavBar />
      <div className="container mt-4 mb-5">
        <div className="row g-4">
          <div className="col-12 col-md-3">
            <SideBar />
          </div>
          <div className="col-12 col-md-9">{children}</div>
        </div>
      </div>
    </div>
  );
};

export default DashboardLayout;
