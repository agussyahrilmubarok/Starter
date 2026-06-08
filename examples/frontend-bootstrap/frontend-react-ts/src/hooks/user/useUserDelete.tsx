import { useMutation } from "@tanstack/react-query";
import { PrivateApi } from "../../services/api";

export const useUserDelete = () => {
  return useMutation({
    mutationFn: async (id: string) => {
      const response = await PrivateApi.delete(`/users/${id}`);
      return response.data;
    },
  });
};
