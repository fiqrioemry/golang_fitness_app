import auth from "@/services/auth";
import { useQuery } from "@tanstack/react-query";
import { useAuthStore } from "@/store/useAuthStore";

export const useAuthMe = () => {
  const setUser = useAuthStore((s) => s.setUser);
  return useQuery({
    queryKey: ["auth", "me"],
    queryFn: auth.getMe,
    onSuccess: (data) => setUser(data),
    onError: () => setUser(null),
    retry: 1,
    refetchOnWindowFocus: false,
  });
};
