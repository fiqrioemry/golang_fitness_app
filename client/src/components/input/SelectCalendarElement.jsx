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

export const SelectCalendarElement = ({
  name,
  label,
  rules = { required: true },
  mode = "future",
  ageLimit = 80,
}) => {
  const { control } = useFormContext();

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
      render={({ field: { value, onChange }, fieldState }) => (
        <div className="space-y-1">
          {label && (
            <label className="block text-sm font-medium text-gray-700">
              {label}
            </label>
          )}

          <Popover>
            <PopoverTrigger asChild>
              <Button
                variant="outline"
                className={cn(
                  "w-full justify-start text-left font-normal",
                  !value && "text-muted-foreground"
                )}
              >
                {value ? format(value, "PPP") : "Pilih tanggal"}
              </Button>
            </PopoverTrigger>

            <PopoverContent className="w-auto p-0">
              <Calendar
                mode="single"
                selected={value}
                onSelect={onChange}
                initialFocus
                fromDate={minDate}
                toDate={maxDate}
              />
            </PopoverContent>
          </Popover>

          {fieldState.error && (
            <p className="text-red-500 text-xs mt-1">
              {fieldState.error.message}
            </p>
          )}
        </div>
      )}
    />
  );
};
