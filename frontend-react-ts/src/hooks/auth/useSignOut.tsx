import { useContext } from "react";
import { useNavigate } from "react-router";
import Cookies from "js-cookie";
import { AuthContext } from "../../context/AuthContext";

export const useSignOut = (): (() => void) => {
  const authContext = useContext(AuthContext);
  const { setIsAuthenticated } = authContext!;
  const navigate = useNavigate();
  const signout = (): void => {
    Cookies.remove("token");
    Cookies.remove("user");
    setIsAuthenticated(false);
    navigate("/login");
  };
  return signout;
};
