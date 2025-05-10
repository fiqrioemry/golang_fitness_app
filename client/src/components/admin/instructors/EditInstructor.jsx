import React from "react";
import { instructorSchema } from "@/lib/schema";
import { useInstructorMutation } from "@/hooks/useInstructor";
import { FormUpdateDialog } from "@/components/form/FormUpdateDialog";
import { InputTextElement } from "@/components/input/InputTextElement";

const EditInstructor = ({ instructor }) => {
  const { updateInstructor } = useInstructorMutation();

  return (
    <FormUpdateDialog
      state={instructor}
      schema={instructorSchema}
      title="Update Instructors"
      loading={updateInstructor.isPending}
      action={updateInstructor.mutateAsync}
    >
      <InputTextElement
        name="userId"
        label="name"
        placeholder="Enter instructor name"
      />
    </FormUpdateDialog>
  );
};

export { EditInstructor };
