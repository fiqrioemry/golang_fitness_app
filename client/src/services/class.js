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

export const uploadClassGallery = async (id, images) => {
  const formData = buildFormData({ images });
  const res = await authInstance.post(`/classes/${id}/gallery`, formData);
  return res.data;
};
