import React from "react";
import { useNavigate } from "react-router-dom";
import { Loading } from "@/components/ui/Loading";
import { usePackagesQuery } from "@/hooks/usePackage";
import { ErrorDialog } from "@/components/ui/ErrorDialog";

const Packages = () => {
  const {
    data: packages = [],
    isLoading,
    isError,
    refetch,
  } = usePackagesQuery();

  const navigate = useNavigate();

  if (isLoading) return <Loading />;

  if (isError) return <ErrorDialog onRetry={refetch} />;

  return (
    <section className="min-h-screen px-4 py-10 max-w-7xl mx-auto">
      <div className="bg-gradient-to-r from-sky-500 to-indigo-600 text-white rounded-xl shadow-md px-6 py-10 text-center space-y-2 mb-8">
        <div className="max-w-4xl mx-auto text-center space-y-2">
          <h3 className="text-3xl font-bold flex items-center justify-center gap-2">
            Choose Your Package
          </h3>
          <p className="text-sm text-blue-100">
            Unlock your fitness journey with the right package — flexible,
            affordable, and tailored for your lifestyle.
          </p>
        </div>
      </div>

      <div className="min-h-screen grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6">
        {packages.map((pkg) => {
          const hasDiscount = pkg.Discount > 0;
          const discountedPrice = pkg.price * (1 - pkg.Discount / 100);

          return (
            <div
              key={pkg.id}
              className="relative bg-white rounded-2xl shadow-sm border hover:shadow-md transition-all duration-200 flex flex-col"
            >
              {/* DISCOUNT BADGE */}
              {hasDiscount && (
                <div className="absolute top-3 left-3 z-10 bg-gradient-to-r from-green-500 to-lime-500 text-white text-[11px] font-semibold px-2 py-1 rounded-full shadow-md">
                  {pkg.Discount} % OFF
                </div>
              )}
              {/* SOLD OUT Ribbon */}
              {!pkg.isActive && (
                <div className="absolute top-3 -left-12 w-40 rotate-[-45deg] bg-red-600 text-white text-center text-xs font-bold py-1 shadow-lg z-10">
                  SOLD OUT
                </div>
              )}
              {/* IMAGE */}
              <div className="relative">
                <img
                  src={pkg.image}
                  alt={pkg.name}
                  className={`w-full h-48 object-cover rounded-t-2xl ${
                    !pkg.isActive ? "opacity-50 grayscale" : ""
                  }`}
                />
              </div>
              {/* CONTENT */}
              <div className="p-5 flex flex-col flex-grow space-y-2">
                <div>
                  <h3 className="text-lg font-semibold line-clamp-1">
                    {pkg.name}
                  </h3>
                  <p className="text-sm text-muted-foreground line-clamp-2">
                    {pkg.description}
                  </p>
                </div>

                <div className="flex items-center gap-3 text-sm mt-1">
                  <p>
                    <span className="text-gray-500">Credit:</span>{" "}
                    <span className="font-medium">{pkg.credit}</span>
                  </p>
                  <p>
                    <span className="text-gray-500">Duration:</span>{" "}
                    <span className="font-medium">{pkg.expired} days</span>
                  </p>
                </div>

                {/* PRICE DISPLAY */}
                <div className="mt-1">
                  {hasDiscount ? (
                    <div className="text-base font-semibold text-green-600">
                      Rp {discountedPrice.toLocaleString("id-ID")}
                      <span className="ml-2 text-sm text-gray-400 line-through font-normal">
                        Rp {pkg.price.toLocaleString("id-ID")}
                      </span>
                    </div>
                  ) : (
                    <div className="text-base font-semibold text-gray-800">
                      Rp {pkg.price.toLocaleString("id-ID")}
                    </div>
                  )}
                </div>

                {/* ADDITIONAL INFO */}
                {pkg.additional.length > 0 && (
                  <ul className="text-xs text-muted-foreground list-disc pl-4 mt-2 space-y-1 flex-1">
                    {pkg.additional.map((item, i) => (
                      <li key={i}>{item}</li>
                    ))}
                  </ul>
                )}

                {/* BUTTON */}
                <button
                  onClick={() => navigate(`/packages/${pkg.id}`)}
                  className={`mt-auto w-full text-sm font-medium rounded-lg px-4 py-2 transition-colors ${
                    pkg.isActive
                      ? "bg-primary text-white hover:bg-primary/90"
                      : "bg-gray-200 text-gray-500 cursor-not-allowed"
                  }`}
                  disabled={!pkg.isActive}
                >
                  {pkg.isActive ? "Buy Now" : "Unavailable"}
                </button>
              </div>
            </div>
          );
        })}
      </div>
    </section>
  );
};

export default Packages;
