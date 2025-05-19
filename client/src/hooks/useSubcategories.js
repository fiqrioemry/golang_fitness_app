import { toast } from "sonner";
import * as subcategoryService from "@/services/subcategories";
import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";

// GET /api/subcategories
export const useSubcategoriesQuery = () =>
  useQuery({
    queryKey: ["subcategories"],
    queryFn: subcategoryService.getAllSubcategories,
    keepPreviousData: true,
  });

// GET /api/subcategories/:id
export const useSubcategoryDetailQuery = (id) =>
  useQuery({
    queryKey: ["subcategory", id],
    queryFn: () => subcategoryService.getSubcategoryById(id),
    enabled: !!id,
  });

// GET /api/subcategories/category/:categoryId
export const useSubcategoriesByCategoryQuery = (categoryId) =>
  useQuery({
    queryKey: ["subcategories", "category", categoryId],
    queryFn: () => subcategoryService.getSubcategoriesByCategory(categoryId),
    enabled: !!categoryId,
  });

export const useSubcategoryMutation = () => {
  const qc = useQueryClient();

  const mutationOpts = (msg, refetchFn) => ({
    onSuccess: (res, vars) => {
      toast.success(res?.message || msg);
      if (typeof refetchFn === "function") refetchFn(vars);
      else qc.invalidateQueries({ queryKey: ["subcategories"] });
    },
    onError: (err) => {
      toast.error(err?.response?.data?.message || "Something went wrong");
    },
  });

  return {
    createOptions: useMutation({
      mutationFn: subcategoryService.createSubcategory,
      ...mutationOpts("Subcategory created successfully"),
    }),

    updateOptions: useMutation({
      mutationFn: ({ id, data }) =>
        subcategoryService.updateSubcategory(id, data),
      ...mutationOpts("Subcategory updated successfully", ({ id }) => {
        qc.invalidateQueries({ queryKey: ["subcategory", id] });
        qc.invalidateQueries({ queryKey: ["subcategories"] });
      }),
    }),

    deleteOptions: useMutation({
      mutationFn: subcategoryService.deleteSubcategory,
      ...mutationOpts("Subcategory deleted successfully"),
    }),
  };
};
