import { CalendarIcon, ClockIcon, UserIcon } from "lucide-react";
import { Badge } from "@/components/ui/badge";
import { format, parseISO } from "date-fns";
import { Card, CardTitle } from "@/components/ui/card";
import { ReviewClass } from "./ReviewClass";

export const PastAttendanceCard = ({ attendance }) => {
  const {
    class: classData,
    instructor,
    date,
    startHour,
    startMinute,
    status,
    checkedAt,
    reviewed,
  } = attendance;

  const formattedDate = format(parseISO(date), "EEEE, dd MMM yyyy");
  const formattedTime = `${String(startHour).padStart(2, "0")}:${String(
    startMinute
  ).padStart(2, "0")}`;
  const checkedTime = checkedAt ? format(parseISO(checkedAt), "HH:mm") : "-";

  return (
    <Card className="flex flex-col sm:flex-row items-start sm:items-center gap-4 p-4 shadow-sm border border-border bg-muted/20">
      <img
        src={classData.image}
        alt={classData.title}
        className="w-full sm:w-48 h-32 object-cover rounded-xl"
      />

      <div className="flex-1 space-y-1">
        <CardTitle className="text-lg font-semibold">
          {classData.title}
        </CardTitle>
        <div className="text-sm text-muted-foreground">
          {classData.duration} min
        </div>
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
          <span>{instructor.fullname}</span>
        </div>

        {!reviewed && <ReviewClass cls={classData} />}
      </div>

      <div className="flex flex-col gap-1 items-end w-full sm:w-48">
        <Badge
          variant={status === "attended" ? "success" : "destructive"}
          className="w-fit"
        >
          {status.toUpperCase()}
        </Badge>
        <div className="text-xs text-muted-foreground">
          Checked at: <span className="font-medium">{checkedTime}</span>
        </div>
      </div>
    </Card>
  );
};
