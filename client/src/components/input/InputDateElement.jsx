// src/components/input/InputDateDropdownElement.jsx

import { Controller, useFormContext } from "react-hook-form";
import { useEffect, useState } from "react";

export const InputDateElement = ({
  name,
  label,
  rules = { required: true },
  startYear = 1970,
  endYear = new Date().getFullYear(),
}) => {
  const { control, setValue } = useFormContext();
  const [day, setDay] = useState("");
  const [month, setMonth] = useState("");
  const [year, setYear] = useState("");

  // Update parent field as string (YYYY-MM-DD)
  useEffect(() => {
    if (day && month && year) {
      const formatted = `${year}-${String(month).padStart(2, "0")}-${String(
        day
      ).padStart(2, "0")}`;
      setValue(name, formatted);
    }
  }, [day, month, year, name, setValue]);

  const days = Array.from({ length: 31 }, (_, i) => i + 1);
  const months = [
    "Januari",
    "Februari",
    "Maret",
    "April",
    "Mei",
    "Juni",
    "Juli",
    "Agustus",
    "September",
    "Oktober",
    "November",
    "Desember",
  ];
  const years = Array.from(
    { length: endYear - startYear + 1 },
    (_, i) => endYear - i
  );

  return (
    <Controller
      name={name}
      control={control}
      rules={rules}
      render={({ fieldState }) => (
        <div className="space-y-1">
          {label && (
            <label className="block text-sm font-medium text-gray-700">
              {label}
            </label>
          )}
          <div className="flex gap-2">
            <select
              value={day}
              onChange={(e) => setDay(e.target.value)}
              className="border p-2 rounded w-1/3"
            >
              <option value="">Tanggal</option>
              {days.map((d) => (
                <option key={d} value={d}>
                  {d}
                </option>
              ))}
            </select>

            <select
              value={month}
              onChange={(e) => setMonth(e.target.value)}
              className="border p-2 rounded w-1/3"
            >
              <option value="">Bulan</option>
              {months.map((m, i) => (
                <option key={m} value={i + 1}>
                  {m}
                </option>
              ))}
            </select>

            <select
              value={year}
              onChange={(e) => setYear(e.target.value)}
              className="border p-2 rounded w-1/3"
            >
              <option value="">Tahun</option>
              {years.map((y) => (
                <option key={y} value={y}>
                  {y}
                </option>
              ))}
            </select>
          </div>

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
