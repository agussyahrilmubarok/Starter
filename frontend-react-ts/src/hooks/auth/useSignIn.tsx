import { useMutation } from "@tanstack/react-query";
import { type SignInResponse } from "../../types/auth";
import { PublicApi } from "../../services/api";

interface SignInRequest {
  email: string;
  password: string;
}

export const useSignIn = () => {
  return useMutation({
    mutationFn: async (payload: SignInRequest) => {
      const response = await PublicApi.post<SignInResponse>("/auth/sign-in", payload);
      return response.data;
    },
  });
};
