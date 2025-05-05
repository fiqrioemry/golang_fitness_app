import React from "react";
import { Badge } from "@/components/ui/badge";
import { format, differenceInDays, isBefore } from "date-fns";
import {
  CalendarCheck,
  Clock,
  AlertTriangle,
  XCircle,
  CheckCircle,
} from "lucide-react";

export const PackageCard = ({ pkgs }) => {
  const today = new Date();

  return (
    <div className="grid gap-6">
      {pkgs.map((item) => {
        const expiredDate = new Date(item.expiredAt);
        const daysLeft = differenceInDays(expiredDate, today);

        const isExpired = isBefore(expiredDate, today);
        const isCreditEmpty = item.remainingCredit === 0;
        const isActive = !isExpired && item.remainingCredit > 0;

        return (
          <div
            key={item.id}
            className={`border rounded-xl p-5 space-y-4 transition-shadow ${
              isActive ? "bg-white" : "bg-gray-100 opacity-60"
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
                  variant={isActive ? "default" : "destructive"}
                  className="text-xs"
                >
                  {item.remainingCredit} CREDIT
                </Badge>
                <div className="text-xs mt-1 flex items-center gap-1">
                  {isActive ? (
                    <>
                      <CheckCircle className="w-4 h-4 text-green-500" />
                      <span className="text-green-600">Active</span>
                    </>
                  ) : isExpired ? (
                    <>
                      <XCircle className="w-4 h-4 text-red-500" />
                      <span className="text-red-500">Expired</span>
                    </>
                  ) : (
                    <>
                      <AlertTriangle className="w-4 h-4 text-yellow-500" />
                      <span className="text-yellow-600">No Credit</span>
                    </>
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
                    <span>Expired {Math.abs(daysLeft)} day(s) ago</span>
                  </>
                ) : (
                  <span>Time left: {daysLeft} day(s)</span>
                )}
              </p>
            </div>
          </div>
        );
      })}
    </div>
  );
};
