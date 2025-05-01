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
      <h2 className="text-3xl font-bold text-center mb-10">
        Choose Your Package
      </h2>

      <div className="min-h-screen grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6">
        {packages.map((pkg) => (
          <div
            key={pkg.id}
            className="bg-white rounded-2xl shadow-md border overflow-hidden flex flex-col"
          >
            <img
              src={pkg.image}
              alt={pkg.name}
              className="w-full h-48 object-cover"
            />
            <div className="p-5 flex flex-col flex-grow">
              <h3 className="text-xl font-semibold mb-1">{pkg.name}</h3>
              <p className="text-sm text-muted-foreground mb-3">
                {pkg.description}
              </p>

              <div className="mb-4 space-y-1">
                <p className="text-sm">
                  <span className="font-medium">Credit:</span> {pkg.credit}
                </p>
                <p className="text-sm">
                  <span className="font-medium">Price:</span> Rp{" "}
                  {pkg.price.toLocaleString("id-ID")}
                </p>
              </div>

              <ul className="text-xs text-muted-foreground list-disc pl-4 mb-4 flex-1">
                {pkg.additional.map((item, i) => (
                  <li key={i}>{item}</li>
                ))}
              </ul>

              <button
                onClick={() => navigate(`/packages/${pkg.id}`)}
                className="mt-auto bg-blue-600 hover:bg-blue-700 text-white font-medium rounded-lg px-4 py-2 transition-colors"
              >
                Buy Now
              </button>
            </div>
          </div>
        ))}
      </div>
    </section>
  );
};

export default Packages;
