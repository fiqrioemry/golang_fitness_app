import { Loader2 } from "lucide-react"; // pastikan pakai lucide-react

const Button = ({
  children,
  type = "button",
  variant = "primary",
  disabled = false,
  isLoading = false,
  className = "",
  ...rest
}) => {
  const baseStyle =
    "btn inline-flex items-center justify-center px-4 py-2 text-sm gap-2 font-medium rounded-lg transition";
  const variants = {
    primary: "btn-primary",
    secondary: "btn-secondary",
    danger: "btn-danger",
    outline: "btn-outline",
  };

  return (
    <button
      type={type}
      disabled={disabled || isLoading}
      className={`${baseStyle} ${variants[variant]} ${
        disabled || isLoading ? "opacity-50 cursor-not-allowed" : ""
      } ${className}`}
      {...rest}
    >
      {isLoading && <Loader2 className="animate-spin mr-2 h-4 w-4" />}
      {children}
    </button>
  );
};

export { Button };
