// src/store/useAuthStore.jsx

import { create } from "zustand";
import auth from "@/services/auth";
import toast from "react-hot-toast";
import { persist } from "zustand/middleware";

export const useAuthStore = create(
  persist(
    (set, get) => ({
      user: null,
      loading: false,
      checkingAuth: true,

      setUser: (user) => set({ user }),

      clearUser: () => set({ user: null }),

      setCheckingAuth: () => set({ checkingAuth: false }),

      authMe: async () => {
        try {
          const { user } = await auth.getMe();
          set({ user });
        } catch (err) {
          set({ user: null });
        } finally {
          set({ checkingAuth: false });
        }
      },

      login: async (formData) => {
        set({ loading: true });
        try {
          const { message } = await auth.login(formData);
          toast.success(message);
          await get().authMe();
        } catch (error) {
          toast.error(error.message);
        } finally {
          set({ loading: false });
        }
      },

      logout: async () => {
        try {
          await auth.logout();
          set({ user: null });
        } catch (error) {
          console.error(error.message);
        }
      },

      register: async (formData, navigate) => {
        set({ loading: true });
        try {
          const step = get().step;
          if (step === 1) {
            const { message } = await auth.sendOTP(formData);
            toast.success(message);
            set({ step: 2 });
          } else if (step === 2) {
            const { message } = await auth.verifyOTP(formData);
            toast.success(message);
            set({ step: 3 });
          } else if (step === 3) {
            const { message } = await auth.register(formData);
            toast.success(message);
            set({ step: 1 });
            navigate("/signin");
          }
        } catch (error) {
          toast.error(error.message);
        } finally {
          set({ loading: false });
        }
      },
    }),

    {
      name: "auth-storage",
      partialize: (state) => ({ user: state.user }),
    }
  )
);
