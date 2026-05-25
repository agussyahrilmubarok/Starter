import type { FC } from "react";
import { Link, useNavigate } from "react-router";
import { APP_NAME } from "../../constants/app";
import { useAuthStore } from "../../stores/auth";

const NavBar: FC = () => {
  const navigate = useNavigate();
  const { user, signOut } = useAuthStore();

  const handleSignOut = () => {
    signOut();
    navigate("/sign-in");
  };

  return (
    <nav className="navbar navbar-expand-lg navbar-dark bg-dark px-3 shadow-sm">
      <Link className="navbar-brand fw-bold fs-5" to="/dashboard">
        ⚡ {APP_NAME}
      </Link>

      <button
        className="navbar-toggler"
        type="button"
        data-bs-toggle="collapse"
        data-bs-target="#navbarSupportedContent"
        aria-controls="navbarSupportedContent"
        aria-expanded="false"
        aria-label="Toggle navigation"
      >
        <span className="navbar-toggler-icon" />
      </button>

      <div className="collapse navbar-collapse" id="navbarSupportedContent">
        <ul className="navbar-nav ms-auto">
          {" "}
          <li className="nav-item dropdown">
            <a
              className="nav-link dropdown-toggle text-light"
              type="button"
              role="button"
              data-bs-toggle="dropdown"
              aria-expanded="false"
            >
              👤 { user?.email || "" }
            </a>
            <ul className="dropdown-menu dropdown-menu-end">
              {" "}
              <li>
                <Link className="dropdown-item" to="/dashboard/profile">
                  My Profile
                </Link>
              </li>
              <li>
                <hr className="dropdown-divider" />
              </li>
              <li>
                <button
                  className="dropdown-item text-danger"
                  onClick={handleSignOut}
                >
                  Sign Out
                </button>
              </li>
            </ul>
          </li>
        </ul>
      </div>
    </nav>
  );
};

export default NavBar;
