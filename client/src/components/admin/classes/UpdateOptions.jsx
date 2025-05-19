import { optionSchema, locationSchema } from "@/lib/schema";
import { useMutationOptions } from "@/hooks/useSelectOptions";
import { FormUpdateDialog } from "@/components/form/FormUpdateDialog";
import { InputTextElement } from "@/components/input/InputTextElement";
import { InputTextareaElement } from "@/components/input/InputTextareaElement";

const UpdateOptions = ({ option, activeTab }) => {
  const { updateOptions } = useMutationOptions(activeTab);
  const schema = activeTab === "location" ? locationSchema : optionSchema;

  return (
    <FormUpdateDialog
      state={option}
      schema={schema}
      title={`Update ${activeTab}`}
      loading={updateOptions.isPending}
      action={updateOptions.mutateAsync}
    >
      <InputTextElement
        name="name"
        label="Name"
        placeholder={`Enter the name for ${activeTab}`}
      />

      {activeTab === "location" && <LocationInputElement />}
    </FormUpdateDialog>
  );
};

export { UpdateOptions };

const LocationInputElement = () => {
  return (
    <>
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
    </>
  );
};
