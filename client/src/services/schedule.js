import { authInstance } from ".";

// =====================
// SCHEDULE TEMPLATE (Admin Only)
// =====================

// PUT /api/schedule-templates/:id
export const getAllRecuringSchedule = async () => {
  const res = await authInstance.get(`/schedule-templates`);
  return res.data.data;
};

// POST /api/schedule-templates/auto-generate
export const autoGenerateSchedules = async (data) => {
  const res = await authInstance.post("/schedule-templates/generate", data);
  return res.data;
};

// POST /api/schedule-templates/recurring
export const createScheduleTemplate = async (data) => {
  const res = await authInstance.post("/schedule-templates", data);
  return res.data;
};

// PUT /api/schedule-templates/:id
export const updateScheduleTemplate = async (id, data) => {
  const res = await authInstance.put(`/schedule-templates/${id}`, data);
  return res.data;
};

// PUT /api/schedule-templates/:id/run
export const runScheduleTemplate = async (id) => {
  const res = await authInstance.post(`/schedule-templates/${id}/run`);
  return res.data;
};

export const stopScheduleTemplate = async (id) => {
  const res = await authInstance.post(`/schedule-templates/${id}/stop`);
  return res.data;
};
