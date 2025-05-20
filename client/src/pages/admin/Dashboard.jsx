import {
  Select,
  SelectItem,
  SelectValue,
  SelectTrigger,
  SelectContent,
} from "@/components/ui/select";

import {
  Table,
  TableRow,
  TableCell,
  TableBody,
  TableHead,
  TableHeader,
} from "@/components/ui/table";
import { useState } from "react";
import { Badge } from "@/components/ui/badge";
import { Card, CardContent } from "@/components/ui/card";
import { ErrorDialog } from "@/components/ui/ErrorDialog";
import { formatDateTime, formatRupiah } from "@/lib/utils";
import { useAdminPaymentsQuery } from "@/hooks/usePayment";
import { SummaryCard } from "@/components/admin/dashboard/SummaryCard";
import { RevenueChart } from "@/components/admin/dashboard/RevenueChart";
import { DashboardSkeleton } from "@/components/loading/DashboardSkeleton";
import { useDashboardSummary, useRevenueStats } from "@/hooks/useDashboard";

const Dashboard = () => {
  const [range, setRange] = useState("daily");
  const { data: revenue } = useRevenueStats(range);
  const { data: response } = useAdminPaymentsQuery();
  const { data: summary, isLoading, isError, refetch } = useDashboardSummary();

  if (isLoading) return <DashboardSkeleton />;

  if (isError) return <ErrorDialog onRetry={refetch} />;

  const transactions = response.payments || [];

  return (
    <div className="p-6 space-y-6">
      {/* Header */}
      <div>
        <h3 className="mb-4">Statistic Summary</h3>
        <div className="grid grid-cols-2 lg:grid-cols-4 gap-4">
          <SummaryCard title="Total Users" value={summary.totalUsers} />
          <SummaryCard title="Active Classes" value={summary.totalClasses} />
          <SummaryCard title="Total Bookings" value={summary.totalBookings} />
          <SummaryCard
            title="Total Revenue"
            value={`Rp${summary.totalRevenue.toLocaleString()}`}
          />
        </div>
      </div>

      <div>
        <h3 className="mb-4">Transaction Volume</h3>
        <div className="bg-background rounded-xl shadow p-6">
          <div className="flex items-center justify-between mb-4">
            <h2 className="text-lg font-semibold">Revenue</h2>
            <Select value={range} onValueChange={setRange}>
              <SelectTrigger className="w-[120px]">
                <SelectValue placeholder="Range" />
              </SelectTrigger>
              <SelectContent>
                <SelectItem value="daily">Daily</SelectItem>
                <SelectItem value="monthly">Monthly</SelectItem>
                <SelectItem value="yearly">Yearly</SelectItem>
              </SelectContent>
            </Select>
          </div>

          <RevenueChart data={revenue?.revenueSeries} range={revenue?.range} />
        </div>
      </div>

      <div>
        <h3 className="mb-4">Recent transaction</h3>
        <Card className="border shadow-sm">
          <CardContent className="overflow-x-auto p-0">
            {/* dekstop view */}
            <div className="bg-background hidden md:block w-full">
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

            {/* mobile view */}
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
          </CardContent>
        </Card>
      </div>
    </div>
  );
};
export default Dashboard;
