import { type FC } from "react";
import { useAuthStore } from "../../../stores/auth";
import DashboardLayout from "../../../layouts/DashboardLayout";
import useDocumentTitle from "../../../hooks/common/useDocumentTitle";

const Profile: FC = () => {
  useDocumentTitle("Profile");

  const user = useAuthStore((state) => state.user);

  if (!user) {
    return (
      <DashboardLayout>
        <div className="card border-0 rounded-4 shadow-sm">
          <div className="card-body text-center text-muted py-5">
            User data not found. Please log in again.
          </div>
        </div>
      </DashboardLayout>
    );
  }

  return (
    <DashboardLayout>
      <div className="card border-0 rounded-4 shadow-sm">
        <div className="card-header fw-semibold">MY PROFILE</div>

        <div className="card-body">
          <div className="text-center mb-4">
            <div
              className="rounded-circle d-inline-flex align-items-center justify-content-center bg-dark text-white fw-bold"
              style={{ width: 80, height: 80, fontSize: 32 }}
            >
              {user.name?.charAt(0).toUpperCase() ?? "?"}
            </div>

            <h5 className="mt-3 mb-0">{user.name}</h5>
            <small className="text-muted">{user.email}</small>
          </div>

          <hr />

          <div className="row g-3">
            <div className="col-12 col-md-6">
              <label className="form-label text-muted small fw-semibold">
                FULL NAME
              </label>

              <p className="form-control-plaintext fw-medium">{user.name}</p>
            </div>

            <div className="col-12 col-md-6">
              <label className="form-label text-muted small fw-semibold">
                EMAIL
              </label>

              <p className="form-control-plaintext fw-medium">{user.email}</p>
            </div>

            <div className="col-12 col-md-6">
              <label className="form-label text-muted small fw-semibold">
                USER ID
              </label>

              <p className="form-control-plaintext text-muted small">
                {user.id}
              </p>
            </div>

            <div className="col-12 col-md-6">
              <label className="form-label text-muted small fw-semibold">
                MEMBER SINCE
              </label>

              <p className="form-control-plaintext fw-medium">
                {new Date(user.created_at).toLocaleDateString("en-US", {
                  day: "numeric",
                  month: "long",
                  year: "numeric",
                })}
              </p>
            </div>
          </div>
        </div>
      </div>
    </DashboardLayout>
  );
};

export default Profile;
