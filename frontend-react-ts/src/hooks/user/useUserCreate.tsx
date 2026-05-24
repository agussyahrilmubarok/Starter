import { useMutation } from "@tanstack/react-query";
import { PrivateApi } from "../../services/api";
import { type UserRequest } from "../../types/user";

export const useUserCreate = () => {
  return useMutation({
    mutationFn: async (payload: UserRequest) => {
      const response = await PrivateApi.post("/users", payload);
      return response.data;
    },
  });
};
