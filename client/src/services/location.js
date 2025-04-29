import { publicInstance, authInstance } from ".";

export const getAllLocations = async () => {
  const res = await publicInstance.get("/locations");
  return res.data;
};

export const getLocationById = async (id) => {
  const res = await publicInstance.get(`/locations/${id}`);
  return res.data;
};

export const createLocation = async (data) => {
  const res = await authInstance.post("/locations", data);
  return res.data;
};

export const updateLocation = async (id, data) => {
  const res = await authInstance.put(`/locations/${id}`, data);
  return res.data;
};

export const deleteLocation = async (id) => {
  const res = await authInstance.delete(`/locations/${id}`);
  return res.data;
};
