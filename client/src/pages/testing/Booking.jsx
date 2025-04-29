import React from "react";
import ClassDisplay from "./ClassDisplay";
import { useLevelsQuery } from "@/hooks/useLevel";
import FormFilter from "@/components/form/FormFilter";
import ErrorDialog from "@/components/ui/ErrorDialog";
import FetchLoading from "@/components/ui/FetchLoading";
import { useCategoriesQuery } from "@/hooks/useCategory";
import { useSubcategoriesQuery } from "@/hooks/useSubcategories";
import { SelectFilterElement } from "@/components/filter/SelectFilterElement";

const Booking = () => {
  const { data: categories } = useCategoriesQuery();
  const { data: subcategories } = useSubcategoriesQuery();
  const { data: levels, isLoading, isError, refetch } = useLevelsQuery();

  if (isLoading) return <FetchLoading />;

  if (isError) return <ErrorDialog onRetry={refetch} />;

  return (
    <div className="container mx-auto px-4 space-y-6">
      <FormFilter>
        <SelectFilterElement
          name="categoryId"
          label="Category"
          options={categories}
        />
        <SelectFilterElement
          name="subcategoriesId"
          label="Subcategories"
          options={subcategories}
        />
        <SelectFilterElement name="levelId" label="Level" options={levels} />
      </FormFilter>

      <ClassDisplay />
    </div>
  );
};

export default Booking;
