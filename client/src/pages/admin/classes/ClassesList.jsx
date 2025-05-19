import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableRow,
  TableHeader,
} from "@/components/ui/table";
import {
  Select,
  SelectItem,
  SelectValue,
  SelectTrigger,
  SelectContent,
} from "@/components/ui/select";
import { useState } from "react";
import { Input } from "@/components/ui/input";
import { useNavigate } from "react-router-dom";
import { Button } from "@/components/ui/button";
import { useDebounce } from "@/hooks/useDebounce";
import { useClassesQuery } from "@/hooks/useClass";
import { Loading } from "@/components/ui/Loading";
import { Pagination } from "@/components/ui/Pagination";
import { Card, CardContent } from "@/components/ui/card";
import { ErrorDialog } from "@/components/ui/ErrorDialog";
import { SectionTitle } from "@/components/header/SectionTitle";
import { DeleteClass } from "@/components/admin/classes/DeleteClass";
import { UpdateClass } from "@/components/admin/classes/UpdateClass";
import { UploadClassGallery } from "@/components/admin/classes/UploadClassGallery";

const ClassesList = () => {
  const navigate = useNavigate();
  const limit = 10;
  const [page, setPage] = useState(1);
  const [search, setSearch] = useState("");
  const [isActive, setIsActive] = useState("");
  const [sort, setSort] = useState("latest");

  const debouncedSearch = useDebounce(search, 500);
  const query = { q: debouncedSearch, isActive, sort, page, limit };

  const { data, isLoading, isError, refetch } = useClassesQuery(query);
  const classes = data?.classes || [];
  const total = data?.total || 0;

  const handleSort = (key) => {
    setSort((prev) => {
      if (key === "title") {
        return prev === "title_asc" ? "title_desc" : "title_asc";
      } else if (key === "created") {
        return prev === "oldest" ? "latest" : "oldest";
      }
      return "latest";
    });
  };

  return (
    <section className="section px-4 py-10 space-y-6 text-foreground">
      <SectionTitle
        title="Class Management"
        description="View, add, and manage training classes available for users."
      />

      <div className="flex flex-col md:flex-row md:items-center md:justify-between gap-4 mt-4">
        <div className="flex flex-col md:flex-row items-start md:items-center gap-4 w-full md:w-2/3">
          <Input
            value={search}
            onChange={(e) => setSearch(e.target.value)}
            placeholder="Search classes by title..."
            className="w-full md:w-96"
          />
        </div>
        <div className="flex items-center gap-4">
          <Select
            onValueChange={(value) => setIsActive(value === "all" ? "" : value)}
            defaultValue="all"
          >
            <SelectTrigger className="w-40">
              <SelectValue placeholder="Filter Status" />
            </SelectTrigger>
            <SelectContent>
              <SelectItem value="all">All</SelectItem>
              <SelectItem value="true">Active</SelectItem>
              <SelectItem value="false">Inactive</SelectItem>
            </SelectContent>
          </Select>

          <Button onClick={() => navigate("/admin/classes/add")}>
            Add Class
          </Button>
        </div>
      </div>

      <Card className="border shadow-sm">
        <CardContent className="overflow-x-auto p-0">
          {isLoading ? (
            <Loading />
          ) : isError ? (
            <ErrorDialog onRetry={refetch} />
          ) : classes.length === 0 ? (
            <div className="py-12 text-center text-gray-500 text-sm">
              No classes found{search && ` for “${search}”`}
            </div>
          ) : (
            <>
              {/* Desktop Table */}
              <div className="hidden md:block max-w-8xl w-full">
                <Table>
                  <TableHeader>
                    <TableRow className="bg-muted/40">
                      <TableHead>Thumbnail</TableHead>
                      <TableHead
                        onClick={() => handleSort("title")}
                        className="cursor-pointer"
                      >
                        Title
                      </TableHead>
                      <TableHead>Duration</TableHead>
                      <TableHead>Status</TableHead>
                      <TableHead
                        onClick={() => handleSort("created")}
                        className="cursor-pointer"
                      >
                        Created At
                      </TableHead>
                      <TableHead className="text-left">Actions</TableHead>
                    </TableRow>
                  </TableHeader>
                  <TableBody>
                    {classes.map((item) => (
                      <TableRow key={item.id}>
                        <TableCell>
                          <img
                            src={item.image}
                            alt={item.title}
                            className="w-14 h-14 rounded-md object-cover border"
                          />
                        </TableCell>
                        <TableCell>{item.title}</TableCell>
                        <TableCell>{item.duration} minutes</TableCell>
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
                        <TableCell>
                          {new Date(item.createdAt).toLocaleDateString()}
                        </TableCell>
                        <TableCell>
                          <div className="flex gap-2">
                            <UpdateClass classes={item} />
                            <DeleteClass classes={item} />
                            <UploadClassGallery classes={item} />
                          </div>
                        </TableCell>
                      </TableRow>
                    ))}
                  </TableBody>
                </Table>
              </div>

              {/* Mobile Card View */}
              <div className="md:hidden w-full space-y-4 p-4">
                {classes.map((item) => (
                  <div
                    key={item.id}
                    className="border rounded-xl shadow-sm overflow-hidden bg-background"
                  >
                    <img
                      src={item.image}
                      alt={item.title}
                      className="w-full h-48 object-cover"
                    />

                    {/* Konten */}
                    <div className="p-4 space-y-3">
                      <div className="flex flex-col space-y-1">
                        <h3 className="text-base font-semibold">
                          {item.title}
                        </h3>
                        <p className="text-sm text-muted-foreground">
                          {item.duration} minutes
                        </p>
                        <div className="flex items-center gap-2 flex-wrap">
                          <span
                            className={`px-2 py-0.5 rounded-full text-xs font-semibold ${
                              item.isActive
                                ? "bg-green-100 text-green-700"
                                : "bg-red-100 text-red-700"
                            }`}
                          >
                            {item.isActive ? "Active" : "Inactive"}
                          </span>
                          <p className="text-xs text-muted-foreground">
                            {new Date(item.createdAt).toLocaleDateString()}
                          </p>
                        </div>
                      </div>

                      <div className="flex justify-end gap-2 flex-wrap pt-1">
                        <UpdateClass classes={item} />
                        <DeleteClass classes={item} />
                        <UploadClassGallery classes={item} />
                      </div>
                    </div>
                  </div>
                ))}
              </div>
            </>
          )}
          {classes.length > 0 && (
            <Pagination
              page={page}
              total={total}
              limit={limit}
              onPageChange={setPage}
            />
          )}
        </CardContent>
      </Card>
    </section>
  );
};

export default ClassesList;
