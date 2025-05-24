import qs from "qs";
import { publicInstance, authInstance } from ".";

// GET /api/payments?q=&page=&limit=&status=&sort=
export const getAllUserPayments = async (params) => {
  const queryString = qs.stringify(params, { skipNulls: true });
  const res = await authInstance.get(`/payments?${queryString}`);
  return res.data;
};

// POST /api/payments
export const createPayment = async (data) => {
  const res = await authInstance.post("/payments", data);
  return res.data;
};

// POST /api/payments/notification (webhook - public)
export const handlePaymentNotification = async (data) => {
  const res = await publicInstance.post("/payments/notification", data);
  return res.data;
};
