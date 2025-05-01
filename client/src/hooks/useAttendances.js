// src/hooks/useAttendances.js
import { toast } from "sonner";
import * as attendance from "@/services/attendance";
import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";

export const useAttendancesQuery = () =>
  useQuery({
    queryKey: ["attendances"],
    queryFn: attendance.getAllAttendances,
    keepPreviousData: true,
  });

export const useMarkAttendanceMutation = () => {
  const queryClient = useQueryClient();
  return useMutation({
    mutationFn: attendance.markAttendance,
    onSuccess: () => {
      toast.success("Attendance marked successfully");
      queryClient.invalidateQueries({ queryKey: ["attendances"] });
    },
    onError: (error) => {
      toast.error(
        error?.response?.data?.message || "Failed to mark attendance"
      );
    },
  });
};

export const useExportAttendances = () =>
  useQuery({
    queryKey: ["attendances", "export"],
    queryFn: attendance.exportAttendances,
    enabled: false, // Only fetch manually
  });
