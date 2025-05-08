import { useScheduleMutation } from "@/hooks/useSchedules";
import { FormDelete } from "@/components/form/FormDelete";

const DeleteClassSchedule = ({ schedule, onClose }) => {
  const { deleteSchedule, isPending } = useScheduleMutation();

  const handleDeleteSchedule = async () => {
    await deleteSchedule.mutateAsync(schedule.id);
    onClose();
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
