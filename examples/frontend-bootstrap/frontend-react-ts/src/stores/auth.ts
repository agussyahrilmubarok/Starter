import { create } from "zustand";
import Cookies from "js-cookie";
import { type AuthState } from "../types/auth";
import { COOKIE_KEYS } from "../constants/cookie";

export const useAuthStore = create<AuthState>((set) => ({
  token: Cookies.get(COOKIE_KEYS.TOKEN) || "",

  user: Cookies.get(COOKIE_KEYS.USER_DATA) ? JSON.parse(Cookies.get(COOKIE_KEYS.USER_DATA) as string) : null,

  setAuth: (token, user) => {
    Cookies.set(COOKIE_KEYS.TOKEN, token);
    Cookies.set(COOKIE_KEYS.USER_DATA, JSON.stringify(user));

    set({ token, user });
  },

  signOut: () => {
    Cookies.remove(COOKIE_KEYS.TOKEN);
    Cookies.remove(COOKIE_KEYS.USER_DATA);

    set({ token: "", user: null });
  },
}));
