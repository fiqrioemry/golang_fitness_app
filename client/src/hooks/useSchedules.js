import { toast } from "sonner";
import * as scheduleService from "@/services/schedule";
import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";

// GET /api/schedules
export const useSchedulesQuery = () =>
  useQuery({
    queryKey: ["schedules"],
    queryFn: scheduleService.getAllClassSchedules,
  });

// GET /api/schedules/status
export const useSchedulesWithStatusQuery = () => {
  return useQuery({
    queryKey: ["schedules", "with-status"],
    queryFn: scheduleService.getAllClassSchedulesWithStatus,
    refetchOnMount: true,
  });
};

export const useScheduleDetailQuery = (id) =>
  useQuery({
    queryKey: ["schedule", id],
    queryFn: () => scheduleService.getClassScheduleDetail(id),
    enabled: !!id,
  });

export const useScheduleMutation = () => {
  const qc = useQueryClient();

  const baseOpts = (msg) => ({
    onSuccess: () => {
      toast.success(msg);
      qc.invalidateQueries({ queryKey: ["schedules"] });
    },
    onError: (err) => {
      toast.error(err?.response?.data?.message || "Something went wrong");
    },
  });

  return {
    createSchedule: useMutation({
      mutationFn: scheduleService.createClassSchedule,
      ...baseOpts("Schedule created"),
    }),
    updateSchedule: useMutation({
      mutationFn: ({ id, data }) =>
        scheduleService.updateClassSchedule(id, data),
      ...baseOpts("Schedule updated"),
    }),
    deleteSchedule: useMutation({
      mutationFn: scheduleService.deleteClassSchedule,
      ...baseOpts("Schedule deleted"),
    }),
  };
};

export const useRecurringTemplatesQuery = () =>
  useQuery({
    queryKey: ["schedule-templates"],
    queryFn: scheduleService.getAllRecuringSchedule,
  });

export const useScheduleTemplateMutation = () => {
  const qc = useQueryClient();

  const mutationOpts = (successMsg, refetch) => ({
    onSuccess: (res, vars) => {
      toast.success(res?.message || successMsg);
      if (typeof refetch === "function") refetch(vars);
      else qc.invalidateQueries({ queryKey: ["schedule-templates"] });
    },
    onError: (err) => {
      toast.error(err?.response?.data?.message || "Something went wrong");
    },
  });

  return {
    updateTemplate: useMutation({
      mutationFn: ({ id, data }) =>
        scheduleService.updateScheduleTemplate(id, data),
      ...mutationOpts("Template updated successfully"),
    }),
    deleteTemplate: useMutation({
      mutationFn: scheduleService.deleteScheduleTemplate,
      ...mutationOpts("Template deleted Successfully"),
    }),

    runTemplate: useMutation({
      mutationFn: scheduleService.runScheduleTemplate,
      ...mutationOpts("Template activated successfully"),
    }),

    stopTemplate: useMutation({
      mutationFn: scheduleService.stopScheduleTemplate,
      ...mutationOpts("Template deactivated successfully"),
    }),
  };
};

// GET /api/schedules/instructor
export const useInstructorSchedulesQuery = (params) =>
  useQuery({
    queryKey: ["instructor-schedules", params],
    queryFn: () => scheduleService.getInstructorSchedules(params),
    staleTime: 1000 * 60 * 45,
  });

// GET /api/schedules/:id/attendance
export const useScheduleAttendanceQuery = (id) =>
  useQuery({
    queryKey: ["schedule", id, "attendance"],
    queryFn: () => scheduleService.getClassAttendances(id),
    enabled: !!id,
  });

// PATCH /api/schedules/:id/open
export const useOpenScheduleMutation = () => {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: scheduleService.openClassSchedule,
    onSuccess: () => {
      toast.success("Schedule opened successfully");
      qc.invalidateQueries({ queryKey: ["instructor-schedules"] });
    },
    onError: (err) => {
      toast.error(err?.response?.data?.message || "Failed to open schedule");
    },
  });
};
