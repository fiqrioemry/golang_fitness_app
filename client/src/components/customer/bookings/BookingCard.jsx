import {
  UserIcon,
  InfoIcon,
  ClockIcon,
  MapPinIcon,
  QrCodeIcon,
  CalendarIcon,
  AlertCircleIcon,
} from "lucide-react";
import {
  useRegenerateQRCode,
  useAttendanceMutation,
} from "@/hooks/useAttendance";
import { toast } from "sonner";
import { format } from "date-fns";
import { useEffect, useState } from "react";
import { Badge } from "@/components/ui/Badge";
import { Button } from "@/components/ui/Button";
import { Card, CardContent } from "@/components/ui/Card";
import {
  formatHour,
  getTimeLeft,
  buildDateTime,
  isAttendanceWindow,
} from "@/lib/utils";
import { Dialog, DialogContent, DialogTrigger } from "@/components/ui/Dialog";

export const BookingCard = ({ booking }) => {
  const [showQR, setShowQR] = useState(false);
  const [timeLeft, setTimeLeft] = useState("");
  const [canAttend, setCanAttend] = useState(false);
  const [isClassOver, setIsClassOver] = useState(false);

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
      setIsClassOver(endTime < new Date());
    };

    updateStatus();
    const timer = setInterval(updateStatus, 1000);
    return () => clearInterval(timer);
  }, [startTime, endTime, booking.duration]);

  const { checkin } = useAttendanceMutation();
  const qrCodeQuery = useRegenerateQRCode(showQR ? booking.id : null);

  const handleAttend = async () => {
    try {
      await checkin.mutateAsync(booking.id);
      setShowQR(true);
    } catch (error) {
      toast.error(error?.response?.data?.error || "Failed to check in");
    }
  };

  const handleShowQR = () => setShowQR(true);

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
          alt={booking.className}
          className="w-full h-48 sm:h-full object-cover"
        />

        <CardContent className="p-5 space-y-5">
          <div className="space-y-4">
            <div>
              <h3 className="text-lg font-bold text-foreground">
                {booking.className}
              </h3>
              <p className="text-sm text-muted-foreground">
                {booking.instructorName} • {booking.duration} mins
              </p>
            </div>
            <Badge className="text-xs" variant="outline">
              {booking.status}
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
                <span>{booking.location}</span>
              </div>
              <div className="flex items-center gap-2">
                <UserIcon className="w-4 h-4" />
                <span>{booking.participant} Participant</span>
              </div>
            </div>
          </div>

          {booking.status !== "checked_in" && (
            <div className="text-sm font-medium text-primary">
              Starts in: {timeLeft} ⏳
            </div>
          )}

          {booking.status !== "checked_in" && (
            <div className="flex items-center text-sm text-muted-foreground gap-2">
              <AlertCircleIcon className="w-4 h-4 text-yellow-500" />
              Attend button will be enabled 15 minutes before class.
            </div>
          )}

          {isClassOver && (
            <div className="flex items-center text-sm text-red-500 gap-2">
              <AlertCircleIcon className="w-4 h-4" />
              This class has already ended. You can no longer attend.
            </div>
          )}

          <div className="pt-2">
            <Dialog open={showQR} onOpenChange={setShowQR}>
              <DialogTrigger asChild>
                <Button
                  type="button"
                  className="w-full py-4"
                  onClick={
                    booking.status === "checked_in"
                      ? handleShowQR
                      : handleAttend
                  }
                  disabled={
                    booking.status === "checked_in" &&
                    (!canAttend || isClassOver)
                  }
                >
                  {booking.status === "checked_in" ? (
                    <>
                      <QrCodeIcon className="w-4 h-4 mr-2" /> Show QR Code
                    </>
                  ) : (
                    <>
                      <InfoIcon className="w-4 h-4 mr-2" /> Attend Now
                    </>
                  )}
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
                  <div>
                    <img
                      alt="QR Code"
                      src={`data:image/png;base64,${qrCodeQuery.data}`}
                      className="mx-auto w-60 h-60"
                    />
                  </div>
                )}
              </DialogContent>
            </Dialog>
          </div>

          <p className="text-xs text-muted-foreground text-right">
            Booked at {format(new Date(booking.bookedAt), "dd MMM yyyy, HH:mm")}
          </p>
        </CardContent>
      </div>
    </Card>
  );
};
