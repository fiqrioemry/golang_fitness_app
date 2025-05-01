// src/services/user.js
import { authInstance } from ".";

// =====================
// USER (Admin Only)
// =====================

/**
 * GET /api/admin/users
 * @param {Object} params - { q, role, sort, page, limit }
 */
export const getAllUsers = async (params) => {
  const res = await authInstance.get("/admin/users", { params });
  console.log(res   );
  return res.data;
};

/**
 * GET /api/admin/users/:id
 */
export const getUserDetail = async (id) => {
  const res = await authInstance.get(`/admin/users/${id}`);
  return res.data;
};

/**
 * GET /api/admin/users/stats
 */
export const getUserStats = async () => {
  const res = await authInstance.get("/admin/users/stats");
  return res.data;
};

export default {
  getAllUsers,
  getUserDetail,
  getUserStats,
};
