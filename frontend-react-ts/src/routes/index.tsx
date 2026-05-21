import { Routes, Route } from "react-router";
import Home from "../views/home/Index";
import SignUp from "../views/auth/SignUp";
import SignIn from "../views/auth/SignIn";

export default function AppRoutes() {
    return (
        <Routes>
            <Route path="/" element={<Home />} />
            <Route path="/sign-up" element={<SignUp />} />
            <Route path="/sign-in" element={<SignIn />} />
        </Routes>
    );
}