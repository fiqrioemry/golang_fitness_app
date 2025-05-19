import { FormDelete } from "@/components/form/FormDelete";
import { useMutationOptions } from "@/hooks/useSelectOptions";

const DeleteOptions = ({ option, activeTab }) => {
  const { deleteOption, isPending } = useMutationOptions(activeTab);

  const handleDeleteOptions = () => {
    deleteOption.mutate(option.id);
  };

  return (
    <FormDelete
      loading={isPending}
      title={`Delete ${activeTab}`}
      onDelete={handleDeleteOptions}
      description={`Are you sure want to delete this ${activeTab}?`}
    />
  );
};

export { DeleteOptions };
