import { authInstance } from ".";
import { buildFormData } from "@/lib/utils";

// GET /api/user/profile
export const getProfile = async () => {
  const res = await authInstance.get("/user/profile");
  return res.data;
};

// PUT /api/user/profile
export const updateProfile = async (data) => {
  const res = await authInstance.put("/user/profile", data);
  return res.data;
};

// PUT /api/user/profile/avatar
export const updateAvatar = async (data) => {
  const formData = buildFormData(data);
  const res = await authInstance.put("/user/profile/avatar", formData);
  return res.data;
};

// GET /api/user/packages
export const getUserPackages = async () => {
  const res = await authInstance.get("/user/packages");
  return res.data;
};

// GET /api/user/transactions
export const getUserTransactions = async (page = 1, limit = 10) => {
  const res = await authInstance.get("/user/transactions", {
    params: { page, limit },
  });
  return res.data;
};

// GET /api/user/bookings
export const getUserBookings = async (page = 1, limit = 10) => {
  const res = await authInstance.get("/user/bookings", {
    params: { page, limit },
  });
  return res.data;
};
