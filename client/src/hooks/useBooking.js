// src/hooks/useBooking.js
import { toast } from "sonner";
import booking from "@/services/booking";
import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";

// =====================
// GET /api/bookings
// =====================
export const useUserBookingsQuery = () =>
  useQuery({
    queryKey: ["bookings"],
    queryFn: booking.getUserBookings,
    keepPreviousData: true,
  });

// =====================
// POST /api/bookings
// =====================
export const useCreateBookingMutation = () => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: booking.createBooking,
    onSuccess: () => {
      toast.success("Booking created successfully");
      queryClient.invalidateQueries({ queryKey: ["bookings"] }); // Refetch daftar booking
    },
    onError: (error) => {
      toast.error(error?.response?.data?.message || "Failed to create booking");
    },
  });
};
