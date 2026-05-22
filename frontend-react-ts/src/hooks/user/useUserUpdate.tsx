import { useMutation } from "@tanstack/react-query";
import { PrivateApi } from "../../services/api";
import { type UserRequest } from "../../types/user";

export const useUserUpdate = () => {
  return useMutation({
    mutationFn: async ({
      id,
      payload,
    }: {
      id: string;
      payload: UserRequest;
    }) => {
      const response = await PrivateApi.put(`/users/${id}`, payload);
      return response.data;
    },
  });
};
