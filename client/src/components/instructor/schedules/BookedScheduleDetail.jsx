import {
  Sheet,
  SheetContent,
  SheetHeader,
  SheetTitle,
  SheetClose,
  SheetDescription,
} from "@/components/ui/Sheet";
import { Badge } from "@/components/ui/Badge";
import { Button } from "@/components/ui/Button";
import { useNavigate, useParams } from "react-router-dom";
import { useBookingDetailQuery } from "@/hooks/useBooking";
import { formatDateTime, formatHour, buildDateTime } from "@/lib/utils";
import { BookedDetailSkeleton } from "@/components/loading/BookedDetailSkeleton";

export const BookedScheduleDetail = () => {
  const { id } = useParams();
  const navigate = useNavigate();

  const { data, isLoading } = useBookingDetailQuery(id);
  console.log(data);
  const startTime = buildDateTime(
    data?.date,
    data?.startHour,
    data?.startMinute
  );
  const endTime = buildDateTime(
    data?.date,
    data?.startHour,
    data?.startMinute + data?.duration
  );
  const isDisabled = data?.checkedIn && data?.checkedOut;

  console.log(data);
  return (
    <Sheet open={true} onOpenChange={() => navigate(-1)}>
      <SheetContent side="right" className="max-w-xl w-full">
        <SheetHeader>
          <SheetTitle>Class Detail</SheetTitle>
          <SheetDescription>Your booking and attendance info</SheetDescription>
        </SheetHeader>

        {isLoading ? (
          <BookedDetailSkeleton />
        ) : (
          <div className="mt-4 space-y-6">
            <img
              src={data?.classImage}
              alt={data?.className}
              className="w-full h-48 rounded-md object-cover border"
            />

            <div className="space-y-1">
              <h3 className="text-lg font-semibold">{data?.className}</h3>
              <p className="text-sm text-muted-foreground">
                {data?.instructorName} • {data?.duration} mins
              </p>
            </div>

            <div className="text-sm space-y-2">
              <DetailRow label="Date">{formatDateTime(startTime)}</DetailRow>
              <DetailRow label="Time">
                {formatHour(startTime)} - {formatHour(endTime)}
              </DetailRow>
              <DetailRow label="Attendance">
                <Badge variant="outline" className="capitalize">
                  {data.attendanceStatus}
                </Badge>
              </DetailRow>
              <DetailRow label="Checked in">
                {data.checkedAt ? formatDateTime(data.checkedAt) : "-"}
              </DetailRow>
              <DetailRow label="Checked out">
                {data.verifiedAt ? formatDateTime(data.verifiedAt) : "-"}
              </DetailRow>
            </div>

            {data.checkedIn && (
              <div className="text-center">
                {" "}
                <a
                  href={data.zoomLink}
                  target="_blank"
                  rel="noopener noreferrer"
                  className="block text-sm text-primary hover:underline"
                >
                  Join Zoom Class
                </a>
              </div>
            )}

            <div className="flex justify-between gap-2">
              <Button
                variant="secondary"
                disabled={!data.isOpened || data.checkedIn}
              >
                Check In
              </Button>

              <Button variant="default" disabled={isDisabled}>
                Check Out
              </Button>
            </div>
          </div>
        )}

        <SheetClose asChild>
          <Button variant="outline" className="w-full mt-6">
            Close
          </Button>
        </SheetClose>
      </SheetContent>
    </Sheet>
  );
};

const DetailRow = ({ label, children }) => (
  <p>
    <span className="font-medium text-muted-foreground">{label}:</span>{" "}
    <span className="text-foreground">{children}</span>
  </p>
);
