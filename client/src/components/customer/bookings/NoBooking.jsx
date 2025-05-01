import React from "react";

export const NoBooking = () => {
  return (
    <div className="text-center py-16 space-y-4 border border-dashed border-gray-300 rounded-xl bg-muted/50">
      <div className="flex items-center justify-center">
        <img className="h-72" src="/no-bookings.webp" alt="no-bookings" />
      </div>
      <p className="text-lg font-medium text-muted-foreground">
        You haven’t booked any classes yet.
      </p>
      <p className="text-sm text-muted-foreground">
        Find a class that suits your fitness goals and start sweating 🔥
      </p>
      <a
        href="/classes"
        className="inline-block px-4 py-2 bg-primary text-white rounded-md hover:bg-primary/90 transition"
      >
        Explore Classes
      </a>
    </div>
  );
};
