// src/hooks/useReview.js
import { toast } from "sonner";
import * as reviewService from "@/services/review";
import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";

export const useClassReviewsQuery = (classId) =>
  useQuery({
    queryKey: ["reviews", classId],
    queryFn: () => reviewService.getReviewsByClass(classId),
    enabled: !!classId,
  });

export const useCreateReviewMutation = () => {
  const qc = useQueryClient();

  return useMutation({
    mutationFn: reviewService.createReview,
    onSuccess: (_, { classId }) => {
      toast.success("Review submitted successfully");
      qc.invalidateQueries({ queryKey: ["reviews", classId] });
      qc.invalidateQueries({ queryKey: ["attendances"] });
    },
    onError: (err) => {
      toast.error(err?.response?.data?.message || "Failed to submit review");
    },
  });
};
