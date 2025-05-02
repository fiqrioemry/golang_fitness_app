// src/components/input/ClassFilterSelect.jsx
import React from "react";
import { useSearchParams } from "react-router-dom";
import { Skeleton } from "@/components/ui/skeleton";
import { useSelectOptions } from "@/hooks/useSelectOptions";

const FilterSelection = ({ paramKey, label, data }) => {
  const { data: options = [], isLoading } = useSelectOptions(data);
  const [searchParams, setSearchParams] = useSearchParams();
  const selected = searchParams.get(paramKey) || "";

  const handleChange = (e) => {
    const value = e.target.value;
    if (value) {
      searchParams.set(paramKey, value);
    } else {
      searchParams.delete(paramKey);
    }
    setSearchParams(searchParams);
  };

  if (isLoading) return <Skeleton className="h-9 rounded-md" />;

  return (
    <div className="space-y-1">
      {label && (
        <label className="block text-sm font-medium text-gray-700">
          {label}
        </label>
      )}
      <select
        value={selected}
        onChange={handleChange}
        className="w-full border p-2 rounded"
      >
        <option value="">All</option>
        {options.map((opt) => (
          <option key={opt.id} value={opt.id}>
            {opt.name}
          </option>
        ))}
      </select>
    </div>
  );
};

export default FilterSelection;
