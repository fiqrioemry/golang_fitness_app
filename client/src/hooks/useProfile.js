// src/hooks/useProfile.js
import { toast } from "sonner";
import * as profileService from "@/services/profile";
import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";

// =====================
// QUERY: GET PROFILE
// =====================

export const useProfileQuery = () =>
  useQuery({
    queryKey: ["user", "profile"],
    queryFn: profileService.getProfile,
  });

// =====================
// QUERY: USER PACKAGES
// =====================

export const useUserPackagesQuery = () =>
  useQuery({
    queryKey: ["user", "packages"],
    queryFn: profileService.getUserPackages,
  });

// =====================
// MUTATION: UPDATE PROFILE
// =====================

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
