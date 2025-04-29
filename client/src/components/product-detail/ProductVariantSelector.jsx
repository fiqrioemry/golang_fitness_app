import clsx from "clsx";
import { Minus, Plus } from "lucide-react";
import { Button } from "@/components/ui/button";
import { useCartMutation } from "@/hooks/useCartMutation";
import React, { useMemo, useState, useEffect } from "react";

const ProductVariantSelector = ({
  product,
  selectedVariant,
  selectedOptions,
  onOptionChange,
}) => {
  const { addToCart } = useCartMutation();
  const [quantity, setQuantity] = useState(1);

  const handleAddToCart = async () => {
    if (quantity > selectedVariant.stock) return;
    await addToCart({
      productId: product.id,
      variantId: selectedVariant.id,
      quantity,
    });
  };

  const optionKeys = useMemo(() => {
    const allOptions = product.variants.flatMap((v) =>
      Object.keys(v.options || {})
    );
    return [...new Set(allOptions)];
  }, [product]);

  const getAvailableOptionValues = (key) => {
    if (key === "clothing size" && selectedOptions?.colors) {
      return [
        ...new Set(
          product.variants
            .filter((v) => v.options?.colors === selectedOptions.colors)
            .map((v) => v.options?.[key])
        ),
      ];
    }

    return [
      ...new Set(product.variants.map((v) => v.options?.[key]).filter(Boolean)),
    ];
  };

  useEffect(() => {
    setQuantity(1);
  }, [selectedVariant]);

  return (
    <>
      {optionKeys.map((key) => {
        const values = getAvailableOptionValues(key);
        return (
          <div key={key} className="mt-4">
            <p className="text-sm font-medium text-muted-foreground mb-1">
              Pilih {key}
            </p>
            <div className="flex gap-2 flex-wrap">
              {values.map((optionValue) => {
                const isActive = selectedOptions?.[key] === optionValue;
                return (
                  <button
                    key={optionValue}
                    onClick={() => onOptionChange(key, optionValue)}
                    className={clsx(
                      "px-3 py-1 border rounded text-sm",
                      isActive
                        ? "bg-primary text-white border-primary"
                        : "bg-white text-gray-700 border-gray-300 hover:border-primary"
                    )}
                  >
                    {optionValue}
                  </button>
                );
              })}
            </div>
          </div>
        );
      })}

      {/* Quantity Selector */}
      <div className="flex items-center gap-3 mt-6">
        <p className="text-sm font-medium text-muted-foreground">Jumlah:</p>
        <div className="flex items-center border rounded-md px-2">
          <button
            onClick={() => setQuantity(Math.max(1, quantity - 1))}
            className="p-1 text-gray-600 hover:text-primary"
          >
            <Minus className="w-4 h-4" />
          </button>
          <span className="px-3 text-sm font-medium">{quantity}</span>
          <button
            onClick={() =>
              setQuantity(Math.min(quantity + 1, selectedVariant.stock))
            }
            className="p-1 text-gray-600 hover:text-primary"
          >
            <Plus className="w-4 h-4" />
          </button>
        </div>
        <span className="text-xs text-muted-foreground">
          Stok: {selectedVariant.stock}
        </span>
      </div>

      <div className="mt-4">
        <Button
          disabled={
            selectedVariant.stock === 0 || quantity > selectedVariant.stock
          }
          onClick={handleAddToCart}
          className="w-full"
        >
          Tambah ke Keranjang
        </Button>
      </div>
    </>
  );
};

export default ProductVariantSelector;
