import {
  Table,
  TableRow,
  TableCell,
  TableBody,
  TableHead,
  TableHeader,
} from "@/components/ui/table";
import { useState } from "react";
import { Input } from "@/components/ui/input";
import { Badge } from "@/components/ui/badge";
import { Loading } from "@/components/ui/Loading";
import { useDebounce } from "@/hooks/useDebounce";
import { Pagination } from "@/components/ui/Pagination";
import { Card, CardContent } from "@/components/ui/card";
import { ErrorDialog } from "@/components/ui/ErrorDialog";
import { formatDateTime, formatRupiah } from "@/lib/utils";
import { useAdminPaymentsQuery } from "@/hooks/usePayment";

const TransactionsList = () => {
  const limit = 10;
  const [page, setPage] = useState(1);
  const [search, setSearch] = useState("");
  const debouncedSearch = useDebounce(search, 500);
  const params = { q: debouncedSearch, page, limit };
  const {
    data: response = {},
    isLoading,
    isError,
    refetch,
  } = useAdminPaymentsQuery(params);

  const transactions = Array.isArray(response.payments)
    ? response.payments
    : [];
  const total = typeof response.total === "number" ? response.total : 0;

  return (
    <section className="section">
      <div className="space-y-1 text-center">
        <h2 className="text-2xl font-bold">Transaction List</h2>
        <p className="text-muted-foreground text-sm">
          Manage all user transactions and monitor payment activities.
        </p>
      </div>

      <div className="flex items-center justify-between gap-4">
        <Input
          value={search}
          onChange={(e) => setSearch(e.target.value)}
          placeholder="Search by name or email..."
          className="w-full md:w-80"
        />
      </div>

      <Card className="border shadow-sm">
        <CardContent className="overflow-x-auto p-0">
          {isLoading ? (
            <Loading />
          ) : isError ? (
            <ErrorDialog onRetry={refetch} />
          ) : transactions.length === 0 ? (
            <div className="py-12 text-center text-gray-500 text-sm">
              No transactions found{search && ` for "${search}"`}
            </div>
          ) : (
            <>
              {/* Desktop Table */}
              <div className="hidden md:block w-full">
                <Table>
                  <TableHeader>
                    <TableRow>
                      <TableHead className="text-center">User</TableHead>
                      <TableHead className="text-center">Email</TableHead>
                      <TableHead className="text-center">Package</TableHead>
                      <TableHead className="text-center">Price</TableHead>
                      <TableHead className="text-center">Method</TableHead>
                      <TableHead className="text-center">Status</TableHead>
                      <TableHead className="text-center">Paid At</TableHead>
                    </TableRow>
                  </TableHeader>
                  <TableBody className="h-12">
                    {transactions.map((tx) => (
                      <TableRow
                        key={tx.id}
                        className="border-t border-border hover:bg-muted transition"
                      >
                        <TableCell>{tx.fullname}</TableCell>
                        <TableCell>{tx.userEmail}</TableCell>
                        <TableCell>{tx.packageName}</TableCell>
                        <TableCell>{formatRupiah(tx.price)}</TableCell>
                        <TableCell>
                          {tx.paymentMethod?.toUpperCase() || "-"}
                        </TableCell>
                        <TableCell>
                          <Badge
                            variant={
                              tx.status === "success"
                                ? "default"
                                : tx.status === "failed"
                                ? "destructive"
                                : "secondary"
                            }
                          >
                            {tx.status}
                          </Badge>
                        </TableCell>
                        <TableCell>
                          {" "}
                          {tx.status === "success"
                            ? formatDateTime(tx.paidAt)
                            : "-"}
                        </TableCell>
                      </TableRow>
                    ))}
                  </TableBody>
                </Table>
              </div>

              {/* Mobile view */}
              <div className="md:hidden space-y-4 p-4  w-full">
                {transactions.map((tx) => (
                  <div
                    key={tx.id}
                    className="border rounded-lg p-4 shadow-sm space-y-2"
                  >
                    <div>
                      <h3 className="text-base font-semibold">{tx.fullname}</h3>
                      <p className="text-sm text-muted-foreground">
                        {tx.userEmail}
                      </p>
                    </div>
                    <div className="text-sm space-y-4">
                      <div className="grid grid-cols-2 gap-4">
                        <p>
                          <strong>{tx.packageName}</strong>
                        </p>
                        <p>{formatRupiah(tx.price)}</p>
                        <p>{tx.paymentMethod?.toUpperCase() || "-"}</p>
                        <p>
                          <Badge
                            variant={
                              tx.status === "success"
                                ? "default"
                                : tx.status === "failed"
                                ? "destructive"
                                : "secondary"
                            }
                          >
                            {tx.status}
                          </Badge>
                        </p>
                      </div>
                      <div className="space-x-4">
                        {" "}
                        <span className="text-muted-foreground">Paid At</span>
                        <span className="text-right whitespace-nowrap">
                          {tx.status === "success"
                            ? formatDateTime(tx.paidAt)
                            : "-"}
                        </span>
                      </div>
                    </div>
                  </div>
                ))}
              </div>
            </>
          )}
          {transactions.length > 0 && (
            <Pagination
              page={page}
              limit={limit}
              total={total}
              onPageChange={(p) => setPage(p)}
            />
          )}
        </CardContent>
      </Card>
    </section>
  );
};

export default TransactionsList;
