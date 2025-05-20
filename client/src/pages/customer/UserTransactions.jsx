import { ErrorDialog } from "@/components/ui/ErrorDialog";
import { useUserTransactionsQuery } from "@/hooks/useProfile";
import { SectionTitle } from "@/components/header/SectionTitle";
import { SectionSkeleton } from "@/components/loading/SectionSkeleton";
import { NoTransaction } from "@/components/customer/transactions/NoTransaction";
import { TransactionCard } from "@/components/customer/transactions/TransactionCard";

const UserTransactions = () => {
  const { data, isLoading, isError, refetch } = useUserTransactionsQuery();

  if (isLoading) return <SectionSkeleton />;

  if (isError) return <ErrorDialog onRetry={refetch} />;

  const transactions = data.transactions || [];

  return (
    <section className="section p-8 space-y-6">
      <SectionTitle
        title="Transaction History"
        description=" All your recent payment activities and package purchases are listed
          here. Stay updated with your transaction status and history."
      />

      {transactions.length === 0 ? (
        <NoTransaction />
      ) : (
        <TransactionCard transactions={transactions} />
      )}
    </section>
  );
};

export default UserTransactions;
