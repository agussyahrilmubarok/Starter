export interface AxiosErrorResponse {
  response?: {
    data?: {
      message?: string;
      errors?: ValidationErrors;
    };
  };
}

export interface ValidationErrors {
  [key: string]: string;
}
