import {
  Table,
  TableRow,
  TableCell,
  TableBody,
  TableHead,
  TableHeader,
} from "@/components/ui/Table";
import { Badge } from "@/components/ui/Badge";
import { formatDate, formatRupiah } from "@/lib/utils";
import { ChevronDown, ChevronUp } from "lucide-react";

export const TransactionCard = ({ transactions, sort, setSort }) => {
  const renderSortIcon = (field) => {
    if (sort === `${field}_asc`)
      return <ChevronUp className="w-4 h-4 inline" />;
    if (sort === `${field}_desc`)
      return <ChevronDown className="w-4 h-4 inline" />;
    return null;
  };

  return (
    <>
      <div className="hidden md:block w-full">
        <Table>
          <TableHeader>
            <TableRow>
              <TableHead
                className="cursor-pointer"
                onClick={() => setSort("name")}
              >
                Fullname {renderSortIcon("name")}
              </TableHead>
              <TableHead
                className="cursor-pointer"
                onClick={() => setSort("email")}
              >
                Email
                {renderSortIcon("email")}
              </TableHead>
              <TableHead>Package</TableHead>
              <TableHead>Price</TableHead>
              <TableHead>Method</TableHead>
              <TableHead>Status</TableHead>
              <TableHead
                className="cursor-pointer"
                onClick={() => setSort("paid_at")}
              >
                Paid At
                {renderSortIcon("paid_at")}
              </TableHead>
            </TableRow>
          </TableHeader>
          <TableBody className="h-12">
            {transactions.map((tx) => (
              <TableRow
                key={tx.id}
                className="border-t border-border hover:bg-muted transition"
              >
                <TableCell>{tx.fullname}</TableCell>
                <TableCell>{tx.userEmail}</TableCell>
                <TableCell>{tx.packageName}</TableCell>
                <TableCell>{formatRupiah(tx.price)}</TableCell>
                <TableCell>{tx.paymentMethod?.toUpperCase() || "-"}</TableCell>
                <TableCell>
                  <Badge
                    variant={
                      tx.status === "success"
                        ? "default"
                        : tx.status === "failed"
                        ? "destructive"
                        : "secondary"
                    }
                  >
                    {tx.status}
                  </Badge>
                </TableCell>
                <TableCell>
                  {" "}
                  {tx.status === "success" ? tx.paidAt : "-"}
                </TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </div>

      {/* Mobile view */}
      <div className="md:hidden space-y-4 p-4  w-full">
        {transactions.map((tx) => (
          <div
            key={tx.id}
            className="border rounded-lg p-4 shadow-sm space-y-2"
          >
            <div>
              <h3 className="text-base font-semibold">{tx.fullname}</h3>
              <p className="text-sm text-muted-foreground">{tx.userEmail}</p>
            </div>
            <div className="text-sm space-y-4">
              <div className="grid grid-cols-2 gap-4">
                <p>
                  <strong>{tx.packageName}</strong>
                </p>
                <p>{formatRupiah(tx.price)}</p>
                <p>{tx.paymentMethod?.toUpperCase() || "-"}</p>
                <p>
                  <Badge
                    variant={
                      tx.status === "success"
                        ? "default"
                        : tx.status === "failed"
                        ? "destructive"
                        : "secondary"
                    }
                  >
                    {tx.status}
                  </Badge>
                </p>
              </div>
              <div className="space-x-4">
                {" "}
                <span className="text-muted-foreground">Paid At</span>
                <span className="text-right whitespace-nowrap">
                  {tx.status === "success" ? formatDate(tx.paidAt) : "-"}
                </span>
              </div>
            </div>
          </div>
        ))}
      </div>
    </>
  );
};
