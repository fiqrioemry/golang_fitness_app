import React from "react";
import { packageSchema } from "@/lib/schema";
import { packageState } from "@/lib/constant";
import { usePackageMutation } from "@/hooks/usePackage";
import { FormInput } from "@/components/form/FormInput";
import { InputFileElement } from "@/components/input/InputFileElement";
import { InputTextElement } from "@/components/input/InputTextElement";
import { InputTagsElement } from "@/components/input/InputTagsElement";
import { InputNumberElement } from "@/components/input/InputNumberElement";
import { MultiSelectElement } from "@/components/input/MultiSelectElement";
import { InputTextareaElement } from "@/components/input/InputTextareaElement";

const PackageAdd = () => {
  const { createPackage } = usePackageMutation();

  return (
    <section className="section">
      <div className="space-y-1 text-center">
        <h2 className="text-2xl font-bold">Add New Package</h2>
        <p className="text-muted-foreground text-sm">
          Complete the package information to offer it to users.
        </p>
      </div>

      <div className="bg-background rounded-xl shadow-sm border p-6">
        <FormInput
          className="w-full md:w-72"
          state={packageState}
          schema={packageSchema}
          text={"Add New Package"}
          isLoading={createPackage.isPending}
          action={createPackage.mutateAsync}
        >
          {/* Main Information */}
          <div className="grid grid-cols-1 md:grid-cols-2 gap-8">
            <div className="space-y-4">
              <InputTextElement
                name="name"
                label="Package Name"
                placeholder="Enter the package name"
              />
              <InputTextareaElement
                name="description"
                label="Package Description"
                placeholder="Minimum 20 characters"
                maxLength={200}
              />

              <div className="grid grid-cols-2 sm:grid-cols-4 gap-4">
                <InputNumberElement
                  name="price"
                  label="Price (Rp)"
                  placeholder="e.g. 500000"
                />
                <InputNumberElement
                  name="credit"
                  label="Total Credit"
                  placeholder="e.g. 5"
                />
                <InputNumberElement
                  name="expired"
                  label="Expired"
                  placeholder="e.g. 60"
                />
                <InputNumberElement
                  name="discount"
                  maxLength={2}
                  label="Discount"
                  placeholder="e.g. 10 (10%)"
                />
              </div>
              <MultiSelectElement
                name="classIds"
                label="Accessible Classes"
                data="class"
              />
              <InputTagsElement
                name="additional"
                label="Additional Information"
                placeholder="e.g. Non-refundable, press Enter"
              />
            </div>

            <div>
              <InputFileElement
                isSingle
                name="image"
                label="Package Thumbnail"
                note="Recommended: 1:1 ratio (400x400px)"
              />
            </div>
          </div>
        </FormInput>
      </div>
    </section>
  );
};

export default PackageAdd;
