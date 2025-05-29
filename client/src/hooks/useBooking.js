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

// GET ALL BOOKED SCHEDULES
export const useMyBookingsQuery = (params) =>
  useQuery({
    queryKey: ["bookings", params],
    queryFn: () => booking.getMyBookings(params),
    keepPreviousData: true,
    staleTime: 1000 * 60 * 60,
  });

// GET BOOKED SCHEDULE DETAIL
export const useBookingDetailQuery = (id) =>
  useQuery({
    queryKey: ["booking", id],
    queryFn: () => booking.getBookingDetail(id),
    enabled: !!id,
  });

// CHECK-IN
export const useCheckinBookingMutation = () => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: booking.checkinBooking,
    onSuccess: () => {
      toast.success("Check-in successful");
      queryClient.invalidateQueries({ queryKey: ["bookings"] });
    },
    onError: (error) => {
      toast.error(error?.response?.data?.message || "Failed to check-in");
    },
  });
};

// CHECK-OUT
export const useCheckoutBookingMutation = () => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: booking.checkoutBooking,
    onSuccess: () => {
      toast.success("Check-out successful");
      queryClient.invalidateQueries({ queryKey: ["bookings"] });
    },
    onError: (error) => {
      toast.error(error?.response?.data?.message || "Failed to check-out");
    },
  });
};

// const { mutate: checkin } = useCheckinBookingMutation();
// checkin(bookingId); // hanya kirim ID

// const { mutate: checkout } = useCheckoutBookingMutation();
// checkout({ id: bookingId, code: "ABCD1234" }); // kirim ID dan kode checkout
