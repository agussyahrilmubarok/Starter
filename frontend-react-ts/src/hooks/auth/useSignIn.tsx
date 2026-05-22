import { useMutation } from "@tanstack/react-query";
import { PublicApi } from "../../services/api";
import { type SignInRequest, type SignInResponse } from "../../types/auth";

export const useSignIn = () => {
  return useMutation({
    mutationFn: async (payload: SignInRequest) => {
      const response = await PublicApi.post<SignInResponse>("/auth/sign-in", payload);
      return response.data;
    },
  });
};
