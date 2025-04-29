import { formatRupiah } from "@/lib/utils";
import { useNavigate } from "react-router-dom";

const ProductCard = ({ product }) => {
  const navigate = useNavigate();
  const hasDiscount = product.discount > 0;
  const finalPrice = hasDiscount
    ? product.price * (1 - product.discount)
    : product.price;

  return (
    <div
      onClick={() => navigate(`/products/${product.slug}`)}
      className="bg-white rounded-xl border shadow-sm hover:shadow-md transition duration-300 cursor-pointer group"
      key={product.id}
    >
      <div className="relative w-full aspect-square overflow-hidden rounded-t-xl">
        {/* Gambar utama */}
        <img
          src={product.images?.[0]}
          alt={product.name}
          className="absolute inset-0 w-full h-full object-cover transition-all duration-500 ease-in-out opacity-100 group-hover:opacity-0 scale-100 group-hover:scale-110"
        />
        {/* Gambar hover */}
        <img
          src={product.images?.[1]}
          alt={product.name}
          className="absolute inset-0 w-full h-full object-cover transition-all duration-500 ease-in-out opacity-0 group-hover:opacity-100 scale-110 group-hover:scale-100"
        />
        {product.isFeatured && (
          <span className="absolute top-2 left-2 text-xs font-semibold bg-yellow-400 text-black px-2 py-1 rounded">
            Featured
          </span>
        )}
        {hasDiscount && (
          <span className="absolute top-2 right-2 text-xs font-semibold bg-red-500 text-white px-2 py-1 rounded">
            -{Math.round(product.discount * 100)}%
          </span>
        )}
      </div>

      <div className="p-4 space-y-1">
        <h3 className="text-base font-semibold text-gray-900">
          {product.name.length > 25
            ? product.name.slice(0, 25) + "..."
            : product.name}
        </h3>
        <p className="text-xs text-muted-foreground">
          {product.category?.name}
        </p>
        <div className="mt-2 text-sm">
          {hasDiscount ? (
            <div className="flex items-center gap-2">
              <span className="text-red-600 font-semibold">
                {formatRupiah(finalPrice)}
              </span>
              <span className="line-through text-gray-400 text-xs">
                {formatRupiah(product.price)}
              </span>
            </div>
          ) : (
            <span className="text-gray-900 font-semibold">
              {formatRupiah(product.price)}
            </span>
          )}
        </div>
      </div>
    </div>
  );
};

export default ProductCard;
