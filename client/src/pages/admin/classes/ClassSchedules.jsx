import { Loading } from "@/components/ui/Loading";
import React, { useState, useEffect } from "react";
import { useSchedulesQuery } from "@/hooks/useClass";
import { AddClassSchedule } from "@/components/admin/classes/AddClassSchedule";
import { ClassScheduleDetail } from "@/components/admin/classes/ClassScheduleDetail";
import { ClassScheduleCalendar } from "@/components/admin/classes/ClassScheduleCalendar";
import { UpdateClassSchedule } from "../../../components/admin/classes/UpdateClassSchedule";

const ClassSchedule = () => {
  const [events, setEvents] = useState([]);
  const [isFormOpen, setIsFormOpen] = useState(false);
  const [updateFormOpen, setUpdateFormOpen] = useState(false);
  const [isDetailOpen, setIsDetailOpen] = useState(false);
  const [selectedDate, setSelectedDate] = useState(null);
  const [selectedEvent, setSelectedEvent] = useState(null);
  const { data: schedules = [], isLoading } = useSchedulesQuery();

  const handleEmptySlotClick = (dateTime) => {
    setSelectedDate(dateTime);
    setIsFormOpen(true);
  };

  const handleSelectUpdate = (event) => {
    setSelectedEvent(event);
    setIsDetailOpen(false);
    setUpdateFormOpen(true);
  };
  const handleSelectEvent = (event) => {
    setSelectedEvent(event);
    setIsDetailOpen(true);
  };

  useEffect(() => {
    if (!schedules || schedules.length === 0) return;

    const mapped = schedules.map((item) => {
      const start = new Date(item.date);
      start.setHours(item.startHour);
      start.setMinutes(item.startMinute);

      const end = new Date(start);
      end.setMinutes(end.getMinutes() + 60);

      return {
        id: item.id,
        title: `${item.class} - ${item.instructor} (${item.bookedCount}/${item.capacity})`,
        start,
        end,
        allDay: false,
        resource: {
          ...item,
        },
      };
    });

    setEvents(mapped);
  }, [schedules]);

  if (isLoading) return <Loading />;
  console.log(selectedEvent);
  return (
    <section>
      <div className="space-y-1 text-center mb-4">
        <h2 className="text-2xl font-bold">Class Schedules Event</h2>
      </div>
      <AddClassSchedule
        open={isFormOpen}
        setOpen={setIsFormOpen}
        defaultDateTime={selectedDate}
      />
      <ClassScheduleDetail
        open={isDetailOpen}
        onUpdate={handleSelectUpdate}
        onClose={() => setIsDetailOpen(false)}
        event={selectedEvent}
      />
      <div className="w-full">
        <ClassScheduleCalendar
          events={events}
          schedules={schedules}
          onSelectEvent={handleSelectEvent}
          onEmptySlotClick={handleEmptySlotClick}
        />
      </div>
    </section>
  );
};

export default ClassSchedule;
