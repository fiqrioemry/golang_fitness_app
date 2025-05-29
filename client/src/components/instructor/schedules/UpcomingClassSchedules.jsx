import {
  CalendarIcon,
  ClockIcon,
  MapPinIcon,
  ArrowRightIcon,
} from "lucide-react";
import { Badge } from "@/components/ui/Badge";
import { Button } from "@/components/ui/Button";
import { Card, CardContent } from "@/components/ui/Card";
import { useLocation, useNavigate } from "react-router-dom";
import { buildDateTime, formatDateTime, formatHour } from "@/lib/utils";

export const UpcomingClassSchedules = ({ schedule }) => {
  const location = useLocation();
  const navigate = useNavigate();

  const openModal = (id) => {
    navigate(`/profile/bookings/${id}`, {
      state: { backgroundLocation: location },
    });
  };

  const startTime = buildDateTime(
    schedule.date,
    schedule.startHour,
    schedule.startMinute
  );
  const endTime = buildDateTime(
    schedule.date,
    schedule.startHour,
    schedule.startMinute + schedule.duration
  );

  return (
    <Card>
      <div className="grid grid-cols-1 sm:grid-cols-[220px_1fr]">
        <img
          src={schedule.classImage}
          alt={schedule.className}
          className="w-full h-48 sm:h-full object-cover"
        />

        <CardContent className="p-5 flex flex-col justify-between gap-5">
          <div className="space-y-4 max-w-xl w-full">
            <div>
              <h3 className="text-lg font-semibold text-foreground">
                {schedule.className}
              </h3>
              <p className="text-sm text-muted-foreground">
                {schedule.instructorName} • {schedule.duration} mins
              </p>
            </div>

            <Badge variant="outline" className="text-xs capitalize">
              {schedule.bookingStatus}
            </Badge>

            <div className="flex justify-between text-sm">
              <div className="space-y-2">
                <div className="flex items-center gap-2">
                  <CalendarIcon className="w-4 h-4" />
                  <span>{formatDateTime(startTime)}</span>
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
              </div>
            </div>
          </div>

          <div className="flex justify-between items-center text-xs text-muted-foreground max-w-xl w-full">
            <span>Booked at {formatDateTime(schedule.bookedAt)}</span>
            <Button
              size="sm"
              variant="outline"
              onClick={() => openModal(schedule.id)}
            >
              See Detail <ArrowRightIcon className="w-4 h-4" />
            </Button>
          </div>
        </CardContent>
      </div>
    </Card>
  );
};
