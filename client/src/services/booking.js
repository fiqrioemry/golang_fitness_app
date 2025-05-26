import qs from "qs";
import { authInstance } from ".";

// POST /api/bookings
export const createBooking = async (data) => {
  const res = await authInstance.post("/bookings", data);
  return res.data;
};

// GET /api/bookings
export const getMyBookings = async (params) => {
  const queryString = qs.stringify(params, { skipNulls: true });
  const res = await authInstance.get(`/bookings?${queryString}`);
  return res.data;
};

export const checkinBookedClass = async (bookingId) => {
  console.log("CHECKIN CLASS", bookingId);
  const res = await authInstance.post(`/bookings/${bookingId}`);
  return res.data.qr;
};

export const regenerateQRCode = async (bookingId) => {
  const res = await authInstance.get(`/bookings/${bookingId}/qr-code`);
  return res.data.qr;
};
