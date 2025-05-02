/* eslint-disable react/prop-types */
import { usePackageMutation } from "@/hooks/usePackage";
import { FormDelete } from "@/components/form/FormDelete";

const DeletePackage = ({ pkg }) => {
  const { deletePackage } = usePackageMutation();
  const { mutateAsync, isPending } = deletePackage;

  const handleDeletePackage = () => {
    mutateAsync(pkg.id);
  };

  return (
    <FormDelete
      loading={isPending}
      title="Delete Package"
      onDelete={handleDeletePackage}
      description="Are you sure want to delete this package ?"
    />
  );
};

export { DeletePackage };
