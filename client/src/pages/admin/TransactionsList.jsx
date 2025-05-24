import { Loading } from "@/components/ui/Loading";
import { useDebounce } from "@/hooks/useDebounce";
import { paymentStatusOptions } from "@/lib/constant";
import { useQueryStore } from "@/store/useQueryStore";
import { Pagination } from "@/components/ui/Pagination";
import { Card, CardContent } from "@/components/ui/Card";
import { ErrorDialog } from "@/components/ui/ErrorDialog";
import { SearchInput } from "@/components/ui/SearchInput";
import { useAdminPaymentsQuery } from "@/hooks/usePayment";
import { SectionTitle } from "@/components/header/SectionTitle";
import { SearchNotFound } from "@/components/ui/SearchNotFound";
import { FilterSelection } from "@/components/ui/FilterSelecction";
import { TransactionCard } from "@/components/admin/transactions/TransactionCard";

const TransactionsList = () => {
  const { page, limit, q, sort, setPage, status, setQ, setSort, setStatus } =
    useQueryStore();

  const debouncedQ = useDebounce(q, 500);
  const { data, isLoading, isError, refetch } = useAdminPaymentsQuery({
    q: debouncedQ,
    status,
    sort,
    page,
    limit,
  });

  const transactions = data?.data || [];
  const pagination = data?.pagination;

  return (
    <section className="section">
      <SectionTitle
        title="Transaction List"
        description="Manage all user transactions and monitor payment activities."
      />

      <div className="flex items-center justify-between gap-4">
        <SearchInput
          q={q}
          setQ={setQ}
          setPage={setPage}
          placeholder={"search by name or email"}
        />

        <FilterSelection
          value={status}
          onChange={setStatus}
          options={paymentStatusOptions}
        />
      </div>

      <Card className="border shadow-sm">
        <CardContent className="overflow-x-auto p-0">
          {isLoading ? (
            <Loading />
          ) : isError ? (
            <ErrorDialog onRetry={refetch} />
          ) : transactions.length === 0 ? (
            <SearchNotFound title="No transactions found" q={q} />
          ) : (
            <TransactionCard
              sort={sort}
              setSort={setSort}
              transactions={transactions}
            />
          )}
          {pagination && (
            <Pagination
              page={pagination.page}
              onPageChange={setPage}
              limit={pagination.limit}
              total={pagination.totalRows}
            />
          )}
        </CardContent>
      </Card>
    </section>
  );
};

export default TransactionsList;
