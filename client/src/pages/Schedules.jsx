import { Link } from "react-router-dom";
import { Button } from "@/components/ui/button";
import React, { useState, useMemo } from "react";
import { id as localeId } from "date-fns/locale";
import { Loading } from "@/components/ui/Loading";
import { useAuthStore } from "@/store/useAuthStore";
import { useSchedulesQuery } from "@/hooks/useClass";
import { format, isSameDay, addDays } from "date-fns";
import { ErrorDialog } from "@/components/ui/ErrorDialog";
import { useSchedulesWithStatusQuery } from "@/hooks/useClass";

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
      ?.map((item) => {
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
    <section className="min-h-screen px-4 py-10 max-w-7xl mx-auto">
      {/* Date Header */}
      <div className="flex items-center justify-between mb-4">
        <button className="text-xl">&#8592;</button>
        <div className="flex gap-2 overflow-x-auto no-scrollbar">
          {dateRange.map((date, i) => {
            const isSelected = isSameDay(date, selectedDate);
            return (
              <div
                key={i}
                onClick={() => setSelectedDate(date)}
                className={`w-14 min-w-[56px] h-16 flex flex-col items-center justify-center rounded-lg cursor-pointer ${
                  isSelected
                    ? "bg-primary text-white font-semibold"
                    : "bg-white text-gray-800 hover:bg-gray-100"
                }`}
              >
                <span className="text-sm">
                  {format(date, "EEE", { locale: localeId })}
                </span>
                <span className="text-lg font-bold">
                  {format(date, "dd", { locale: localeId })}
                </span>
              </div>
            );
          })}
        </div>
        <button className="text-xl">&#8594;</button>
      </div>

      {/* Summary */}
      <div className="text-sm text-gray-500 mb-4">
        <strong className="text-gray-800">
          {format(selectedDate, "EEEE, dd MMM", { locale: localeId })}
        </strong>{" "}
        • {filteredSchedules.length} classes
      </div>

      {/* Class Schedule List */}
      <div className="space-y-4">
        {filteredSchedules.length > 0 ? (
          filteredSchedules.map((s) => (
            <div
              key={s.id}
              className="bg-secondary rounded-xl p-4 shadow-sm flex flex-col sm:flex-row sm:items-center sm:justify-between"
            >
              <div className="space-y-1">
                <p className="text-sm font-semibold text-[#56684c]">
                  {format(s.startTime, "h:mm a", { locale: localeId })} •{" "}
                  {Math.round((s.endTime - s.startTime) / 60000)} mins
                </p>
                <p className="font-medium">{s.class}</p>
                <p className="text-sm text-gray-500">{s.instructor}</p>
              </div>
              <div className="mt-3 sm:mt-0 sm:text-right">
                <p className="text-sm text-gray-500">
                  {s.capacity - s.bookedCount > 0
                    ? `${s.capacity - s.bookedCount} left`
                    : "0 in waitlist"}
                </p>
                {s.isBooked ? (
                  <Link to="/profile/bookings">
                    <Button
                      variant="outline"
                      className="bg-green-100 text-green-700 border-green-500 hover:bg-green-200"
                    >
                      Booked
                    </Button>
                  </Link>
                ) : s.capacity - s.bookedCount > 0 ? (
                  <Link to={`/schedules/${s.id}`}>
                    <Button>Book Now</Button>
                  </Link>
                ) : (
                  <button className="mt-2 px-4 py-2 bg-gray-200 text-gray-700 rounded-full text-sm font-medium">
                    Join Waitlist
                  </button>
                )}
              </div>
            </div>
          ))
        ) : (
          <p className="text-center text-gray-500 text-lg mt-20">
            📭 No schedules on this day.
          </p>
        )}
      </div>
    </section>
  );
};

export default Schedules;
