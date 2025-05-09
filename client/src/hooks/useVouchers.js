import { toast } from "sonner";
import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";
import * as voucherService from "@/services/voucher";

// =====================
// QUERIES
// =====================

export const useVouchersQuery = () =>
  useQuery({
    queryKey: ["vouchers"],
    queryFn: voucherService.getAllVouchers,
    staleTime: 1000 * 60 * 10,
  });

// =====================
// MUTATIONS
// =====================

export const useVoucherMutation = () => {
  const qc = useQueryClient();

  const mutationOpts = (msg, refetchFn) => ({
    onSuccess: (res, vars) => {
      toast.success(res?.message || msg);
      if (typeof refetchFn === "function") refetchFn(vars);
      else qc.invalidateQueries({ queryKey: ["vouchers"] });
    },
    onError: (err) => {
      toast.error(err?.response?.data?.message || "Something went wrong");
    },
  });

  return {
    createVoucher: useMutation({
      mutationFn: voucherService.createVoucher,
      ...mutationOpts("Voucher created successfully"),
    }),

    updateVoucher: useMutation({
      mutationFn: ({ id, data }) => voucherService.updateVoucher({ id, data }),
      ...mutationOpts("Voucher updated successfully", () => {
        qc.invalidateQueries({ queryKey: ["vouchers"] });
      }),
    }),

    deleteVoucher: useMutation({
      mutationFn: voucherService.deleteVoucher,
      ...mutationOpts("Voucher deleted successfully"),
    }),

    applyVoucher: useMutation({
      mutationFn: voucherService.applyVoucher,
      onError: (err) => {
        toast.error(err?.response?.data?.message || "Failed to apply voucher");
      },
    }),
  };
};
