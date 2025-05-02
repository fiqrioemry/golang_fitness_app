import React from "react";
import { Skeleton } from "@/components/ui/skeleton";
import { useSelectOptions } from "@/hooks/useSelectOptions";
import { Controller, useFormContext } from "react-hook-form";

const SelectOptionsElement = ({
  name,
  label,
  data = "category",
  disabled = false,
  rules = { required: true },
  placeholder = "Select an option",
}) => {
  const { control } = useFormContext();
  const { data: options = [], isLoading } = useSelectOptions(data);

  if (isLoading) return <Skeleton className="h-9 rounded-md" />;

  return (
    <Controller
      name={name}
      control={control}
      rules={rules}
      render={({ field, fieldState }) => (
        <div className="space-y-1">
          {label && (
            <label
              htmlFor={name}
              className="block text-sm font-medium text-gray-700"
            >
              {label}
            </label>
          )}
          <select
            id={name}
            {...field}
            disabled={disabled}
            className="w-full border p-2 rounded disabled:bg-gray-100"
          >
            <option value="">{placeholder}</option>
            {options.map((option) => (
              <option
                key={option.id || option.value || option}
                value={option.id || option.value || option}
              >
                {option.name || option.label || option}
              </option>
            ))}
          </select>
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

export { SelectOptionsElement };
