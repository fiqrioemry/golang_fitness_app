import {
  Select,
  SelectItem,
  SelectValue,
  SelectTrigger,
  SelectContent,
} from "@/components/ui/Select";
import { useState } from "react";
import { formatRupiah } from "@/lib/utils";
import { Card, CardContent } from "@/components/ui/Card";
import { ErrorDialog } from "@/components/ui/ErrorDialog";
import { useAdminPaymentsQuery } from "@/hooks/usePayment";
import { SummaryCard } from "@/components/admin/dashboard/SummaryCard";
import { RevenueChart } from "@/components/admin/dashboard/RevenueChart";
import { DashboardSkeleton } from "@/components/loading/DashboardSkeleton";
import { useDashboardSummary, useRevenueStats } from "@/hooks/useDashboard";
import { TransactionCard } from "@/components/admin/transactions/TransactionCard";

const Dashboard = () => {
  const [range, setRange] = useState("daily");
  const { data: revenue } = useRevenueStats(range);
  const { data: response } = useAdminPaymentsQuery({ limit: 5 });

  const { data: summary, isLoading, isError, refetch } = useDashboardSummary();

  if (isLoading) return <DashboardSkeleton />;

  if (isError) return <ErrorDialog onRetry={refetch} />;

  const transactions = response?.data || [];

  return (
    <div className="p-6 space-y-6">
      <div>
        <h3 className="mb-4">Statistic Summary</h3>
        <div className="grid grid-cols-2 lg:grid-cols-4 gap-4">
          <SummaryCard title="Total Users" value={summary?.totalUsers} />
          <SummaryCard title="Active Classes" value={summary?.totalClasses} />
          <SummaryCard title="Total Bookings" value={summary?.totalBookings} />
          <SummaryCard
            title="Total Renue"
            value={formatRupiah(summary?.totalRevenue)}
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
        <Card className="border">
          <CardContent className=" p-0">
            <TransactionCard transactions={transactions} />
          </CardContent>
        </Card>
      </div>
    </div>
  );
};
export default Dashboard;
