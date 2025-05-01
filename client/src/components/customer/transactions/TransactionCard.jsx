import React from "react";
import { Badge } from "@/components/ui/badge";
import { Card, CardContent } from "@/components/ui/card";
import { formatRupiah, formatDateTime } from "@/lib/utils";

export const TransactionCard = ({ transactions }) => {
  return (
    <div className="space-y-4">
      {transactions.map((tx) => (
        <Card key={tx.id} className="shadow-sm border rounded-2xl">
          <CardContent className="p-5">
            <div className="flex justify-between items-center">
              <div>
                <h3 className="text-lg font-medium">{tx.packageName}</h3>
                <p className="text-sm text-muted-foreground">
                  Paid at: {formatDateTime(tx.paidAt)}
                </p>
                <p className="text-sm text-muted-foreground">
                  Method: {tx.paymentMethod.toUpperCase()}
                </p>
              </div>
              <div className="text-right">
                <p className="text-lg font-semibold text-primary">
                  {formatRupiah(tx.price)}
                </p>
                <Badge variant="success" className="mt-1 capitalize">
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
