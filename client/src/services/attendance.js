// src/services/attendance.js
import { authInstance } from ".";

// =====================
// ATTENDANCE (Auth Required)
// =====================

// POST /api/attendances
export const markAttendance = async (data) => {
  const res = await authInstance.post("/attendances", data);
  return res.data;
};

// GET /api/attendances
export const getAllAttendances = async () => {
  const res = await authInstance.get("/attendances");
  return res.data;
};

// GET /api/attendances/export
export const exportAttendances = async () => {
  const res = await authInstance.get("/attendances/export", {
    responseType: "blob", // download as Excel file
  });
  return res;
};
