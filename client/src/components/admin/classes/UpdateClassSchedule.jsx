import React from "react";
import { operationHours } from "@/lib/constant";
import { updateScheduleSchema } from "@/lib/schema";
import { useScheduleMutation } from "@/hooks/useSchedules";
import { SelectElement } from "@/components/input/SelectElement";
import { FormUpdateSheet } from "@/components/form/FormUpdateSheet";
import { ColorPickerElement } from "@/components/input/ColorPickerElement";
import { InputNumberElement } from "@/components/input/InputNumberElement";
import { SelectOptionsElement } from "@/components/input/SelectOptionsElement";
import { SelectCalendarElement } from "@/components/input/SelectCalendarElement";

const UpdateClassSchedule = ({ schedule, onClose }) => {
  const { updateSchedule } = useScheduleMutation();
  const { isPending, mutateAsync } = updateSchedule;

  const handleUpdateSchedule = async (data) => {
    await mutateAsync({ id: schedule.id, ...data });
    onClose();
  };

  return (
    <FormUpdateSheet
      icon={false}
      state={schedule}
      loading={isPending}
      title="Update Schedule"
      schema={updateScheduleSchema}
      action={handleUpdateSchedule}
    >
      <SelectCalendarElement name="date" label="End Date" />
      <ColorPickerElement name="colorCode" label="Cardboard Color" />
      <SelectOptionsElement
        data="class"
        name="classId"
        label="Class"
        placeholder="Select option for class"
      />
      <SelectOptionsElement
        name="instructorId"
        data="instructor"
        label="Instructor"
        placeholder="select option for instructor"
      />
      <div className="grid grid-cols-2 gap-4">
        <SelectElement
          name="startHour"
          label="Start Hour"
          isNumeric={true}
          options={operationHours}
          placeholder="Select Hour"
        />

        <SelectElement
          name="startMinute"
          label="Start Minute"
          isNumeric={true}
          options={operationHours}
          placeholder="Select Minute"
        />
      </div>
      <InputNumberElement name="capacity" label="Capacity" min={1} />
    </FormUpdateSheet>
  );
};

export { UpdateClassSchedule };
