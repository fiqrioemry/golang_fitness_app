import React from "react";
import { packageSchema } from "@/lib/schema";
import { packageState } from "@/lib/constant";
import { useNavigate } from "react-router-dom";
import { usePackageMutation } from "@/hooks/usePackage";
import { FormInput } from "@/components/form/FormInput";
import { SwitchElement } from "@/components/input/SwitchElement";
import { InputFileElement } from "@/components/input/InputFileElement";
import { InputTextElement } from "@/components/input/InputTextElement";
import { InputTagsElement } from "@/components/input/InputTagsElement";
import { InputNumberElement } from "@/components/input/InputNumberElement";
import { InputTextareaElement } from "@/components/input/InputTextareaElement";

const AddPackage = () => {
  const { createPackage } = usePackageMutation();

  return (
    <section className="max-w-8xl mx-auto px-4 py-8 space-y-6">
      <div className="space-y-1 text-center">
        <h2 className="text-2xl font-bold">Add New Package</h2>
        <p className="text-muted-foreground text-sm">
          Complete the package information to offer it to users.
        </p>
      </div>

      <div className="bg-white rounded-xl shadow-sm border p-6">
        <FormInput
          className="w-72"
          state={packageState}
          schema={packageSchema}
          text={"Submit New Package"}
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

              <div className="grid grid-cols-1 sm:grid-cols-3 gap-4">
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
                  label="Expiry (days)"
                  placeholder="e.g. 60"
                />
              </div>

              <InputTagsElement
                name="additional"
                label="Additional Information"
                placeholder="e.g. Non-refundable, press Enter"
              />

              <SwitchElement
                name="isActive"
                label="Active Status"
                description="Enable to display this package on the user purchase page."
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

export default AddPackage;
