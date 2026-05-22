import type { FC } from "react";
import { Link } from "react-router";

const SideBar: FC = () => {
  return (
    <div className="card border-0 rounded-4 shadow-sm">
      <div className="card-header">MAIN MENU</div>
      <div className="card-body">
        <div className="list-group">
          <Link
            to="/dashboard"
            className="list-group-item list-group-item-action"
          >
            Dashboard
          </Link>

          <Link
            to="/dashboard/users"
            className="list-group-item list-group-item-action"
          >
            Users
          </Link>
        </div>
      </div>
    </div>
  );
};

export default SideBar;
