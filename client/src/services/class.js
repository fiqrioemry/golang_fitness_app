// src/services/class.js
import { publicInstance, authInstance } from ".";
import { buildFormData } from "@/lib/utils";

export const getAllClasses = async (params) => {
  const res = await publicInstance.get("/classes", { params });
  return res.data;
};

// GET /api/classes/active
export const getActiveClasses = async () => {
  const res = await publicInstance.get("/classes/active");
  return res.data;
};

// GET /api/classes/:id
export const getClassById = async (id) => {
  const res = await publicInstance.get(`/classes/${id}`);
  return res.data;
};

// POST /api/classes
export const createClass = async (data) => {
  const formData = buildFormData(data);
  const res = await authInstance.post("/classes", formData);
  return res.data;
};

// PUT /api/classes/:id
export const updateClass = async (id, data) => {
  const formData = buildFormData(data);
  const res = await authInstance.put(`/classes/${id}`, formData);
  return res.data;
};

// DELETE /api/classes/:id
export const deleteClass = async (id) => {
  const res = await authInstance.delete(`/classes/${id}`);
  return res.data;
};

export const uploadClassGallery = async (id, gallery) => {
  console.log(id);
  console.log(gallery);
  const formData = buildFormData(gallery);
  const res = await authInstance.post(`/classes/${id}/gallery`, formData);
  return res.data;
};

export const deleteClassGallery = async (id, galleryId) => {
  const res = await authInstance.delete(`/classes/${id}/gallery/${galleryId}`);
  return res.data;
};

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

// GET /api/schedules/:id
export const getClassScheduleDetail = async (id) => {
  const res = await publicInstance.get(`/schedules/${id}`);
  return res.data;
};

// POST /api/schedules (Admin Only)
export const createClassSchedule = async (data) => {
  const res = await authInstance.post("/schedules", data);
  return res.data;
};

// PUT /api/schedules/:id (Admin Only)
export const updateClassSchedule = async (id, data) => {
  const res = await authInstance.put(`/schedules/${id}`, data);
  return res.data;
};

// DELETE /api/schedules/:id (Admin Only)
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
