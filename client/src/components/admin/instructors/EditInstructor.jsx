import React from "react";
import { instructorSchema } from "@/lib/schema";
import { FormDialog } from "@/components/form/FormDialog";
import { useInstructorMutation } from "@/hooks/useInstructor";
import { InputTextElement } from "@/components/input/InputTextElement";

const EditInstructor = ({ instructor }) => {
  const { updateInstructor } = useInstructorMutation();
  const { isPending, mutateAsync } = updateInstructor;

  return (
    <FormDialog
      state={instructor}
      loading={isPending}
      action={mutateAsync}
      schema={instructorSchema}
      resourceId={instructor.id}
      title="Update Instructors"
    >
      <InputTextElement
        name="userId"
        label="name"
        placeholder="Enter instructor name"
      />
    </FormDialog>
  );
};

export { EditInstructor };
