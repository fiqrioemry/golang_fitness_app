import React from "react";
import { updateScheduleSchema } from "@/lib/schema";
import { SelectElement } from "@/components/input/SelectElement";
import { useScheduleTemplateMutation } from "@/hooks/useSchedules";
import { FormUpdateDialog } from "@/components/form/FormUpdateDialog";
import { DaySelectorElement } from "@/components/input/DaySelectorElement";
import { InputNumberElement } from "@/components/input/InputNumberElement";
import { SelectOptionsElement } from "@/components/input/SelectOptionsElement";
import { SelectCalendarElement } from "@/components/input/SelectCalendarElement";

const UpdateTemplate = ({ template }) => {
  const { updateTemplate } = useScheduleTemplateMutation();
  const { isPending, mutateAsync } = updateTemplate;

  return (
    <FormUpdateDialog
      icon={false}
      state={template}
      title="Update Recurence Schedule"
      loading={isPending}
      schema={updateScheduleSchema}
      action={mutateAsync}
    >
      <SelectOptionsElement
        data="class"
        label="Class"
        name="classId"
        placeholder="Select class"
      />
      <SelectOptionsElement
        data="instructor"
        label="Instructor"
        name="instructorId"
        placeholder="Select instructor"
      />
      <SelectCalendarElement name="endDate" label="Event Date" />

      <DaySelectorElement name="dayOfWeeks" label="Recurring Days" />
      <div className="grid grid-cols-3 gap-4">
        <SelectElement
          name="startHour"
          label="Start Hour"
          isNumeric={true}
          placeholder="Hour"
          options={[...Array(10)].map((_, i) => 8 + i)} // 8–17
        />
        <SelectElement
          isNumeric={true}
          name="startMinute"
          label="Start Minute"
          placeholder="Minute"
          options={[0, 15, 30, 45]}
        />
        <InputNumberElement name="capacity" label="Capacity" />
      </div>
    </FormUpdateDialog>
  );
};

export { UpdateTemplate };
