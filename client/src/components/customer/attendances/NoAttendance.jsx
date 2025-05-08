import React from "react";
import { Link } from "react-router-dom";
import { Button } from "@/components/ui/button";

const NoAttendance = () => {
  return (
    <div className="text-center py-16 px-6 bg-muted/50 border border-dashed border-border rounded-xl space-y-4">
      <div className="flex justify-center">
        <img
          src="/no-bookings.webp"
          alt="No Bookings"
          className="h-60 md:h-72 object-contain"
        />
      </div>
      <div className="mb-8">
        <h2 className="text-lg font-semibold text-foreground">
          You haven’t booked any classes yet.
        </h2>

        <p className="text-sm text-muted-foreground ">
          Find a class that suits your fitness goals and start sweating 🔥
        </p>
      </div>
      <Link to="/profile/bookings">
        <Button>See my booking</Button>
      </Link>
    </div>
  );
};

export { NoAttendance };
