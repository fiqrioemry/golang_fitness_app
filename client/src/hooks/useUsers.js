// src/hooks/useUsers.js
import * as userService from "@/services/users";
import { useQuery } from "@tanstack/react-query";

// GET /api/admin/users
export const useUsersQuery = (params = {}) =>
  useQuery({
    queryKey: ["admin-users", params],
    queryFn: () => userService.getAllUsers(params),
    keepPreviousData: true,
    staleTime: 0,
  });

// GET /api/admin/users/:id
export const useUserDetailQuery = (id) =>
  useQuery({
    queryKey: ["admin-user", id],
    queryFn: () => userService.getUserDetail(id),
    enabled: !!id,
  });

// GET /api/admin/users/stats
export const useUserStatsQuery = () =>
  useQuery({
    queryKey: ["admin-user-stats"],
    queryFn: userService.getUserStats,
  });
