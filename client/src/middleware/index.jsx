import { Fragment, useEffect } from "react";
import { useNavigate } from "react-router-dom";
import { useAuthStore } from "@/store/useAuthStore";

export const AuthRoute = ({ children }) => {
  const navigate = useNavigate();
  const { user } = useAuthStore();

  useEffect(() => {
    if (!user || user.role !== "customer") {
      navigate("/");
    }
  }, [user, navigate]);

  if (!user || user.role !== "customer") return null;

  return <Fragment>{children}</Fragment>;
};

export const NonAuthRoute = ({ children }) => {
  const navigate = useNavigate();
  const { user } = useAuthStore();

  useEffect(() => {
    if (user) {
      navigate("/");
    }
  }, [user, navigate]);

  if (user) return null;

  return <Fragment>{children}</Fragment>;
};
