import React from "react";
import ErrorDialog from "@/components/ui/ErrorDialog";
import FetchLoading from "@/components/ui/FetchLoading";
import { useCategoriesQuery } from "@/hooks/useCategory";

const Categories = () => {
  const {
    data: categories,
    isError,
    refetch,
    isLoading,
  } = useCategoriesQuery();

  if (isLoading) return <FetchLoading />;

  if (isError) return <ErrorDialog onRetry={refetch} />;

  console.log(categories);

  return (
    <div className="space-y-4 container mx-auto px-4">
      <div>
        <h3>GET CATEGORIES</h3>
        <div></div>
      </div>

      <div>
        <h3>GET CATEGORY BY ID</h3>
        <div></div>
      </div>

      <div>
        <h3>CREATE CATEGORY</h3>
        <div></div>
      </div>

      <div>
        <h3>DELETE CATEGORY</h3>
        <div></div>
      </div>
    </div>
  );
};

export default Categories;
