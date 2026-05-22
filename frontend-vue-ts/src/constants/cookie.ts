// Storage keys for Cookies, LocalStorage, or SessionStorage
export const COOKIE_KEYS = {
  TOKEN: "frontend-react-ts_auth_token", // Key for JWT Access Token
  USER_DATA: "frontend-react-ts_user_profile", // Key for storing basic user JSON
};

// Cookie Options (Security Configuration)
export const COOKIE_OPTIONS = {
  expires: 7, // Cookies expire in 7 days
  secure: true, // Only sent over HTTPS
  sameSite: "strict", // Protects against CSRF attacks
  path: "/", // Available across the whole site
};
