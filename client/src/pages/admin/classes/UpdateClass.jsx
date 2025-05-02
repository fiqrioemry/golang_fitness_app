import React from "react";
import { classSchema } from "@/lib/schema";
import { useClassMutation } from "@/hooks/useClass";
import { FormDialog } from "@/components/form/FormDialog";
import { SwitchElement } from "@/components/input/SwitchElement";
import { InputTextElement } from "@/components/input/InputTextElement";
import { InputTagsElement } from "@/components/input/InputTagsElement";
import { InputFileElement } from "@/components/input/InputFileElement";
import { InputNumberElement } from "@/components/input/InputNumberElement";
import { InputTextareaElement } from "@/components/input/InputTextareaElement";
import { SelectOptionsElement } from "@/components/input/SelectOptionsElement";
import { Pencil } from "lucide-react";

const UpdateClass = ({ classes }) => {
  const { updateClass, isLoading } = useClassMutation();

  return (
    <FormDialog
      loading={isLoading}
      resourceId={classes.id}
      title="Update Class"
      state={classes}
      schema={classSchema}
      action={updateClass.mutateAsync}
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
        name="title"
        label="Class Title"
        placeholder="Enter class title"
      />
      <div className="grid grid-cols-2 gap-4">
        <InputFileElement name="image" label="Class Thumbnail" isSingle />
        <div>
          <SelectOptionsElement
            label="Location"
            data="location"
            name="locationId"
            placeholder="Select class location"
          />
          <InputNumberElement
            name="duration"
            label="Duration"
            placeholder="Class duration in minutes"
          />

          <InputTextareaElement
            maxLength={200}
            name="description"
            label="Class Description"
            placeholder="Enter class description (min. 20 characters)"
          />
        </div>
      </div>
      <div className="grid grid-cols-2 gap-4">
        <SelectOptionsElement
          name="categoryId"
          data="category"
          label="Category"
          placeholder="Select class category"
        />
        <SelectOptionsElement
          name="levelId"
          data="level"
          label="Level"
          placeholder="Select difficulty level"
        />
        <SelectOptionsElement
          name="subcategoryId"
          data="subcategory"
          label="Subcategory"
          placeholder="Select class subcategory"
        />
        <SelectOptionsElement
          data="type"
          name="typeId"
          label="Class Type"
          placeholder="Select class type"
        />
      </div>
      <InputTagsElement
        name="additional"
        label="Additional Information"
        placeholder="Enter info, press enter to add"
      />
      <InputFileElement name="images" label="Gallery (Optional)" />
      <SwitchElement name="isActive" label="Set as active class" />
    </FormDialog>
  );
};

export default UpdateClass;
