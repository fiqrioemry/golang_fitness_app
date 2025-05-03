import { useEffect, useState } from "react";
import { Controller, useFormContext } from "react-hook-form";
export const InputDateElement = ({
  name,
  label,
  rules = { required: true },
  mode = "past",
  ageLimit = 80,
}) => {
  const { control, setValue } = useFormContext();
  const [day, setDay] = useState("");
  const [month, setMonth] = useState("");
  const [year, setYear] = useState("");
  const [hasInteracted, setHasInteracted] = useState(false);

  const today = new Date();
  const startDay = today.getDate();
  const startMonth = today.getMonth() + 1;
  const startYearVal = today.getFullYear();

  const startYear = mode === "future" ? startYearVal : startYearVal - ageLimit;
  const endYear = mode === "future" ? startYearVal + 5 : startYearVal;

  const getDaysInMonth = (month, year) => {
    if (!month || !year) return 31;
    return new Date(year, month, 0).getDate();
  };

  useEffect(() => {
    if (!day && !month && !year) {
      setDay(startDay);
      setMonth(startMonth);
      setYear(startYearVal);
    }
  }, [startDay, startMonth, startYearVal, day, month, year]);

  useEffect(() => {
    const maxDay = getDaysInMonth(month, year);
    if (day > maxDay) setDay("");
  }, [month, year]);

  useEffect(() => {
    if (day && month && year) {
      const formatted = `${year}-${String(month).padStart(2, "0")}-${String(
        day
      ).padStart(2, "0")}`;
      setValue(name, formatted);
    }
  }, [day, month, year, name, setValue]);

  const daysInMonth = getDaysInMonth(month, year);
  const isSameMonthAndYear =
    Number(month) === startMonth && Number(year) === startYearVal;

  const minValidDay = (() => {
    if (mode === "future" && isSameMonthAndYear) return startDay;
    return 1;
  })();

  const maxValidDay = (() => {
    if (mode === "past" && isSameMonthAndYear && hasInteracted) return startDay;
    if (mode === "past" && !hasInteracted) return daysInMonth;
    return daysInMonth;
  })();
  const days = Array.from(
    { length: maxValidDay - minValidDay + 1 },
    (_, i) => i + minValidDay
  );

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
    (_, i) => startYear + i
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
            {/* Day */}
            <select
              value={day}
              onChange={(e) => setDay(Number(e.target.value))}
              className="border p-2 rounded w-1/3"
            >
              <option value="">Tanggal</option>
              {days.map((d) => (
                <option key={d} value={d}>
                  {d}
                </option>
              ))}
            </select>

            {/* Month */}
            <select
              value={month}
              onChange={(e) => {
                setMonth(Number(e.target.value));
                setHasInteracted(true);
              }}
              className="border p-2 rounded w-1/3"
            >
              <option value="">Bulan</option>
              {months.map((m, i) => {
                const monthVal = i + 1;
                const isDisabledFuture =
                  mode === "future" &&
                  Number(year) === startYearVal &&
                  monthVal < startMonth;
                const isDisabledPast =
                  mode === "past" &&
                  Number(year) === startYearVal &&
                  monthVal > startMonth;
                return (
                  <option
                    key={m}
                    value={monthVal}
                    disabled={isDisabledFuture || isDisabledPast}
                  >
                    {m}
                  </option>
                );
              })}
            </select>

            {/* Year */}
            <select
              value={year}
              onChange={(e) => {
                setYear(Number(e.target.value));
                setHasInteracted(true);
              }}
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
