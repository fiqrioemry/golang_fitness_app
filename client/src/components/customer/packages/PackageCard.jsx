import React from "react";
import { Badge } from "@/components/ui/badge";
import { format, differenceInDays, isBefore } from "date-fns";
import { CalendarCheck, Clock, AlertTriangle } from "lucide-react";

export const PackageCard = ({ pkgs }) => {
  const today = new Date();

  return (
    <div className="grid gap-6">
      {pkgs.map((item) => {
        const expiredDate = new Date(item.expiredAt);
        const daysLeft = differenceInDays(expiredDate, today);
        const isExpired = isBefore(expiredDate, today);

        return (
          <div
            key={item.id}
            className={`border rounded-xl p-5 space-y-4 transition-shadow ${
              isExpired ? "bg-gray-100 opacity-60" : "bg-white"
            } shadow-sm hover:shadow-md`}
          >
            <div className="flex justify-between items-start">
              <div className="space-y-1">
                <h3 className="text-lg font-semibold text-gray-900">
                  {item.packageName}
                </h3>
                <p className="text-sm text-muted-foreground">
                  Purchased on{" "}
                  {format(new Date(item.purchasedAt), "dd MMM yyyy")}
                </p>
              </div>

              <div className="text-right space-y-1">
                <Badge
                  variant={isExpired ? "destructive" : "default"}
                  className="text-xs"
                >
                  {item.remainingCredit} session
                </Badge>
                <div className="text-xs mt-1">
                  {isExpired ? (
                    <span className="text-red-500">Expired</span>
                  ) : (
                    <span className="text-green-600">Active</span>
                  )}
                </div>
              </div>
            </div>

            <div className="grid grid-cols-1 sm:grid-cols-2 gap-4 text-sm text-muted-foreground">
              <p className="flex items-center gap-2">
                <CalendarCheck className="w-4 h-4" />
                Expires on: {format(expiredDate, "dd MMM yyyy")}
              </p>
              <p className="flex items-center gap-2">
                <Clock className="w-4 h-4" />
                {isExpired ? (
                  <>
                    <AlertTriangle className="w-4 h-4 text-red-500" />
                    Expired {Math.abs(daysLeft)} day(s) ago
                  </>
                ) : (
                  <>Time left: {daysLeft} day(s)</>
                )}
              </p>
            </div>
          </div>
        );
      })}
    </div>
  );
};
