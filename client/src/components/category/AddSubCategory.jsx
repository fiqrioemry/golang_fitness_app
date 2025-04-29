// src/components/category/AddSubCategory.jsx
import React, { useState } from "react";
import { PlusCircle } from "lucide-react";
import { subCategorySchema } from "@/lib/schema";
import { subCategoryState } from "@/lib/constant";
import { FormDialog } from "@/components/form/FormDialog";
import { InputElement } from "@/components/input/InputElement";
import { UploadElement } from "@/components/input/UploadElement";
import { useCategoryMutation } from "@/hooks/useCategoryMutation";
import { useFormContext } from "react-hook-form";

const AddSubCategory = () => {
  const { watch } = useFormContext();
  const [loading, setLoading] = useState(false);
  const selectedCategoryId = watch("categoryId");
  const { createSubcategory } = useCategoryMutation();

  const handleAddSubcategory = async (data) => {
    setLoading(true);
    try {
      await createSubcategory.mutateAsync({
        categoryId: selectedCategoryId,
        data,
      });
      toast.success("Subcategory added successfully");
    } catch (error) {
      toast.error(
        error?.response?.data?.message || "Failed to add subcategory"
      );
    } finally {
      setLoading(false);
    }
  };

  return (
    <FormDialog
      loading={loading}
      title="Add Subcategory"
      state={subCategoryState}
      schema={subCategorySchema}
      action={handleAddSubcategory}
      buttonText={
        <button className="btn btn-primary gap-4">
          <PlusCircle size={18} />
          <span>Category</span>
        </button>
      }
    >
      <InputElement
        name="name"
        label="Subcategory Name"
        placeholder="Enter subcategory name"
      />
      <UploadElement name="image" label="Subategory Image" maxImages={1} />
    </FormDialog>
  );
};

export default AddSubCategory;
