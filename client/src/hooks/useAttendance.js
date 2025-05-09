import { toast } from "sonner";
import * as attendService from "@/services/attendance";
import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";

export const useAttendancesQuery = () =>
  useQuery({
    queryKey: ["attendances"],
    queryFn: attendService.getAllAttendances,
    staleTime: 1000 * 60 * 5,
  });

export const useAttendanceDetailQuery = (scheduleId) =>
  useQuery({
    queryKey: ["attendance", scheduleId],
    queryFn: () => attendService.getAttendanceDetail(scheduleId),
    enabled: !!scheduleId,
  });

export const useAttendanceMutation = () => {
  const qc = useQueryClient();

  const mutationOpts = (msg) => ({
    onSuccess: () => {
      toast.success(msg);
      qc.invalidateQueries({ queryKey: ["attendances"] });
      qc.invalidateQueries({ queryKey: ["user", "bookings"] });
    },
    onError: (err) => {
      toast.error(err?.response?.data?.message || "Something went wrong");
    },
  });

  return {
    checkin: useMutation({
      mutationFn: attendService.checkinAttendance,
      ...mutationOpts("Check-in successfully"),
    }),
    validateQR: useMutation({
      mutationFn: attendService.validateQRCodeScan,
      ...mutationOpts("QR validated"),
    }),
  };
};

export const useRegenerateQRCode = (bookingId) =>
  useQuery({
    queryKey: ["attendances", bookingId, "qr-code"],
    queryFn: () => attendService.regenerateQRCode(bookingId),
    enabled: !!bookingId,
  });
