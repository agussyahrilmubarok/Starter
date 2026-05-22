import { create } from "zustand";
import Cookies from "js-cookie";
import { type AuthState } from "../types/auth";
import { APP_CONSTANT } from "../utils/constant";

export const useAuthStore = create<AuthState>((set) => ({
  token: Cookies.get("token") || "",

  user: Cookies.get("user") ? JSON.parse(Cookies.get("user") as string) : null,

  setAuth: (token, user) => {
    Cookies.set(APP_CONSTANT.COOKIES_TOKEN, token);
    Cookies.set(APP_CONSTANT.COOKIES_USER, JSON.stringify(user));

    set({ token, user });
  },

  signOut: () => {
    Cookies.remove(APP_CONSTANT.COOKIES_TOKEN);
    Cookies.remove(APP_CONSTANT.COOKIES_USER);

    set({ token: "", user: null });
  },
}));
