// src/components/form/SubmitButton.jsx
import { Loader2 } from "lucide-react";

const SubmitButton = ({
  text = "Submit",
  isLoading = false,
  disabled = false,
  className = "",
  ...props
}) => {
  const baseClass =
    "inline-flex items-center justify-center px-4 py-2 text-sm font-medium rounded-md bg-primary text-white hover:bg-primary/90 transition disabled:opacity-50 disabled:cursor-not-allowed";

  return (
    <button
      type="submit"
      disabled={disabled || isLoading}
      className={`${baseClass} ${className}`}
      {...props}
    >
      {isLoading && <Loader2 className="animate-spin h-4 w-4 mr-2" />}
      {text}
    </button>
  );
};

export { SubmitButton };
