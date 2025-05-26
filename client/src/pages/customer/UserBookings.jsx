import { isBefore, parseISO } from "date-fns";
import { useQueryStore } from "@/store/useQueryStore";
import { useMyBookingsQuery } from "@/hooks/useBooking";
import { ErrorDialog } from "@/components/ui/ErrorDialog";
import { SectionTitle } from "@/components/header/SectionTitle";
import { SectionSkeleton } from "@/components/loading/SectionSkeleton";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/Tabs";
import { NoClassSchedule } from "@/components/customer/bookings/NoClassSchedule";
import { PastClassSchedules } from "@/components/customer/bookings/PastClassSchedules";
import { UpcomingClassSchedules } from "@/components/customer/bookings/UpcomingClassSchedules";

const UserBookings = () => {
  const { page, limit, sort, status } = useQueryStore();

  const { data, isError, refetch, isLoading } = useMyBookingsQuery({
    status,
    sort,
    page,
    limit,
  });

  if (isLoading) return <SectionSkeleton />;

  if (isError) return <ErrorDialog onRetry={refetch} />;

  const now = new Date();

  const bookings = data?.data || [];

  const pagination = data?.pagination || [];

  const upcoming = bookings.filter((a) => {
    const end = new Date(parseISO(a.date));
    end.setHours(a.startHour + 1, a.startMinute);
    return isBefore(now, end);
  });

  const past = bookings.filter((a) => {
    const end = new Date(parseISO(a.date));
    end.setHours(a.startHour + 1, a.startMinute);
    return !isBefore(now, end);
  });

  return (
    <section className="section p-8 space-y-6">
      <SectionTitle
        title="My Bookings"
        description="View and manage your upcoming and past class bookings."
      />

      <Tabs defaultValue="upcoming" className="w-full">
        <TabsList className="mb-4">
          <TabsTrigger value="upcoming">Upcoming</TabsTrigger>
          <TabsTrigger value="past">Past</TabsTrigger>
        </TabsList>

        <TabsContent value="upcoming">
          {upcoming.length === 0 ? (
            <NoClassSchedule type="upcoming" />
          ) : (
            <div className="grid gap-6">
              {upcoming.map((schedule) => (
                <UpcomingClassSchedules key={schedule.id} schedule={schedule} />
              ))}
            </div>
          )}
        </TabsContent>

        <TabsContent value="past">
          {past.length === 0 ? (
            <NoClassSchedule type="past" />
          ) : (
            <div className="grid gap-6">
              {past.map((schedule) => (
                <PastClassSchedules key={schedule.id} schedule={schedule} />
              ))}
            </div>
          )}
        </TabsContent>
      </Tabs>
    </section>
  );
};
export default UserBookings;
