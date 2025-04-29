// components/productDetail/ProductInfo.jsx
import React from "react";
import { formatRupiah } from "../../lib/utils";

const ProductInfo = ({ product, selectedVariant }) => {
  return (
    <div className="space-y-4">
      <div>
        <h1 className="text-3xl font-bold tracking-tight text-foreground">
          {product.name}
        </h1>
        <p className="text-sm text-muted-foreground">
          {product.category?.name} &raquo; {product.subcategory?.name}
        </p>
      </div>

      <div className="text-primary text-2xl font-semibold">
        {formatRupiah(selectedVariant.price)}
      </div>

      <p className="text-sm text-muted-foreground leading-relaxed">
        {product.description}
      </p>
    </div>
  );
};

export default ProductInfo;
