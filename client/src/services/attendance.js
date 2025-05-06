import { authInstance } from ".";

// ✅ GET /api/attendances
export const getAllAttendances = async () => {
  const res = await authInstance.get("/attendances");
  return res.data;
};

// ✅ GET /api/attendances/:bookingId/qr
export const getQRCode = async (bookingId) => {
  const res = await authInstance.get(`/attendances/${bookingId}/qr`);
  return res.data.qr;
};

// ✅ POST /api/attendances/:bookingId/checkin
export const checkinAttendance = async (bookingId) => {
  const res = await authInstance.post(`/attendances/${bookingId}/checkin`);
  return res.data.qr;
};

// GET /api/attendances/export
export const exportAttendances = async () => {
  const res = await authInstance.get("/attendances/export", {
    responseType: "blob", // download as Excel file
  });
  return res;
};

export const validateQRCode = async (qrCode) => {
  const res = await authInstance.post("/attendances/validate", { qrCode });
  return res.data;
};
