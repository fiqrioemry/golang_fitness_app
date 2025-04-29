// src/hooks/usePackage.js
import { toast } from "sonner";
import * as packageService from "@/services/package";
import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";

// =====================
// QUERIES
// =====================

// GET /api/packages
export const usePackagesQuery = () =>
  useQuery({
    queryKey: ["packages"],
    queryFn: packageService.getAllPackages,
    keepPreviousData: true,
  });

// GET /api/packages/:id
export const usePackageDetailQuery = (id) =>
  useQuery({
    queryKey: ["package", id],
    queryFn: () => packageService.getPackageById(id),
    enabled: !!id,
  });

// =====================
// MUTATIONS (Admin Only)
// =====================

export const usePackageMutation = () => {
  const qc = useQueryClient();

  const mutationOpts = (msg, refetchFn) => ({
    onSuccess: (res, vars) => {
      toast.success(res?.message || msg);
      if (typeof refetchFn === "function") refetchFn(vars);
      else qc.invalidateQueries({ queryKey: ["packages"] });
    },
    onError: (err) => {
      toast.error(err?.response?.data?.message || "Something went wrong");
    },
  });

  return {
    createPackage: useMutation({
      mutationFn: packageService.createPackage,
      ...mutationOpts("Package created successfully"),
    }),

    updatePackage: useMutation({
      mutationFn: ({ id, formData }) =>
        packageService.updatePackage(id, formData),
      ...mutationOpts("Package updated successfully", ({ id }) => {
        qc.invalidateQueries({ queryKey: ["package", id] });
        qc.invalidateQueries({ queryKey: ["packages"] });
      }),
    }),

    deletePackage: useMutation({
      mutationFn: packageService.deletePackage,
      ...mutationOpts("Package deleted successfully"),
    }),
  };
};
