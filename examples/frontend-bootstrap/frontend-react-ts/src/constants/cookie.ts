// Storage keys for Cookies, LocalStorage, or SessionStorage
export const COOKIE_KEYS = {
  TOKEN: `${import.meta.env.VITE_APP_NAME || "frontend-react-ts"}_auth_token`,
  USER_DATA: `${import.meta.env.VITE_APP_NAME || "frontend-react-ts"}_user_profile`,
};

// Cookie Options (Security Configuration)
export const COOKIE_OPTIONS = {
  expires: 7, // Cookies expire in 7 days
  secure: true, // Only sent over HTTPS
  sameSite: "strict", // Protects against CSRF attacks
  path: "/", // Available across the whole site
};
