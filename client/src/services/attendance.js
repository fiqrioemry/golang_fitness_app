import { authInstance } from ".";

export const getAllAttendances = async () => {
  const res = await authInstance.get("/attendances");
  return res.data;
};

export const getAttendanceDetail = async (id) => {
  const res = await authInstance.get(`/attendances/${id}`);
  return res.data;
};

export const checkinAttendance = async (bookingId) => {
  const res = await authInstance.post(`/attendances/${bookingId}`);
  return res.data.qr;
};

export const regenerateQRCode = async (bookingId) => {
  const res = await authInstance.get(`/attendances/${bookingId}/qr-code`);
  return res.data.qr;
};

export const validateQRCodeScan = async ({ qr }) => {
  const res = await authInstance.post("/attendances/validate", { qr });
  return res.data;
};
