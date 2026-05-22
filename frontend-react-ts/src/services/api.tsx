import axios from "axios";
import Cookies from "js-cookie";
import { APP_CONSTANT } from "../utils/constant";

const baseURL = "http://localhost:8080/api/v1";

export const PublicApi = axios.create({
  baseURL,
  headers: { "Content-Type": "application/json" },
});

export const PrivateApi = axios.create({
  baseURL,
  headers: { "Content-Type": "application/json" },
});

PrivateApi.interceptors.request.use((config) => {
  const token = Cookies.get(APP_CONSTANT.COOKIES_TOKEN);
  if (token) {
    config.headers.Authorization = `Bearer ${token}`;
  }
  return config;
});