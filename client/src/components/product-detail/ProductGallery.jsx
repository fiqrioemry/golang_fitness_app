// components/productDetail/ProductGallery.jsx
import React from "react";
import clsx from "clsx";

const ProductGallery = ({ images = [], selectedImage, onSelectImage }) => {
  return (
    <div>
      <div className="aspect-square overflow-hidden rounded-xl border mb-4 group">
        <img
          src={selectedImage}
          alt="Selected"
          className="w-full h-full object-cover transition-transform duration-300 group-hover:scale-110"
        />
      </div>
      <div className="flex gap-3">
        {images.map((img, idx) => (
          <button
            key={idx}
            onClick={() => onSelectImage(img)}
            className={clsx(
              "w-16 h-16 rounded overflow-hidden border transition",
              selectedImage === img
                ? "border-primary ring-2 ring-primary"
                : "border-gray-300 hover:border-primary"
            )}
          >
            <img
              src={img}
              alt={`Thumbnail ${idx}`}
              className="object-cover w-full h-full"
            />
          </button>
        ))}
      </div>
    </div>
  );
};

export default ProductGallery;
