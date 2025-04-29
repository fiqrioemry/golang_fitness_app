import { publicInstance, authInstance } from ".";

// =====================
// REVIEW
// =====================

// POST /api/reviews (Auth Required)
export const createReview = async (data) => {
  const res = await authInstance.post("/reviews", data);
  return res.data;
};

// GET /api/reviews/:classId (Public)
export const getReviewsByClass = async (classId) => {
  const res = await publicInstance.get(`/reviews/${classId}`);
  return res.data;
};

export default {
  createReview,
  getReviewsByClass,
};
