// src/hooks/useLevel.js
import { toast } from "sonner";
import * as levelService from "@/services/level";
import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";

// =====================
// QUERIES
// =====================

// GET /api/levels
export const useLevelsQuery = () =>
  useQuery({
    queryKey: ["levels"],
    queryFn: levelService.getAllLevels,
    keepPreviousData: true,
  });

// GET /api/levels/:id
export const useLevelDetailQuery = (id) =>
  useQuery({
    queryKey: ["level", id],
    queryFn: () => levelService.getLevelById(id),
    enabled: !!id,
  });

// =====================
// MUTATIONS (Admin Only)
// =====================

export const useLevelMutation = () => {
  const qc = useQueryClient();

  const mutationOpts = (msg, refetchFn) => ({
    onSuccess: (res, vars) => {
      toast.success(res?.message || msg);
      if (typeof refetchFn === "function") refetchFn(vars);
      else qc.invalidateQueries({ queryKey: ["levels"] });
    },
    onError: (err) => {
      toast.error(err?.response?.data?.message || "Something went wrong");
    },
  });

  return {
    createOptions: useMutation({
      mutationFn: levelService.createLevel,
      ...mutationOpts("Level created successfully"),
    }),

    updateOptions: useMutation({
      mutationFn: ({ id, data }) => levelService.updateLevel(id, data),
      ...mutationOpts("Level updated successfully", ({ id }) => {
        qc.invalidateQueries({ queryKey: ["level", id] });
        qc.invalidateQueries({ queryKey: ["levels"] });
      }),
    }),

    deleteOptions: useMutation({
      mutationFn: levelService.deleteLevel,
      ...mutationOpts("Level deleted successfully"),
    }),
  };
};
