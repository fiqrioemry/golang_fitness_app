import { Loader } from "lucide-react";
import React from "react";

const Loading = () => {
  return (
    <div className="h-screen flex items-center justify-center">
      <Loader size={40} className="animate-spin" />
    </div>
  );
};

export default Loading;
