// src/hooks/useProfile.js
import { toast } from "sonner";
import * as profileService from "@/services/profile";
import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";

export const useProfileQuery = () =>
  useQuery({
    queryKey: ["user", "profile"],
    queryFn: profileService.getProfile,
    staleTime: 0,
  });

export const useUserPackagesQuery = () =>
  useQuery({
    queryKey: ["user-packages"],
    queryFn: profileService.getUserPackages,
    staleTime: 0,
  });

export const useUserClassPackagesQuery = (id) =>
  useQuery({
    queryKey: ["user-packages", id],
    queryFn: () => profileService.getUserPackagesByClassID(id),
    enabled: !!id,
  });

export const useUserTransactionsQuery = (page = 1, limit = 10) =>
  useQuery({
    queryKey: ["user", "transactions", page, limit],
    queryFn: () => profileService.getUserTransactions(page, limit),
  });

export const useUserBookingsQuery = (page = 1, limit = 10) =>
  useQuery({
    queryKey: ["user", "bookings", page, limit],
    queryFn: () => profileService.getUserBookings(page, limit),
    staleTime: 0,
  });

export const useProfileMutation = () => {
  const qc = useQueryClient();

  const mutationOpts = (msg) => ({
    onSuccess: () => {
      toast.success(msg);
      qc.invalidateQueries({ queryKey: ["user", "profile"] });
    },
    onError: (err) => {
      toast.error(err?.response?.data?.message || "Something went wrong");
    },
  });

  return {
    updateProfile: useMutation({
      mutationFn: profileService.updateProfile,
      ...mutationOpts("Profile updated"),
    }),

    updateAvatar: useMutation({
      mutationFn: profileService.updateAvatar,
      ...mutationOpts("Avatar updated"),
    }),
  };
};
