// src/hooks/useFormSchema.jsx
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";

export function useFormSchema({ state, schema, action }) {
  const methods = useForm({
    resolver: zodResolver(schema),
    defaultValues: state,
    mode: "onChange",
  });

  const handleSubmit = methods.handleSubmit(async (data) => {
    const res = await action(data);

    if (res !== false) methods.reset();
  });

  return { methods, handleSubmit };
}
