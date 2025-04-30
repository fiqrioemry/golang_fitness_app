// src/pages/admin/OptionsDisplay.jsx
import AddOptions from "./AddOptions";
import React, { useState } from "react";
import UpdateOptions from "./UpdateOptions";
import DeleteOptions from "./DeleteOptions";
import { Loading } from "@/components/ui/loading";
import { ErrorDialog } from "@/components/ui/ErrorDialog";
import { useSelectOptions } from "@/hooks/useSelectOptions";

const tabs = [
  { label: "Type", type: "type" },
  { label: "Level", type: "level" },
  { label: "Category", type: "category" },
  { label: "Subcategory", type: "subcategory" },
  { label: "Location", type: "location" },
];

const OptionsDisplay = () => {
  const [activeTab, setActiveTab] = useState(tabs[0].type);
  const {
    data: options = [],
    isLoading,
    isError,
    refetch,
  } = useSelectOptions(activeTab);

  if (isLoading) return <Loading />;
  if (isError) return <ErrorDialog onRetry={refetch} />;

  const renderMap = (geoLocation) => {
    const [lat, lng] = geoLocation.split(",");
    const mapUrl = `https://maps.google.com/maps?q=${lat},${lng}&hl=es&z=14&output=embed`;
    return (
      <div className="rounded overflow-hidden border">
        <iframe
          src={mapUrl}
          width="100%"
          height="375"
          allowFullScreen
          loading="lazy"
          className="rounded-md"
          title="location-map"
        ></iframe>
      </div>
    );
  };

  return (
    <section className="p-6 space-y-6">
      {/* Header */}
      <div className="flex justify-between items-center">
        <div>
          <h2 className="text-2xl font-bold">Class Options</h2>
          <p className="text-gray-600 text-sm">
            Kelola semua data statis seperti kategori, lokasi, tipe dan lainnya
          </p>
        </div>
      </div>

      {/* Tabs */}
      <div className="flex gap-2 border-b overflow-x-auto">
        {tabs.map((tab) => (
          <button
            key={tab.type}
            onClick={() => setActiveTab(tab.type)}
            className={`px-4 py-2 text-sm font-medium whitespace-nowrap ${
              activeTab === tab.type
                ? "border-b-2 border-primary text-primary"
                : "text-gray-500 hover:text-primary"
            } transition`}
          >
            {tab.label}
          </button>
        ))}
      </div>

      {/* Content */}
      <div className="space-y-4">
        <div className="flex justify-between items-center">
          <h3 className="text-lg font-semibold capitalize">{activeTab} List</h3>
          <AddOptions activeTab={activeTab} />
        </div>

        <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-4">
          {options.map((item) => (
            <div
              key={item.id}
              className="flex flex-col justify-between gap-2 border rounded-md px-4 py-3 shadow-sm bg-white hover:shadow transition"
            >
              <div className="flex-1 space-y-1">
                <span className="font-medium text-gray-800 block truncate">
                  {item.name}
                </span>
                {activeTab === "location" && (
                  <>
                    <p className="text-gray-600 text-sm truncate">
                      {item.address}
                    </p>
                    {item.geoLocation && renderMap(item.geoLocation)}
                  </>
                )}
              </div>
              <div className="flex justify-end gap-2">
                <DeleteOptions option={item} activeTab={activeTab} />
                <UpdateOptions option={item} activeTab={activeTab} />
              </div>
            </div>
          ))}

          {options.length === 0 && (
            <p className="text-sm text-gray-500 col-span-full">
              Data tidak ditemukan.
            </p>
          )}
        </div>
      </div>
    </section>
  );
};

export default OptionsDisplay;
