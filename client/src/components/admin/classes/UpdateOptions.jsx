// src/components/address/UpdateClass.jsx
import React from "react";
import { Pencil } from "lucide-react";
import { FormDialog } from "@/components/form/FormDialog";
import { optionSchema, locationSchema } from "@/lib/schema";
import { useMutationOptions } from "@/hooks/useSelectOptions";
import { InputTextElement } from "@/components/input/InputTextElement";
import { InputTextareaElement } from "@/components/input/InputTextareaElement";

const UpdateOptions = ({ option, activeTab }) => {
  const { updateOptions } = useMutationOptions(activeTab);

  return (
    <FormDialog
      state={option}
      schema={activeTab === "location" ? locationSchema : optionSchema}
      resourceId={option.id}
      title={`Update ${activeTab}`}
      loading={updateOptions.isPending}
      action={updateOptions.mutateAsync}
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
        label="Name"
        placeholder={`Enter the name for ${activeTab}`}
      />

      {activeTab === "location" && (
        <>
          <InputTextareaElement
            name="address"
            label="Address"
            placeholder={`Enter the address for ${activeTab}`}
          />
          <InputTextElement
            name="geoLocation"
            label="Geo Coordinate"
            placeholder={`Enter the address for ${activeTab}`}
          />
        </>
      )}
    </FormDialog>
  );
};

export { UpdateOptions };
