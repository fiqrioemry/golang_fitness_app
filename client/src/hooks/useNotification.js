// src/hooks/useNotification.js
import { toast } from "sonner";
import * as notifService from "@/services/notification";
import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";

export const useNotificationSettingsQuery = () =>
  useQuery({
    queryKey: ["notification-settings"],
    queryFn: notifService.getNotificationSettings,
    staleTime: 1000 * 60 * 10,
  });

export const useBrowserNotificationsQuery = () =>
  useQuery({
    queryKey: ["browser-notifications"],
    queryFn: notifService.getAllBrowserNotifications,
    staleTime: 0,
  });

export const useUpdateNotificationSetting = () => {
  const qc = useQueryClient();

  return useMutation({
    mutationFn: notifService.updateNotificationSetting,
    onSuccess: () => {
      qc.invalidateQueries({ queryKey: ["notification-settings"] });
      toast.success("Settings updated");
    },
    onError: (err) => {
      toast.error(err?.response?.data?.message || "Failed to update setting");
    },
  });
};

export const useSendPromoNotification = () => {
  return useMutation({
    mutationFn: notifService.sendPromoNotification,
    onSuccess: () => {
      toast.success("Promo notification sent successfully");
    },
    onError: (err) => {
      toast.error(err?.response?.data?.message || "Failed to send promo");
    },
  });
};
export const useMarkAllNotificationsAsRead = () => {
  return useMutation({
    mutationFn: notifService.markAllNotificationsAsRead,
    onError: (err) => {
      console.log(err);
      toast.error(err?.response?.data?.message || "Failed to mark all as read");
    },
  });
};
