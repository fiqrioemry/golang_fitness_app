// src/services/instructor.js
import { publicInstance, authInstance } from ".";

export const getAllInstructors = async () => {
  const res = await publicInstance.get("/instructors");
  return res.data;
};

export const getInstructorById = async (id) => {
  const res = await publicInstance.get(`/instructors/${id}`);
  return res.data;
};

export const createInstructor = async (data) => {
  const res = await authInstance.post("/instructors", data);
  return res.data;
};

export const updateInstructor = async (id, data) => {
  const res = await authInstance.put(`/instructors/${id}`, data);
  return res.data;
};

export const deleteInstructor = async (id) => {
  const res = await authInstance.delete(`/instructors/${id}`);
  return res.data;
};

export default {
  deleteInstructor,
  updateInstructor,
  createInstructor,
  getInstructorById,
  getAllInstructors,
};
