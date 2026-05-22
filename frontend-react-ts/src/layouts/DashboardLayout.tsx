import type React from "react";
import SideBar from "../components/dashboard/SideBar";
import NavBar from "../components/dashboard/NavBar";
import { useAuthStore } from "../stores/auth";

type DashboardLayoutProps = {
  children: React.ReactNode;
};

const DashboardLayout: React.FC<DashboardLayoutProps> = ({ children }) => {
  const { user, signOut } = useAuthStore();

  return (
    <div>
      <NavBar 
        userEmail={ user?.email || "User" }
        onSignOut={ signOut }
      />
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
