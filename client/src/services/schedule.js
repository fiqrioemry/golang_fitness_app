import { authInstance } from ".";

// =====================
// SCHEDULE TEMPLATE (Admin Only)
// =====================

// POST /api/schedule-templates
export const createTemplate = async (data) => {
  const res = await authInstance.post("/schedule-templates", data);
  return res.data;
};

// POST /api/schedule-templates/auto-generate
export const autoGenerateSchedules = async (data) => {
  const res = await authInstance.post(
    "/schedule-templates/auto-generate",
    data
  );
  return res.data;
};

export default {
  createTemplate,
  autoGenerateSchedules,
};
