// src/services/class.js
import { publicInstance, authInstance } from ".";
import { buildFormData } from "@/lib/utils";

// =====================
// CLASS (Public + Admin)
// =====================

export const getAllClasses = async (params) => {
  const res = await publicInstance.get("/classes", { params });
  return res.data;
};

export const getActiveClasses = async () => {
  const res = await publicInstance.get("/classes/active");
  return res.data;
};

export const getClassById = async (id) => {
  const res = await publicInstance.get(`/classes/${id}`);
  return res.data;
};

export const createClass = async (data) => {
  const formData = buildFormData(data);
  const res = await authInstance.post("/classes", formData);
  return res.data;
};

export const updateClass = async (id, data) => {
  const formData = buildFormData(data);
  const res = await authInstance.put(`/classes/${id}`, formData);
  return res.data;
};

export const deleteClass = async (id) => {
  const res = await authInstance.delete(`/classes/${id}`);
  return res.data;
};

export const uploadClassGallery = async (id, formData) => {
  const res = await authInstance.post(`/classes/${id}/gallery`, formData);
  return res.data;
};

export const deleteClassGallery = async (id, galleryId) => {
  const res = await authInstance.delete(`/classes/${id}/gallery/${galleryId}`);
  return res.data;
};

// =====================
// CLASS SCHEDULE (Public + Admin)
// =====================

// GET /api/schedules
export const getAllClassSchedules = async () => {
  const res = await publicInstance.get("/schedules");
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
