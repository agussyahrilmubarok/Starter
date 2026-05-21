import { useState, type FC, type FormEvent } from "react";
import { useNavigate } from "react-router";
import toast from "react-hot-toast";
import { useSignUp } from "../../hooks/auth/useSignUp";
import type { ValidationErrors } from "../../utils/error";

const SignUp: FC = () => {
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
        onSuccess: () => {
          toast.success("Sign up successfully");
          navigate("/sign-in");
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
    <div className="row justify-content-center">
      <div className="row justify-content-center">
        <div className="col-md-6">
          <div className="card border-0 rounded-4 shadow-sm">
            <div className="card-body">
              <h4 className="fw-bold text-center">Sign Up</h4>
              <hr />

              <form onSubmit={handleSignUp}>
                <div className="row">
                  <div className="col-md-12 mb-3">
                    <div className="form-group">
                      <label className="form-label fw-semibold">
                        Full Name
                      </label>
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
                  <div className="col-md-6 mb-3">
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

                  <div className="col-md-6 mb-3">
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
                        <div className="invalid-feedback">
                          {errors.password}
                        </div>
                      )}
                    </div>
                  </div>
                </div>

                <button
                  type="submit"
                  className="btn btn-primary w-100 rounded-4"
                  disabled={isPending}
                >
                  {isPending ? "Loading..." : "Sign Up"}
                </button>

              </form>

            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default SignUp;
