// src/components/address/UpdateClass.jsx
import React from "react";
import { PlusCircle } from "lucide-react";
import { optionSchema } from "@/lib/schema";
import { optionState } from "@/lib/constant";
import { locationSchema } from "@/lib/schema";
import { locationState } from "@/lib/constant";
import { FormDialog } from "@/components/form/FormDialog";
import { useMutationOptions } from "@/hooks/useSelectOptions";
import { InputTextElement } from "@/components/input/InputTextElement";
import { InputTextareaElement } from "@/components/input/InputTextareaElement";

const AddOptions = ({ activeTab }) => {
  const { createOptions, isLoading } = useMutationOptions(activeTab);

  return (
    <FormDialog
      state={activeTab === "location" ? locationState : optionState}
      loading={isLoading}
      title={`Add New ${activeTab}`}
      schema={activeTab === "location" ? locationSchema : optionSchema}
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
        label={`Nama ${activeTab}`}
        placeholder={`Masukkan nama ${activeTab}`}
      />
      {activeTab === "location" && (
        <div>
          <InputTextareaElement
            name="address"
            label="Address"
            placeholder={`Masukkan alamat lengkap`}
          />
          <InputTextElement
            name="geoLocation"
            label="Geo Location"
            placeholder="Masukkan coordinate geolocation"
          />
        </div>
      )}
    </FormDialog>
  );
};

export default AddOptions;
