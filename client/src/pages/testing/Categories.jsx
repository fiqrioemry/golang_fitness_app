import React from "react";
import { createCategorySchema } from "@/lib/schema";
import { createCategoryState } from "@/lib/constant";
import { FormInput } from "@/components/form/FormInput";
import { useCategoryMutation } from "@/hooks/useCategory";
import { SubmitButton } from "@/components/form/SubmitButton";
import { InputElement } from "@/components/input/InputElement";
import { UploadElement } from "@/components/input/UploadElement";

const Categories = () => {
  const { mutate: createCategory, isLoading } = useCategoryMutation();

  return (
    <div className="space-y-4 container mx-auto px-4">
      <div className="max-w-sm">
        <h3 className="text-center mb-4">Create New Category</h3>
        <FormInput
          action={createCategory}
          state={createCategoryState}
          schema={createCategorySchema}
        >
          <InputElement
            name="name"
            label="Category"
            placeholder="Masukkan nama category"
          />
          <UploadElement name="image" isSingle />
          <SubmitButton className="w-full" isLoading={isLoading} />
        </FormInput>
      </div>
    </div>
  );
};

export default Categories;
