import React from "react";
import { Plus } from "lucide-react";
import UpdatePackage from "./UpdatePackage";
import DeletePackage from "./DeletePackage";
import { useNavigate } from "react-router-dom";
import { Loading } from "@/components/ui/Loading";
import { usePackagesQuery } from "@/hooks/usePackage";
import { ErrorDialog } from "@/components/ui/ErrorDialog";

const PackagesDisplay = () => {
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
    <section className="p-6">
      <div className="flex justify-between items-center mb-6">
        <h2 className="text-2xl font-bold">Manajemen Packages</h2>
        <button
          onClick={() => navigate("/packages/add")}
          className="flex items-center gap-2 bg-primary text-white px-4 py-2 rounded-md hover:bg-primary/90 transition"
        >
          <Plus className="w-4 h-4" />
          Tambah Package
        </button>
      </div>

      <div className="overflow-x-auto">
        <table className="min-w-full bg-white rounded-md shadow-sm text-sm">
          <thead className="bg-gray-100 text-gray-700">
            <tr>
              <th className="p-3 text-left">Thumbnail</th>
              <th className="p-3 text-left">Nama Package</th>
              <th className="p-3 text-left">Deskripsi</th>
              <th className="p-3 text-left">Harga</th>
              <th className="p-3 text-left">Credit</th>
              <th className="p-3 text-left">Status</th>
              <th className="p-3 text-center">Aksi</th>
            </tr>
          </thead>
          <tbody>
            {packages.map((pkg) => (
              <tr key={pkg.id} className="border-t hover:bg-gray-50">
                <td className="p-3">
                  <img
                    src={pkg.image}
                    alt={pkg.name}
                    className="w-16 h-16 object-cover rounded-md"
                  />
                </td>
                <td className="p-3 font-medium">{pkg.name}</td>
                <td className="p-3 max-w-xs truncate" title={pkg.description}>
                  {pkg.description}
                </td>
                <td className="p-3 font-semibold text-primary">
                  Rp {pkg.price.toLocaleString("id-ID")}
                </td>
                <td className="p-3">{pkg.credit} sesi</td>
                <td className="p-3">
                  <span
                    className={`px-2 py-1 rounded-full text-xs font-semibold ${
                      pkg.isActive
                        ? "bg-green-100 text-green-700"
                        : "bg-red-100 text-red-700"
                    }`}
                  >
                    {pkg.isActive ? "Aktif" : "Tidak Aktif"}
                  </span>
                </td>
                <td className="p-3 text-left">
                  <div className="flex gap-2">
                    <UpdatePackage pkg={pkg} />
                    <DeletePackage pkg={pkg} />
                  </div>
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </section>
  );
};

export default PackagesDisplay;
