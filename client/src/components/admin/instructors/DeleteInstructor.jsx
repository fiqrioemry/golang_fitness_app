/* eslint-disable react/prop-types */
import { FormDelete } from "@/components/form/FormDelete";
import { useInstructorMutation } from "@/hooks/useInstructor";

const DeleteInstructor = ({ instructor }) => {
  const { deleteInstructor } = useInstructorMutation();

  const handleDelete = () => {
    deleteInstructor.mutateAsync(instructor.id);
  };

  return (
    <FormDelete
      title="Delete Instrutors"
      onDelete={handleDelete}
      loading={deleteInstructor.isPending}
      description="Are you sure want to remove this Instructors ?"
    />
  );
};

export { DeleteInstructor };
