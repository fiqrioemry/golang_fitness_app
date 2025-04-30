// src/store/useAuthStore.jsx

import { create } from "zustand";
import { persist } from "zustand/middleware";
import toast from "react-hot-toast";
import auth from "@/services/auth";

export const useAuthStore = create(
  persist(
    (set, get) => ({
      step: 1,
      user: null,
      loading: false,
      checkingAuth: true,

      setUser: (user) => set({ user }),

      resetStep: () => set({ step: 1 }),

      clearUser: () => set({ user: null }),

      authMe: async () => {
        try {
          const { user } = await auth.getMe();
          set({ user });
        } catch {
          set({ user: null });
        } finally {
          set({ checkingAuth: false });
        }
      },

      sendOTP: async (email) => {
        set({ loading: true });
        try {
          const { message } = await auth.sendOTP(email);
          toast.success(message);
        } catch (error) {
          toast.error(error.message);
        } finally {
          set({ loading: false });
        }
      },

      login: async (formData) => {
        set({ loading: true });
        try {
          const { message } = await auth.login(formData);
          await get().authMe();
          toast.success(message);
        } catch (error) {
          toast.error(error.message);
        } finally {
          set({ loading: false });
        }
      },

      logout: async () => {
        try {
          const { message } = await auth.logout();
          toast.success(message);
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
