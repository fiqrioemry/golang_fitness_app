import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";

export function useFormSchema({ state, schema, action }) {
  const methods = useForm({
    resolver: zodResolver(schema),
    defaultValues: state,
    mode: "onChange",
  });

  const handleSubmit = methods.handleSubmit(async (data) => {
    await action(data);
  });

  return { methods, handleSubmit };
}
