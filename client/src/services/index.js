// src/services/index.js

import axios from "axios";
import { useAuthStore } from "@/store/useAuthStore";

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

let isRefreshing = false;
let failedQueue = [];

const processQueue = (error, token = null) => {
  failedQueue.forEach((prom) => {
    if (error) {
      prom.reject(error);
    } else {
      prom.resolve(token);
    }
  });
  failedQueue = [];
};

authInstance.interceptors.response.use(
  (response) => response,
  async (error) => {
    const originalRequest = error.config;

    if (
      error.response?.status === 401 &&
      !originalRequest._retry &&
      !originalRequest.url.includes("/auth/refresh-token")
    ) {
      originalRequest._retry = true;

      if (isRefreshing) {
        return new Promise((resolve, reject) => {
          failedQueue.push({ resolve, reject });
        })
          .then(() => authInstance(originalRequest))
          .catch((err) => Promise.reject(err));
      }

      isRefreshing = true;

      try {
        await authInstance.post("/auth/refresh-token");
        processQueue(null);
        return authInstance(originalRequest);
      } catch (err) {
        processQueue(err);
        useAuthStore.getState().clearUser();
        return Promise.reject(err);
      } finally {
        isRefreshing = false;
      }
    }

    return Promise.reject(error);
  }
);
