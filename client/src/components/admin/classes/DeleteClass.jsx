/* eslint-disable react/prop-types */
// src/pages/classes/DeleteClass.jsx
import { useClassMutation } from "@/hooks/useClass";
import { FormDelete } from "@/components/form/FormDelete";

const DeleteClass = ({ classes }) => {
  const { deleteClass, isLoading } = useClassMutation();

  const handleDeleteClass = () => {
    deleteClass.mutateAsync(classes.id);
  };

  return (
    <FormDelete
      loading={isLoading}
      title="Delete Class"
      onDelete={handleDeleteClass}
      description="Are you sure want to delete this classes ?"
    />
  );
};

export { DeleteClass };
