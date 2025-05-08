// src/pages/customer/UserAttendances.jsx
import React from "react";
import { Loading } from "@/components/ui/Loading";
import { ErrorDialog } from "@/components/ui/ErrorDialog";
import { useAttendancesQuery } from "@/hooks/useAttendance";
import { NoAttendance } from "@/components/customer/attendances/NoAttendance";
import { AttendanceCard } from "@/components/customer/attendances/AttendanceCard";

const UserAttendances = () => {
  const { data, isError, refetch, isLoading } = useAttendancesQuery();

  if (isLoading) return <Loading />;

  if (isError) return <ErrorDialog onRetry={refetch} />;

  const attendances = data || [];

  console.log(attendances);

  return (
    <section className="section p-8 space-y-6">
      <div className="space-y-1 text-center">
        <h2 className="text-2xl font-bold">My Attendance</h2>
        <p className="text-muted-foreground text-sm">
          View and attend your scheduled fitness classes.
        </p>
      </div>

      {attendances.length === 0 ? (
        <NoAttendance />
      ) : (
        <div className="grid gap-6">
          {attendances.map((a) => (
            <AttendanceCard key={a.id} attendance={a} />
          ))}
        </div>
      )}
    </section>
  );
};

export default UserAttendances;
