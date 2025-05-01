import React from "react";
import { Badge } from "@/components/ui/badge";
import { Loading } from "@/components/ui/Loading";
import { Card, CardContent } from "@/components/ui/card";
import { ErrorDialog } from "@/components/ui/ErrorDialog";
import { formatRupiah, formatDateTime } from "@/lib/utils";
import { useUserTransactionsQuery } from "@/hooks/useProfile";

const UserTransactions = () => {
  const {
    data: transactions = [],
    isLoading,
    isError,
    refetch,
  } = useUserTransactionsQuery();

  if (isLoading) return <Loading />;
  if (isError) return <ErrorDialog onRetry={refetch} />;

  return (
    <section className="max-w-3xl mx-auto px-4 py-6">
      <h2 className="text-2xl font-semibold mb-6">Transaction History</h2>

      {transactions.length === 0 ? (
        <p className="text-muted-foreground">No transactions yet.</p>
      ) : (
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
      )}
    </section>
  );
};

export default UserTransactions;
