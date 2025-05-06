import React from "react";
import { Pencil } from "lucide-react";
import { Button } from "@/components/ui/button";
import { DeleteClassSchedule } from "./DeleteClassSchedule";
import { Dialog, DialogContent, DialogTitle } from "@/components/ui/dialog";

const ClassScheduleDetail = ({ open, onClose, event, onUpdate }) => {
  if (!event) return null;

  const { title, start, end, resource } = event;

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
          <strong>Capacity :</strong> {resource?.booked || 0}/
          {resource?.capacity}
        </p>
        <div className="mt-4 flex gap-4 justify-end">
          <DeleteClassSchedule onUpdate={onUpdate} schedule={resource} />
          <Button className="w-full" onClick={onUpdate} variant="secondary">
            <Pencil />
            <span>Update</span>
          </Button>
        </div>
      </DialogContent>
    </Dialog>
  );
};

export { ClassScheduleDetail };
