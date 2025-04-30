/* eslint-disable react/prop-types */
import { FormDelete } from "@/components/form/FormDelete";
import { usePackageMutation } from "@/hooks/usePackage";

const DeletePackage = ({ pkg }) => {
  const { deletePackage, isLoading } = usePackageMutation();

  const handleDeletePackage = () => {
    deletePackage.mutateAsync(pkg.id);
  };

  return (
    <FormDelete
      loading={isLoading}
      title="Delete Package"
      onDelete={handleDeletePackage}
      description="Are you sure want to delete this package ?"
    />
  );
};

export default DeletePackage;
