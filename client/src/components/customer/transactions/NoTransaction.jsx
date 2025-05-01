import React from "react";

export const NoTransaction = () => {
  return (
    <div className="border border-dashed border-gray-300 bg-muted/40 rounded-xl p-10 text-center space-y-4">
      <div className="flex justify-center">
        <img
          src="/no-transactions.webp"
          alt="no-transactions"
          className="h-72"
        />
      </div>
      <h3 className="text-xl font-semibold text-gray-800">
        No Transactions Found
      </h3>
      <p className="text-sm text-muted-foreground max-w-md mx-auto">
        You haven't made any transactions yet. Start exploring our packages and
        enjoy various training sessions today!
      </p>
      <a
        href="/packages"
        className="inline-block px-4 py-2 text-sm bg-primary text-white rounded-md hover:bg-primary/90 transition"
      >
        Browse Packages
      </a>
    </div>
  );
};
