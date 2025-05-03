import { useScheduleMutation } from "@/hooks/useClass";
import { FormDelete } from "@/components/form/FormDelete";

const DeleteClassSchedule = ({ schedule }) => {
  const { deleteSchedule, isPending } = useScheduleMutation();

  const handleDeleteSchedule = () => {
    deleteSchedule.mutateAsync(schedule.id);
  };

  return (
    <FormDelete
      icon={false}
      loading={isPending}
      title="Delete Schedule"
      onDelete={handleDeleteSchedule}
      description="Are you sure want to delete this Schedule?"
    />
  );
};

export { DeleteClassSchedule };
