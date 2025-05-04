// src/services/booking.js
import { authInstance } from ".";

// =====================
// BOOKING (Auth Required)
// =====================

// POST /api/bookings
export const createBooking = async (data) => {
  const res = await authInstance.post("/bookings", data);
  return res.data;
};

// GET /api/bookings
export const getUserBookings = async () => {
  const res = await authInstance.get("/bookings");
  return res.data.bookings;
};
