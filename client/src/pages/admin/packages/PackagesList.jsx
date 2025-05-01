import React from "react";
import { CirclePlus } from "lucide-react";
import { useNavigate } from "react-router-dom";
import { Button } from "@/components/ui/button";
import { Loading } from "@/components/ui/Loading";
import { usePackagesQuery } from "@/hooks/usePackage";
import { ErrorDialog } from "@/components/ui/ErrorDialog";
import DeletePackage from "@/pages/admin/packages/DeletePackage";
import UpdatePackage from "@/pages/admin/packages/UpdatePackage";

const PackagesList = () => {
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
    <section className="max-w-8xl mx-auto px-4 py-8 space-y-6">
      <div className="space-y-1 text-center">
        <h2 className="text-2xl font-bold">Package Management</h2>
        <p className="text-muted-foreground text-sm">
          View, add, and manage training packages available for purchase by
          users.
        </p>
      </div>
      <div className="flex justify-end">
        <Button onClick={() => navigate("/admin/packages/add")}>
          <CirclePlus className="w-4 h-4 mr-2" />
          Add Package
        </Button>
      </div>

      {/* Desktop Table */}
      <div className="hidden md:block overflow-x-auto border rounded-xl shadow-sm">
        <table className="min-w-full bg-white text-sm">
          <thead className="bg-muted/40 text-muted-foreground text-xs uppercase">
            <tr>
              <th className="p-3 text-left">Thumbnail</th>
              <th className="p-3 text-left">Package Name</th>
              <th className="p-3 text-left">Description</th>
              <th className="p-3 text-left">Price</th>
              <th className="p-3 text-left">Credit</th>
              <th className="p-3 text-left">Status</th>
              <th className="p-3 text-center">Actions</th>
            </tr>
          </thead>
          <tbody>
            {packages.map((pkg) => (
              <tr key={pkg.id} className="border-t hover:bg-accent/30">
                <td className="p-3">
                  <img
                    src={pkg.image}
                    alt={pkg.name}
                    className="w-14 h-14 rounded-md object-cover border"
                  />
                </td>
                <td className="p-3 font-medium">{pkg.name}</td>
                <td className="p-3 max-w-xs truncate" title={pkg.description}>
                  {pkg.description}
                </td>
                <td className="p-3 text-primary font-semibold">
                  Rp {pkg.price.toLocaleString("id-ID")}
                </td>
                <td className="p-3">{pkg.credit} sessions</td>
                <td className="p-3">
                  <span
                    className={`px-2 py-1 rounded-full text-xs font-semibold ${
                      pkg.isActive
                        ? "bg-green-100 text-green-700"
                        : "bg-red-100 text-red-700"
                    }`}
                  >
                    {pkg.isActive ? "Active" : "Inactive"}
                  </span>
                </td>
                <td className="p-3 text-center">
                  <div className="flex justify-center gap-2">
                    <UpdatePackage pkg={pkg} />
                    <DeletePackage pkg={pkg} />
                  </div>
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>

      {/* Mobile Card View */}
      <div className="md:hidden space-y-4">
        {packages.map((pkg) => (
          <div key={pkg.id} className="border rounded-lg p-4 shadow-sm">
            <div className="flex items-center gap-4 mb-3">
              <img
                src={pkg.image}
                alt={pkg.name}
                className="w-16 h-16 rounded-md object-cover border"
              />
              <div className="flex-1">
                <h3 className="text-base font-semibold">{pkg.name}</h3>
                <p className="text-xs text-muted-foreground">
                  {pkg.description}
                </p>
              </div>
            </div>

            <div className="text-sm space-y-1 mb-3">
              <p>
                <span className="font-medium">Price:</span> Rp{" "}
                {pkg.price.toLocaleString("id-ID")}
              </p>
              <p>
                <span className="font-medium">Credit:</span> {pkg.credit}{" "}
                sessions
              </p>
              <p>
                <span className="font-medium">Status:</span>{" "}
                <span
                  className={`ml-1 px-2 py-0.5 rounded-full text-xs font-medium ${
                    pkg.isActive
                      ? "bg-green-100 text-green-700"
                      : "bg-red-100 text-red-700"
                  }`}
                >
                  {pkg.isActive ? "Active" : "Inactive"}
                </span>
              </p>
            </div>

            <div className="flex justify-end gap-2">
              <UpdatePackage pkg={pkg} />
              <DeletePackage pkg={pkg} />
            </div>
          </div>
        ))}
      </div>
    </section>
  );
};

export default PackagesList;
