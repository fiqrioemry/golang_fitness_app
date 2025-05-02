import { publicInstance, authInstance } from ".";
import { buildFormData } from "@/lib/utils";

// =====================
// PACKAGE (Public + Admin)
// =====================

export const getAllPackages = async () => {
  const res = await publicInstance.get("/packages");
  return res.data;
};

export const getPackageById = async (id) => {
  const res = await publicInstance.get(`/packages/${id}`);
  return res.data;
};

export const createPackage = async (data) => {
  const formData = buildFormData(data);
  const res = await authInstance.post("/packages", formData);
  return res.data;
};

export const updatePackage = async (id, data) => {
  const formData = buildFormData(data);
  const res = await authInstance.put(`/packages/${id}`, formData);
  return res.data;
};

export const deletePackage = async (id) => {
  const res = await authInstance.delete(`/packages/${id}`);
  return res.data;
};
