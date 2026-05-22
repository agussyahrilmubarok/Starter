import { useState, type FC, type FormEvent } from "react";
import { Link, useNavigate } from "react-router";
import toast from "react-hot-toast";
import { useSignIn } from "../../hooks/auth/useSignIn";
import type { ValidationErrors } from "../../utils/error";

const SignIn: FC = () => {
  const navigate = useNavigate();
  const { mutate, isPending } = useSignIn();

  const [email, setEmail] = useState<string>("");
  const [password, setPassword] = useState<string>("");
  const [errors, setErrors] = useState<ValidationErrors>({});

  const handleSignIn = async (e: FormEvent) => {
    e.preventDefault();

    mutate(
      {
        email,
        password,
      },
      {
        onSuccess: (data: any) => {
          toast.success(data?.message || "Sign in successfully");
          navigate("/dashboard");
        },
        onError: (error: any) => {
          console.log(error.response?.data?.errors);
          const validationErrors = error.response?.data?.errors;
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
            <h3 className="fw-bold text-center mb-2">Sign In</h3>
            <p className="text-secondary text-center mb-4">
              Welcome back! Please enter your details
            </p>

            <form onSubmit={handleSignIn}>
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
                {isPending ? "Loading..." : "Sign In"}
              </button>

              <div className="text-center mt-3">
                <span className="text-secondary">Don't have an account?</span>
                <Link
                  to="/sign-up"
                  className="text-decoration-none fw-semibold ms-1"
                >
                  Sign Up
                </Link>
              </div>
            </form>
          </div>
        </div>
      </div>
    </div>
  );
};

export default SignIn;
