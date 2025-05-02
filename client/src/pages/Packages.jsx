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
        {packages.map((pkg) => (
          <div
            key={pkg.id}
            className="relative bg-white rounded-2xl shadow-md border overflow-hidden flex flex-col transition-transform hover:scale-[1.02]"
          >
            {/* SOLD OUT Ribbon */}
            {!pkg.isActive && (
              <div className="absolute top-3 -left-12 w-40 rotate-[-45deg] bg-red-600 text-white text-center text-xs font-bold py-1 shadow-lg z-10">
                SOLD OUT
              </div>
            )}

            {/* Image */}
            <img
              src={pkg.image}
              alt={pkg.name}
              className={`w-full h-48 object-cover ${
                !pkg.isActive ? "opacity-50" : ""
              }`}
            />

            {/* Content */}
            <div className="p-5 flex flex-col flex-grow">
              <h3 className="text-lg font-bold mb-1 line-clamp-1">
                {pkg.name}
              </h3>
              <p className="text-sm text-gray-500 mb-3 line-clamp-2">
                {pkg.description}
              </p>

              <div className="grid grid-cols-2 gap-2 text-sm mb-4">
                <p>
                  <span className="font-medium text-gray-700">Credit:</span>{" "}
                  {pkg.credit}
                </p>
                <p>
                  <span className="font-medium text-gray-700">Price:</span> Rp{" "}
                  {pkg.price.toLocaleString("id-ID")}
                </p>
              </div>

              {pkg.additional.length > 0 && (
                <ul className="text-xs text-muted-foreground list-disc pl-4 mb-4 space-y-1 flex-1">
                  {pkg.additional.map((item, i) => (
                    <li key={i}>{item}</li>
                  ))}
                </ul>
              )}

              <button
                onClick={() => navigate(`/packages/${pkg.id}`)}
                className={`mt-auto font-medium rounded-lg px-4 py-2 transition-colors ${
                  pkg.isActive
                    ? "bg-blue-600 hover:bg-blue-700 text-white"
                    : "bg-gray-300 text-gray-500 cursor-not-allowed"
                }`}
                disabled={!pkg.isActive}
              >
                {pkg.isActive ? "Buy Now" : "Unavailable"}
              </button>
            </div>
          </div>
        ))}
      </div>
    </section>
  );
};

export default Packages;
