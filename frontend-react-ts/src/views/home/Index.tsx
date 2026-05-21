import { type FC } from "react";
import { Link } from "react-router";

const Home: FC = () => {
  return (
    <div className="p-5 mb-4 bg-light rounded-5 shadow-sm">
      <div className="container-fluid py-5">
        <h1 className="display-5 fw-bold">frontend-react-ts</h1>
        <p className="col-md-12 fs-4">
          Frontend development with React TypeScript
        </p>
        <hr />
        <Link to="/sign-up" className="btn btn-primary btn-lg me-3">
          Sign Up
        </Link>
        <Link to="/sign-in" className="btn btn-secondary btn-lg">
          Sign In
        </Link>
      </div>
    </div>
  );
};

export default Home;
