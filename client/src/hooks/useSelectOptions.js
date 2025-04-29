// src/hooks/useSelectOptions.js

import { useTypesQuery } from "@/hooks/useType";
import { useLevelsQuery } from "@/hooks/useLevel";
import { useLocationsQuery } from "@/hooks/useLocation";
import { useCategoriesQuery } from "@/hooks/useCategory";
import { useSubcategoriesQuery } from "@/hooks/useSubcategories";

export const useSelectOptions = (type) => {
  switch (type) {
    case "category":
      return useCategoriesQuery();
    case "level":
      return useLevelsQuery();
    case "location":
      return useLocationsQuery();
    case "subcategory":
      return useSubcategoriesQuery();
    case "type":
      return useTypesQuery();
    default:
      throw new Error(`Unknown select type: ${type}`);
  }
};
