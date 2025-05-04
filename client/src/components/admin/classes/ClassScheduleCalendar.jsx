import React from "react";
import { localizer } from "@/lib/utils";
import { Calendar, Views } from "react-big-calendar";
import ClassScheduleHead from "./ClassScheduleHead";
import "react-big-calendar/lib/css/react-big-calendar.css";

const ClassScheduleCalendar = ({ onEmptySlotClick, onSelectEvent, events }) => {
  return (
    <div className="w-full overflow-x-auto">
      <div className="min-w-[700px]">
        <Calendar
          components={{
            toolbar: ClassScheduleHead,
          }}
          className="rounded border shadow-sm"
          selectable
          localizer={localizer}
          events={events}
          startAccessor="start"
          endAccessor="end"
          style={{ height: "calc(100vh - 100px)" }}
          defaultView={Views.WEEK}
          views={["week"]}
          min={new Date(1970, 1, 1, 8, 0)}
          max={new Date(1970, 1, 1, 17, 0)}
          showMultiDayTimes={true}
          allDayAccessor={false}
          showAllDayEvents={false}
          onSelectSlot={(slotInfo) => {
            const dateTime = new Date(slotInfo.start);
            onEmptySlotClick(dateTime);
          }}
          onSelectEvent={onSelectEvent}
          eventPropGetter={(event) => {
            console.log();
            const bgColor = event?.resource?.color || "#4f46e5";
            return {
              style: {
                backgroundColor: bgColor,
                color: "#fff",
                borderRadius: "4px",
                padding: "4px",
              },
            };
          }}
        />
      </div>
    </div>
  );
};

export { ClassScheduleCalendar };
// onSelect={(date) => {
//   if (date) onChange(date.toISOString().split("T")[0]);
// }}
