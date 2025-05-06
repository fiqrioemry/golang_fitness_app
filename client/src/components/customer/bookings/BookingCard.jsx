import { useEffect, useState } from "react";
import { Badge } from "@/components/ui/badge";
import { format } from "date-fns";
import { Card, CardContent } from "@/components/ui/card";
import { buildDateTime, getTimeLeft, isAttendanceWindow } from "@/lib/utils";
import { CalendarIcon, ClockIcon, MapPinIcon, UserIcon } from "lucide-react";
import { useQRCodeQuery, useCheckinAttendance } from "@/hooks/useAttendances";
import { Dialog, DialogContent, DialogTrigger } from "@/components/ui/dialog";
import { Button } from "@/components/ui/button";
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
      <Card className="p-4 bg-card border text-destructive border-border shadow-sm">
        <p>Invalid booking time.</p>
      </Card>
    );
  }
  return (
    <Card className="overflow-hidden border border-border bg-card shadow-md hover:shadow-xl transition">
      <div className="grid grid-cols-1 sm:grid-cols-[215px_1fr]">
        <img
          src={booking.classImage}
          alt={booking.classTitle}
          className="w-full h-48 sm:h-full object-cover"
        />

        <CardContent className="p-5 space-y-4">
          {/* Header Title & Status */}
          <div className="flex justify-between items-center space-x-4">
            <h3 className="text-base font-semibold text-foreground">
              {booking.classTitle}
            </h3>
            <Badge variant="outline" className="capitalize text-xs">
              {booking.status}
            </Badge>
          </div>

          {/* Instructor & Duration */}
          <p className="text-sm text-muted-foreground">
            {booking.instructor} • {booking.duration} mins
          </p>

          {/* Countdown */}
          <p className="text-sm font-medium text-primary">
            Starts in: {timeLeft} ⏳🔥
          </p>

          {/* Schedule Info */}
          <div className="grid grid-cols-2 gap-8 text-sm text-muted-foreground ">
            <div className="space-y-4">
              <div className="flex items-center gap-2">
                <CalendarIcon className="w-4 h-4" />
                <span>{format(startTime, "EEE, dd MMM yyyy")}</span>
              </div>
              <div className="flex items-center gap-2">
                <ClockIcon className="w-4 h-4" />
                <span>
                  {format(startTime, "HH:mm")} - {format(endTime, "HH:mm")}
                </span>
              </div>
            </div>
            <div className="space-y-4">
              <div className="flex items-center gap-2">
                <MapPinIcon className="w-4 h-4" />
                <span>{booking.location}</span>
              </div>
              <div className="flex items-center gap-2">
                <UserIcon className="w-4 h-4" />
                <span>{booking.participant} Participant</span>
              </div>
            </div>
          </div>

          {/* Attend Button & QR Dialog */}
          {!canAttend && (
            <Dialog open={showQR} onOpenChange={setShowQR}>
              <DialogTrigger asChild>
                <Button
                  type="button"
                  className="py-5 w-60"
                  onClick={booking.attended ? handleShowQR : handleAttend}
                >
                  {booking.attended ? "Show QR Code" : "Attend Now"}
                </Button>
              </DialogTrigger>

              <DialogContent className="text-center space-y-4">
                <h4 className="text-lg font-semibold">Scan QR Code</h4>
                {qrCodeQuery.isLoading ? (
                  <p className="text-sm text-muted-foreground">
                    Loading QR Code...
                  </p>
                ) : qrCodeQuery.isError || !qrCodeQuery.data ? (
                  <p className="text-destructive text-sm">
                    Failed to load QR Code
                  </p>
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

          {/* Footer */}
          <p className="text-xs text-muted-foreground text-right">
            Booked at {format(new Date(booking.bookedAt), "dd MMM yyyy, HH:mm")}
          </p>
        </CardContent>
      </div>
    </Card>
  );
};
