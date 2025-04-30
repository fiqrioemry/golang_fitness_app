/* eslint-disable react/prop-types */
import { FormDelete } from "@/components/form/FormDelete";
import { useMutationOptions } from "@/hooks/useSelectOptions";

const DeleteOptions = ({ option, activeTab }) => {
  const { deleteOptions, isLoading } = useMutationOptions(activeTab);

  const handleDeleteOptions = () => {
    deleteOptions.mutateAsync(option.id);
  };

  return (
    <FormDelete
      loading={isLoading}
      title={`Delete ${activeTab}`}
      onDelete={handleDeleteOptions}
      description={`Are you sure want to delete this ${activeTab}?`}
    />
  );
};

export default DeleteOptions;
