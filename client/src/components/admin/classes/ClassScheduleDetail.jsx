import React from "react";
import { DeleteClassSchedule } from "./DeleteClassSchedule";
import { UpdateClassSchedule } from "./UpdateClassSchedule";
import { Dialog, DialogContent, DialogTitle } from "@/components/ui/dialog";

const ClassScheduleDetail = ({ open, onClose, event }) => {
  if (!event) return null;

  const { title, start, end, resource } = event;

  const isPast = new Date(resource?.date) < new Date();
  const hasBooking = resource?.bookedCount > 0;

  return (
    <Dialog open={open} onOpenChange={onClose}>
      <DialogContent>
        <DialogTitle>{title}</DialogTitle>
        <p>
          <strong>Date :</strong> {new Date(start).toLocaleString()} -{" "}
          {new Date(end).toLocaleTimeString()}
        </p>
        <p>
          <strong>Instructor :</strong> {resource?.instructor}
        </p>
        <p>
          <strong>Capacity :</strong> {resource?.bookedCount || 0}/
          {resource?.capacity}
        </p>
        {!isPast && (
          <div className="mt-4 flex gap-4 justify-end">
            {!hasBooking && (
              <DeleteClassSchedule onClose={onClose} schedule={resource} />
            )}
            <UpdateClassSchedule onClose={onClose} schedule={resource} />
          </div>
        )}
      </DialogContent>
    </Dialog>
  );
};

export { ClassScheduleDetail };
