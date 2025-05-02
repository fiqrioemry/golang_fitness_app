// src/hooks/useBooking.js
import { toast } from "sonner";
import * as booking from "@/services/booking";
import { useMutation, useQueryClient } from "@tanstack/react-query";

// =====================
// POST /api/bookings
// =====================
export const useCreateBookingMutation = () => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: booking.createBooking,
    onSuccess: () => {
      toast.success("Booking created successfully");
      queryClient.invalidateQueries({ queryKey: ["bookings"] });
    },
    onError: (error) => {
      toast.error(error?.response?.data?.message || "Failed to create booking");
    },
  });
};
