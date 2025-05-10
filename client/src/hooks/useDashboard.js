import { useQuery } from "@tanstack/react-query";
import * as dashboardService from "@/services/dashboard";

// GET summary stats
export const useDashboardSummary = () =>
  useQuery({
    queryKey: ["dashboard", "summary"],
    queryFn: dashboardService.getDashboardSummary,
  });

// GET revenue stats by range
export const useRevenueStats = (range = "daily") =>
  useQuery({
    queryKey: ["dashboard", "revenue", range],
    queryFn: () => dashboardService.getRevenueStats(range),
  });
