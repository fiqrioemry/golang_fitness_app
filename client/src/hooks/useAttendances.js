import { toast } from "sonner";
import * as attendance from "@/services/attendance";
import { useQuery, useMutation } from "@tanstack/react-query";

export const useAttendancesQuery = () =>
  useQuery({
    queryKey: ["attendances"],
    queryFn: attendance.getAllAttendances,
    keepPreviousData: true,
  });

export const useValidateQRCode = () => {
  return useMutation({
    mutationFn: attendance.validateQRCode,
  });
};

export const useCheckinAttendance = () =>
  useMutation({
    mutationFn: attendance.checkinAttendance,
    onSuccess: () => {
      toast.success("Check-in berhasil, QR Code siap ditampilkan.");
    },
    onError: (err) => {
      toast.error(err?.response?.data?.error || "Gagal check-in kelas.");
    },
  });

export const useQRCodeQuery = (bookingId) =>
  useQuery({
    queryKey: ["attendances", "qr", bookingId],
    queryFn: () => attendance.getQRCode(bookingId),
    enabled: !!bookingId,
  });

export const useExportAttendances = () =>
  useQuery({
    queryKey: ["attendances", "export"],
    queryFn: attendance.exportAttendances,
    enabled: false,
  });
