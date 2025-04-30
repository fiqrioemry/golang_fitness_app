import React, { useState } from "react";
import { Plus, Pencil, Trash2 } from "lucide-react";
import { useSelectOptions } from "@/hooks/useSelectOptions";
import { Loading } from "@/components/ui/Loading";
import { ErrorDialog } from "@/components/ui/ErrorDialog";

const tabs = [
  { label: "Category", type: "category" },
  { label: "Subcategory", type: "subcategory" },
  { label: "Location", type: "location" },
  { label: "Type", type: "type" },
  { label: "Level", type: "level" },
];

const CategoriesDisplay = () => {
  const [activeTab, setActiveTab] = useState(tabs[0].type);

  const {
    data: options = [],
    isLoading,
    isError,
    refetch,
  } = useSelectOptions(activeTab);

  const handleCreate = () => {
    console.log("Tambah", activeTab);
    // Misal: navigate(`/admin/${activeTab}/create`)
  };

  const handleEdit = (id) => {
    console.log("Edit", activeTab, id);
    // Misal: navigate(`/admin/${activeTab}/edit/${id}`)
  };

  const handleDelete = (id) => {
    console.log("Delete", activeTab, id);
    // Konfirmasi lalu delete
  };

  if (isLoading) return <Loading />;
  if (isError) return <ErrorDialog onRetry={refetch} />;

  return (
    <section className="p-6 space-y-6">
      {/* Header */}
      <div className="flex justify-between items-center">
        <div>
          <h2 className="text-2xl font-bold">Manajemen Data Master</h2>
          <p className="text-gray-600 text-sm">
            Kelola kategori, lokasi, tipe, dan level dari 1 tempat
          </p>
        </div>
      </div>

      {/* Tabs */}
      <div className="flex gap-2 border-b overflow-x-auto">
        {tabs.map((tab) => (
          <button
            key={tab.type}
            onClick={() => setActiveTab(tab.type)}
            className={`px-4 py-2 text-sm font-medium ${
              activeTab === tab.type
                ? "border-b-2 border-primary text-primary"
                : "text-gray-500 hover:text-primary"
            } transition`}
          >
            {tab.label}
          </button>
        ))}
      </div>

      {/* Table */}
      <div className="bg-white rounded-md shadow p-6">
        <div className="flex justify-between items-center mb-4">
          <h3 className="text-lg font-semibold capitalize">{activeTab} List</h3>
          <button
            onClick={handleCreate}
            className="flex items-center gap-2 bg-primary text-white px-4 py-2 rounded-md hover:bg-primary/90 transition"
          >
            <Plus className="w-4 h-4" />
            Tambah {activeTab}
          </button>
        </div>

        <div className="overflow-x-auto">
          <table className="min-w-full text-sm">
            <thead className="bg-gray-100 text-gray-700">
              <tr>
                <th className="p-3 text-left">Nama</th>
                <th className="p-3 text-center">Aksi</th>
              </tr>
            </thead>
            <tbody>
              {options.length === 0 ? (
                <tr>
                  <td colSpan="2" className="p-6 text-center text-gray-400">
                    Tidak ada data
                  </td>
                </tr>
              ) : (
                options.map((item) => (
                  <tr key={item.id} className="border-t hover:bg-gray-50">
                    <td className="p-3">{item.name}</td>
                    <td className="p-3 flex justify-center gap-2">
                      <button
                        type="button"
                        onClick={() => handleEdit(item.id)}
                        className="text-primary hover:text-blue-600 transition"
                      >
                        <Pencil className="w-4 h-4" />
                      </button>
                      <button
                        type="button"
                        onClick={() => handleDelete(item.id)}
                        className="text-red-500 hover:text-red-700 transition"
                      >
                        <Trash2 className="w-4 h-4" />
                      </button>
                    </td>
                  </tr>
                ))
              )}
            </tbody>
          </table>
        </div>
      </div>
    </section>
  );
};

export default CategoriesDisplay;
