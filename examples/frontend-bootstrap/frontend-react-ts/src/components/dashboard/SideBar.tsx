import type { FC } from "react";
import { NavLink } from "react-router";

const SideBar: FC = () => {
  const navLinkClass = ({ isActive }: { isActive: boolean }) =>
    `list-group-item list-group-item-action${isActive ? " active" : ""}`;

  return (
    <div className="card border-0 rounded-4 shadow-sm">
      <div className="card-header">MAIN MENU</div>

      <div className="card-body">
        <div className="list-group">
          <NavLink to="/dashboard" end className={navLinkClass}>
            Dashboard
          </NavLink>

          <NavLink to="/dashboard/users" className={navLinkClass}>
            Users
          </NavLink>
        </div>
      </div>
    </div>
  );
};

export default SideBar;
