import { format } from "date-fns";
import React, { useEffect } from "react";
import { Button } from "@/components/ui/button";
import { Loading } from "@/components/ui/Loading";
import { ErrorDialog } from "@/components/ui/ErrorDialog";
import { CalendarClock, Edit, Trash2 } from "lucide-react";
import { useRecurringTemplatesQuery } from "@/hooks/useSchedules";
import { RunTemplate } from "@/components/admin/classes/RunTemplate";
import { StopTemplate } from "@/components/admin/classes/StopTemplate";

const weekdays = [
  "Sunday",
  "Monday",
  "Tuesday",
  "Wednesday",
  "Thursday",
  "Friday",
  "Saturday",
];

const ClassRecuring = () => {
  const {
    data: templates = [],
    isLoading,
    isError,
    refetch,
  } = useRecurringTemplatesQuery();

  useEffect(() => {
    refetch();
  }, []);

  if (isLoading) return <Loading />;
  if (isError) return <ErrorDialog onRetry={refetch} />;

  return (
    <section className="section max-w-7xl mx-auto px-6 py-12">
      <div className="text-center mb-10">
        <h2 className="text-3xl font-bold tracking-tight text-gray-800">
          Recurring Schedule Templates
        </h2>
        <p className="text-gray-500 text-sm">
          Manage, activate and generate recurring class schedules
        </p>
      </div>

      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
        {templates.map((t) => (
          <div
            key={t.id}
            className="rounded-2xl border border-gray-200 bg-white p-6 shadow-sm hover:shadow-md transition duration-200"
          >
            <div className="flex items-center justify-between mb-4">
              <div className="flex items-center gap-3">
                <CalendarClock className="w-5 h-5 text-indigo-600" />
                <h3 className="text-lg font-semibold text-gray-800 truncate max-w-[200px]">
                  {t.className}
                </h3>
              </div>
              <span
                className={`text-xs px-2 py-0.5 rounded-full font-medium ${
                  t.isActive
                    ? "bg-green-100 text-green-700"
                    : "bg-yellow-100 text-yellow-700"
                }`}
              >
                {t.isActive ? "Active" : "Inactive"}
              </span>
            </div>

            <ul className="text-sm text-gray-600 space-y-1 mb-4">
              <li>
                <strong>Instructor:</strong> {t.instructor}
              </li>
              <li>
                <strong>Capacity:</strong> {t.capacity}
              </li>
              <li>
                <strong>Time:</strong> {t.startHour}:
                {t.startMinute.toString().padStart(2, "0")}
              </li>
              <li>
                <strong>Days:</strong>{" "}
                {t.dayOfWeeks.map((d) => weekdays[d]).join(", ")}
              </li>
              <li>
                <strong>End Date:</strong>{" "}
                {format(new Date(t.endDate), "yyyy-MM-dd")}
              </li>
              <li>
                <strong>Frequency:</strong> {t.frequency}
              </li>
            </ul>

            <div className="flex flex-wrap gap-2">
              {t.isActive ? (
                <StopTemplate template={t} />
              ) : (
                <>
                  <RunTemplate template={t} />
                  <Button
                    variant="outline"
                    size="sm"
                    className="w-full"
                    onClick={() => handleUpdate(t.id)}
                  >
                    <Edit className="w-4 h-4 mr-2" /> Update
                  </Button>
                  <Button
                    variant="destructive"
                    size="sm"
                    className="w-full"
                    onClick={() => handleDelete(t.id)}
                  >
                    <Trash2 className="w-4 h-4 mr-2" /> Delete
                  </Button>
                </>
              )}
            </div>
          </div>
        ))}
      </div>
    </section>
  );
};

const handleRun = (id) => console.log("Run", id);
const handleStop = (id) => console.log("Stop", id);
const handleGenerate = (id) => console.log("Generate", id);
const handleUpdate = (id) => console.log("Update", id);
const handleDelete = (id) => console.log("Delete", id);

export default ClassRecuring;
