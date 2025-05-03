/* eslint-disable react/prop-types */
// src/pages/classes/DeleteClass.jsx
import { useClassMutation } from "@/hooks/useClass";
import { FormDelete } from "@/components/form/FormDelete";

const DeleteClassSchedule = ({ schedule }) => {
  const { deleteClass, isLoading } = useClassMutation();

  const handleDeleteClass = () => {
    deleteClass.mutateAsync(schedule.id);
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

export { DeleteClassSchedule };
