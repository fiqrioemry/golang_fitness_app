// src/components/input/InputTextareaElement.jsx

import { Controller, useFormContext } from "react-hook-form";

const InputTextareaElement = ({
  name,
  label,
  rows = 4,
  maxLength,
  placeholder = "",
  disabled = false,
  rules = { required: true },
}) => {
  const { control } = useFormContext();

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
          <textarea
            id={name}
            {...field}
            placeholder={placeholder}
            disabled={disabled}
            rows={rows}
            maxLength={maxLength}
            className="w-full border p-2 rounded resize-none disabled:bg-gray-100"
          />
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

export { InputTextareaElement };
