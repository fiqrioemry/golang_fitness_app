import React from "react";
import { PlusCircle } from "lucide-react";
import { FormDialog } from "@/components/form/FormDialog";
import { useMutationOptions } from "@/hooks/useSelectOptions";
import { InputTextElement } from "@/components/input/InputTextElement";
import { locationSchema, subcategorySchema, optionSchema } from "@/lib/schema";
import { locationState, subcategoryState, optionState } from "@/lib/constant";
import { InputTextareaElement } from "@/components/input/InputTextareaElement";
import { SelectOptionsElement } from "@/components/input/SelectOptionsElement";

const AddOptions = ({ activeTab }) => {
  const { createOptions } = useMutationOptions(activeTab);

  return (
    <FormDialog
      loading={createOptions.isPending}
      title={`Add New ${activeTab}`}
      state={
        activeTab === "location"
          ? locationState
          : activeTab === "subcategory"
          ? subcategoryState
          : optionState
      }
      schema={
        activeTab === "location"
          ? locationSchema
          : activeTab === "subcategory"
          ? subcategorySchema
          : optionSchema
      }
      action={createOptions.mutateAsync}
      buttonText={
        <button type="button" className="btn btn-primary gap-2">
          <PlusCircle />
          <span>Add New {activeTab}</span>
        </button>
      }
    >
      <InputTextElement
        name="name"
        label={`Name of ${activeTab}`}
        placeholder={`Enter ${activeTab} name`}
      />
      {activeTab === "subcategory" && (
        <div>
          <SelectOptionsElement
            name="categoryId"
            label="Category"
            placeholder="Select Category options"
          />
        </div>
      )}
      {activeTab === "location" && (
        <div>
          <InputTextareaElement
            name="address"
            label="Address"
            placeholder="Enter full address"
          />
          <InputTextElement
            name="geoLocation"
            label="Geo Location"
            placeholder="Enter geolocation coordinates"
          />
        </div>
      )}
    </FormDialog>
  );
};

export { AddOptions };
