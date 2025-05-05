import { useEffect, useState } from "react";
import { Badge } from "@/components/ui/badge";
import { format } from "date-fns";
import { useAttendNow } from "@/hooks/useAttendances";
import { Card, CardContent } from "@/components/ui/card";
import { buildDateTime, getTimeLeft, isAttendanceWindow } from "@/lib/utils";
import { CalendarIcon, ClockIcon, MapPinIcon, UserIcon } from "lucide-react";

export const BookingCard = ({ booking }) => {
  const [timeLeft, setTimeLeft] = useState("");
  const [canAttend, setCanAttend] = useState(false);
  const attendNow = useAttendNow();

  const startTime = buildDateTime(
    booking.date,
    booking.startHour,
    booking.startMinute
  );
  const endTime = startTime
    ? new Date(startTime.getTime() + booking.duration * 60000)
    : null;

  useEffect(() => {
    if (!startTime || !endTime) return;

    const updateStatus = () => {
      setTimeLeft(getTimeLeft(startTime));
      setCanAttend(isAttendanceWindow(startTime, booking.duration));
    };

    updateStatus();
    const timer = setInterval(updateStatus, 1000);
    return () => clearInterval(timer);
  }, [startTime, endTime]);

  const handleAttend = () => {
    console.log("masyukk");
    // attendNow(booking.id);
  };

  if (!startTime || !endTime || isNaN(startTime.getTime())) {
    return (
      <Card className="p-4 shadow-lg">
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
            <h2 className="text-xl font-semibold">{booking.classTitle}</h2>
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

          {canAttend && (
            <button
              onClick={handleAttend}
              className="mt-2 px-4 py-2 bg-green-600 text-white rounded-md hover:bg-green-700"
            >
              Attend Now
            </button>
          )}

          <div className="text-xs text-muted-foreground text-right mt-2">
            Booked at {format(new Date(booking.bookedAt), "dd MMM yyyy, HH:mm")}
          </div>
        </CardContent>
      </div>
    </Card>
  );
};
