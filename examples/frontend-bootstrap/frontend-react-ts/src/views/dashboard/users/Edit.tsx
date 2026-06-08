import { type FC, useState, type FormEvent } from "react";
import { useNavigate, useParams, Link } from "react-router";
import { type AxiosError } from "axios";
import toast from "react-hot-toast";
import { useUserById } from "../../../hooks/user/useUserById";
import { useUserUpdate } from "../../../hooks/user/useUserUpdate";
import { type ValidationErrors } from "../../../types/common";
import { type UserResponse, type User } from "../../../types/user";
import useDocumentTitle from "../../../hooks/common/useDocumentTitle";
import DashboardLayout from "../../../layouts/DashboardLayout";

const UserEditForm: FC<{ user: User }> = ({ user }) => {
  useDocumentTitle("Edit User");

  const navigate = useNavigate();
  const { mutate, isPending } = useUserUpdate();

  const [name, setName] = useState<string>(user.name);
  const [email, setEmail] = useState<string>(user.email);
  const [password, setPassword] = useState<string>("");
  const [errors, setErrors] = useState<ValidationErrors>({});

  const handleUpdate = async (e: FormEvent) => {
    e.preventDefault();

    mutate(
      {
        id: user.id,
        payload: { name, email, password },
      },
      {
        onSuccess: (data: UserResponse) => {
          toast.success(data?.message || "User updated successfully");
          navigate("/dashboard/users");
        },
        onError: (error: Error) => {
          const axiosError = error as AxiosError<{
            errors: Record<string, string>;
          }>;
          const validationErrors = axiosError.response?.data?.errors;
          if (validationErrors) {
            setErrors(validationErrors);
          } else {
            toast.error("Something went wrong");
          }
        },
      },
    );
  };

  return (
    <form onSubmit={handleUpdate}>
      <div className="form-group mb-3">
        <label className="mb-1 fw-bold">Full Name</label>
        <input
          type="text"
          value={name}
          onChange={(e) => setName(e.target.value)}
          className={`form-control ${errors.name ? "is-invalid" : ""}`}
          placeholder="Full Name"
        />
        {errors.name && <div className="invalid-feedback">{errors.name}</div>}
      </div>

      <div className="form-group mb-3">
        <label className="mb-1 fw-bold">Email Address</label>
        <input
          type="email"
          value={email}
          onChange={(e) => setEmail(e.target.value)}
          className={`form-control ${errors.email ? "is-invalid" : ""}`}
          placeholder="Email Address"
        />
        {errors.email && <div className="invalid-feedback">{errors.email}</div>}
      </div>

      <div className="form-group mb-3">
        <label className="mb-1 fw-bold">Password</label>
        <input
          type="password"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
          className={`form-control ${errors.password ? "is-invalid" : ""}`}
          placeholder="Leave blank to keep current password"
        />
        {errors.password && (
          <div className="invalid-feedback">{errors.password}</div>
        )}
      </div>

      <button
        type="submit"
        className="btn btn-md btn-primary rounded-4 shadow-sm border-0"
        disabled={isPending}
      >
        {isPending ? "Updating..." : "Update"}
      </button>

      <Link
        to="/dashboard/users"
        className="btn btn-md btn-secondary rounded-4 shadow-sm border-0 ms-2"
      >
        Cancel
      </Link>
    </form>
  );
};

const UserEdit: FC = () => {
  useDocumentTitle("Edit");

  const { id } = useParams();
  const { data: user } = useUserById(String(id));

  return (
    <DashboardLayout>
      <div className="card border-0 rounded-4 shadow-sm">
        <div className="card-header">EDIT USER</div>
        <div className="card-body">
          {user ? <UserEditForm user={user} /> : <p>Loading...</p>}
        </div>
      </div>
    </DashboardLayout>
  );
};

export default UserEdit;
