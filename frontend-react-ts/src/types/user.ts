export interface User {
  id: string;
  name: string;
  email: string;
  created_at: string;
  updated_at: string;
}

export interface UserRequest {
  name: string;
  email: string;
  password?: string;
}

export interface UserResponse {
  message: string;
  data: {
    id: string;
    name: string;
    email: string;
    created_at: string;
    updated_at: string;
  };
}
