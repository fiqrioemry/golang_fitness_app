// src/store/useAuthStore.jsx
import { toast } from "sonner";
import { create } from "zustand";
import auth from "@/services/auth";
import { persist } from "zustand/middleware";
import { queryClient } from "@/lib/react-query";

export const useAuthStore = create(
  persist(
    (set, get) => ({
      user: null,
      loading: false,
      checkingAuth: true,
      rememberMe: false,
      resetStep: () => set({ step: 1 }),

      setUser: (user) => set({ user }),

      clearUser: () => set({ user: null }),

      setCheckingAuth: () => set({ checkingAuth: false }),

      setRememberMe: (remember) => set({ rememberMe: remember }),

      authMe: async () => {
        try {
          const user = await auth.getMe();
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

          if (formData.rememberMe) {
            get().setRememberMe(true);
          } else {
            get().setRememberMe(false);
          }

          await get().authMe();
        } catch (error) {
          console.log(error);
          toast.error(error.response?.data?.message || "Login failed");
        } finally {
          set({ loading: false });
        }
      },

      googleLogin: async (idToken) => {
        try {
          const { message } = await auth.googleSignIn(idToken);
          toast.success(message);
          await get().authMe();
        } catch (error) {
          console.error(error);
          toast.error(error.response?.data?.message || "Google Sign-In failed");
        }
      },

      logout: async () => {
        try {
          await auth.logout();
          get().clearUser();
          queryClient.clear();
        } catch (error) {
          console.error(error.message);
        }
      },

      sendOTP: async (formData) => {
        set({ loading: true });
        try {
          const { message } = await auth.sendOTP(formData);
          toast.success(message);
        } catch (error) {
          toast.error(error.response.data.message);
        } finally {
          set({ loading: false });
        }
      },

      register: async (formData) => {
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
            await get().authMe();
            set({ step: 1 });
          }
        } catch (error) {
          toast.error(error.response.data.message);
        } finally {
          set({ loading: false });
        }
      },
    }),
    {
      name: "auth-storage",
      partialize: (state) => ({
        user: state.user,
        rememberMe: state.rememberMe,
      }),
    }
  )
);
