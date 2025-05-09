import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableRow,
  TableHeader,
} from "@/components/ui/table";
import React, { useEffect } from "react";
import { CirclePlus } from "lucide-react";
import { useNavigate } from "react-router-dom";
import { Button } from "@/components/ui/button";
import { Loading } from "@/components/ui/Loading";
import { usePackagesQuery } from "@/hooks/usePackage";
import { ErrorDialog } from "@/components/ui/ErrorDialog";
import { Card, CardContent } from "@/components/ui/card";
import { PackageDelete } from "@/components/admin/packages/PackageDelete";
import { PackageUpdate } from "@/components/admin/packages/PackageUpdate";

const PackagesList = () => {
  const {
    data: packages = [],
    isLoading,
    isError,
    refetch,
  } = usePackagesQuery();
  const navigate = useNavigate();

  useEffect(() => {
    refetch();
  }, []);

  console.log(packages);
  if (isLoading) return <Loading />;
  if (isError) return <ErrorDialog onRetry={refetch} />;

  return (
    <section className="section">
      {/* Header */}
      <div className="space-y-1 text-center">
        <h2 className="text-2xl font-bold text-foreground">
          Package Management
        </h2>
        <p className="text-sm text-muted-foreground">
          View, add, and manage training packages available for purchase by
          users.
        </p>
      </div>

      {/* Add Button */}
      <div className="flex justify-end mt-4">
        <Button size="nav" onClick={() => navigate("/admin/packages/add")}>
          <CirclePlus className="w-4 h-4 mr-2" />
          Add Package
        </Button>
      </div>

      {/* Desktop Table */}
      <Card className="border shadow-sm">
        <CardContent className="overflow-x-auto p-0">
          <div className="hidden md:block max-w-8xl w-full">
            <Table>
              <TableHeader>
                <TableRow className="bg-muted/40">
                  <TableHead>Thumbnail</TableHead>
                  <TableHead>Package Name</TableHead>
                  <TableHead>Description</TableHead>
                  <TableHead>Price</TableHead>
                  <TableHead>Credit</TableHead>
                  <TableHead>Status</TableHead>
                  <TableHead className="text-center">Actions</TableHead>
                </TableRow>
              </TableHeader>
              <TableBody className="h-12">
                {packages.map((pkg) => (
                  <TableRow
                    key={pkg.id}
                    className="border-t border-border hover:bg-muted transition"
                  >
                    <TableCell>
                      <img
                        src={pkg.image}
                        alt={pkg.name}
                        className="w-14 h-14 rounded-md object-cover border"
                      />
                    </TableCell>
                    <TableCell className="font-medium text-foreground">
                      {pkg.name}
                    </TableCell>
                    <TableCell
                      className="max-w-xs truncate text-muted-foreground"
                      title={pkg.description}
                    >
                      {pkg.description}
                    </TableCell>
                    <TableCell className="text-primary font-semibold">
                      Rp {pkg.price.toLocaleString("id-ID")}
                    </TableCell>
                    <TableCell>{pkg.credit} sessions</TableCell>
                    <TableCell>
                      <span
                        className={`px-2 py-1 rounded-full text-xs font-semibold ${
                          pkg.isActive
                            ? "bg-green-100 text-green-700"
                            : "bg-red-100 text-red-700"
                        }`}
                      >
                        {pkg.isActive ? "Active" : "Inactive"}
                      </span>
                    </TableCell>
                    <TableCell className="text-center">
                      <div className="flex justify-center gap-2">
                        <PackageUpdate pkg={pkg} />
                        <PackageDelete pkg={pkg} />
                      </div>
                    </TableCell>
                  </TableRow>
                ))}
              </TableBody>
            </Table>
          </div>

          {/* Mobile View */}
          <div className="md:hidden w-full space-y-4 p-4">
            {packages.map((pkg) => (
              <div
                key={pkg.id}
                className="border border-border rounded-lg p-4 shadow-sm bg-background"
              >
                <div className="flex items-center gap-4 mb-3">
                  <img
                    src={pkg.image}
                    alt={pkg.name}
                    className="w-16 h-16 rounded-md object-cover border"
                  />
                  <div className="flex-1">
                    <h3 className="text-base text-start font-semibold text-foreground">
                      {pkg.name}
                    </h3>
                    <p className="text-xs text-start text-muted-foreground">
                      {pkg.description}
                    </p>
                  </div>
                </div>

                <div className="text-sm text-start text-muted-foreground space-y-1 mb-3">
                  <p>
                    <span className="font-medium text-foreground">Price :</span>{" "}
                    Rp {pkg.price.toLocaleString("id-ID")}
                  </p>
                  <p>
                    <span className="font-medium text-foreground">
                      Credit :
                    </span>{" "}
                    {pkg.credit} sessions
                  </p>
                  <p>
                    <span className="font-medium text-foreground">
                      Status :
                    </span>{" "}
                    <span
                      className={`ml-1 px-2 py-0.5 rounded-full text-xs font-medium ${
                        pkg.isActive
                          ? "bg-green-100 text-green-700"
                          : "bg-red-100 text-red-700"
                      }`}
                    >
                      {pkg.isActive ? "Active" : "Inactive"}
                    </span>
                  </p>
                </div>

                <div className="flex justify-end gap-2">
                  <PackageUpdate pkg={pkg} />
                  <PackageDelete pkg={pkg} />
                </div>
              </div>
            ))}
          </div>
        </CardContent>
      </Card>
    </section>
  );
};

export default PackagesList;
