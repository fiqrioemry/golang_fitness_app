// src/hooks/useScheduleTemplate.js
import { toast } from "sonner";
import { useMutation } from "@tanstack/react-query";
import * as scheduleService from "@/services/schedules";

// =====================
// MUTATIONS: ADMIN ONLY
// =====================

// POST /api/schedule-templates
export const useCreateTemplateMutation = () =>
  useMutation({
    mutationFn: scheduleService.createTemplate,
    onSuccess: () => {
      toast.success("Schedule template created successfully");
    },
    onError: (err) => {
      toast.error(err?.response?.data?.message || "Failed to create template");
    },
  });

// POST /api/schedule-templates/auto-generate
export const useAutoGenerateScheduleMutation = () =>
  useMutation({
    mutationFn: scheduleService.autoGenerateSchedules,
    onSuccess: () => {
      toast.success("Schedules generated successfully");
    },
    onError: (err) => {
      toast.error(
        err?.response?.data?.message || "Failed to generate schedules"
      );
    },
  });
