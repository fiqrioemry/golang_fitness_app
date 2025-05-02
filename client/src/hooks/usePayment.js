// src/hooks/usePayment.js
import { toast } from "sonner";
import * as paymentService from "@/services/payment";
import { useMutation, useQuery } from "@tanstack/react-query";

// GET /api/payments?q=&page=&limit= (admin only)
export const useAdminPaymentsQuery = (params) => {
  return useQuery({
    queryKey: ["admin-payments", params],
    queryFn: () => paymentService.getAllUserPayments(params),
    keepPreviousData: true,
    staleTime: 0,
  });
};

// POST /api/payments (auth required)
export const useCreatePaymentMutation = () =>
  useMutation({
    mutationFn: paymentService.createPayment,
    onSuccess: (res) => {
      console.log(res);
      toast.success(res?.message || "Payment created successfully");
    },
    onError: (err) => {
      console.log(err);
      toast.error(err?.response?.data?.message || "Failed to create payment");
    },
  });

// POST /api/payments/notification (webhook - public)
export const useHandlePaymentNotification = () =>
  useMutation({
    mutationFn: paymentService.handlePaymentNotification,
    onSuccess: () => {
      toast.success("Payment notification handled");
    },
    onError: (err) => {
      toast.error(
        err?.response?.data?.message || "Failed to handle notification"
      );
    },
  });
