import { type FC, useState, type FormEvent } from "react";
import { useNavigate, Link } from "react-router";
import { type AxiosError } from "axios";
import toast from "react-hot-toast";
import { useUserCreate } from "../../../hooks/user/useUserCreate";
import { type ValidationErrors } from "../../../types/common";
import useDocumentTitle from "../../../hooks/common/useDocumentTitle";
import DashboardLayout from "../../../layouts/DashboardLayout";
import { type UserResponse } from "../../../types/user";

const UserCreate: FC = () => {
  useDocumentTitle("Users");

  const navigate = useNavigate();
  const { mutate, isPending } = useUserCreate();

  const [name, setName] = useState<string>("");
  const [email, setEmail] = useState<string>("");
  const [password, setPassword] = useState<string>("");
  const [errors, setErrors] = useState<ValidationErrors>({});

  const createUser = async (e: FormEvent) => {
    e.preventDefault();
    mutate(
      {
        name,
        email,
        password,
      },
      {
        onSuccess: (data: UserResponse) => {
          toast.success(data?.message || "Create user successfully");
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
    <DashboardLayout>
      <div className="card border-0 rounded-4 shadow-sm">
        <div className="card-header">ADD USER</div>
        <div className="card-body">
          <form onSubmit={createUser}>
            <div className="form-group mb-3">
              <label className="mb-1 fw-bold">Full Name</label>
              <input
                type="text"
                value={name}
                onChange={(e) => setName(e.target.value)}
                className="form-control"
                placeholder="Full Name"
              />
              {errors.name && (
                <div className="alert alert-danger mt-2 rounded-4">
                  {errors.name}
                </div>
              )}
            </div>

            <div className="form-group mb-3">
              <label className="mb-1 fw-bold">Email address</label>
              <input
                type="email"
                value={email}
                onChange={(e) => setEmail(e.target.value)}
                className="form-control"
                placeholder="Email Address"
              />
              {errors.email && (
                <div className="alert alert-danger mt-2 rounded-4">
                  {errors.email}
                </div>
              )}
            </div>

            <div className="form-group mb-3">
              <label className="mb-1 fw-bold">Password</label>
              <input
                type="password"
                value={password}
                onChange={(e) => setPassword(e.target.value)}
                className="form-control"
                placeholder="Password"
              />
              {errors.password && (
                <div className="alert alert-danger mt-2 rounded-4">
                  {errors.password}
                </div>
              )}
            </div>

            <button
              type="submit"
              className="btn btn-md btn-primary rounded-4 shadow-sm border-0"
              disabled={isPending}
            >
              {isPending ? "Saving..." : "Save"}
            </button>

            <Link
              to="/dashboard/users"
              className="btn btn-md btn-secondary rounded-4 shadow-sm border-0 ms-2"
            >
              Cancel
            </Link>
          </form>
        </div>
      </div>
    </DashboardLayout>
  );
};

export default UserCreate;
