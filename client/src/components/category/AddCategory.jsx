// src/components/category/AddCategory.jsx
import React, { useState } from "react";
import { PlusCircle } from "lucide-react";
import { categorySchema } from "@/lib/schema";
import { categoryState } from "@/lib/constant";
import { FormDialog } from "@/components/form/FormDialog";
import { InputElement } from "@/components/input/InputElement";
import { UploadElement } from "@/components/input/UploadElement";
import { useCategoryMutation } from "@/hooks/useCategoryMutation";

const AddCategory = () => {
  const createCategory = useCategoryMutation();
  return (
    <FormDialog
      loading={createCategory.pending}
      state={categoryState}
      schema={categorySchema}
      title="Add NewCategory"
      action={createCategory}
      buttonText={
        <button className="btn btn-primary gap-4">
          <PlusCircle size={18} />
          <span>Category</span>
        </button>
      }
    >
      <InputElement
        name="name"
        label="Category"
        placeholder="Masukkan Nama Category"
      />
      <UploadElement
        name="image"
        isSingle
        maxImages={1}
        label="Category Image"
      />
    </FormDialog>
  );
};

export default AddCategory;
