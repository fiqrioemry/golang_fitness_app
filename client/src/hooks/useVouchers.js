import { toast } from "sonner";
import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";
import * as voucherService from "@/services/voucher";

export const useVouchersQuery = () =>
  useQuery({
    queryKey: ["vouchers"],
    queryFn: voucherService.getAllVouchers,
    staleTime: 1000 * 60 * 10,
  });

export const useCreateVoucherMutation = () => {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: voucherService.createVoucher,
    onSuccess: () => {
      toast.success("Voucher created successfully");
      qc.invalidateQueries({ queryKey: ["vouchers"] });
    },
    onError: (err) => {
      toast.error(err?.response?.data?.message);
    },
  });
};

export const useApplyVoucherMutation = () =>
  useMutation({
    mutationFn: voucherService.applyVoucher,
    onError: (err) => {
      toast.error(err?.response?.data?.message);
    },
  });

export const useUpdateVoucherMutation = () => {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: voucherService.updateVoucher,
    onSuccess: () => {
      toast.success("Voucher updated successfully");
      qc.invalidateQueries({ queryKey: ["vouchers"] });
    },
    onError: (err) => {
      toast.error(err?.response?.data?.message || "Failed to update voucher");
    },
  });
};

export const useDeleteVoucherMutation = () => {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: voucherService.deleteVoucher,
    onSuccess: () => {
      toast.success("Voucher deleted successfully");
      qc.invalidateQueries({ queryKey: ["vouchers"] });
    },
    onError: (err) => {
      toast.error(err?.response?.data?.message || "Failed to delete voucher");
    },
  });
};
