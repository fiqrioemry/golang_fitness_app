import {
  format,
  formatDuration,
  intervalToDuration,
  differenceInSeconds,
} from "date-fns";
import { useEffect, useState } from "react";
import { Badge } from "@/components/ui/badge";
import { Card, CardContent } from "@/components/ui/card";
import { CalendarIcon, ClockIcon, MapPinIcon, UserIcon } from "lucide-react";

const buildDateTime = (dateStr, hour, minute) => {
  if (!dateStr || hour === undefined || minute === undefined) return null;
  const date = new Date(dateStr);
  date.setHours(hour);
  date.setMinutes(minute);
  date.setSeconds(0);
  return date;
};

export const BookingCard = ({ booking }) => {
  const [timeLeft, setTimeLeft] = useState("");

  const startTime = buildDateTime(
    booking.date,
    booking.startHour,
    booking.startMinute
  );
  const endTime = startTime
    ? new Date(startTime.getTime() + booking.duration * 60 * 1000)
    : null;

  useEffect(() => {
    if (!startTime) return;

    const updateCountdown = () => {
      const seconds = differenceInSeconds(startTime, new Date());
      if (seconds > 0) {
        const duration = intervalToDuration({ start: 0, end: seconds * 1000 });
        const formatted = formatDuration(duration, {
          format: ["days", "hours", "minutes", "seconds"],
        });
        setTimeLeft(formatted);
      } else {
        setTimeLeft("Ongoing or passed");
      }
    };

    updateCountdown();
    const timer = setInterval(updateCountdown, 1000);
    return () => clearInterval(timer);
  }, [startTime]);

  if (!startTime || isNaN(startTime.getTime()) || !endTime) {
    return (
      <Card className="overflow-hidden shadow-lg transition hover:shadow-xl p-4">
        <p className="text-red-500">Invalid booking time.</p>
      </Card>
    );
  }

  return (
    <Card className="overflow-hidden shadow-lg transition hover:shadow-xl">
      <div className="grid grid-cols-1 sm:grid-cols-[150px_1fr]">
        <img
          src={booking.classImage}
          alt={booking.classTitle}
          className="h-full w-full object-cover sm:h-full sm:w-full"
        />
        <CardContent className="p-4 space-y-2">
          <div className="flex justify-between items-start">
            <h2 className="text-xl font-semibold leading-tight">
              {booking.classTitle}
            </h2>
            <Badge variant="outline" className="capitalize">
              {booking.status}
            </Badge>
          </div>
          <p className="text-sm text-muted-foreground">
            {booking.instructor} • {booking.duration} mins
          </p>

          <div className="text-sm text-blue-500 font-medium">
            Starts in: {timeLeft} ⏳🔥
          </div>

          <div className="flex items-center gap-2 text-sm text-muted-foreground">
            <CalendarIcon className="h-4 w-4" />
            <span>{format(startTime, "EEE, dd MMM yyyy")}</span>
          </div>
          <div className="flex items-center gap-2 text-sm text-muted-foreground">
            <ClockIcon className="h-4 w-4" />
            <span>
              {format(startTime, "HH:mm")} - {format(endTime, "HH:mm")}
            </span>
          </div>
          <div className="flex items-center gap-2 text-sm text-muted-foreground">
            <MapPinIcon className="h-4 w-4" />
            <span>{booking.location}</span>
          </div>
          <div className="flex items-center gap-2 text-sm text-muted-foreground">
            <UserIcon className="h-4 w-4" />
            <span>{booking.participant} Participant</span>
          </div>
          <div className="text-xs text-muted-foreground text-right mt-2">
            Booked at {format(new Date(booking.bookedAt), "dd MMM yyyy, HH:mm")}
          </div>
        </CardContent>
      </div>
    </Card>
  );
};
