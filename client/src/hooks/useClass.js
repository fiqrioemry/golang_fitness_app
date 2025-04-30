import { toast } from "sonner";
import * as classService from "@/services/class";
import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";

// ============================
// QUERY HOOKS
// ============================

// GET /api/classes
export const useClassesQuery = (params = {}) =>
  useQuery({
    queryKey: ["classes", params],
    queryFn: () => classService.getAllClasses(params),
    keepPreviousData: true,
  });

// GET /api/classes/active
export const useActiveClassesQuery = () =>
  useQuery({
    queryKey: ["classes", "active"],
    queryFn: classService.getActiveClasses,
  });

// GET /api/classes/:id
export const useClassDetailQuery = (id) =>
  useQuery({
    queryKey: ["class", id],
    queryFn: () => classService.getClassById(id),
    enabled: !!id,
  });

// GET /api/schedules
export const useSchedulesQuery = () =>
  useQuery({
    queryKey: ["schedules"],
    queryFn: classService.getAllClassSchedules,
  });

// ============================
// MUTATION HOOKS
// ============================

export const useClassMutation = () => {
  const qc = useQueryClient();

  const mutationOpts = (msg, refetch) => ({
    onSuccess: (res, vars) => {
      toast.success(res?.message || msg);
      if (typeof refetch === "function") refetch(vars);
      else qc.invalidateQueries({ queryKey: ["classes"] });
    },
    onError: (err) => {
      toast.error(err?.response?.data?.message || "Something went wrong");
    },
  });

  return {
    createClass: useMutation({
      mutationFn: classService.createClass,
      ...mutationOpts("Class created successfully"),
    }),

    updateClass: useMutation({
      mutationFn: ({ id, data }) => classService.updateClass(id, data),

      ...mutationOpts("Class updated", ({ id }) => {
        qc.invalidateQueries({ queryKey: ["class", id] });
        qc.invalidateQueries({ queryKey: ["classes"] });
      }),
    }),

    deleteClass: useMutation({
      mutationFn: classService.deleteClass,
      ...mutationOpts("Class deleted"),
    }),

    uploadGallery: useMutation({
      mutationFn: ({ id, data }) => classService.uploadClassGallery(id, data),
      ...mutationOpts("Gallery uploaded", ({ id }) => {
        qc.invalidateQueries({ queryKey: ["class", id] });
      }),
    }),

    deleteGallery: useMutation({
      mutationFn: ({ id, galleryId }) =>
        classService.deleteClassGallery(id, galleryId),
      ...mutationOpts("Gallery deleted", ({ id }) => {
        qc.invalidateQueries({ queryKey: ["class", id] });
      }),
    }),
  };
};

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
      mutationFn: classService.createClassSchedule,
      ...baseOpts("Schedule created"),
    }),
    updateSchedule: useMutation({
      mutationFn: ({ id, data }) => classService.updateClassSchedule(id, data),
      ...baseOpts("Schedule updated"),
    }),
    deleteSchedule: useMutation({
      mutationFn: classService.deleteClassSchedule,
      ...baseOpts("Schedule deleted"),
    }),
  };
};

export const useScheduleTemplateMutation = () => {
  const qc = useQueryClient();

  const baseOpts = (msg) => ({
    onSuccess: () => toast.success(msg),
    onError: (err) =>
      toast.error(err?.response?.data?.message || "Something went wrong"),
  });

  return {
    createTemplate: useMutation({
      mutationFn: classService.createTemplate,
      ...baseOpts("Template created"),
    }),
    autoGenerateSchedules: useMutation({
      mutationFn: classService.autoGenerateSchedules,
      ...baseOpts("Schedule auto-generated"),
    }),
  };
};
