/* eslint-disable react/prop-types */
// src/pages/classes/DeleteClass.jsx
import { FormToggle } from "@/components/form/FormToggle";
import { useScheduleTemplateMutation } from "@/hooks/useSchedules";

const StopTemplate = ({ template }) => {
  const { stopTemplate, isPending } = useScheduleTemplateMutation();

  const handleStopTemplate = () => {
    stopTemplate.mutateAsync(template.id);
  };

  return (
    <FormToggle
      type="stop"
      loading={isPending}
      title="Stop Recuring Schedule"
      onToggle={handleStopTemplate}
      description="Are you sure want to Stop this recuring schedule ?"
    />
  );
};

export { StopTemplate };
