import { defineStore } from "pinia";
import Cookies from "js-cookie";
import { COOKIE_KEYS } from "../constants/cookie";
import { type AuthUser } from "../types/auth";

export const useAuthStore = defineStore("auth", {
  state: () => ({
    token: Cookies.get(COOKIE_KEYS.TOKEN) || "",
    user: Cookies.get(COOKIE_KEYS.USER_DATA)
      ? JSON.parse(Cookies.get(COOKIE_KEYS.USER_DATA)!)
      : null,
  }),

  getters: {
    isAuthenticated: (state) => !!state.token,
  },
  actions: {
    setAuth(token: string, user: AuthUser) {
      this.token = token;
      this.user = user;

      Cookies.set(COOKIE_KEYS.TOKEN, token);
      Cookies.set(COOKIE_KEYS.USER_DATA, JSON.stringify(user));
    },

    signOut() {
      this.token = "";
      this.user = null;

      Cookies.remove(COOKIE_KEYS.TOKEN);
      Cookies.remove(COOKIE_KEYS.USER_DATA);
    },
  },
});
