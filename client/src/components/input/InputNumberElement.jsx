// src/components/input/InputNumberElement.jsx

import { Controller, useFormContext } from "react-hook-form";

const InputNumberElement = ({
  name,
  label,
  placeholder = "",
  disabled = false,
  rules = { required: true },
  max,
  min,
}) => {
  const { control } = useFormContext();

  return (
    <Controller
      name={name}
      control={control}
      rules={rules}
      render={({ field, fieldState }) => {
        const handleKeyDown = (e) => {
          const allowedKeys = [
            "Backspace",
            "Tab",
            "Delete",
            "ArrowLeft",
            "ArrowRight",
          ];
          if (allowedKeys.includes(e.key)) return;
          if (!/^[0-9]$/.test(e.key)) {
            e.preventDefault();
          }
        };

        return (
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
              type="text"
              value={field.value ?? ""}
              onChange={(e) => {
                const value = e.target.value;
                field.onChange(value === "" ? null : Number(value));
              }}
              onKeyDown={handleKeyDown}
              inputMode="numeric"
              placeholder={placeholder}
              disabled={disabled}
              min={min}
              max={max}
              className="w-full border p-2 rounded disabled:bg-gray-100"
            />
            {fieldState.error && (
              <p className="text-red-500 text-xs mt-1">
                {fieldState.error.message}
              </p>
            )}
          </div>
        );
      }}
    />
  );
};

export { InputNumberElement };
