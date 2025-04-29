// src/components/input/InputElement.jsx
import { cn } from "@/lib/utils";
import { Controller, useFormContext } from "react-hook-form";

const InputElement = ({
  name,
  label,
  rows = 4,
  maxLength,
  type = "text",
  placeholder = "",
  className = "",
  isNumber = false,
  disabled = false,
  isTextArea = false,
  isNumeric = false,
  rules = { required: true },
}) => {
  const { control } = useFormContext();

  return (
    <Controller
      control={control}
      name={name}
      rules={rules}
      render={({ field, fieldState }) => {
        const handleKeyDown = (e) => {
          if (isNumeric) {
            // Allow basic control keys
            if (
              [
                "Backspace",
                "Tab",
                "Delete",
                "ArrowLeft",
                "ArrowRight",
              ].includes(e.key)
            )
              return;
            // Prevent non-numeric keys
            if (!/^[0-9]$/.test(e.key)) {
              e.preventDefault();
            }
          }
        };

        return (
          <div className={(cn("space-y-1"), className)}>
            {label && (
              <label
                htmlFor={name}
                className="block text-sm font-medium text-gray-700"
              >
                {label}
              </label>
            )}

            {isTextArea ? (
              <textarea
                id={name}
                {...field}
                value={field.value ?? ""}
                placeholder={placeholder}
                disabled={disabled}
                rows={rows}
                maxLength={maxLength || undefined}
                className="w-full border p-2 rounded resize-none disabled:bg-gray-100"
              />
            ) : (
              <input
                id={name}
                {...field}
                value={field.value ?? ""}
                onChange={(e) => field.onChange(e.target.value)}
                onKeyDown={handleKeyDown}
                type={isNumber ? "text" : type}
                placeholder={placeholder}
                disabled={disabled}
                inputMode={isNumber ? "numeric" : undefined}
                maxLength={maxLength || undefined}
                className="w-full border p-2 rounded disabled:bg-gray-100"
              />
            )}

            {fieldState.error && (
              <p className="text-red-500 text-xs mt-1">
                {fieldState.error.message || "This field is required"}
              </p>
            )}
          </div>
        );
      }}
    />
  );
};

export { InputElement };
