// src/components/address/UpdateClass.jsx
import React from "react";
import { Pencil } from "lucide-react";
import { optionSchema } from "@/lib/schema";
import { FormDialog } from "@/components/form/FormDialog";
import { useMutationOptions } from "@/hooks/useSelectOptions";
import { InputTextElement } from "@/components/input/InputTextElement";

const UpdateOptions = ({ option, activeTab }) => {
  const { updateOptions, isLoading } = useMutationOptions(activeTab);

  return (
    <FormDialog
      state={option}
      loading={isLoading}
      resourceId={option.id}
      title={`Update ${activeTab}`}
      schema={optionSchema}
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
        label={`Nama ${activeTab}`}
        placeholder={`Masukkan nama ${activeTab}`}
      />
    </FormDialog>
  );
};

export default UpdateOptions;
