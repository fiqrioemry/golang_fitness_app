import { Link } from "react-router-dom";
import { Button } from "@/components/ui/button";
import React, { useState, useMemo } from "react";
import { id as localeId } from "date-fns/locale";
import { Loading } from "@/components/ui/Loading";
import { useAuthStore } from "@/store/useAuthStore";
import {
  useSchedulesQuery,
  useSchedulesWithStatusQuery,
} from "@/hooks/useClass";
import { format, isSameDay, addDays } from "date-fns";
import { ErrorDialog } from "@/components/ui/ErrorDialog";

const Schedules = () => {
  const today = new Date();
  const { user } = useAuthStore();
  const [selectedDate, setSelectedDate] = useState(today);

  const { data, isLoading, isError, refetch } = user?.id
    ? useSchedulesWithStatusQuery()
    : useSchedulesQuery();

  const schedules = data || [];

  const dateRange = useMemo(() => {
    return Array.from({ length: 14 }, (_, i) => addDays(today, i));
  }, [today]);

  const filteredSchedules = useMemo(() => {
    return schedules
      .map((item) => {
        const start = new Date(item.date);
        start.setHours(item.startHour, item.startMinute, 0, 0);
        const end = new Date(start.getTime() + 60 * 60 * 1000);
        return { ...item, startTime: start, endTime: end };
      })
      .filter((item) => isSameDay(item.startTime, selectedDate));
  }, [schedules, selectedDate]);

  if (isLoading) return <Loading />;
  if (isError) return <ErrorDialog onRetry={refetch} />;

  return (
    <section className="min-h-screen px-4 py-10 max-w-7xl mx-auto text-foreground">
      {/* Date Picker */}
      <div className="flex items-center justify-between mb-6">
        <button className="text-xl text-muted-foreground">&#8592;</button>
        <div className="flex gap-2 overflow-x-auto no-scrollbar">
          {dateRange.map((date, i) => {
            const isSelected = isSameDay(date, selectedDate);
            return (
              <button
                key={i}
                onClick={() => setSelectedDate(date)}
                className={`w-14 min-w-[56px] h-16 flex flex-col items-center justify-center rounded-lg transition font-medium ${
                  isSelected
                    ? "bg-primary text-primary-foreground"
                    : "bg-muted hover:bg-accent text-muted-foreground"
                }`}
              >
                <span className="text-xs">
                  {format(date, "EEE", { locale: localeId })}
                </span>
                <span className="text-lg font-bold">
                  {format(date, "dd", { locale: localeId })}
                </span>
              </button>
            );
          })}
        </div>
        <button className="text-xl text-muted-foreground">&#8594;</button>
      </div>

      {/* Summary */}
      <div className="text-sm text-muted-foreground mb-6">
        <strong className="text-foreground">
          {format(selectedDate, "EEEE, dd MMM", { locale: localeId })}
        </strong>{" "}
        • {filteredSchedules.length} classes
      </div>

      {/* Schedule List */}
      <div className="space-y-4">
        {filteredSchedules.length > 0 ? (
          filteredSchedules.map((s) => (
            <div
              key={s.id}
              className="bg-card border border-border rounded-xl px-6 py-8 shadow-sm flex flex-col md:flex-row md:items-center md:justify-between"
            >
              {/* Left section */}
              <div className="flex flex-col space-y-2">
                <p className="text-sm font-semibold">
                  {format(s.startTime, "h:mm a", { locale: localeId })} •{" "}
                  {Math.round((s.endTime - s.startTime) / 60000)} mins
                </p>
                <p className="font-medium text-base">{s.class}</p>
                <p className="text-sm text-muted-foreground">{s.instructor}</p>
              </div>

              {/* Right section */}
              <div className="mt-4 md:mt-0 md:text-right flex flex-col items-end gap-2">
                <p className="text-sm text-muted-foreground">
                  {s.capacity - s.bookedCount > 0
                    ? `${s.capacity - s.bookedCount} left`
                    : "0 in waitlist"}
                </p>

                {s.isBooked ? (
                  <Link to="/profile/bookings">
                    <Button
                      variant="outline"
                      className="text-green-700 border-green-500 bg-green-100 hover:bg-green-200"
                    >
                      Booked
                    </Button>
                  </Link>
                ) : s.capacity - s.bookedCount > 0 ? (
                  <Link to={`/schedules/${s.id}`}>
                    <Button>Book Now</Button>
                  </Link>
                ) : (
                  <Button
                    variant="secondary"
                    className="cursor-not-allowed"
                    disabled
                  >
                    Join Waitlist
                  </Button>
                )}
              </div>
            </div>
          ))
        ) : (
          <p className="text-center text-muted-foreground text-lg mt-20">
            📭 No schedules available for this day.
          </p>
        )}
      </div>
    </section>
  );
};

export default Schedules;
