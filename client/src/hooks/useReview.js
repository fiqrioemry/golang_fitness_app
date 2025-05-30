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
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: reviewService.createReview,
    onSuccess: (_, variables) => {
      toast.success("Review Submitted Successfully");
      queryClient.invalidateQueries(["booking", variables.id]);
    },
    onError: (err) => {
      toast.error(err?.response?.data?.message || "Failed to submit review");
    },
  });
};
