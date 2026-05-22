import { create } from "zustand";
import Cookies from "js-cookie";
import { type AuthState } from "../types/auth";
import { APP_CONSTANT } from "../utils/constant";

export const useAuthStore = create<AuthState>((set) => ({
  token: Cookies.get(APP_CONSTANT.COOKIES_TOKEN) || "",

  user: Cookies.get(APP_CONSTANT.COOKIES_USER) ? JSON.parse(Cookies.get(APP_CONSTANT.COOKIES_USER) as string) : null,

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
