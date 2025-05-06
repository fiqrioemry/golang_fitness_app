import {
  Popover,
  PopoverTrigger,
  PopoverContent,
} from "@/components/ui/popover";
import { format } from "date-fns";
import { cn } from "@/lib/utils";
import { Button } from "@/components/ui/button";
import { Calendar } from "@/components/ui/calendar";
import { useFormContext, Controller } from "react-hook-form";
import { useState } from "react";

export const SelectCalendarElement = ({
  name,
  label,
  rules = { required: true },
  mode = "future",
  ageLimit = 80,
}) => {
  const { control } = useFormContext();
  const [open, setOpen] = useState(false);

  const today = new Date();
  const minDate =
    mode === "past" ? new Date(today.getFullYear() - ageLimit, 0, 1) : today;
  const maxDate =
    mode === "past" ? today : new Date(today.getFullYear() + 5, 11, 31);

  return (
    <Controller
      name={name}
      control={control}
      rules={rules}
      render={({ field: { value, onChange }, fieldState }) => {
        const selectedDate = value ? new Date(value) : undefined;

        return (
          <div className="space-y-1">
            {label && <label className="label">{label}</label>}

            <Popover open={open} onOpenChange={setOpen}>
              <PopoverTrigger asChild>
                <Button
                  variant="outline"
                  type="button"
                  className={cn(
                    "input text-left font-normal px-3 py-2",
                    !value && "text-muted-foreground"
                  )}
                >
                  {value ? format(new Date(value), "PPP") : "Pilih tanggal"}
                </Button>
              </PopoverTrigger>

              <PopoverContent className="w-auto p-0 bg-popover text-popover-foreground border border-border rounded-xl shadow-lg">
                <Calendar
                  mode="single"
                  selected={selectedDate}
                  onSelect={(date) => {
                    if (date) {
                      onChange(date.toISOString());
                      setOpen(false);
                    }
                  }}
                  initialFocus
                  fromDate={minDate}
                  toDate={maxDate}
                />
              </PopoverContent>
            </Popover>

            {fieldState.error && (
              <p className="error-message">{fieldState.error.message}</p>
            )}
          </div>
        );
      }}
    />
  );
};
