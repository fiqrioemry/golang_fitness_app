/* eslint-disable react/prop-types */
// src/pages/classes/DeleteClass.jsx
import { FormDelete } from "@/components/form/FormDelete";
import { useClassMutation } from "@/hooks/useClass";

const DeleteClass = ({ classes }) => {
  const { deleteClass, isLoading } = useClassMutation();

  return (
    <FormDelete
      loading={isLoading}
      title="Hapus Kelas"
      onClick={deleteClass.mutateAsync}
      description="Apa kamu yakin untuk menghapus kelas ini ?"
    />
  );
};

export default DeleteClass;
