import Cookies from "js-cookie";
import { useRouter } from "vue-router";
import { COOKIE_KEYS } from "../../constants/cookie";

export const useSignOut = (): (() => void) => {
  const router = useRouter();
  const signout = (): void => {
    Cookies.remove(COOKIE_KEYS.TOKEN);
    Cookies.remove(COOKIE_KEYS.USER_DATA);
    router.push({ name: "signin" });
  };
  return signout;
};
