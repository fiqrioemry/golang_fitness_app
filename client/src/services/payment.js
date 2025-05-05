// src/services/payment.js
import { publicInstance, authInstance } from ".";

// =====================
// PAYMENT
// =====================

// GET /api/payments?q=&page=&limit=
export const getAllUserPayments = async (params) => {
  const res = await authInstance.get("/payments", { params });
  return res.data;
};

// POST /api/payments (auth required)
export const createPayment = async (data) => {
  console.log(data);
  const res = await authInstance.post("/payments", data);
  return res.data;
};

// POST /api/payments/notification (webhook - public)
export const handlePaymentNotification = async (data) => {
  const res = await publicInstance.post("/payments/notification", data);
  return res.data;
};
