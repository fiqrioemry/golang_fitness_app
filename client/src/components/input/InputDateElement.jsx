// src/components/input/InputDateElement.jsx

import { Controller, useFormContext } from "react-hook-form";

const InputDateElement = ({
  name,
  label,
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
          <input
            id={name}
            type="date"
            {...field}
            placeholder={placeholder}
            disabled={disabled}
            className="w-full border p-2 rounded disabled:bg-gray-100"
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

export { InputDateElement };
