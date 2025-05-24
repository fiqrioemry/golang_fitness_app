import { ReviewClass } from "./ReviewClass";
import { Badge } from "@/components/ui/Badge";
import { Card, CardTitle } from "@/components/ui/Card";
import { formatHour, formatDate, formatTime } from "@/lib/utils";
import { CalendarIcon, ClockIcon, UserIcon } from "lucide-react";

export const PastAttendanceCard = ({ attendance }) => {
  const formattedTime = formatTime(
    attendance.startHour,
    attendance.startMinute
  );
  const checkedTime = attendance.checkedAt
    ? formatHour(attendance.checkedAt)
    : "-";
  const formattedDate = formatDate(attendance.checkedAt);

  console.log(attendance);
  return (
    <Card className="flex flex-col sm:flex-row items-start sm:items-center gap-4 p-4 shadow-sm border border-border bg-muted/20">
      <img
        src={attendance.classImage}
        alt={attendance.className}
        className="w-full sm:w-48 h-32 object-cover rounded-xl"
      />

      <div className="flex-1">
        <div className=" flex justify-between">
          <div>
            <CardTitle className="text-lg font-semibold">
              {attendance.className}
            </CardTitle>

            <div className="flex items-center gap-2 text-sm text-muted-foreground">
              <CalendarIcon className="w-4 h-4" />
              <span>{formattedDate}</span>
            </div>
            <div className="flex items-center gap-2 text-sm text-muted-foreground">
              <ClockIcon className="w-4 h-4" />
              <span>{formattedTime}</span>
            </div>
            <div className="flex items-center gap-2 text-sm text-muted-foreground">
              <UserIcon className="w-4 h-4" />
              <span>{attendance.instructorName}</span>
            </div>
          </div>

          <div className="flex flex-col gap-1 items-end w-full sm:w-48">
            <Badge
              variant={
                attendance.status === "attended" ? "success" : "destructive"
              }
              className="w-fit"
            >
              {attendance.status.toUpperCase()}
            </Badge>
            <div className="text-xs text-muted-foreground">
              Checked at: <span className="font-medium">{checkedTime}</span>
            </div>
          </div>
        </div>
        <div className="flex justify-end mt-4">
          {!attendance.reviewed && <ReviewClass attendance={attendance} />}
        </div>
      </div>
    </Card>
  );
};
