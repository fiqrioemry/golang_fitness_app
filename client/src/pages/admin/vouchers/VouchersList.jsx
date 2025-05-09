import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableRow,
  TableHeader,
} from "@/components/ui/table";
import VoucherAdd from "./VouchersAdd";
import React, { useEffect } from "react";
import { CirclePlus } from "lucide-react";
import { useNavigate } from "react-router-dom";
import { Button } from "@/components/ui/button";
import { Loading } from "@/components/ui/Loading";
import { useVouchersQuery } from "@/hooks/useVouchers";
import { ErrorDialog } from "@/components/ui/ErrorDialog";
import { Card, CardContent } from "@/components/ui/card";
import { VoucherUpdate } from "@/components/admin/vouchers/VoucherUpdate";
import { VoucherDelete } from "@/components/admin/vouchers/VoucherDelete";

const VouchersList = () => {
  const {
    data: vouchers = [],
    isLoading,
    isError,
    refetch,
  } = useVouchersQuery();
  const navigate = useNavigate();

  useEffect(() => {
    refetch();
  }, []);

  if (isLoading) return <Loading />;

  if (isError) return <ErrorDialog onRetry={refetch} />;

  return (
    <section className="section">
      {/* Header */}
      <div className="space-y-1 text-center">
        <h2 className="text-2xl font-bold text-foreground">
          Vouchers Management
        </h2>
        <p className="text-sm text-muted-foreground">
          View, add, and manage training packages available for purchase by
          users.
        </p>
      </div>

      {/* Add Button */}
      <div className="flex justify-end mt-4">
        <Button size="nav" onClick={() => navigate("/admin/vouchers/add")}>
          <CirclePlus className="w-4 h-4 mr-2" />
          New Voucher
        </Button>
      </div>

      <Card className="border shadow-sm">
        <CardContent className="overflow-x-auto p-0">
          <div className="hidden md:block max-w-8xl w-full">
            <Table>
              <TableHeader>
                <TableRow className="bg-muted/40">
                  <TableHead className="whitespace-nowrap">Code</TableHead>
                  <TableHead className="whitespace-nowrap">
                    Description
                  </TableHead>
                  <TableHead className="whitespace-nowrap">Type</TableHead>
                  <TableHead className="whitespace-nowrap">Discount</TableHead>
                  <TableHead className="whitespace-nowrap">
                    Max Discount
                  </TableHead>
                  <TableHead className="whitespace-nowrap">Quota</TableHead>
                  <TableHead className="whitespace-nowrap">
                    Expired At
                  </TableHead>
                  <TableHead className="text-center whitespace-nowrap">
                    Actions
                  </TableHead>
                </TableRow>
              </TableHeader>
              <TableBody>
                {vouchers.map((voucher) => (
                  <TableRow key={voucher.id} className="hover:bg-muted">
                    <TableCell className="text-sm font-medium whitespace-nowrap">
                      {voucher.code}
                    </TableCell>
                    <TableCell
                      className="text-sm text-muted-foreground max-w-xs truncate"
                      title={voucher.description}
                    >
                      {voucher.description}
                    </TableCell>
                    <TableCell className="text-sm capitalize whitespace-nowrap">
                      {voucher.discountType}
                    </TableCell>
                    <TableCell className="text-sm whitespace-nowrap">
                      {voucher.discountType === "percentage"
                        ? `${voucher.discount}%`
                        : `Rp ${voucher.discount.toLocaleString("id-ID")}`}
                    </TableCell>
                    <TableCell className="text-sm whitespace-nowrap">
                      {voucher.maxDiscount
                        ? `Rp ${voucher.maxDiscount.toLocaleString("id-ID")}`
                        : "-"}
                    </TableCell>
                    <TableCell className="text-sm whitespace-nowrap">
                      {voucher.quota}
                    </TableCell>
                    <TableCell className="text-sm whitespace-nowrap">
                      {new Date(voucher.expiredAt).toLocaleDateString("id-ID")}
                    </TableCell>
                    <TableCell className="text-center whitespace-nowrap">
                      <div className="flex justify-center gap-2">
                        <VoucherDelete voucher={voucher} />
                        <VoucherUpdate voucher={voucher} />
                      </div>
                    </TableCell>
                  </TableRow>
                ))}
              </TableBody>
            </Table>
          </div>

          {/* Mobile View */}
          <div className="md:hidden w-full space-y-4 p-4">
            {vouchers.map((voucher) => (
              <div
                key={voucher.id}
                className="border border-border rounded-lg p-4 shadow-sm bg-background"
              >
                <div className="space-y-1 mb-3">
                  <h3 className="text-base font-semibold text-foreground">
                    {voucher.code}
                  </h3>
                  <p className="text-sm text-muted-foreground">
                    {voucher.description}
                  </p>
                </div>

                <div className="text-sm text-muted-foreground space-y-1 mb-3">
                  <p>
                    <span className="font-medium text-foreground">Type:</span>{" "}
                    {voucher.discountType}
                  </p>
                  <p>
                    <span className="font-medium text-foreground">
                      Discount:
                    </span>{" "}
                    {voucher.discountType === "percentage"
                      ? `${voucher.discount}%`
                      : `Rp ${voucher.discount.toLocaleString("id-ID")}`}
                  </p>
                  <p>
                    <span className="font-medium text-foreground">
                      Max Discount:
                    </span>{" "}
                    {voucher.maxDiscount
                      ? `Rp ${voucher.maxDiscount.toLocaleString("id-ID")}`
                      : "-"}
                  </p>
                  <p>
                    <span className="font-medium text-foreground">Quota:</span>{" "}
                    {voucher.quota}
                  </p>
                  <p>
                    <span className="font-medium text-foreground">
                      Expired:
                    </span>{" "}
                    {new Date(voucher.expiredAt).toLocaleDateString("id-ID")}
                  </p>
                </div>

                <div className="flex justify-end gap-2">
                  <VoucherUpdate voucher={voucher} />
                  <VoucherAdd voucher={voucher} />
                </div>
              </div>
            ))}
          </div>
        </CardContent>
      </Card>
    </section>
  );
};

export default VouchersList;
