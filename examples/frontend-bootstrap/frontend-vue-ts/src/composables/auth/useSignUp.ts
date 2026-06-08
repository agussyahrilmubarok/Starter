import { useMutation } from "@tanstack/vue-query";
import { PublicApi } from "../../services/api";
import { type SignUpResponse, type SignUpRequest } from "../../types/auth";

export const useSignUp = () => {
  return useMutation({
    mutationFn: async (payload: SignUpRequest) => {
      const response = await PublicApi.post<SignUpResponse>(
        "/auth/sign-up",
        payload,
      );
      return response.data;
    },
  });
};
