import { useState } from "react";
import { PlusCircle, X } from "lucide-react";
import { Controller, useFormContext } from "react-hook-form";
import { toast } from "sonner";

const UploadElement = ({
  name,
  label,
  maxImages = 5,
  maxSizeMB = 2,
  rules = { required: true },
  isSingle = false, // 🆕 mode single/multi
}) => {
  const { control } = useFormContext();
  const [isDragging, setIsDragging] = useState(false);

  const actualMaxImages = isSingle ? 1 : maxImages;

  return (
    <Controller
      name={name}
      control={control}
      rules={rules}
      render={({ field, fieldState }) => {
        const handleFiles = (files) => {
          const fileArray = Array.from(files);
          const validFiles = fileArray.filter((file) => {
            const isValidSize = file.size / (1024 * 1024) <= maxSizeMB;
            if (!isValidSize) {
              toast.warning(
                `${file.name} exceeds maximum size of ${maxSizeMB}MB`
              );
            }
            return isValidSize;
          });

          if (validFiles.length === 0) return;

          let updated = [];

          if (isSingle) {
            updated = [validFiles[0]]; // Single: replace
          } else {
            updated = [...(field.value || []), ...validFiles];
            if (updated.length > actualMaxImages) {
              toast.warning(
                `You can only upload up to ${actualMaxImages} images`
              );
              return;
            }
          }

          field.onChange(updated);
        };

        const handleDrop = (e) => {
          e.preventDefault();
          setIsDragging(false);
          if (e.dataTransfer.files && e.dataTransfer.files.length > 0) {
            handleFiles(e.dataTransfer.files);
            e.dataTransfer.clearData();
          }
        };

        const handleDragOver = (e) => {
          e.preventDefault();
          setIsDragging(true);
        };

        const handleDragLeave = (e) => {
          e.preventDefault();
          setIsDragging(false);
        };

        const handleRemoveImage = (img) => {
          const updated = (field.value || []).filter((file) => file !== img);
          field.onChange(updated);
        };

        return (
          <div className="space-y-2">
            {label && (
              <label className="block text-sm font-medium text-gray-700">
                {label}
              </label>
            )}

            {/* Upload Area */}
            <div
              onDrop={handleDrop}
              onDragOver={handleDragOver}
              onDragLeave={handleDragLeave}
              className={`${
                isSingle
                  ? "relative w-full h-64 border-2 rounded-md flex justify-center items-center overflow-hidden"
                  : "flex flex-wrap gap-4 p-4 border-2 rounded-md"
              } ${
                isDragging
                  ? "border-primary bg-primary/10"
                  : "border-dashed border-primary"
              } transition`}
            >
              {/* Single mode */}
              {isSingle ? (
                <>
                  {field.value && field.value.length > 0 ? (
                    <div className="relative w-full h-full">
                      <img
                        src={URL.createObjectURL(field.value[0])}
                        alt="preview"
                        className="object-cover w-full h-full"
                      />
                      <button
                        type="button"
                        className="absolute top-2 right-2 bg-white rounded-full p-1 shadow-md hover:bg-red-500 hover:text-white transition"
                        onClick={() => handleRemoveImage(field.value[0])}
                      >
                        <X className="w-4 h-4" />
                      </button>
                    </div>
                  ) : (
                    <>
                      <label
                        htmlFor={`${name}-upload`}
                        className="flex flex-col items-center justify-center w-full h-full cursor-pointer hover:bg-primary/10"
                      >
                        <PlusCircle className="text-primary mb-2" />
                        <span className="text-primary text-sm">
                          Select Image
                        </span>
                      </label>
                      <input
                        id={`${name}-upload`}
                        type="file"
                        accept="image/*"
                        multiple={false}
                        onChange={(e) => handleFiles(e.target.files)}
                        className="hidden"
                      />
                    </>
                  )}
                </>
              ) : (
                // Multi mode
                <>
                  {(field.value || []).map((img, idx) => (
                    <div
                      key={idx}
                      className="relative w-32 h-32 rounded-md overflow-hidden border border-border"
                    >
                      <img
                        src={URL.createObjectURL(img)}
                        alt="preview"
                        className="object-cover w-full h-full"
                      />
                      <button
                        type="button"
                        className="absolute top-1 right-1 bg-white rounded-full p-1 shadow-md hover:bg-red-500 hover:text-white transition"
                        onClick={() => handleRemoveImage(img)}
                      >
                        <X className="w-4 h-4" />
                      </button>
                    </div>
                  ))}

                  {(!field.value || field.value.length < actualMaxImages) && (
                    <>
                      <label
                        htmlFor={`${name}-upload`}
                        className="flex flex-col items-center justify-center w-32 h-32 border-2 border-dashed border-primary rounded-md cursor-pointer hover:bg-primary/10 transition"
                      >
                        <PlusCircle className="text-primary mb-2" />
                        <span className="text-primary text-sm">
                          Select Images
                        </span>
                      </label>
                      <input
                        id={`${name}-upload`}
                        type="file"
                        accept="image/*"
                        multiple={true}
                        onChange={(e) => handleFiles(e.target.files)}
                        className="hidden"
                      />
                    </>
                  )}
                </>
              )}
            </div>

            {/* Error Message */}
            {fieldState.error && (
              <p className="text-red-500 text-xs mt-1">
                {fieldState.error.message || "This field is required"}
              </p>
            )}
          </div>
        );
      }}
    />
  );
};

export { UploadElement };
