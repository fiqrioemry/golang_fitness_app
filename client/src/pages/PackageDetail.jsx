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
import { MidtransScriptLoader } from "@/components/midtrans/MidtransScriptLoader";

const PackageDetail = () => {
  const { id } = useParams();
  const navigate = useNavigate();
  const { user } = useAuthStore();
  const { mutate: createPayment, isPending } = useCreatePaymentMutation();
  const { data: pkg, isLoading, isError, refetch } = usePackageDetailQuery(id);

  const handleBuyNow = () => {
    if (!user) return navigate("/signin");

    createPayment(
      { packageId: id },
      {
        onSuccess: (res) => {
          if (res.snapToken && window.snap) {
            window.snap.pay(res.snapToken, {
              onSuccess: function (result) {
                toast.success("Payment successful!");
                console.log("Success", result);
                navigate("/transactions");
              },
              onPending: function (result) {
                toast("Waiting for payment confirmation...");
                console.log("Pending", result);
              },
              onError: function (result) {
                toast.error("Payment failed.");
                console.error("Error", result);
              },
              onClose: function () {
                navigate("/user/trans");
                toast.info(
                  "You closed the payment popup before completing the transaction."
                );
              },
            });
          } else {
            toast.error("Failed to load Snap UI.");
          }
        },
        onError: () => {
          toast.error("Failed to create transaction.");
        },
      }
    );
  };

  if (isLoading) return <Loading />;

  if (isError) return <ErrorDialog onRetry={refetch} />;

  const discountedPrice =
    pkg.Discount && pkg.Discount > 0
      ? pkg.price * (1 - pkg.Discount / 100)
      : pkg.price;

  const tax = discountedPrice * 0.1;

  const totalPrice = discountedPrice + tax;

  return (
    <>
      <MidtransScriptLoader />
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
              <div className="mt-2 flex items-center gap-2 flex-wrap">
                <Badge variant="outline">
                  {pkg.credit} Credits •{" "}
                  {pkg.Discount > 0 ? (
                    <>
                      <span className="line-through text-red-500 ml-1">
                        Rp {pkg.price.toLocaleString("id-ID")}
                      </span>
                      <span className="ml-2 text-green-600 font-semibold">
                        Rp {discountedPrice.toLocaleString("id-ID")}
                      </span>
                    </>
                  ) : (
                    <span className="ml-1">
                      Rp {pkg.price.toLocaleString("id-ID")}
                    </span>
                  )}
                </Badge>
                {pkg.Discount > 0 && (
                  <Badge className="bg-green-600 text-white">
                    {pkg.Discount}% Discount
                  </Badge>
                )}
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
            {pkg.classes && pkg.classes.length > 0 && (
              <div className="mt-6">
                <h3 className="text-lg font-semibold mb-3">
                  🧘 Classes Included in This Package
                </h3>
                <div className="grid grid-cols-1 sm:grid-cols-2 gap-4">
                  {pkg.classes.map((cls) => (
                    <div
                      key={cls.id}
                      className="flex items-center gap-4 bg-gray-50 border rounded-xl p-3 shadow-sm"
                    >
                      <img
                        src={cls.image}
                        alt={cls.title}
                        className="w-16 h-16 rounded object-cover border"
                      />
                      <div className="flex-1">
                        <p className="font-medium text-sm">{cls.title}</p>
                        <p className="text-muted-foreground text-xs">
                          {cls.duration} minutes
                        </p>
                      </div>
                    </div>
                  ))}
                </div>
              </div>
            )}
          </div>

          <div className="bg-white border shadow-md rounded-2xl p-5 sticky top-24 space-y-4">
            <h3 className="text-xl font-semibold mb-1">Checkout</h3>

            <div className="text-sm text-muted-foreground flex justify-between">
              <span>Base Price</span>
              {pkg.Discount > 0 ? (
                <span>
                  <span className="line-through text-red-500 mr-1">
                    Rp {pkg.price.toLocaleString("id-ID")}
                  </span>
                  <span className="text-green-600 font-semibold">
                    Rp {discountedPrice.toLocaleString("id-ID")}
                  </span>
                </span>
              ) : (
                <span>Rp {pkg.price.toLocaleString("id-ID")}</span>
              )}
            </div>

            {pkg.Discount > 0 && (
              <div className="text-sm text-green-600 flex justify-between">
                <span> Discount</span>
                <span>-{pkg.Discount}%</span>
              </div>
            )}

            <div className="text-sm text-muted-foreground flex justify-between">
              <span>Tax (10%)</span>
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
    </>
  );
};

export default PackageDetail;
