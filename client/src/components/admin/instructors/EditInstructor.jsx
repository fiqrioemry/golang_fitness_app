import React from "react";
import { instructorSchema } from "@/lib/schema";
import { useInstructorMutation } from "@/hooks/useInstructor";
import { FormUpdateDialog } from "@/components/form/FormUpdateDialog";
import { InputTextElement } from "@/components/input/InputTextElement";

const EditInstructor = ({ instructor }) => {
  const { updateInstructor } = useInstructorMutation();
  const { isPending, mutateAsync } = updateInstructor;

  return (
    <FormUpdateDialog
      state={instructor}
      loading={isPending}
      action={mutateAsync}
      schema={instructorSchema}
      title="Update Instructors"
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
