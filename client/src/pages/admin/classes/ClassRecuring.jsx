import { format } from "date-fns";
import React, { useEffect } from "react";
import { CalendarClock, PlusCircle } from "lucide-react";
import { Loading } from "@/components/ui/Loading";
import { ErrorDialog } from "@/components/ui/ErrorDialog";
import { useRecurringTemplatesQuery } from "@/hooks/useSchedules";
import { RunTemplate } from "@/components/admin/classes/RunTemplate";
import { StopTemplate } from "@/components/admin/classes/StopTemplate";
import { DeleteTemplate } from "@/components/admin/classes/DeleteTemplate";
import { UpdateTemplate } from "@/components/admin/classes/UpdateTemplate";
import { useNavigate } from "react-router-dom";
import { Button } from "@/components/ui/button";
import { Badge } from "@/components/ui/badge";

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
  const { data, isLoading, isError, refetch } = useRecurringTemplatesQuery();
  const navigate = useNavigate();

  useEffect(() => {
    refetch();
  }, []);

  if (isLoading) return <Loading />;
  if (isError) return <ErrorDialog onRetry={refetch} />;

  const templates = data || [];

  return (
    <section className="section max-w-7xl mx-auto px-4 py-10 space-y-6 text-foreground">
      <div className="text-center space-y-1">
        <h2 className="text-2xl font-bold">Recurring Schedule Templates</h2>
        <p className="text-muted-foreground text-sm">
          Manage, activate, and generate recurring class schedules.
        </p>
      </div>

      {templates.length === 0 ? (
        <div className="flex flex-col items-center justify-center py-16 text-center">
          <img
            src="/no-bookings.webp"
            alt="No templates"
            className="w-72 h-72 mb-6 object-contain"
          />
          <h3 className="text-lg font-semibold text-foreground">
            No templates yet
          </h3>
          <p className="text-muted-foreground mb-6">
            You haven’t created any recurring schedule template. Let’s start
            now!
          </p>
          <Button onClick={() => navigate("/admin/classes/schedules")}>
            <PlusCircle className="w-4 h-4 mr-2" />
            Add Template
          </Button>
        </div>
      ) : (
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
          {templates.map((t) => (
            <div
              key={t.id}
              className="border rounded-xl p-5 bg-background shadow-sm hover:shadow transition"
            >
              <div className="flex items-center justify-between mb-4">
                <div className="flex items-center gap-2">
                  <CalendarClock className="w-4 h-4 text-primary" />
                  <h3 className="font-semibold text-base truncate max-w-[200px]">
                    {t.className}
                  </h3>
                </div>
                <Badge variant={t.isActive ? "success" : "outline"}>
                  {t.isActive ? "Active" : "Inactive"}
                </Badge>
              </div>

              <ul className="text-sm text-muted-foreground space-y-1 mb-4">
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
                    <UpdateTemplate template={t} />
                    <DeleteTemplate template={t} />
                  </>
                )}
              </div>
            </div>
          ))}
        </div>
      )}
    </section>
  );
};

export default ClassRecuring;
