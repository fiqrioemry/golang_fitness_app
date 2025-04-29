// src/hooks/useReview.js
import { toast } from "sonner";
import * as reviewService from "@/services/review";
import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";

// =====================
// QUERY: GET REVIEWS BY CLASS ID
// =====================

export const useClassReviewsQuery = (classId) =>
  useQuery({
    queryKey: ["reviews", classId],
    queryFn: () => reviewService.getReviewsByClass(classId),
    enabled: !!classId,
  });

// =====================
// MUTATION: CREATE REVIEW (Auth Required)
// =====================

export const useCreateReviewMutation = () => {
  const qc = useQueryClient();

  return useMutation({
    mutationFn: reviewService.createReview,
    onSuccess: (_, { classId }) => {
      toast.success("Review submitted successfully");
      qc.invalidateQueries({ queryKey: ["reviews", classId] });
    },
    onError: (err) => {
      toast.error(err?.response?.data?.message || "Failed to submit review");
    },
  });
};
