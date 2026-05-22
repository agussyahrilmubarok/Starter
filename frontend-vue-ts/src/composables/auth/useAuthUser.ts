import Cookies from "js-cookie";
import { COOKIE_KEYS } from "../../constants/cookie";
import { type AuthUser } from "../../types/auth";

export const useAuthUser = (): AuthUser | null => {
  const user = Cookies.get(COOKIE_KEYS.USER_DATA);
  return user ? (JSON.parse(user) as AuthUser) : null;
};
