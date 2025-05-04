import { toast } from "sonner";
import * as scheduleService from "@/services/schedule";
import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";

export const useRecurringTemplatesQuery = () =>
  useQuery({
    queryKey: ["schedule-templates"],
    queryFn: scheduleService.getAllRecuringSchedule,
  });

// hooks/useSchedules.js

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
    createRecurring: useMutation({
      mutationFn: scheduleService.createScheduleTemplates,
      ...mutationOpts("Template created Successfully"),
    }),

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
