import { Controller, useFormContext } from "react-hook-form";

const SelectElement = ({
  name,
  label,
  placeholder = "Select an option",
  options = [],
  disabled = false,
  rules = { required: true },
  isNumeric = false, // ✅ tambahkan opsi numeric
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
          <select
            id={name}
            value={field.value ?? ""}
            onChange={(e) => {
              const value = e.target.value;
              field.onChange(isNumeric ? Number(value) : value);
            }}
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

export { SelectElement };
