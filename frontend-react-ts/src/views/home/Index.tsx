import { type FC } from "react";
import { Link } from "react-router";
import useDocumentTitle from "../../hooks/common/useDocumentTitle";

const Home: FC = () => {
  useDocumentTitle("Home");

  return (
    <div className="min-vh-100 d-flex align-items-center justify-content-center">
      <div className="container">
        <div className="row justify-content-center">
          <div className="col-lg-8 col-sm-6">
            <div className="card border-0 rounded-5 shadow-lg">
              <div className="card-body p-5 text-center">
                <span className="badge bg-primary px-3 py-2 mb-3 rounded-pill">
                  React + TypeScript
                </span>
                <h1 className="display-3 fw-bold mb-3">frontend-react-ts</h1>
                <div className="d-flex flex-column flex-sm-row gap-3 justify-content-center">
                  <Link
                    to="/sign-up"
                    className="btn btn-primary btn-lg px-5 rounded-pill"
                  >
                    Get Started
                  </Link>
                  <Link
                    to="/sign-in"
                    className="btn btn-dark btn-lg px-5 rounded-pill"
                  >
                    Sign In
                  </Link>
                </div>
                <div className="mt-4 text-dark opacity-50 small">
                  starter.agussyahrilmubarok.github.io
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Home;
