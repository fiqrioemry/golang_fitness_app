import { format } from "date-fns";
import { Badge } from "@/components/ui/Badge";
import { Card, CardContent } from "@/components/ui/Card";
import { formatHour, buildDateTime } from "@/lib/utils";
import { UserIcon, ClockIcon, MapPinIcon, CalendarIcon } from "lucide-react";

export const PastClassSchedules = ({ schedule }) => {
  const startTime = buildDateTime(
    schedule.date,
    schedule.startHour,
    schedule.startMinute
  );
  const endTime = startTime
    ? new Date(startTime.getTime() + Number(schedule.duration || 0) * 60000)
    : null;

  return (
    <Card className="overflow-hidden border border-border bg-card shadow-md hover:shadow-xl transition">
      <div className="grid grid-cols-1 sm:grid-cols-[215px_1fr]">
        <img
          src={schedule.classImage}
          alt={schedule.className}
          className="w-full h-48 sm:h-full object-cover"
        />

        <CardContent className="p-5 space-y-5">
          <div className="space-y-4">
            <div>
              <h3 className="text-lg font-bold text-foreground">
                {schedule.className}
              </h3>
              <p className="text-sm text-muted-foreground">
                {schedule.instructorName} • {schedule.duration} mins
              </p>
            </div>
            <Badge className="text-xs" variant="outline">
              {schedule.bookingStatus}
            </Badge>
          </div>

          <div className="grid grid-cols-2 gap-6 text-sm">
            <div className="space-y-2">
              <div className="flex items-center gap-2">
                <CalendarIcon className="w-4 h-4" />
                <span>{format(startTime, "EEEE, dd MMM yyyy")}</span>
              </div>
              <div className="flex items-center gap-2">
                <ClockIcon className="w-4 h-4" />
                <span>
                  {formatHour(startTime)} - {formatHour(endTime)}
                </span>
              </div>
            </div>
            <div className="space-y-2">
              <div className="flex items-center gap-2">
                <MapPinIcon className="w-4 h-4" />
                <span>{schedule.location}</span>
              </div>
              <div className="flex items-center gap-2">
                <UserIcon className="w-4 h-4" />
                <span>{schedule.participant} Participant</span>
              </div>
            </div>
          </div>

          <p className="text-xs text-muted-foreground text-right">
            Booked at{" "}
            {format(new Date(schedule.bookedAt), "dd MMM yyyy, HH:mm")}
          </p>
        </CardContent>
      </div>
    </Card>
  );
};
