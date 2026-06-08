import axios from "axios";
import Cookies from "js-cookie";
import { API_BASE_URL } from "../constants/api";
import { COOKIE_KEYS } from "../constants/cookie";

const baseURL = API_BASE_URL;

export const PublicApi = axios.create({
  baseURL,
  headers: { "Content-Type": "application/json" },
});

export const PrivateApi = axios.create({
  baseURL,
  headers: { "Content-Type": "application/json" },
});

PrivateApi.interceptors.request.use((config) => {
  const token = Cookies.get(COOKIE_KEYS.TOKEN);
  if (token) {
    config.headers.Authorization = `Bearer ${token}`;
  }
  return config;
});