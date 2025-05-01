// src/components/address/UpdatePackage.jsx
import React from "react";
import { Pencil } from "lucide-react";
import { packageSchema } from "@/lib/schema";
import { usePackageMutation } from "@/hooks/usePackage";
import { FormDialog } from "@/components/form/FormDialog";
import { SwitchElement } from "@/components/input/SwitchElement";
import { InputTextElement } from "@/components/input/InputTextElement";
import { InputTagsElement } from "@/components/input/InputTagsElement";
import { InputFileElement } from "@/components/input/InputFileElement";
import { InputNumberElement } from "@/components/input/InputNumberElement";
import { InputTextareaElement } from "@/components/input/InputTextareaElement";

const UpdatePackage = ({ pkg }) => {
  const { updatePackage, isLoading } = usePackageMutation();

  return (
    <FormDialog
      state={pkg}
      loading={isLoading}
      resourceId={pkg.id}
      title="Update Package"
      schema={packageSchema}
      action={updatePackage.mutateAsync}
      buttonText={
        <button
          type="button"
          className="text-primary hover:text-blue-600 transition"
        >
          <Pencil className="w-4 h-4" />
        </button>
      }
    >
      <InputTextElement
        name="name"
        label="Package Name"
        placeholder="Enter package name"
      />
      <InputTextareaElement
        maxLength={200}
        name="description"
        label="Package Description"
        placeholder="Enter package description (min. 20 characters)"
      />
      <InputNumberElement
        name="price"
        label="Package Price"
        placeholder="Price in IDR"
      />
      <InputNumberElement
        name="credit"
        label="Credit"
        placeholder="Total credits in unit"
      />
      <InputNumberElement
        name="expired"
        label="Expiration Time"
        placeholder="Package expiration duration in days"
      />
      <InputTagsElement
        name="additional"
        label="Additional Information"
        placeholder="Enter info, press enter to add"
      />
      <InputFileElement name="image" label="Thumbnail Image" isSingle />
      <SwitchElement name="isActive" label="Set as active package" />
    </FormDialog>
  );
};

export default UpdatePackage;
