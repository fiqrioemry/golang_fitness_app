import { toast } from "sonner";
import * as paymentService from "@/services/payment";
import { useMutation, useQuery } from "@tanstack/react-query";

export const useAdminPaymentsQuery = (params) =>
  useQuery({
    queryKey: ["payments", params],
    queryFn: () => paymentService.getAllUserPayments(params),
    keepPreviousData: true,
    staleTime: 1000 * 60 * 2,
  });

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
