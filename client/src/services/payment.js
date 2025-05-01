// src/services/payment.js
import { publicInstance, authInstance } from ".";

// =====================
// PAYMENT
// =====================

// POST /api/payments (auth required)
export const createPayment = async (data) => {
  const res = await authInstance.post("/payments", data);
  return res.data;
};

// POST /api/payments/notification (webhook - public)
export const handlePaymentNotification = async (data) => {
  const res = await publicInstance.post("/payments/notification", data);
  return res.data;
};

export default {
  createPayment,
  handlePaymentNotification,
};
