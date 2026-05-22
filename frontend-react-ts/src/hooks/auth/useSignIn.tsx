import { useMutation } from "@tanstack/react-query";
import Cookies from "js-cookie";
import Api from "../../services/api";

interface SignInRequest {
  email: string;
  password: string;
}

export const useSignIn = () => {
  return useMutation({
    mutationFn: async (payload: SignInRequest) => {
      const response = await Api.post("/auth/sign-in", payload);
      const { data } = response.data;

      const token: string = data.token;
      const user = {
        id: data.user.id,
        name: data.user.name,
        email: data.user.email,
      };

      Cookies.set("token", token);
      Cookies.set("user", JSON.stringify(user));
      
      return response.data;
    },
  });
};
