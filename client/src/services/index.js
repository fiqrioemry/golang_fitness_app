// src/services/index.js

import axios from "axios";
import auth from "@/services/auth";

export const publicInstance = axios.create({
  baseURL: import.meta.env.VITE_API_SERVICES,
  headers: {
    "X-API-Key": import.meta.env.VITE_API_KEY,
  },
});

export const authInstance = axios.create({
  baseURL: import.meta.env.VITE_API_SERVICES,
  withCredentials: true,
  headers: {
    "X-API-Key": import.meta.env.VITE_API_KEY,
  },
});

authInstance.interceptors.response.use(
  (response) => response,
  async (error) => {
    if (error.response?.status === 401) {
      try {
        await auth.refreshToken();
        return authInstance(error.config);
      } catch (refreshError) {
        await auth.logout();
        return Promise.reject(refreshError);
      }
    }

    return Promise.reject(error);
  }
);
