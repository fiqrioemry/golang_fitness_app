// src/pages/customer/UserBookings.jsx
import React from "react";
import { Loading } from "@/components/ui/Loading";
import { ErrorDialog } from "@/components/ui/ErrorDialog";
import { useUserBookingsQuery } from "@/hooks/useProfile";
import { BookingCard } from "@/components/customer/bookings/BookingCard";
import { NoBooking } from "@/components/customer/bookings/NoBooking";

const UserBookings = () => {
  const {
    data: bookings = [],
    isError,
    refetch,
    isLoading,
  } = useUserBookingsQuery();

  if (isLoading) return <Loading />;

  if (isError) return <ErrorDialog onRetry={refetch} />;

  return (
    <section className="max-w-6xl mx-auto px-4 py-8 space-y-6">
      <div className="space-y-1 text-center">
        <h2 className="text-2xl font-bold">My Bookings</h2>
        <p className="text-muted-foreground text-sm">
          View and manage your upcoming and past class bookings. Make sure to
          attend on time and track your fitness journey with ease.
        </p>
      </div>
      {bookings.length === 0 ? (
        <NoBooking />
      ) : (
        bookings.map((booking) => (
          <BookingCard key={booking.id} booking={booking} />
        ))
      )}
    </section>
  );
};

export default UserBookings;
