import { authInstance, publicInstance } from ".";

// GET /api/schedules
export const getAllClassSchedules = async () => {
  const res = await publicInstance.get("/schedules");
  return res.data;
};

// GET /api/schedules/status
export const getAllClassSchedulesWithStatus = async () => {
  const res = await authInstance.get("/schedules/status");
  return res.data;
};

// GET /api/schedules/booked
export const getAllBookedSchedule = async () => {
  const res = await authInstance.get("/schedules/booked");
  return res.data;
};

// GET /api/schedules/:id
export const getClassScheduleDetail = async (id) => {
  const res = await authInstance.get(`/schedules/${id}`);
  return res.data;
};

// POST /api/schedules
export const createClassSchedule = async (data) => {
  const res = await authInstance.post("/schedules", data);
  return res.data;
};

// PUT /api/schedules/:id
export const updateClassSchedule = async (id, data) => {
  const res = await authInstance.put(`/schedules/${id}`, data);
  return res.data;
};

// DELETE /api/schedules/:id
export const deleteClassSchedule = async (id) => {
  const res = await authInstance.delete(`/schedules/${id}`);
  return res.data;
};

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

// GET /api/schedule-templates
export const getAllRecuringSchedule = async () => {
  const res = await authInstance.get(`/schedule-templates`);
  return res.data;
};

// PUT /api/schedule-templates/:id
export const updateScheduleTemplate = async (id, data) => {
  const res = await authInstance.put(`/schedule-templates/${id}`, data);
  return res.data;
};

// DELETE /api/schedule-templates/:id
export const deleteScheduleTemplate = async (id) => {
  const res = await authInstance.delete(`/schedule-templates/${id}`);
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

// TODO : belum diimplementasikan - next feature
// // POST /api/schedule-templates/recurring
// export const createScheduleTemplate = async (data) => {
//   const res = await authInstance.post("/schedule-templates", data);
//   return res.data;
// };
