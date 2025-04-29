import { SearchX } from "lucide-react";

const NoProductResult = () => {
  return (
    <div className="text-center py-16 px-4">
      <div className="flex justify-center mb-4">
        <SearchX size={48} className="text-gray-400" />
      </div>
      <h2 className="text-lg font-semibold text-gray-700">
        Produk tidak ditemukan
      </h2>
      <p className="text-sm text-muted-foreground mt-1">
        Coba periksa kata kunci pencarian atau filter yang digunakan.
      </p>
    </div>
  );
};

export default NoProductResult;
