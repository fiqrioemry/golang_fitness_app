import { toast } from "sonner";
import * as booking from "@/services/booking";
import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";

export const useCreateBookingMutation = () => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: booking.createBooking,
    onSuccess: () => {
      toast.success("Booking created successfully");
      queryClient.invalidateQueries({ queryKey: ["bookings"] });
      queryClient.invalidateQueries({ queryKey: ["schedules", "with-status"] });
    },
    onError: (error) => {
      toast.error(error?.response?.data?.message || "Failed to create booking");
    },
  });
};

export const useCheckinBookedClassMutation = () => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: booking.checkinBookedClass,
    onSuccess: () => {
      toast.success("Check-in successful");
      queryClient.invalidateQueries({ queryKey: ["bookings"] });
    },
    onError: (error) => {
      toast.error(error?.response?.data?.message || "Failed to check-in");
    },
  });
};

export const useRegenerateQRMutation = () => {
  return useMutation({
    mutationFn: booking.regenerateQRCode,
    onSuccess: () => {
      toast.success("QR code regenerated");
    },
    onError: (error) => {
      toast.error(
        error?.response?.data?.message || "Failed to regenerate QR code"
      );
    },
  });
};

export const useMyBookingsQuery = (params) =>
  useQuery({
    queryKey: ["bookings", params],
    queryFn: () => booking.getMyBookings(params),
    keepPreviousData: true,
    staleTime: 1000 * 60 * 60, // stale 1 jam karena jadwal booking tdk bisa saling bertabrakan
  });
