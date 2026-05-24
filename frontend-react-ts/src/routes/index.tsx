import { Routes, Route } from "react-router";
import { ProtectedRoute, PublicRoute } from "../components/common/AuthGuard";
import Home from "../views/home/Index";
import SignUp from "../views/auth/SignUp";
import SignIn from "../views/auth/SignIn";
import Dashboard from "../views/dashboard/Index";
import Profile from "../views/dashboard/profile/Index";
import UsersIndex from "../views/dashboard/users/Index";
import UserCreate from "../views/dashboard/users/Create";
import UserEdit from "../views/dashboard/users/Edit";
import NotFound from "../views/common/NotFound";

export default function AppRoutes() {
  return (
    <Routes>
      <Route path="/" element={<Home />} />

      <Route element={<PublicRoute />}>
        <Route path="/sign-up" element={<SignUp />} />
        <Route path="/sign-in" element={<SignIn />} />
      </Route>

      <Route element={<ProtectedRoute />}>
        <Route path="/dashboard" element={<Dashboard />} />
        <Route path="/dashboard/profile" element={<Profile />} />
        <Route path="/dashboard/users" element={<UsersIndex />} />
        <Route path="/dashboard/users/create" element={<UserCreate />} />
        <Route path="/dashboard/users/edit/:id" element={<UserEdit />} />
      </Route>

      <Route path="*" element={<NotFound />} />
    </Routes>
  );
}
