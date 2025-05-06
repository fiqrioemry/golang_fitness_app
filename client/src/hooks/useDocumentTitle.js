// src/hooks/useDocumentTitle.ts
import { useEffect } from "react";

export const useDocumentTitle = (title, siteName = "Sweat up") => {
  useEffect(() => {
    document.title = `${title} | ${siteName}`;
  }, [title, siteName]);
};
