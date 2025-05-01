import React from "react";
import { Loading } from "@/components/ui/Loading";
import { ErrorDialog } from "@/components/ui/ErrorDialog";
import { useUserTransactionsQuery } from "@/hooks/useProfile";
import { NoTransaction } from "@/components/customer/transactions/NoTransaction";
import { TransactionCard } from "@/components/customer/transactions/TransactionCard";

const UserTransactions = () => {
  const {
    data: response,
    isLoading,
    isError,
    refetch,
  } = useUserTransactionsQuery();

  if (isLoading) return <Loading />;

  if (isError) return <ErrorDialog onRetry={refetch} />;

  const transactions = response.transactions || [];

  return (
    <section className="max-w-6xl mx-auto px-4 py-8 space-y-6">
      <div className="space-y-1 text-center">
        <h2 className="text-2xl font-bold">Transaction History</h2>
        <p className="text-muted-foreground text-sm">
          All your recent payment activities and package purchases are listed
          here. Stay updated with your transaction status and history.
        </p>
      </div>

      {transactions.length === 0 ? (
        <NoTransaction />
      ) : (
        <TransactionCard transactions={transactions} />
      )}
    </section>
  );
};

export default UserTransactions;
