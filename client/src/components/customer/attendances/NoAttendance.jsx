import React from "react";
import { Link } from "react-router-dom";
import { Button } from "@/components/ui/button";

const NoAttendance = ({ type = "upcoming" }) => {
  const isUpcoming = type === "upcoming";

  return (
    <div className="text-center py-16 px-6 bg-muted/40 border border-dashed border-border rounded-xl space-y-4">
      <div className="flex justify-center">
        <img
          src={isUpcoming ? "/no-bookings.webp" : "/no-bookings.webp"}
          alt="No Classes"
          className="h-60 md:h-72 object-contain"
        />
      </div>

      <div className="mb-6 space-y-1">
        <h2 className="text-lg font-semibold text-foreground">
          {isUpcoming
            ? "You don’t have any upcoming classes"
            : "You haven’t attended any classes yet"}
        </h2>
        <p className="text-sm text-muted-foreground">
          {isUpcoming
            ? "Find and book classes to stay active 💪"
            : "Start joining classes and track your progress 📈"}
        </p>
      </div>

      {isUpcoming && (
        <Link to="/profile/bookings">
          <Button>See My Bookings</Button>
        </Link>
      )}
    </div>
  );
};

export { NoAttendance };
