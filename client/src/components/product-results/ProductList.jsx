import { Link } from "react-router-dom";

const ProductList = ({ product }) => {
  const image =
    product.images?.[0] || "https://placehold.co/300x300?text=No+Image";
  console.log(product);
  return (
    <div className="border rounded-lg shadow-sm hover:shadow-md transition overflow-hidden mb-4">
      <div className="flex flex-col md:flex-row gap-4">
        {/* Product Image */}
        <div className="md:w-40 w-full">
          <img
            src={image}
            alt={product.name}
            className="h-full w-full object-cover md:rounded-l-lg md:rounded-none rounded-t-lg"
          />
        </div>

        {/* Product Info */}
        <div className="flex-1 px-4 py-3 flex flex-col justify-between">
          <div>
            <h3 className="text-base md:text-lg font-semibold text-gray-800">
              {product.name}
            </h3>
            <p className="text-sm text-muted-foreground mt-1">
              {product.description?.slice(0, 100)}
            </p>
            <p className="text-sm text-gray-500 mt-1">
              {product.category?.name} / {product.subcategory?.name}
            </p>
          </div>

          <div className="flex items-center justify-between mt-4">
            <p className="text-primary font-semibold text-base">
              Rp {product.price?.toLocaleString()}
            </p>
            <Link
              to={`/products/${product.slug}`}
              className="text-sm font-medium text-white bg-primary hover:bg-primary/90 px-4 py-1.5 rounded-md transition"
            >
              View Detail
            </Link>
          </div>
        </div>
      </div>
    </div>
  );
};

export default ProductList;
