import React from "react";
import { Loading } from "@/components/ui/Loading";
import { ErrorDialog } from "@/components/ui/ErrorDialog";
import { useUserPackagesQuery } from "@/hooks/useProfile";
import { NoPackage } from "@/components/customer/packages/NoPackage";
import { PackageCard } from "@/components/customer/packages/PackageCard";

const UserPackages = () => {
  const {
    data: response,
    isError,
    refetch,
    isLoading,
  } = useUserPackagesQuery();

  if (isLoading) return <Loading />;

  if (isError) return <ErrorDialog onRetry={refetch} />;

  const pkgs = response.packages || [];

  return (
    <section className="max-w-6xl mx-auto px-4 py-8 space-y-6">
      <div className="space-y-1 text-center">
        <h2 className="text-2xl font-bold">My Packages</h2>
        <p className="text-muted-foreground text-sm">
          List of your purchased packages & remaining available sessions
        </p>
      </div>
      {pkgs.length === 0 ? <NoPackage /> : <PackageCard pkgs={pkgs} />}
    </section>
  );
};

export default UserPackages;
