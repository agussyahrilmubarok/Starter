import { useMutation } from "@tanstack/react-query";
import Api from "../../services/api";

interface SignUpRequest {
  name: string;
  email: string;
  password: string;
}

export const useSignUp = () => {
  return useMutation({
    mutationFn: async (payload: SignUpRequest) => {
      const response = await Api.post("/auth/sign-up", payload);
      return response.data;
    },
  });
};
