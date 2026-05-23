import { Navigate, Outlet } from "react-router";
import { useAuthStore } from "../../stores/auth";

export const ProtectedRoute = () => {
  const { token } = useAuthStore();
  return token ? <Outlet /> : <Navigate to="/sign-in" replace />;
};

export const PublicRoute = () => {
  const isAuthenticated = useAuthStore((state) => !!state.token);
  return isAuthenticated ? <Outlet /> : <Navigate to="/sign-in" replace />;
};
