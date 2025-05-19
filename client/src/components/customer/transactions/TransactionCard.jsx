import React from "react";
import { Badge } from "@/components/ui/badge";
import { Card, CardContent } from "@/components/ui/card";
import { formatRupiah, formatDateTime } from "@/lib/utils";

export const TransactionCard = ({ transactions }) => {
  return (
    <div className="space-y-4">
      {transactions.map((tx) => (
        <Card
          key={tx.id}
          className="border border-border bg-card shadow-sm hover:shadow-md transition"
        >
          <CardContent className="p-4">
            <div className="flex flex-col sm:flex-row sm:items-center sm:justify-between gap-2 sm:gap-6 w-full ">
              <div className="text-sm font-medium text-foreground truncate">
                {tx.packageName}
              </div>

              <div className="text-sm text-muted-foreground whitespace-nowrap">
                {formatDateTime(tx.paidAt)}
              </div>

              <div className="text-sm text-muted-foreground whitespace-nowrap uppercase">
                {tx.paymentMethod}
              </div>

              <div className="text-sm font-semibold text-primary whitespace-nowrap">
                {formatRupiah(tx.price)}
              </div>

              <div className="whitespace-nowrap">
                <Badge
                  variant={tx.status === "success" ? "success" : "outline"}
                  className="capitalize text-xs"
                >
                  {tx.status}
                </Badge>
              </div>
            </div>
          </CardContent>
        </Card>
      ))}
    </div>
  );
};
