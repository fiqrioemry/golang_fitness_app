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

export const BookingCard = ({ booking }) => {
  const [timeLeft, setTimeLeft] = useState("");

  useEffect(() => {
    const updateCountdown = () => {
      const seconds = differenceInSeconds(
        new Date(booking.startTime),
        new Date()
      );
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
  }, [booking.startTime]);

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
            {booking.instructorName} • {booking.duration} mins
          </p>

          {/* Countdown */}
          <div className="text-sm text-blue-500 font-medium">
            Starts in: {timeLeft} ⏳🔥
          </div>

          <div className="flex items-center gap-2 text-sm text-muted-foreground">
            <CalendarIcon className="h-4 w-4" />
            <span>
              {format(new Date(booking.startTime), "EEE, dd MMM yyyy")}
            </span>
          </div>
          <div className="flex items-center gap-2 text-sm text-muted-foreground">
            <ClockIcon className="h-4 w-4" />
            <span>
              {format(new Date(booking.startTime), "HH:mm")} -{" "}
              {format(new Date(booking.endTime), "HH:mm")}
            </span>
          </div>
          <div className="flex items-center gap-2 text-sm text-muted-foreground">
            <MapPinIcon className="h-4 w-4" />
            <span>
              {booking.locationName} - {booking.locationAddress}
            </span>
          </div>
          <div className="flex items-center gap-2 text-sm text-muted-foreground">
            <UserIcon className="h-4 w-4" />
            <span>{booking.participantCount} Participant</span>
          </div>
          <div className="text-xs text-muted-foreground text-right mt-2">
            Booked at {format(new Date(booking.bookedAt), "dd MMM yyyy, HH:mm")}
          </div>
        </CardContent>
      </div>
    </Card>
  );
};
