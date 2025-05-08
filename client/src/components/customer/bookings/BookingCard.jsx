import {
  CalendarIcon,
  ClockIcon,
  MapPinIcon,
  UserIcon,
  QrCodeIcon,
  InfoIcon,
  AlertCircleIcon,
} from "lucide-react";
import {
  useRegenerateQRCode,
  useAttendanceMutation,
} from "@/hooks/useAttendance";
import { toast } from "sonner";
import { format } from "date-fns";
import { useEffect, useState } from "react";
import { Badge } from "@/components/ui/badge";
import { Button } from "@/components/ui/button";
import { Card, CardContent } from "@/components/ui/card";
import { buildDateTime, getTimeLeft, isAttendanceWindow } from "@/lib/utils";
import { Dialog, DialogContent, DialogTrigger } from "@/components/ui/dialog";

export const BookingCard = ({ booking }) => {
  const [timeLeft, setTimeLeft] = useState("");
  const [showQR, setShowQR] = useState(false);
  const [canAttend, setCanAttend] = useState(false);

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

  console.log();
  return (
    <Card className="overflow-hidden border border-border bg-card shadow-md hover:shadow-xl transition">
      <div className="grid grid-cols-1 sm:grid-cols-[215px_1fr]">
        <img
          src={booking.classImage}
          alt={booking.classTitle}
          className="w-full h-48 sm:h-full object-cover"
        />

        <CardContent className="p-5 space-y-5">
          {/* Header */}
          <div className="flex items-center justify-between">
            <div>
              <h3 className="text-lg font-bold text-foreground">
                {booking.classTitle}
              </h3>
              <p className="text-sm text-muted-foreground">
                {booking.instructor} • {booking.duration} mins
              </p>
            </div>
            <Badge className="text-xs" variant="outline">
              {booking.status}
            </Badge>
          </div>

          {/* Timing */}
          <div className="grid grid-cols-2 gap-6 text-sm">
            <div className="space-y-2">
              <div className="flex items-center gap-2">
                <CalendarIcon className="w-4 h-4" />
                <span>{format(startTime, "EEEE, dd MMM yyyy")}</span>
              </div>
              <div className="flex items-center gap-2">
                <ClockIcon className="w-4 h-4" />
                <span>
                  {format(startTime, "HH:mm")} - {format(endTime, "HH:mm")}
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

          {/* Countdown */}
          <div className="text-sm font-medium text-primary">
            Starts in: {timeLeft} ⏳
          </div>

          {/* Attend Info Notice */}
          {!canAttend && (
            <div className="flex items-center text-sm text-muted-foreground gap-2">
              <AlertCircleIcon className="w-4 h-4 text-yellow-500" />
              Attend button will be enabled 15 minutes before class.
            </div>
          )}

          {/* Button */}
          <div className="pt-2">
            <Dialog open={showQR} onOpenChange={setShowQR}>
              <DialogTrigger asChild>
                <Button
                  type="button"
                  className="w-full py-4"
                  onClick={booking.attended ? handleShowQR : handleAttend}
                  disabled={canAttend}
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

          {/* Footer */}
          <p className="text-xs text-muted-foreground text-right">
            Booked at {format(new Date(booking.bookedAt), "dd MMM yyyy, HH:mm")}
          </p>
        </CardContent>
      </div>
    </Card>
  );
};
