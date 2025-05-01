// src/hooks/useInstructor.js
import { toast } from "sonner";
import * as instructor from "@/services/instructor";
import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";

// =====================
// QUERIES
// =====================

// GET /api/instructors
export const useInstructorsQuery = () =>
  useQuery({
    queryKey: ["instructors"],
    queryFn: instructor.getAllInstructors,
    keepPreviousData: true,
  });

// GET /api/instructors/:id
export const useInstructorDetailQuery = (id) =>
  useQuery({
    queryKey: ["instructor", id],
    queryFn: () => instructor.getInstructorById(id),
    enabled: !!id,
  });

// =====================
// MUTATIONS (Admin Only)
// =====================

export const useInstructorMutation = () => {
  const qc = useQueryClient();

  const mutationOptions = (successMessage, refetchFn) => ({
    onSuccess: (res, vars) => {
      toast.success(res?.message || successMessage);
      if (typeof refetchFn === "function") refetchFn(vars);
      else qc.invalidateQueries({ queryKey: ["instructors"] });
    },
    onError: (err) => {
      toast.error(err?.response?.data?.message || "Something went wrong");
    },
  });

  return {
    createInstructor: useMutation({
      mutationFn: instructor.createInstructor,
      ...mutationOptions("Instructor created successfully"),
    }),

    updateInstructor: useMutation({
      mutationFn: ({ id, data }) => instructor.updateInstructor(id, data),
      ...mutationOptions("Instructor updated successfully", ({ id }) => {
        qc.invalidateQueries({ queryKey: ["instructor", id] });
        qc.invalidateQueries({ queryKey: ["instructors"] });
      }),
    }),

    deleteInstructor: useMutation({
      mutationFn: instructor.deleteInstructor,
      ...mutationOptions("Instructor deleted successfully"),
    }),
  };
};
