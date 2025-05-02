// src/components/address/UpdatePackage.jsx
import React from "react";
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
  const { updatePackage } = usePackageMutation();
  const { isPending, mutateAsync } = updatePackage;

  return (
    <FormDialog
      state={pkg}
      loading={isPending}
      resourceId={pkg.id}
      action={mutateAsync}
      title="Update Package"
      schema={packageSchema}
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
      <InputNumberElement
        name="discount"
        label="Discount"
        min={0}
        max={99}
        placeholder="Discount in percent"
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

export { UpdatePackage };
