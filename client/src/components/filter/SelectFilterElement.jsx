import React from "react";

const SelectFilterElement = ({
  name,
  label,
  value,
  onChange,
  options = [],
  disabled = false,
  placeholder = "Select an option",
}) => {
  return (
    <div className="space-y-1">
      {label && (
        <label htmlFor={name} className="block text-sm font-medium">
          {label}
        </label>
      )}
      <select
        id={name}
        name={name}
        value={value}
        onChange={(e) => onChange && onChange(e)}
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
    </div>
  );
};

export { SelectFilterElement };
