import { useFormContext, Controller } from "react-hook-form";
import AddCategory from "@/components/category/AddCategory";
import { useCategoriesQuery } from "@/hooks/useCategoryQuery";
import AddSubCategory from "@/components/category/AddSubCategory";

export const SelectCategoryElement = () => {
  const selectedCategoryId = watch("categoryId");
  const { control, watch, setValue } = useFormContext();

  const { data: categories = [] } = useCategoriesQuery();

  const selectedCategory = categories.find(
    (cat) => cat.ID === selectedCategoryId
  );
  const subcategories = selectedCategory?.Subcategories || [];

  const handleCategoryChange = (e, field) => {
    const newCategoryId = e.target.value;
    field.onChange(newCategoryId);
    setValue("subcategoryId", ""); // Reset subcategory kalau ganti kategori
  };

  return (
    <div className="space-y-4">
      {/* Category */}
      <div className="space-y-1">
        <label className="block text-sm font-medium">Category</label>
        <div className="flex items-center gap-2">
          <Controller
            control={control}
            name="categoryId"
            rules={{ required: true }}
            render={({ field, fieldState }) => (
              <select
                {...field}
                onChange={(e) => handleCategoryChange(e, field)}
                className="w-full border p-2 rounded disabled:bg-gray-100"
              >
                <option value="">Select Category</option>
                {categories.map((cat) => (
                  <option key={cat.ID} value={cat.ID}>
                    {cat.name}
                  </option>
                ))}
              </select>
            )}
          />
          {/* Button to Add Category */}
          <AddCategory />
        </div>
      </div>

      {/* Subcategory */}
      {selectedCategoryId && subcategories.length > 0 && (
        <div className="space-y-1">
          <label className="block text-sm font-medium">Subcategory</label>
          <div className="flex items-center gap-2">
            <Controller
              control={control}
              name="subcategoryId"
              rules={{ required: true }}
              render={({ field, fieldState }) => (
                <select
                  {...field}
                  className="w-full border p-2 rounded disabled:bg-gray-100"
                >
                  <option value="">Select Subcategory</option>
                  {subcategories.map((sub) => (
                    <option key={sub.ID} value={sub.ID}>
                      {sub.name}
                    </option>
                  ))}
                </select>
              )}
            />
            {/* Button to Add Subcategory */}
            <AddSubCategory />
          </div>
        </div>
      )}
    </div>
  );
};
