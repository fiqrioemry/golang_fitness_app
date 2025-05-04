import { scheduleSchema } from "@/lib/schema";
import { useScheduleMutation } from "@/hooks/useClass";
import { FormSheet } from "@/components/form/FormSheet";
import { useFormContext, useWatch } from "react-hook-form";
import { SelectElement } from "@/components/input/SelectElement";
import { SwitchElement } from "@/components/input/SwitchElement";
import { DaySelectorElement } from "@/components/input/DaySelectorElement";
import { ColorPickerElement } from "@/components/input/ColorPickerElement";
import { InputNumberElement } from "@/components/input/InputNumberElement";
import { SelectOptionsElement } from "@/components/input/SelectOptionsElement";
import { SelectCalendarElement } from "@/components/input/SelectCalendarElement";

const AddClassSchedule = ({ open, setOpen, defaultDateTime }) => {
  const { createSchedule } = useScheduleMutation();
  const { isPending, mutateAsync } = createSchedule;

  const initialState = {
    isRecurring: false,
    classId: "",
    instructorId: "",
    capacity: 0,
    color: "#4ade80",
    date: defaultDateTime ? defaultDateTime.toISOString().split("T")[0] : "",
    startHour: defaultDateTime?.getHours(),
    startMinute: defaultDateTime?.getMinutes(),
    recurringDays: [],
    endType: "never",
    endDate: "",
  };

  return (
    <FormSheet
      open={open}
      setOpen={setOpen}
      loading={isPending}
      state={initialState}
      action={mutateAsync}
      schema={scheduleSchema}
      title="Create Class Schedule"
    >
      <RecurringSection />
      <ColorPickerElement name="color" label="Cardboard Color" />
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
      <div className="grid grid-cols-2 gap-4">
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
      </div>
      <InputNumberElement name="capacity" label="Capacity" />
    </FormSheet>
  );
};

const RecurringSection = () => {
  const { control } = useFormContext();
  const isRecurring = useWatch({ control, name: "isRecurring" });

  return (
    <>
      <SwitchElement name="isRecurring" label="Repeat weekly?" />

      {!isRecurring && <SelectCalendarElement name="date" label="Event Date" />}

      {isRecurring && (
        <>
          <DaySelectorElement name="recurringDays" label="Recurring Days" />
          <SelectElement
            name="endType"
            label="End Type"
            options={[
              { value: "never", label: "Never" },
              { value: "until", label: "Until date" },
            ]}
            placeholder="Select end type"
          />
          <SelectCalendarElement name="endDate" label="End Date (optional)" />
        </>
      )}
    </>
  );
};

export { AddClassSchedule };
