import { useParams } from "react-router-dom";
import { usePackageDetailQuery } from "@/hooks/usePackage";
import { Loading } from "@/components/ui/Loading";
import { ErrorDialog } from "@/components/ui/ErrorDialog";
import { Button } from "@/components/ui/button";
import { Badge } from "@/components/ui/badge";
import { ArrowLeft } from "lucide-react";

const PackageDetail = () => {
  const { id } = useParams();
  const { data: pkg, isLoading, isError, refetch } = usePackageDetailQuery(id);

  if (isLoading) return <Loading />;
  if (isError) return <ErrorDialog onRetry={refetch} />;

  return (
    <section className="px-4 py-10 max-w-5xl mx-auto">
      <div className="mb-6">
        <button
          onClick={() => history.back()}
          className="flex items-center text-sm text-muted-foreground hover:text-primary transition"
        >
          <ArrowLeft className="w-4 h-4 mr-1" />
          Back to Packages
        </button>
      </div>

      <div className="grid grid-cols-1 md:grid-cols-2 gap-10 items-start">
        <div className="w-full">
          <img
            src={pkg.image}
            alt={pkg.name}
            className="rounded-xl w-full h-[400px] object-cover border shadow-sm"
          />
        </div>

        <div className="flex flex-col gap-5">
          <div>
            <h2 className="text-3xl font-bold mb-1">{pkg.name}</h2>
            <p className="text-muted-foreground text-sm">{pkg.description}</p>
            <div className="mt-2">
              <Badge variant="outline">
                {pkg.credit} Credits • Rp {pkg.price.toLocaleString("id-ID")}
              </Badge>
            </div>
          </div>

          <div className="space-y-2 text-sm">
            <p>
              <span className="font-medium">Duration:</span> Valid for{" "}
              <span className="text-primary">{pkg.expired}</span> days
            </p>
            <ul className="list-disc text-muted-foreground pl-5">
              {pkg.additional?.map((item, idx) => (
                <li key={idx}>{item}</li>
              ))}
            </ul>
          </div>

          <Button size="lg" className="mt-4 w-full md:w-fit px-10">
            Buy Now
          </Button>
        </div>
      </div>
    </section>
  );
};

export default PackageDetail;
