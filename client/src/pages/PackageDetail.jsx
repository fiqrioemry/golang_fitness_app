import React from "react";
import { toast } from "sonner";
import { ArrowLeft } from "lucide-react";
import { Badge } from "@/components/ui/badge";
import { Button } from "@/components/ui/button";
import { Loading } from "@/components/ui/Loading";
import { useAuthStore } from "@/store/useAuthStore";
import { ErrorDialog } from "@/components/ui/ErrorDialog";
import { useParams, useNavigate } from "react-router-dom";
import { usePackageDetailQuery } from "@/hooks/usePackage";
import { useCreatePaymentMutation } from "@/hooks/usePayment";

const PackageDetail = () => {
  const { id } = useParams();
  const navigate = useNavigate();
  const { user } = useAuthStore();

  const { data: pkg, isLoading, isError, refetch } = usePackageDetailQuery(id);
  const { mutate: createPayment, isPending } = useCreatePaymentMutation();

  const handleBuyNow = () => {
    if (!user) return navigate("/signin");

    createPayment(
      { packageId: id },
      {
        onSuccess: (res) => {
          if (res.snapToken && window.snap) {
            window.snap.pay(res.snapToken, {
              onSuccess: function (result) {
                toast.success("Pembayaran berhasil!");
                console.log("Success", result);
                navigate("/transactions");
              },
              onPending: function (result) {
                toast("Menunggu pembayaran...");
                console.log("Pending", result);
              },
              onError: function (result) {
                toast.error("Pembayaran gagal.");
                console.error("Error", result);
              },
              onClose: function () {
                toast.info("Popup ditutup tanpa menyelesaikan pembayaran.");
              },
            });
          } else {
            toast.error("Gagal memuat Snap UI");
          }
        },
        onError: () => {
          toast.error("Gagal membuat transaksi.");
        },
      }
    );
  };

  if (isLoading) return <Loading />;
  if (isError) return <ErrorDialog onRetry={refetch} />;

  const tax = pkg.price * 0.05;
  const totalPrice = pkg.price + tax;

  return (
    <section className="px-4 py-10 max-w-7xl mx-auto">
      <div className="mb-6">
        <button
          onClick={() => history.back()}
          className="flex items-center text-sm text-muted-foreground hover:text-primary transition"
        >
          <ArrowLeft className="w-4 h-4 mr-1" />
          Back to Packages
        </button>
      </div>

      <div className="grid grid-cols-1 md:grid-cols-3 gap-10 items-start">
        <div className="md:col-span-2 space-y-5">
          <img
            src={pkg.image}
            alt={pkg.name}
            className="rounded-xl w-full h-[400px] object-cover border shadow-sm"
          />

          <div>
            <h2 className="text-3xl font-bold mb-1">{pkg.name}</h2>
            <p className="text-muted-foreground text-sm">{pkg.description}</p>
            <div className="mt-2">
              <Badge variant="outline">
                {pkg.credit} Credits • Rp {pkg.price.toLocaleString("id-ID")}
              </Badge>
            </div>
          </div>

          <div className="space-y-2 text-sm">
            <p>
              <span className="font-medium">Duration:</span> Valid for{" "}
              <span className="text-primary">{pkg.expired}</span> days
            </p>
            <ul className="list-disc text-muted-foreground pl-5">
              {pkg.additional?.map((item, idx) => (
                <li key={idx}>{item}</li>
              ))}
            </ul>
          </div>
        </div>

        <div className="bg-white border shadow-md rounded-2xl p-5 sticky top-24 space-y-4">
          <h3 className="text-xl font-semibold mb-1">Checkout</h3>

          <div className="text-sm text-muted-foreground flex justify-between">
            <span>Base Price</span>
            <span>Rp {pkg.price.toLocaleString("id-ID")}</span>
          </div>
          <div className="text-sm text-muted-foreground flex justify-between">
            <span>Tax (5%)</span>
            <span>Rp {tax.toLocaleString("id-ID")}</span>
          </div>

          <hr />

          <div className="text-base font-semibold flex justify-between">
            <span>Total</span>
            <span className="text-primary">
              Rp {totalPrice.toLocaleString("id-ID")}
            </span>
          </div>

          <Button
            onClick={handleBuyNow}
            disabled={isPending}
            size="lg"
            className="w-full mt-2"
          >
            {isPending ? "Processing..." : "Buy Now"}
          </Button>
        </div>
      </div>
    </section>
  );
};

export default PackageDetail;
