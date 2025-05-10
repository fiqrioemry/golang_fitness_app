import React from "react";
import { updateScheduleSchema } from "@/lib/schema";
import { SelectElement } from "@/components/input/SelectElement";
import { operationHours, operationMinutes } from "@/lib/constant";
import { useScheduleTemplateMutation } from "@/hooks/useSchedules";
import { FormUpdateDialog } from "@/components/form/FormUpdateDialog";
import { DaySelectorElement } from "@/components/input/DaySelectorElement";
import { InputNumberElement } from "@/components/input/InputNumberElement";
import { SelectOptionsElement } from "@/components/input/SelectOptionsElement";
import { SelectCalendarElement } from "@/components/input/SelectCalendarElement";

const UpdateTemplate = ({ template }) => {
  const { updateTemplate } = useScheduleTemplateMutation();

  return (
    <FormUpdateDialog
      icon={false}
      state={template}
      schema={updateScheduleSchema}
      title="Update Recurence Schedule"
      loading={updateTemplate.isPending}
      action={updateTemplate.mutateAsync}
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
          options={operationHours}
        />
        <SelectElement
          isNumeric={true}
          name="startMinute"
          label="Start Minute"
          placeholder="Minute"
          options={operationMinutes}
        />
        <InputNumberElement name="capacity" label="Capacity" />
      </div>
    </FormUpdateDialog>
  );
};

export { UpdateTemplate };
