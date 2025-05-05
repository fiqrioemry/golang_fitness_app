import {
  Dialog,
  DialogTitle,
  DialogHeader,
  DialogContent,
} from "@/components/ui/dialog";
import { Loading } from "@/components/ui/Loading";
import { Button } from "@/components/ui/button";
import React, { useState, useEffect } from "react";
import { ErrorDialog } from "@/components/ui/ErrorDialog";
import { useParams, useNavigate } from "react-router-dom";
import { useScheduleDetailQuery } from "@/hooks/useClass";
import { formatDateTime, formatRupiah } from "@/lib/utils";
import { useCreateBookingMutation } from "@/hooks/useBooking";
import { useUserClassPackagesQuery } from "@/hooks/useProfile";

const ScheduleDetail = () => {
  const { id } = useParams();
  const navigate = useNavigate();
  const [openDialog, setOpenDialog] = useState(false);
  const [selectedPackageId, setSelectedPackageId] = useState(null);

  const {
    data: schedule,
    isLoading,
    isError,
    refetch,
  } = useScheduleDetailQuery(id);

  const { data: userPackages = [], isLoading: isLoadingPackages } =
    useUserClassPackagesQuery(schedule?.classId, {
      enabled: !!schedule?.classId && openDialog,
    });

  const bookingMutation = useCreateBookingMutation();

  useEffect(() => {
    if (bookingMutation.isSuccess || schedule?.isBooked) {
      navigate("/profile/bookings");
    }
  }, [bookingMutation.isSuccess, schedule, navigate]);

  const handleBooking = () => {
    if (!selectedPackageId || !schedule?.id) return;
    bookingMutation.mutate({
      scheduleId: schedule.id,
      packageId: selectedPackageId,
    });
    setOpenDialog(false);
  };

  if (isLoading) return <Loading />;
  if (isError) return <ErrorDialog onRetry={refetch} />;

  return (
    <section className="min-h-screen px-4 py-10 max-w-3xl mx-auto space-y-6">
      <div className="text-3xl font-bold">{schedule.class}</div>
      <div className="text-muted-foreground">
        {formatDateTime(schedule.date)}
      </div>

      <div className="mt-6">
        <Button onClick={() => setOpenDialog(true)}>Book Now</Button>
      </div>

      <Dialog open={openDialog} onOpenChange={setOpenDialog}>
        <DialogContent>
          <DialogHeader>
            <DialogTitle>Choose Package</DialogTitle>
          </DialogHeader>

          {isLoadingPackages ? (
            <p>Loading packages...</p>
          ) : userPackages && userPackages.length > 0 ? (
            <div className="space-y-4">
              {userPackages.map((p) => (
                <div
                  key={p.id}
                  onClick={() => setSelectedPackageId(p.packageId)}
                  className={`border p-3 rounded cursor-pointer transition ${
                    selectedPackageId === p.packageId
                      ? "border-primary bg-muted"
                      : ""
                  }`}
                >
                  <div className="font-semibold">{p.packageName}</div>
                  <div className="text-sm text-muted-foreground">
                    Credit: {p.remainingCredit} • Expired:{" "}
                    {formatDateTime(p.expiredAt)}
                  </div>
                </div>
              ))}
              <Button
                className="mt-4 w-full"
                onClick={handleBooking}
                disabled={!selectedPackageId}
              >
                Confirm Booking
              </Button>
            </div>
          ) : (
            <div className="space-y-4">
              <p className="text-muted-foreground text-sm">
                You don’t have a package for this class. Please buy one:
              </p>
              {schedule.packages?.map((p) => (
                <div
                  key={p.id}
                  onClick={() => navigate(`/packages/${p.id}`)}
                  className="border p-3 rounded cursor-pointer hover:bg-muted transition"
                >
                  <div className="font-semibold">{p.name}</div>
                  <div className="text-sm text-muted-foreground">
                    {formatRupiah(p.price)}
                  </div>
                </div>
              ))}
            </div>
          )}
        </DialogContent>
      </Dialog>
    </section>
  );
};

export default ScheduleDetail;
