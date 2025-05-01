import React from "react";
import { format } from "date-fns";
import { Badge } from "@/components/ui/badge";
import { CalendarCheck, Clock } from "lucide-react";

export const PackageCard = ({ pkgs }) => {
  return (
    <div className="grid gap-6">
      {pkgs.map((item) => (
        <div
          key={item.id}
          className="border rounded-xl shadow-sm p-5 bg-white space-y-3"
        >
          <div className="flex justify-between items-start">
            <div className="space-y-1">
              <h3 className="text-lg font-semibold text-gray-900">
                {item.packageName}
              </h3>
              <p className="text-sm text-muted-foreground">
                Purchased on {format(new Date(item.purchasedAt), "dd MMM yyyy")}
              </p>
            </div>
            <Badge
              variant={item.remainingCredit > 0 ? "default" : "destructive"}
            >
              {item.remainingCredit} sessions
            </Badge>
          </div>

          <div className="grid grid-cols-1 sm:grid-cols-2 gap-4 text-sm text-gray-700">
            <p className="flex items-center gap-2">
              <CalendarCheck className="w-4 h-4" />
              Expires on: {item.expiredAt}
            </p>
            <p className="flex items-center gap-2">
              <Clock className="w-4 h-4" />
              Time left: {item.expiredInDays} days
            </p>
          </div>
        </div>
      ))}
    </div>
  );
};
