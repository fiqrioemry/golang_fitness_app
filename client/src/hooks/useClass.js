import { toast } from "sonner";
import * as classService from "@/services/class";
import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";

export const useClassesQuery = (params = {}) =>
  useQuery({
    queryKey: ["classes", params],
    queryFn: () => classService.getAllClasses(params),
    keepPreviousData: true,
  });

export const useActiveClassesQuery = () =>
  useQuery({
    queryKey: ["classes", "active"],
    queryFn: classService.getActiveClasses,
  });

export const useClassDetailQuery = (id) =>
  useQuery({
    queryKey: ["class", id],
    queryFn: () => classService.getClassById(id),
    enabled: !!id,
  });

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
      mutationFn: ({ id, images }) =>
        classService.uploadClassGallery(id, images),
      ...mutationOpts("Gallery uploaded", ({ id }) => {
        qc.invalidateQueries({ queryKey: ["class", id] });
        qc.invalidateQueries({ queryKey: ["classes"] });
      }),
    }),
  };
};
