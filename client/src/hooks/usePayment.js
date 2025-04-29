// src/hooks/usePayment.js
import { toast } from "sonner";
import { useMutation } from "@tanstack/react-query";
import * as paymentService from "@/services/payment";

// =====================
// MUTATION HOOKS
// =====================

// POST /api/payments (auth required)
export const useCreatePaymentMutation = () =>
  useMutation({
    mutationFn: paymentService.createPayment,
    onSuccess: (res) => {
      toast.success(res?.message || "Payment created successfully");
    },
    onError: (err) => {
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
