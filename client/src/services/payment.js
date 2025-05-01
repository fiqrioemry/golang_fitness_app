// src/services/payment.js
import { publicInstance, authInstance } from ".";

// =====================
// PAYMENT
// =====================

// POST /api/payments (auth required)
export const createPayment = async (data) => {
  // console.log("data", data);
  // const response = {
  //   paymentId: "123",
  //   snapToken: "https://www.ahmadfiqrioemry.com",
  // };
  // return response;
  console.log("purchasement", data);
  const res = await authInstance.post("/payments", data);
  console.log("response payment", res);
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
