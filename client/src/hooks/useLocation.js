// src/hooks/useLocation.js
import { toast } from "sonner";
import * as locationService from "@/services/location";
import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";

// =====================
// QUERIES
// =====================

// GET /api/locations
export const useLocationsQuery = () =>
  useQuery({
    queryKey: ["locations"],
    queryFn: locationService.getAllLocations,
    keepPreviousData: true,
  });

// GET /api/locations/:id
export const useLocationDetailQuery = (id) =>
  useQuery({
    queryKey: ["location", id],
    queryFn: () => locationService.getLocationById(id),
    enabled: !!id,
  });

// =====================
// MUTATIONS (Admin Only)
// =====================

export const useLocationMutation = () => {
  const qc = useQueryClient();

  const mutationOpts = (msg, refetchFn) => ({
    onSuccess: (res, vars) => {
      toast.success(res?.message || msg);
      if (typeof refetchFn === "function") refetchFn(vars);
      else qc.invalidateQueries({ queryKey: ["locations"] });
    },
    onError: (err) => {
      toast.error(err?.response?.data?.message || "Something went wrong");
    },
  });

  return {
    createOptions: useMutation({
      mutationFn: locationService.createLocation,
      ...mutationOpts("Location created successfully"),
    }),

    updateOptions: useMutation({
      mutationFn: ({ id, data }) => locationService.updateLocation(id, data),
      ...mutationOpts("Location updated", ({ id }) => {
        qc.invalidateQueries({ queryKey: ["location", id] });
        qc.invalidateQueries({ queryKey: ["locations"] });
      }),
    }),

    deleteOptions: useMutation({
      mutationFn: locationService.deleteLocation,
      ...mutationOpts("Location deleted"),
    }),
  };
};
