// src/components/input/InputTagsElement.jsx

import { Controller, useFormContext } from "react-hook-form";
import { useState } from "react";
import { X } from "lucide-react";

const InputTagsElement = ({
  name,
  label,
  placeholder = "Enter a tag and press Enter",
  disabled = false,
  rules = { required: true },
}) => {
  const { control } = useFormContext();
  const [inputValue, setInputValue] = useState("");

  return (
    <Controller
      name={name}
      control={control}
      rules={rules}
      render={({ field, fieldState }) => {
        const addTag = () => {
          const value = inputValue.trim();
          if (value && !(field.value || []).includes(value)) {
            field.onChange([...(field.value || []), value]);
          }
          setInputValue("");
        };

        const removeTag = (tag) => {
          const updated = (field.value || []).filter((t) => t !== tag);
          field.onChange(updated);
        };

        const handleKeyDown = (e) => {
          if (e.key === "Enter") {
            e.preventDefault();
            addTag();
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
            <div className="flex flex-wrap gap-2 p-2 border rounded">
              {(field.value || []).map((tag, idx) => (
                <div
                  key={idx}
                  className="flex items-center gap-1 px-2 py-1 bg-primary/10 rounded-full text-sm"
                >
                  {tag}
                  <button
                    type="button"
                    onClick={() => removeTag(tag)}
                    className="hover:text-red-500"
                  >
                    <X className="w-3 h-3" />
                  </button>
                </div>
              ))}
              <input
                id={name}
                type="text"
                value={inputValue}
                onChange={(e) => setInputValue(e.target.value)}
                onKeyDown={handleKeyDown}
                placeholder={placeholder}
                disabled={disabled}
                className="flex-1 min-w-[120px] border-none outline-none bg-transparent"
              />
            </div>
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

export { InputTagsElement };
