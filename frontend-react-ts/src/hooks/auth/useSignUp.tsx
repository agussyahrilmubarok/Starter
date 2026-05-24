import { useMutation } from "@tanstack/react-query";
import { PublicApi } from "../../services/api";
import { type SignUpRequest, type SignUpResponse } from "../../types/auth";

export const useSignUp = () => {
  return useMutation({
    mutationFn: async (payload: SignUpRequest) => {
      const response = await PublicApi.post<SignUpResponse>("/auth/sign-up", payload);
      return response.data;
    },
  });
};
