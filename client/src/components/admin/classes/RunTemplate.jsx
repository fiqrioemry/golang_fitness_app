/* eslint-disable react/prop-types */
// src/pages/classes/DeleteClass.jsx
import { FormToggle } from "@/components/form/FormToggle";
import { useScheduleTemplateMutation } from "@/hooks/useSchedules";

const RunTemplate = ({ template }) => {
  const { runTemplate, isPending } = useScheduleTemplateMutation();

  const handleRunTemplate = () => {
    runTemplate.mutateAsync(template.id);
  };

  return (
    <FormToggle
      text="Start"
      loading={isPending}
      title="Run Recuring Schedule"
      onToggle={handleRunTemplate}
      description="Are you sure want to Run this recuring schedule ?"
    />
  );
};

export { RunTemplate };
