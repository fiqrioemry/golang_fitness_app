// src/pages/customer/UserBookings.jsx
import React from "react";
import { Loading } from "@/components/ui/Loading";
import { useUserBookingsQuery } from "@/hooks/useProfile";
import { ErrorDialog } from "@/components/ui/ErrorDialog";
import { NoBooking } from "@/components/customer/bookings/NoBooking";
import { BookingCard } from "@/components/customer/bookings/BookingCard";

const UserBookings = () => {
  const { data, isError, refetch, isLoading } = useUserBookingsQuery();

  if (isLoading) return <Loading />;

  if (isError) return <ErrorDialog onRetry={refetch} />;

  const bookings = data || [];

  return (
    <section className="section p-8 space-y-6">
      <div className="space-y-1 text-center">
        <h2 className="text-2xl font-bold">My Bookings</h2>
        <p className="text-muted-foreground text-sm">
          View and manage your upcoming and past class bookings.
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
