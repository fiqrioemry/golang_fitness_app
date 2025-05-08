import React from "react";
import { Button } from "@/components/ui/button";
import { Link } from "react-router-dom";

export const NoBooking = () => {
  return (
    <div className="text-center py-16 px-6 bg-muted/50 border border-dashed border-border rounded-xl space-y-4">
      <div className="flex justify-center">
        <img
          src="/no-bookings.webp"
          alt="No Bookings"
          className="h-60 md:h-72 object-contain"
        />
      </div>

      <h2 className="text-lg font-semibold text-foreground">
        You haven’t booked any classes yet.
      </h2>

      <p className="text-sm text-muted-foreground">
        Find a class that suits your fitness goals and start sweating 🔥
      </p>
      <Link to="/classes">
        <Button>Explore Classes</Button>
      </Link>
    </div>
  );
};
