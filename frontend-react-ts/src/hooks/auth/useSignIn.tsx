import { useMutation } from "@tanstack/react-query";
import Api from "../../services/api";
import type { SignInResponse } from "../../types/auth";

interface SignInRequest {
  email: string;
  password: string;
}

export const useSignIn = () => {
  return useMutation({
    mutationFn: async (payload: SignInRequest) => {
      const response = await Api.post<SignInResponse>("/auth/sign-in", payload);
      
      return response.data;
    },
  });
};
