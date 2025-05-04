import React from "react";
import { useNavigate } from "react-router-dom";
import { usePackagesQuery } from "@/hooks/usePackage";
import { Loading } from "@/components/ui/Loading";
import { ErrorDialog } from "@/components/ui/ErrorDialog";
import { Button } from "@/components/ui/button";
import { CheckCircle2 } from "lucide-react";

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
      <div className="text-center mb-10">
        <h2 className="text-3xl font-bold text-gray-800">Choose a Package</h2>
        <p className="text-gray-500 mt-2 text-sm">
          Find the right plan that matches your fitness goals and schedule.
        </p>
      </div>

      <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-8">
        {packages.map((pkg) => {
          const hasDiscount = pkg.discount > 0;
          const discountedPrice = pkg.price * (1 - pkg.discount / 100);

          return (
            <div
              key={pkg.id}
              className={`relative rounded-2xl border shadow-sm overflow-hidden bg-white flex flex-col`}
            >
              {/* Banner Sold Out (tetap terang) */}
              {!pkg.isActive && (
                <div className="absolute top-4 -left-12 w-40 rotate-[-45deg] bg-red-600 text-white text-center text-xs font-bold py-1 shadow-lg z-10">
                  SOLD OUT
                </div>
              )}

              {/* Diskon */}
              {hasDiscount && (
                <span className="absolute top-4 right-4 bg-green-500 text-white text-xs px-2 py-1 rounded-full shadow z-10">
                  {pkg.discount}% OFF
                </span>
              )}

              {/* Gambar */}
              <img
                src={pkg.image}
                alt={pkg.name}
                className={`w-full h-44 object-cover ${
                  !pkg.isActive ? "opacity-60 grayscale" : ""
                }`}
              />

              {/* Konten */}
              <div
                className={`p-5 flex flex-col items-center text-center flex-grow ${
                  !pkg.isActive ? "opacity-60" : ""
                }`}
              >
                <h3 className="text-xl font-semibold line-clamp-1">
                  {pkg.name}
                </h3>
                <p className="text-sm text-gray-500 line-clamp-2 mt-1">
                  {pkg.description}
                </p>

                <div className="mt-2 flex items-center text-sm text-gray-700">
                  <p>
                    <strong>Credit:</strong> {pkg.credit}
                  </p>
                </div>

                <div className="mt-2 text-lg font-bold">
                  {hasDiscount ? (
                    <>
                      <span className="text-red-600">
                        Rp {discountedPrice.toLocaleString("id-ID")}
                      </span>
                      <span className="ml-2 text-sm text-gray-400 line-through">
                        Rp {pkg.price.toLocaleString("id-ID")}
                      </span>
                    </>
                  ) : (
                    <span>Rp {pkg.price.toLocaleString("id-ID")}</span>
                  )}
                </div>

                {/* Konten tengah diberi flex-grow agar dorong tombol ke bawah */}
                <div className="flex-grow w-full h-24 mt-4">
                  {pkg.additional?.length > 0 && (
                    <ul className="text-left text-sm text-gray-600 space-y-2 w-full">
                      {pkg.additional.slice(0, 4).map((item, index) => (
                        <li key={index} className="flex items-start gap-2">
                          <CheckCircle2 className="w-4 h-4 text-green-500 mt-1" />
                          <span>{item}</span>
                        </li>
                      ))}
                    </ul>
                  )}
                </div>

                {/* Tombol tetap di bawah */}
                <Button
                  onClick={() => navigate(`/packages/${pkg.id}`)}
                  className="mt-auto w-full"
                  disabled={!pkg.isActive}
                >
                  {pkg.isActive ? "Buy Now" : "Unavailable"}
                </Button>
              </div>
            </div>
          );
        })}
      </div>
    </section>
  );
};

export default Packages;
