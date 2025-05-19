import { formatDateTime } from "@/lib/utils";
import { Button } from "@/components/ui/button";
import { Dialog } from "@/components/ui/dialog";
import { Loading } from "@/components/ui/Loading";
import React, { useState, useEffect } from "react";
import { ErrorDialog } from "@/components/ui/ErrorDialog";
import { useParams, useNavigate } from "react-router-dom";
import { useScheduleDetailQuery } from "@/hooks/useSchedules";
import { BookingDialog } from "@/components/schedule/BookingDialog";

const ScheduleDetail = () => {
  const { id } = useParams();
  const navigate = useNavigate();
  const [openDialog, setOpenDialog] = useState(false);

  const {
    data: schedule,
    isLoading,
    isError,
    refetch,
  } = useScheduleDetailQuery(id);

  useEffect(() => {
    if (schedule?.isBooked) {
      navigate("/profile/bookings");
    }
  }, [schedule, navigate]);

  if (isLoading) return <Loading />;

  if (isError) return <ErrorDialog onRetry={refetch} />;

  return (
    <section className="section py-24 text-foreground">
      <div className="space-y-4 text-center">
        <h2 className="font-bold">{schedule.className}</h2>
        <p className="text-muted-foreground text-sm">
          {formatDateTime(schedule.date)}
        </p>
      </div>

      <div className="border rounded-xl p-6 flex flex-col md:flex-row gap-6 items-start">
        <img
          src={schedule.classImage}
          alt={schedule.className}
          className="w-full md:w-64 aspect-square rounded-lg object-cover border"
        />

        <div className="flex-1 space-y-4">
          <div className="flex justify-between items-center">
            <span className="text-sm text-muted-foreground">Duration</span>
            <span className="font-medium">{schedule.duration} minutes</span>
          </div>
          <div className="flex justify-between items-center">
            <span className="text-sm text-muted-foreground">Start Time</span>
            <span className="font-medium">
              {String(schedule.startHour).padStart(2, "0")}:
              {String(schedule.startMinute).padStart(2, "0")}
            </span>
          </div>
          <div className="flex justify-between items-center">
            <span className="text-sm text-muted-foreground">Slots</span>
            <span className="font-medium">
              {schedule.bookedCount} / {schedule.capacity} booked
            </span>
          </div>
          <div className="flex items-center gap-4 pt-4 border-t mt-4">
            <div className="w-10 h-10 rounded-full bg-muted flex items-center justify-center text-lg font-bold uppercase">
              {schedule.instructorName.charAt(0)}
            </div>
            <div>
              <div className="font-semibold">{schedule.instructorName}</div>
            </div>
          </div>
        </div>
      </div>

      <div className="text-center">
        <Button onClick={() => setOpenDialog(true)}>Book This Class</Button>
      </div>

      <Dialog open={openDialog} onOpenChange={setOpenDialog}>
        <BookingDialog
          schedule={schedule}
          openDialog={openDialog}
          setOpenDialog={setOpenDialog}
        />
      </Dialog>
    </section>
  );
};

export default ScheduleDetail;
