import { useState, type FC, type FormEvent } from "react";
import { Link, useNavigate } from "react-router";
import { type AxiosError } from "axios";
import toast from "react-hot-toast";
import { useSignUp } from "../../hooks/auth/useSignUp";
import { type ValidationErrors } from "../../types/common";
import { type SignUpResponse } from "../../types/auth";
import useDocumentTitle from "../../hooks/common/useDocumentTitle";

const SignUp: FC = () => {
  useDocumentTitle("Sign Up");

  const navigate = useNavigate();
  const { mutate, isPending } = useSignUp();

  const [name, setName] = useState<string>("");
  const [email, setEmail] = useState<string>("");
  const [password, setPassword] = useState<string>("");
  const [errors, setErrors] = useState<ValidationErrors>({});

  const handleSignUp = async (e: FormEvent) => {
    e.preventDefault();

    mutate(
      {
        name,
        email,
        password,
      },
      {
        onSuccess: (data: SignUpResponse) => {
          toast.success(data?.message || "Signed in successfully");
          navigate("/sign-in");
        },
        onError: (error: Error) => {
          const axiosError = error as AxiosError<{ errors: Record<string, string> }>;
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
    <div className="min-vh-100 d-flex align-items-center justify-content-center">
      <div className="col-11 col-sm-10 col-md-8 col-lg-5">
        <div className="card border-0 rounded-4 shadow-lg">
          <div className="card-body p-4 p-md-5">
            <h3 className="fw-bold text-center mb-2">Create Account</h3>
            <p className="text-secondary text-center mb-4">
              Sign up to continue
            </p>

            <form onSubmit={handleSignUp}>
              <div className="row">
                <div className="col-md-12 mb-3">
                  <div className="form-group">
                    <label className="form-label fw-semibold">Full Name</label>
                    <input
                      type="text"
                      value={name}
                      onChange={(e) => setName(e.target.value)}
                      className={`form-control rounded-3 ${
                        errors.name ? "is-invalid" : ""
                      }`}
                      placeholder="Enter your full name"
                    />
                    {errors.name && (
                      <div className="invalid-feedback">{errors.name}</div>
                    )}
                  </div>
                </div>
              </div>

              <div className="row">
                <div className="col-md-12 mb-3">
                  <div className="form-group">
                    <label className="form-label fw-semibold">
                      Email Address
                    </label>
                    <input
                      type="email"
                      value={email}
                      onChange={(e) => setEmail(e.target.value)}
                      className={`form-control rounded-3 ${
                        errors.email ? "is-invalid" : ""
                      }`}
                      placeholder="Enter your email"
                    />

                    {errors.email && (
                      <div className="invalid-feedback">{errors.email}</div>
                    )}
                  </div>
                </div>
              </div>

              <div className="row">
                <div className="form-group">
                  <label className="form-label fw-semibold">Password</label>
                  <input
                    type="password"
                    value={password}
                    onChange={(e) => setPassword(e.target.value)}
                    className={`form-control rounded-3 ${
                      errors.password ? "is-invalid" : ""
                    }`}
                    placeholder="Enter your password"
                  />
                  {errors.password && (
                    <div className="invalid-feedback">{errors.password}</div>
                  )}
                </div>
              </div>

              <button
                type="submit"
                className="btn btn-primary w-100 rounded-4 mt-4"
                disabled={isPending}
              >
                {isPending ? "Loading..." : "Sign Up"}
              </button>

              <div className="text-center mt-3">
                <span className="text-secondary">Already have an account?</span>
                <Link
                  to="/sign-in"
                  className="text-decoration-none fw-semibold ms-1"
                >
                  Sign In
                </Link>
              </div>
            </form>
          </div>
        </div>
      </div>
    </div>
  );
};

export default SignUp;
