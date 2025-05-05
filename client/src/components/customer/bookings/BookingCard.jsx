import { useEffect, useState } from "react";
import { Badge } from "@/components/ui/badge";
import { format } from "date-fns";
import { Card, CardContent } from "@/components/ui/card";
import { buildDateTime, getTimeLeft, isAttendanceWindow } from "@/lib/utils";
import { CalendarIcon, ClockIcon, MapPinIcon, UserIcon } from "lucide-react";
import { useQRCodeQuery, useCheckinAttendance } from "@/hooks/useAttendances";
import { Dialog, DialogContent, DialogTrigger } from "@/components/ui/dialog";
import { toast } from "sonner";

export const BookingCard = ({ booking }) => {
  const [timeLeft, setTimeLeft] = useState("");
  const [canAttend, setCanAttend] = useState(false);
  const [showQR, setShowQR] = useState(false);

  const startTime = buildDateTime(
    booking.date,
    booking.startHour,
    booking.startMinute
  );
  const endTime = startTime
    ? new Date(startTime.getTime() + Number(booking.duration || 0) * 60000)
    : null;

  useEffect(() => {
    if (!startTime || !endTime) return;

    const updateStatus = () => {
      setTimeLeft(getTimeLeft(startTime));
      setCanAttend(
        isAttendanceWindow(startTime, Number(booking.duration || 0))
      );
    };

    updateStatus();
    const timer = setInterval(updateStatus, 1000);
    return () => clearInterval(timer);
  }, [startTime, endTime, booking.duration]);

  const checkinMutation = useCheckinAttendance();
  const qrCodeQuery = useQRCodeQuery(showQR ? booking.id : null);

  const handleAttend = async () => {
    try {
      await checkinMutation.mutateAsync(booking.id);
      setShowQR(true);
    } catch (error) {
      toast.error(error?.response?.data?.error || "Failed to check in");
    }
  };

  const handleShowQR = () => {
    setShowQR(true);
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
            <Dialog open={showQR} onOpenChange={setShowQR}>
              <DialogTrigger asChild>
                <button
                  onClick={booking.attended ? handleShowQR : handleAttend}
                  className="mt-2 px-4 py-2 bg-green-600 text-white rounded-md hover:bg-green-700"
                >
                  {booking.attended ? "Show QR Code" : "Attend Now"}
                </button>
              </DialogTrigger>
              <DialogContent>
                <h3 className="text-lg font-semibold mb-2">Scan QR Code</h3>
                {qrCodeQuery.isLoading ? (
                  <p>Loading QR Code...</p>
                ) : qrCodeQuery.isError || !qrCodeQuery.data ? (
                  <p className="text-red-500">Failed to load QR Code</p>
                ) : (
                  <img
                    src={`data:image/png;base64,${qrCodeQuery.data}`}
                    alt="QR Code"
                    className="mx-auto w-64 h-64"
                  />
                )}
              </DialogContent>
            </Dialog>
          )}

          <div className="text-xs text-muted-foreground text-right mt-2">
            Booked at {format(new Date(booking.bookedAt), "dd MMM yyyy, HH:mm")}
          </div>
        </CardContent>
      </div>
    </Card>
  );
};
