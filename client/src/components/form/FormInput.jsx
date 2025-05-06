// src/components/form/FormInput.jsx
import React from "react";
import { SubmitButton } from "./SubmitButton";
import { FormProvider } from "react-hook-form";
import { useFormSchema } from "@/hooks/useFormSchema";

const FormInput = ({
  action,
  state,
  schema,
  text,
  className,
  isLoading,
  children,
}) => {
  const { methods, handleSubmit } = useFormSchema({ state, schema, action });

  return (
    <FormProvider {...methods}>
      <form onSubmit={handleSubmit} className="grid-cols-2 space-y-2 space-y-4">
        {children}
        <SubmitButton
          text={text}
          className={className}
          isLoading={isLoading}
          disabled={!methods.formState.isValid}
        />
      </form>
    </FormProvider>
  );
};

export { FormInput };
