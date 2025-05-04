import React from "react";
import { scheduleSchema } from "@/lib/schema";
import { useScheduleMutation } from "@/hooks/useClass";
import { FormSheet } from "@/components/form/FormSheet";
import { SelectElement } from "@/components/input/SelectElement";
import { SwitchElement } from "@/components/input/SwitchElement";
import { ColorPickerElement } from "@/components/input/ColorPickerElement";
import { InputNumberElement } from "@/components/input/InputNumberElement";
import { SelectOptionsElement } from "@/components/input/SelectOptionsElement";
import { SelectCalendarElement } from "@/components/input/SelectCalendarElement";

const UpdateClassSchedule = ({ open, setOpen, schedule }) => {
  const { updateSchedule } = useScheduleMutation();

  const { isPending, mutateAsync } = updateSchedule;

  return (
    <FormSheet
      open={open}
      setOpen={setOpen}
      state={schedule}
      loading={isPending}
      schema={scheduleSchema}
      action={mutateAsync}
      title="Update Schedule"
      resourceId={schedule?.id}
    >
      <SelectCalendarElement name="date" label="Event Date" />

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
          placeholder="Select Hour"
          options={[8, 9, 10, 11, 12, 13, 14, 15, 16, 17]}
        />

        <SelectElement
          name="startMinute"
          label="Start Minute"
          isNumeric={true}
          placeholder="Select Minute"
          options={[0, 15, 30, 45]}
        />
      </div>
      <InputNumberElement name="capacity" label="Capacity" min={1} />
      <SwitchElement name="isActive" label="Set as active" />
    </FormSheet>
  );
};

export { UpdateClassSchedule };
