import { toast } from "sonner";
import { useNavigate } from "react-router-dom";
import { useAuthStore } from "@/store/useAuthStore";
import { useCreatePaymentMutation } from "./usePayment";

export function usePackageTransaction(packageId) {
  const navigate = useNavigate();
  const { user } = useAuthStore();
  const { mutate: createPayment, isPending } = useCreatePaymentMutation();

  const handleBuyNow = () => {
    if (!user) return navigate("/signin");

    createPayment(
      { packageId },
      {
        onSuccess: (res) => {
          if (res.snapToken && window.snap) {
            window.snap.pay(res.snapToken, {
              onSuccess: () => {
                toast.success("Payment successful!");
                navigate("/transactions");
              },
              onPending: () => {
                toast("Waiting for payment confirmation...");
              },
              onError: () => {
                toast.error("Payment failed.");
              },
              onClose: () => {
                navigate("/user/trans");
                toast.info("You closed the payment popup.");
              },
            });
          } else {
            toast.error("Failed to load Snap UI.");
          }
        },
        onError: () => toast.error("Failed to create transaction."),
      }
    );
  };

  return { handleBuyNow, isPending };
}
