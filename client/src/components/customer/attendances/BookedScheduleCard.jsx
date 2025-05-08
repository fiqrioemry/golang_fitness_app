import {
  useRegenerateQRCode,
  useCheckinAttendance,
} from "@/hooks/useSchedules";
import { toast } from "sonner";
import { useState } from "react";
import { format } from "date-fns";
import { Badge } from "@/components/ui/badge";
import { Button } from "@/components/ui/button";
import { Dialog, DialogContent, DialogTrigger } from "@/components/ui/dialog";

export const BookedScheduleCard = ({ schedule }) => {
  const [showQR, setShowQR] = useState(false);
  const [attended, setAttended] = useState(schedule.attended);
  const qrQuery = useRegenerateQRCode(showQR && attended ? schedule.id : null);
  const checkinMutation = useCheckinAttendance();

  const handleAttend = async () => {
    try {
      await checkinMutation.mutateAsync(schedule.id);
      setAttended(true);
      setShowQR(true);
    } catch (error) {
      toast.error("Failed to check in");
    }
  };

  return (
    <div className="border p-5 rounded-lg shadow-sm bg-card">
      <div className="flex gap-4">
        <img
          src={schedule.class.image}
          alt={schedule.class.title}
          className="w-32 h-32 rounded-md object-cover"
        />
        <div className="flex-1 space-y-2">
          <h3 className="text-lg font-semibold">{schedule.class.title}</h3>
          <p className="text-muted-foreground text-sm">
            {schedule.instructor.fullname} • {schedule.class.duration} mins
          </p>
          <div className="text-sm text-muted-foreground">
            {format(new Date(schedule.date), "dd MMM yyyy")} —{" "}
            {schedule.startHour}:
            {schedule.startMinute.toString().padStart(2, "0")}
          </div>
          <Badge variant="secondary">Booked</Badge>
          <div className="mt-4">
            <Dialog open={showQR} onOpenChange={setShowQR}>
              <DialogTrigger asChild>
                <Button
                  onClick={attended ? () => setShowQR(true) : handleAttend}
                >
                  {attended ? "Show QR Code" : "Attend Now"}
                </Button>
              </DialogTrigger>
              <DialogContent className="text-center space-y-4">
                <h4 className="text-lg font-semibold">
                  QR Code for Attendance
                </h4>
                {qrQuery.isLoading ? (
                  <p>Loading...</p>
                ) : qrQuery.isError ? (
                  <p className="text-red-500">Failed to load QR</p>
                ) : (
                  <img
                    src={`data:image/png;base64,${qrQuery.data}`}
                    alt="QR Code"
                    className="w-64 h-64 mx-auto"
                  />
                )}
                <p className="text-sm text-muted-foreground">
                  {schedule.class.title}
                </p>
              </DialogContent>
            </Dialog>
          </div>
        </div>
      </div>
    </div>
  );
};
