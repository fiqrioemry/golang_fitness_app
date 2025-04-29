// src/store/useAuthStore.jsx
import { create } from "zustand";
import { persist } from "zustand/middleware";

export const useAuthStore = create(
  persist(
    (set) => ({
      user: null,
      loading: false,
      setUser: (user) => set({ user }),
      clearUser: () => set({ user: null }),
      login: async (data) => {
        await authApi.login(data);
        const user = await authApi.me();
        set({ user });
      },
      logout: async () => {
        await authApi.logout();
        set({ user: null });
      },
      register: async (data) => {
        await authApi.register(data);
        const user = await authApi.me();
        set({ user });
      },
    }),
    { name: "auth-storage", partialize: (s) => ({ user: s.user }) }
  )
);
