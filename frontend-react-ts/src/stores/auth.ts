import { create } from "zustand";
import Cookies from "js-cookie";
import { type AuthState } from "../types/auth";

export const useAuthStore = create<AuthState>((set) => ({
  token: Cookies.get("token") || "",

  user: Cookies.get("user") ? JSON.parse(Cookies.get("user") as string) : null,

  setAuth: (token, user) => {
    Cookies.set("token", token);
    Cookies.set("user", JSON.stringify(user));

    set({ token, user });
  },

  signOut: () => {
    Cookies.remove("token");
    Cookies.remove("user");

    set({ token: "", user: null });
  },
}));
