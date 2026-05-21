import { useMutation } from '@tanstack/react-query';
import Api from '../../services/api';

interface SignInRequest {
    email: string;
    password: string;
}

export const useSignIn = () => {
    return useMutation({
        mutationFn: async (payload: SignInRequest) => {
            const response = await Api.post("/auth/sign-in", payload);
            return response.data;
        }
    });
};