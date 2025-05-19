// src/pages/customer/UserBookings.jsx
import { Loading } from "@/components/ui/Loading";
import { useUserBookingsQuery } from "@/hooks/useProfile";
import { ErrorDialog } from "@/components/ui/ErrorDialog";
import { SectionTitle } from "@/components/header/SectionTitle";
import { NoBooking } from "@/components/customer/bookings/NoBooking";
import { BookingCard } from "@/components/customer/bookings/BookingCard";

const UserBookings = () => {
  const { data, isError, refetch, isLoading } = useUserBookingsQuery();

  if (isLoading) return <Loading />;

  if (isError) return <ErrorDialog onRetry={refetch} />;

  const bookings = data || [];

  return (
    <section className="section p-8 space-y-6">
      <SectionTitle
        title="My Bookings"
        description="View and manage your upcoming and past class bookings."
      />
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
