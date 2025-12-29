type SpinnerProps = {
    size?: "sm" | "md" | "lg";
    color?: string;
  };
  
  export default function Spinner({
    size = "md",
    color = "border-blue-600",
  }: SpinnerProps) {
    const sizeClass = {
      sm: "w-4 h-4 border-2",
      md: "w-6 h-6 border-3",
      lg: "w-10 h-10 border-4",
    };
  
    return (
      <div
        className={`animate-spin rounded-full border-t-transparent ${color} ${sizeClass[size]}`}
      />
    );
  }
  