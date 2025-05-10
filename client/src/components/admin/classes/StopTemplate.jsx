import { FormToggle } from "@/components/form/FormToggle";
import { useScheduleTemplateMutation } from "@/hooks/useSchedules";

const StopTemplate = ({ template }) => {
  const { stopTemplate } = useScheduleTemplateMutation();

  const handleStopTemplate = () => {
    stopTemplate.mutate(template.id);
  };

  return (
    <FormToggle
      type="stop"
      loading={stopTemplate.isPending}
      title="Stop Recuring Schedule"
      onToggle={handleStopTemplate}
      description="Are you sure want to Stop this recuring schedule ?"
    />
  );
};

export { StopTemplate };
