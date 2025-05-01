import React from "react";
import { format } from "date-fns";
import { Badge } from "@/components/ui/badge";
import { Loading } from "@/components/ui/Loading";
import { CalendarCheck, Clock } from "lucide-react";

import { ErrorDialog } from "@/components/ui/ErrorDialog";
import { useUserPackagesQuery } from "../../hooks/useProfile";

const UserPackages = () => {
  const {
    data: userPackages = [],
    isError,
    refetch,
    isLoading,
  } = useUserPackagesQuery();

  if (isLoading) return <Loading />;

  if (isError) return <ErrorDialog onRetry={refetch} />;

  console.log(userPackages);

  return (
    <section className="max-w-4xl mx-auto px-4 py-8 space-y-6">
      <div className="space-y-1 text-center">
        <h2 className="text-2xl font-bold">Paket Saya</h2>
        <p className="text-muted-foreground text-sm">
          Daftar paket yang Anda miliki & sisa sesi yang tersedia
        </p>
      </div>

      {userPackages.length === 0 ? (
        <div className="text-center py-16 space-y-4 border border-dashed border-gray-300 rounded-xl bg-muted/50">
          <div className="flex items-center justify-center">
            <img className="h-72" src="/no-packages.webp" alt="no-packages" />
          </div>
          <p className="text-lg font-semibold text-muted-foreground">
            Anda belum memiliki paket latihan.
          </p>
          <p className="text-sm text-muted-foreground">
            Beli paket sekarang untuk mulai mengikuti berbagai kelas favoritmu!
          </p>
          <a
            href="/packages"
            className="inline-block px-4 py-2 bg-primary text-white rounded-md hover:bg-primary/90 transition"
          >
            Lihat Paket
          </a>
        </div>
      ) : (
        <div className="grid gap-6">
          {userPackages.map((item) => (
            <div
              key={item.id}
              className="border rounded-xl shadow-sm p-5 bg-white space-y-3"
            >
              <div className="flex justify-between items-start">
                <div className="space-y-1">
                  <h3 className="text-lg font-semibold text-gray-900">
                    {item.packageName}
                  </h3>
                  <p className="text-sm text-muted-foreground">
                    Dibeli pada{" "}
                    {format(new Date(item.purchasedAt), "dd MMM yyyy")}
                  </p>
                </div>
                <Badge
                  variant={item.remainingCredit > 0 ? "default" : "destructive"}
                >
                  {item.remainingCredit} sesi
                </Badge>
              </div>

              <div className="grid grid-cols-1 sm:grid-cols-2 gap-4 text-sm text-gray-700">
                <p className="flex items-center gap-2">
                  <CalendarCheck className="w-4 h-4" />
                  Expired: {item.expiredAt}
                </p>
                <p className="flex items-center gap-2">
                  <Clock className="w-4 h-4" />
                  Sisa waktu: {item.expiredInDays} hari
                </p>
              </div>
            </div>
          ))}
        </div>
      )}
    </section>
  );
};
// lorem
export default UserPackages;
