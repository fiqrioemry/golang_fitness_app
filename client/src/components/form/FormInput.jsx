// src/components/form/FormInput.jsx
import React from "react";
import { FormProvider } from "react-hook-form";
import { useFormSchema } from "@/hooks/useFormSchema";

const FormInput = ({ action, state, schema, children }) => {
  const { methods, handleSubmit } = useFormSchema({ state, schema, action });

  return (
    <FormProvider {...methods}>
      <form onSubmit={handleSubmit} className="grid-cols-2 space-y-2">
        {typeof children === "function" ? children(methods) : children}
      </form>
    </FormProvider>
  );
};

export { FormInput };
