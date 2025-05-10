import { FormDelete } from "@/components/form/FormDelete";
import { useScheduleTemplateMutation } from "@/hooks/useSchedules";

const DeleteTemplate = ({ template }) => {
  const { deleteTemplate } = useScheduleTemplateMutation();

  const handleDeleteTemplate = () => {
    deleteTemplate.mutate(template.id);
  };

  return (
    <FormDelete
      icon={false}
      title="Delete Recuring Schedule"
      onDelete={handleDeleteTemplate}
      loading={deleteTemplate.isPending}
      description="Are you sure want to delete this Recuring Schedule ?"
    />
  );
};

export { DeleteTemplate };
