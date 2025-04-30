// src/hooks/useCategory.js
import { toast } from "sonner";
import category from "@/services/categories";
import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";

// =====================
// QUERIES
// =====================

// GET /api/categories
export const useCategoriesQuery = () =>
  useQuery({
    queryKey: ["categories"],
    queryFn: category.getAllCategories,
    keepPreviousData: true,
  });

// GET /api/categories/:id
export const useCategoryDetailQuery = (id) =>
  useQuery({
    queryKey: ["category", id],
    queryFn: () => category.getCategoryById(id),
    enabled: !!id,
  });

// =====================
// MUTATIONS (ADMIN)
// =====================

export const useCategoryMutation = () => {
  const queryClient = useQueryClient();

  const mutationOptions = (successMsg, refetchFn) => ({
    onSuccess: (res, vars) => {
      toast.success(res?.message || successMsg);
      if (typeof refetchFn === "function") {
        refetchFn(vars);
      } else {
        queryClient.invalidateQueries({ queryKey: ["categories"] });
      }
    },
    onError: (err) => {
      toast.error(err?.response?.data?.message || "Something went wrong");
    },
  });

  return {
    createOptions: useMutation({
      mutationFn: category.createCategory,
      ...mutationOptions("Category created successfully"),
    }),

    updateOptions: useMutation({
      mutationFn: ({ id, data }) => category.updateCategory(id, data),
      ...mutationOptions("Category updated successfully", ({ id }) => {
        queryClient.invalidateQueries({ queryKey: ["category", id] });
        queryClient.invalidateQueries({ queryKey: ["categories"] });
      }),
    }),

    deleteOptions: useMutation({
      mutationFn: category.deleteCategory,
      ...mutationOptions("Category deleted successfully"),
    }),
  };
};
