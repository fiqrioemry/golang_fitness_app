import React from "react";
import { useSearchParams } from "react-router-dom";
import { useClassesQuery } from "@/hooks/useClass";
import FetchLoading from "@/components/ui/FetchLoading";
import ErrorDialog from "@/components/ui/ErrorDialog";

const ClassDisplay = () => {
  const [searchParams] = useSearchParams();
  const filters = Object.fromEntries([...searchParams]);

  const {
    data: classes,
    isLoading,
    isError,
    refetch,
  } = useClassesQuery(filters);

  if (isLoading) {
    return (
      <div className="grid gap-4">
        {Array.from({ length: 6 }).map((_, idx) => (
          <div key={idx} className="p-4 border rounded animate-pulse">
            <div className="h-6 bg-gray-300 rounded w-2/3 mb-2"></div>
            <div className="h-4 bg-gray-300 rounded w-1/2"></div>
          </div>
        ))}
      </div>
    );
  }

  if (isError) return <ErrorDialog onRetry={refetch} />;

  return (
    <div className="grid gap-4">
      {classes?.data?.length === 0 && <p>No classes found.</p>}
      {classes?.data?.map((cls) => (
        <div key={cls.id} className="p-4 border rounded">
          <h3 className="font-bold">{cls.title}</h3>
          <p>{cls.description}</p>
        </div>
      ))}
    </div>
  );
};

export default ClassDisplay;
