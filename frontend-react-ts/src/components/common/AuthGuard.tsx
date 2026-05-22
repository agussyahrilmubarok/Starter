import { Navigate, Outlet } from "react-router";
import { useAuthStore } from "../../stores/auth";

export const ProtectedRoute = () => {
  const { token } = useAuthStore();
  return token ? <Outlet /> : <Navigate to="/sign-in" replace />;
};

export const PublicRoute = () => {
  const { token } = useAuthStore();
  return !token ? <Outlet /> : <Navigate to="/dashboard" replace />;
};