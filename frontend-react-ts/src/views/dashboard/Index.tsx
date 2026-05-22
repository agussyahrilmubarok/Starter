import { type FC } from "react";
import DashboardLayout from "../../layouts/DashboardLayout";

const Dashboard: FC = () => {
  return (
    <DashboardLayout>
      <div className="card border-0 rounded-4 shadow-sm">
        <div className="card-header">DASHBOARD</div>
        <div className="card-body">
          Selamat Datang, <strong></strong>
        </div>
      </div>
    </DashboardLayout>
  );
};

export default Dashboard;
