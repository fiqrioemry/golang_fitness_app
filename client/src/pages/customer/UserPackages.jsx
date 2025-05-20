import { ErrorDialog } from "@/components/ui/ErrorDialog";
import { useUserPackagesQuery } from "@/hooks/useProfile";
import { SectionTitle } from "@/components/header/SectionTitle";
import { NoPackage } from "@/components/customer/packages/NoPackage";
import { SectionSkeleton } from "@/components/loading/SectionSkeleton";
import { PackageCard } from "@/components/customer/packages/PackageCard";

const UserPackages = () => {
  const { data, isError, refetch, isLoading } = useUserPackagesQuery();

  if (isLoading) return <SectionSkeleton />;

  if (isError) return <ErrorDialog onRetry={refetch} />;

  const pkgs = data.packages || [];

  return (
    <section className="section p-8 space-y-6">
      <SectionTitle
        title="My Packages"
        description="List of your purchased packages & remaining available sessions."
      />
      {pkgs.length === 0 ? <NoPackage /> : <PackageCard pkgs={pkgs} />}
    </section>
  );
};

export default UserPackages;
