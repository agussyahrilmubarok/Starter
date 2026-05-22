import { useQuery } from "@tanstack/react-query";
import { PrivateApi } from "../../services/api";
import { type User } from "../../types/user";

export const useUserById = (id: string) => {
  return useQuery<User, Error>({
    queryKey: ["user", id],
    queryFn: async () => {
      const response = await PrivateApi.get(`/users/${id}`);
      return response.data.data as User;
    },
  });
};
