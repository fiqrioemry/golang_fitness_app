import { QueryClient } from "@tanstack/react-query";

// Konfigurasi optimal untuk production
export const queryClient = new QueryClient({
  defaultOptions: {
    queries: {
      staleTime: 1000 * 60 * 5, // 5 menit dianggap fresh
      cacheTime: 1000 * 60 * 30, // 30 menit disimpan di memory sebelum GC
      retry: 1, // Retry sekali jika error
      refetchOnWindowFocus: false, // Tidak refetch saat tab aktif
      refetchOnMount: false, // Tidak refetch ulang saat mount ulang
      refetchOnReconnect: true, // Refetch saat internet reconnect
    },
    mutations: {
      retry: 0, // Mutasi biasanya tidak perlu auto-retry
    },
  },
});
