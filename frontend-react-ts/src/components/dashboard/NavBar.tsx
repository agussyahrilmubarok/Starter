import type { FC } from "react";
import { Link } from "react-router";

type NavBarProps = {
  userEmail?: string;
  onSignOut?: () => void;
};

const NavBar: FC<NavBarProps> = ({ userEmail = "John Doe", onSignOut }) => (
  <nav className="navbar navbar-expand-lg navbar-dark bg-dark px-3 shadow-sm">
    <Link className="navbar-brand fw-bold fs-5" to="/dashboard">
      ⚡ frontend-react-ts
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
            href="#"
            role="button"
            data-bs-toggle="dropdown"
            aria-expanded="false"
          >
            👤 {userEmail}
          </a>
          <ul className="dropdown-menu dropdown-menu-end">
            {" "}
            <li>
              <Link className="dropdown-item" to="/dashboard/profile">
                Profil Saya
              </Link>
            </li>
            <li>
              <hr className="dropdown-divider" />
            </li>
            <li>
              <button className="dropdown-item text-danger" onClick={onSignOut}>
                Sign Out
              </button>
            </li>
          </ul>
        </li>
      </ul>
    </div>
  </nav>
);

export default NavBar;
