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

// POST /api/schedule-templates/recurring
export const createRecurringScheduleTemplate = async (data) => {
  const res = await authInstance.post("/schedule-templates/recurring", data);
  return res.data;
};

// PUT /api/schedule-templates/:id
export const updateScheduleTemplate = async (id, data) => {
  const res = await authInstance.put(`/schedule-templates/${id}`, data);
  return res.data;
};
