import { type FC } from "react";
import { useAuthStore } from "../../stores/auth";
import DashboardLayout from "../../layouts/DashboardLayout";
import useDocumentTitle from "../../hooks/common/useDocumentTitle";

const Dashboard: FC = () => {
  useDocumentTitle("Dashboard");

  const { user } = useAuthStore();

  return (
    <DashboardLayout>
      <div className="card border-0 rounded-4 shadow-sm">
        <div className="card-header">DASHBOARD</div>
        <div className="card-body">
          Selamat Datang, <strong>{ user?.name || "" }</strong>
        </div>
      </div>
    </DashboardLayout>
  );
};

export default Dashboard;
