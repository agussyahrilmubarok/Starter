export interface AuthUser {
  id: string;
  name: string;
  email: string;
  created_at: string;
  updated_at: string;
}

export interface AuthState {
  token: string;
  user: AuthUser | null;
  setAuth: (token: string, user: AuthUser) => void;
  signOut: () => void;
}

export interface SignUpRequest {
  email: string;
  password: string;
}

export interface SignUpResponse {
  message: string;
  data: {
    token: string;
    user: {
      id: string;
      name: string;
      email: string;
      created_at: string;
      updated_at: string;
    };
  };
}

export interface SignInRequest {
  email: string;
  password: string;
}

export interface SignInResponse {
  message: string;
  data: {
    token: string;
    user: {
      id: string;
      name: string;
      email: string;
      created_at: string;
      updated_at: string;
    };
  };
}
