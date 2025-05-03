import React from "react";
import { Button } from "@/components/ui/button";
import { Dialog, DialogContent, DialogTitle } from "@/components/ui/dialog";
import { Pencil, Trash } from "lucide-react";

const ClassScheduleDetail = ({ open, onClose, event }) => {
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
          <Button variant="destructive">
            <Trash />
            Delete Schedule
          </Button>
          <Button variant="outline">
            <Pencil />
            Update Schedule
          </Button>
        </div>
      </DialogContent>
    </Dialog>
  );
};

export { ClassScheduleDetail };
