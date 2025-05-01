import React from "react";

export const NoPackage = () => {
  return (
    <div className="text-center py-16 space-y-4 border border-dashed border-gray-300 rounded-xl bg-muted/50">
      <div className="flex items-center justify-center">
        <img className="h-72" src="/no-packages.webp" alt="no-packages" />
      </div>
      <p className="text-lg font-semibold text-muted-foreground">
        You don’t have any active packages yet.
      </p>
      <p className="text-sm text-muted-foreground">
        Purchase a package now to start joining your favorite classes!
      </p>
      <a
        href="/packages"
        className="inline-block px-4 py-2 bg-primary text-white rounded-md hover:bg-primary/90 transition"
      >
        View Packages
      </a>
    </div>
  );
};
