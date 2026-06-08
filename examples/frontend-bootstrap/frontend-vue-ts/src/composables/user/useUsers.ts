import { useQuery } from "@tanstack/vue-query";
import { PrivateApi } from "../../services/api";
import { type User } from "../../types/user";

export const useUsers = () => {
  return useQuery<User[], Error>({
    queryKey: ["users"],
    queryFn: async () => {
      const response = await PrivateApi.get("/users");
      return response.data.data as User[];
    },
  });
};
