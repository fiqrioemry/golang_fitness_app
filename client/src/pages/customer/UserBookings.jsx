// src/pages/customer/UserBookings.jsx
import React from "react";
import { format } from "date-fns";
import { Badge } from "@/components/ui/badge";
import { Loading } from "@/components/ui/Loading";
import { Card, CardContent } from "@/components/ui/card";
import { ErrorDialog } from "@/components/ui/ErrorDialog";
import { useUserBookingsQuery } from "@/hooks/useBooking";
import { CalendarIcon, ClockIcon, MapPinIcon, UserIcon } from "lucide-react";

import { useEffect, useState } from "react";
import {
  differenceInSeconds,
  formatDuration,
  intervalToDuration,
} from "date-fns";

const BookingCard = ({ booking }) => {
  const [timeLeft, setTimeLeft] = useState("");

  useEffect(() => {
    const updateCountdown = () => {
      const seconds = differenceInSeconds(
        new Date(booking.startTime),
        new Date()
      );
      if (seconds > 0) {
        const duration = intervalToDuration({ start: 0, end: seconds * 1000 });
        const formatted = formatDuration(duration, {
          format: ["days", "hours", "minutes", "seconds"],
        });
        setTimeLeft(formatted);
      } else {
        setTimeLeft("Ongoing or passed");
      }
    };

    updateCountdown();
    const timer = setInterval(updateCountdown, 1000);
    return () => clearInterval(timer);
  }, [booking.startTime]);

  return (
    <Card className="overflow-hidden shadow-lg transition hover:shadow-xl">
      <div className="grid grid-cols-1 sm:grid-cols-[150px_1fr]">
        <img
          src={booking.classImage}
          alt={booking.classTitle}
          className="h-full w-full object-cover sm:h-full sm:w-full"
        />
        <CardContent className="p-4 space-y-2">
          <div className="flex justify-between items-start">
            <h2 className="text-xl font-semibold leading-tight">
              {booking.classTitle}
            </h2>
            <Badge variant="outline" className="capitalize">
              {booking.status}
            </Badge>
          </div>
          <p className="text-sm text-muted-foreground">
            {booking.instructorName} • {booking.duration} mins
          </p>

          {/* Countdown */}
          <div className="text-sm text-blue-500 font-medium">
            Starts in: {timeLeft} ⏳🔥
          </div>

          <div className="flex items-center gap-2 text-sm text-muted-foreground">
            <CalendarIcon className="h-4 w-4" />
            <span>
              {format(new Date(booking.startTime), "EEE, dd MMM yyyy")}
            </span>
          </div>
          <div className="flex items-center gap-2 text-sm text-muted-foreground">
            <ClockIcon className="h-4 w-4" />
            <span>
              {format(new Date(booking.startTime), "HH:mm")} -{" "}
              {format(new Date(booking.endTime), "HH:mm")}
            </span>
          </div>
          <div className="flex items-center gap-2 text-sm text-muted-foreground">
            <MapPinIcon className="h-4 w-4" />
            <span>
              {booking.locationName} - {booking.locationAddress}
            </span>
          </div>
          <div className="flex items-center gap-2 text-sm text-muted-foreground">
            <UserIcon className="h-4 w-4" />
            <span>{booking.participantCount} Participant</span>
          </div>
          <div className="text-xs text-muted-foreground text-right mt-2">
            Booked at {format(new Date(booking.bookedAt), "dd MMM yyyy, HH:mm")}
          </div>
        </CardContent>
      </div>
    </Card>
  );
};

const UserBookings = () => {
  const {
    data: bookings = [],
    isError,
    refetch,
    isLoading,
  } = useUserBookingsQuery();

  if (isLoading) return <Loading />;
  if (isError) return <ErrorDialog onRetry={refetch} />;

  return (
    <div className="p-4 space-y-4 max-w-4xl mx-auto">
      <h1 className="text-2xl font-bold mb-4">My Bookings</h1>
      {bookings.length === 0 ? (
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
      ) : (
        bookings.map((booking) => (
          <BookingCard key={booking.id} booking={booking} />
        ))
      )}
    </div>
  );
};

export default UserBookings;
