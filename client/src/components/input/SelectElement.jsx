// src/components/input/SelectElement.jsx

import { Controller, useFormContext } from "react-hook-form";

const SelectElement = ({
  name,
  label,
  options = [],
  disabled = false,
  placeholder = "Select an option",
}) => {
  const { control } = useFormContext();

  return (
    <Controller
      control={control}
      name={name}
      rules={{
        required: true,
        setValueAs: (value) => (value ? Number(value) : undefined),
      }}
      render={({ field, fieldState }) => (
        <div className="space-y-1">
          {label && (
            <label className="block text-sm font-medium">{label}</label>
          )}
          <select
            {...field}
            disabled={disabled}
            className="w-full border p-2 rounded disabled:bg-gray-100"
          >
            <option value="">{placeholder}</option>
            {options.map((option) => (
              <option
                key={option.id || option.name || option}
                value={option.id || option.name || option}
              >
                {option.name || option}
              </option>
            ))}
          </select>

          {fieldState.error && (
            <p className="text-red-500 text-xs mt-1">
              {fieldState.error.message || "This field is required"}
            </p>
          )}
        </div>
      )}
    />
  );
};

export { SelectElement };
