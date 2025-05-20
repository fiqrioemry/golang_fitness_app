import { isBefore, parseISO } from "date-fns";
import { ErrorDialog } from "@/components/ui/ErrorDialog";
import { useAttendancesQuery } from "@/hooks/useAttendance";
import { SectionTitle } from "@/components/header/SectionTitle";
import { SectionSkeleton } from "@/components/loading/SectionSkeleton";
import { NoAttendance } from "@/components/customer/attendances/NoAttendance";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/Tabs";
import { AttendanceCard } from "@/components/customer/attendances/AttendanceCard";
import { PastAttendanceCard } from "@/components/customer/attendances/PastAttendanceCard";

const UserAttendances = () => {
  const { data, isError, refetch, isLoading } = useAttendancesQuery();

  if (isLoading) return <SectionSkeleton />;

  if (isError) return <ErrorDialog onRetry={refetch} />;

  const now = new Date();
  const attendances = data || [];

  const upcoming = attendances.filter((a) => {
    const end = new Date(parseISO(a.date));
    end.setHours(a.startHour + 1, a.startMinute);
    return isBefore(now, end);
  });

  const past = attendances.filter((a) => {
    const end = new Date(parseISO(a.date));
    end.setHours(a.startHour + 1, a.startMinute);
    return !isBefore(now, end);
  });

  return (
    <section className="section p-8 space-y-6">
      <SectionTitle
        title="My Attendance"
        description="View and attend your scheduled fitness classes."
      />

      <Tabs defaultValue="upcoming" className="w-full">
        <TabsList className="mb-4">
          <TabsTrigger value="upcoming">Upcoming</TabsTrigger>
          <TabsTrigger value="past">Past</TabsTrigger>
        </TabsList>

        <TabsContent value="upcoming">
          {upcoming.length === 0 ? (
            <NoAttendance type="upcoming" />
          ) : (
            <div className="grid gap-6">
              {upcoming.map((a) => (
                <AttendanceCard key={a.id} attendance={a} />
              ))}
            </div>
          )}
        </TabsContent>

        <TabsContent value="past">
          {past.length === 0 ? (
            <NoAttendance type="past" />
          ) : (
            <div className="grid gap-6">
              {past.map((a) => (
                <PastAttendanceCard key={a.id} attendance={a} />
              ))}
            </div>
          )}
        </TabsContent>
      </Tabs>
    </section>
  );
};

export default UserAttendances;
