import React from "react";
import { Badge } from "@/components/ui/badge";
import { Card, CardContent } from "@/components/ui/card";
import { formatRupiah, formatDateTime } from "@/lib/utils";

export const TransactionCard = ({ transactions }) => {
  return (
    <div className="space-y-4">
      {transactions.map((tx) => (
        <Card key={tx.id} className="card card-hover p-0">
          <CardContent className="p-5">
            <div className="flex justify-between items-start gap-4">
              <div className="space-y-1">
                <h3 className="text-base font-semibold text-foreground">
                  {tx.packageName}
                </h3>
                <p className="text-subtitle">
                  Paid at: {formatDateTime(tx.paidAt)}
                </p>
                <p className="text-subtitle">
                  Method: {tx.paymentMethod.toUpperCase()}
                </p>
              </div>

              <div className="text-right space-y-1">
                <p className="text-lg font-semibold text-primary">
                  {formatRupiah(tx.price)}
                </p>
                <Badge variant="success" className="capitalize text-xs">
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
