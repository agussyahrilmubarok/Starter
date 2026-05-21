import { Routes, Route } from "react-router";
import Home from "../views/home/Index";

export default function AppRoutes() {
    return (
        <Routes>
            <Route path="/" element={<Home />} />
        </Routes>
    );
}