import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableRow,
  TableHeader,
} from "@/components/ui/table";
import React, { useEffect } from "react";
import { PlusCircle } from "lucide-react";
import { useNavigate } from "react-router-dom";
import { Button } from "@/components/ui/button";
import { Loading } from "@/components/ui/Loading";
import { useClassesQuery } from "@/hooks/useClass";
import { ErrorDialog } from "@/components/ui/ErrorDialog";
import { Card, CardContent } from "@/components/ui/card";
import { DeleteClass } from "@/components/admin/classes/DeleteClass";
import { UpdateClass } from "@/components/admin/classes/UpdateClass";

const ClassesList = () => {
  const navigate = useNavigate();
  const { data, isLoading, isError, refetch } = useClassesQuery();

  useEffect(() => {
    refetch();
  }, []);

  if (isLoading) return <Loading />;
  if (isError) return <ErrorDialog onRetry={refetch} />;

  const classes = data?.classes || [];

  return (
    <section className="section">
      {/* Header */}
      <div className="space-y-1 text-center">
        <h2 className="text-2xl font-bold text-foreground">Class Management</h2>
        <p className="text-sm text-muted-foreground">
          View, add, and manage training classes available for users.
        </p>
      </div>

      {/* Add Button */}
      <div className="flex justify-end mt-4">
        <Button size="nav" onClick={() => navigate("/admin/classes/add")}>
          <PlusCircle className="w-4 h-4 mr-2" />
          Add Class
        </Button>
      </div>

      {/* Desktop Table */}
      <Card className="border shadow-sm mt-6">
        <CardContent className="overflow-x-auto p-0">
          <div className="hidden md:block max-w-8xl w-full">
            <Table>
              <TableHeader>
                <TableRow className="bg-muted/40">
                  <TableHead>Thumbnail</TableHead>
                  <TableHead>Title</TableHead>
                  <TableHead>Duration</TableHead>
                  <TableHead>Status</TableHead>
                  <TableHead>Created At</TableHead>
                  <TableHead className="text-left">Actions</TableHead>
                </TableRow>
              </TableHeader>
              <TableBody>
                {classes.map((item) => (
                  <TableRow
                    key={item.id}
                    className="border-t border-border hover:bg-muted transition"
                  >
                    <TableCell>
                      <img
                        src={item.image}
                        alt={item.title}
                        className="w-14 h-14 rounded-md object-cover border"
                      />
                    </TableCell>
                    <TableCell className="font-medium text-foreground">
                      {item.title}
                    </TableCell>
                    <TableCell className="text-muted-foreground">
                      {item.duration} minutes
                    </TableCell>
                    <TableCell>
                      <span
                        className={`px-2 py-1 rounded-full text-xs font-semibold ${
                          item.isActive
                            ? "bg-green-100 text-green-700"
                            : "bg-red-100 text-red-700"
                        }`}
                      >
                        {item.isActive ? "Active" : "Inactive"}
                      </span>
                    </TableCell>
                    <TableCell className="text-muted-foreground">
                      {new Date(item.createdAt).toLocaleDateString()}
                    </TableCell>
                    <TableCell>
                      <div className="flex gap-2">
                        <UpdateClass classes={item} />
                        <DeleteClass classes={item} />
                      </div>
                    </TableCell>
                  </TableRow>
                ))}
              </TableBody>
            </Table>
          </div>

          {/* Mobile View */}
          <div className="md:hidden w-full space-y-4 p-4">
            {classes.map((item) => (
              <div
                key={item.id}
                className="border border-border rounded-lg p-4 shadow-sm bg-background"
              >
                <div className="flex items-center gap-4 mb-3">
                  <img
                    src={item.image}
                    alt={item.title}
                    className="w-16 h-16 rounded-md object-cover border"
                  />
                  <div className="flex-1">
                    <h3 className="text-base font-semibold text-foreground">
                      {item.title}
                    </h3>
                    <p className="text-xs text-muted-foreground">
                      {item.duration} minutes
                    </p>
                  </div>
                </div>

                <div className="text-sm text-muted-foreground space-y-1 mb-3">
                  <p>
                    <span className="font-medium text-foreground">
                      Status :
                    </span>{" "}
                    <span
                      className={`ml-1 px-2 py-0.5 rounded-full text-xs font-medium ${
                        item.isActive
                          ? "bg-green-100 text-green-700"
                          : "bg-red-100 text-red-700"
                      }`}
                    >
                      {item.isActive ? "Active" : "Inactive"}
                    </span>
                  </p>
                  <p>
                    <span className="font-medium text-foreground">
                      Created :
                    </span>{" "}
                    {new Date(item.createdAt).toLocaleDateString()}
                  </p>
                </div>

                <div className="flex justify-end gap-2">
                  <UpdateClass classes={item} />
                  <DeleteClass classes={item} />
                </div>
              </div>
            ))}
          </div>
        </CardContent>
      </Card>
    </section>
  );
};

export default ClassesList;
