import {
  Table,
  TableCell,
  TableRow,
  TableBody,
  TableHead,
  TableHeader,
} from "@/components/ui/Table";
import {
  Select,
  SelectItem,
  SelectTrigger,
  SelectValue,
  SelectContent,
} from "@/components/ui/Select";
import { useState } from "react";
import { formatDateTime } from "@/lib/utils";
import { Badge } from "@/components/ui/Badge";
import { Input } from "@/components/ui/Input";
import { Button } from "@/components/ui/Button";
import { useUsersQuery } from "@/hooks/useUsers";
import { useDebounce } from "@/hooks/useDebounce";
import { Loading } from "@/components/ui/Loading";
import { Eye, ArrowDown, ArrowUp } from "lucide-react";
import { Pagination } from "@/components/ui/Pagination";
import { Card, CardContent } from "@/components/ui/Card";
import { ErrorDialog } from "@/components/ui/ErrorDialog";
import { SectionTitle } from "@/components/header/SectionTitle";
import { UserDetailDialog } from "@/components/admin/users/UserDetailDialog";

const UsersList = () => {
  const limit = 10;
  const [page, setPage] = useState(1);
  const [role, setRole] = useState("");
  const [search, setSearch] = useState("");
  const [sort, setSort] = useState("latest");

  const debouncedSearch = useDebounce(search, 500);
  const params = { q: debouncedSearch, role, sort, page, limit };
  const { data = {}, isLoading, isError, refetch } = useUsersQuery(params);
  const users = Array.isArray(data?.data) ? data.data : [];
  const total = typeof data?.total === "number" ? data.total : 0;

  const handleSort = (key) => {
    setSort((prev) => {
      if (key === "name") {
        return prev === "name_asc" ? "name_desc" : "name_asc";
      } else if (key === "created") {
        return prev === "oldest" ? "latest" : "oldest";
      }
      return "latest";
    });
  };

  return (
    <section className="section px-4 py-10 space-y-6 text-foreground">
      <SectionTitle
        title="User List"
        description="Manage all registered users, monitor their activities, and maintain
          platform integrity efficiently."
      />

      <div className="flex flex-col md:flex-row md:items-center md:justify-between gap-4 mt-4">
        <div className="flex flex-col md:flex-row items-start md:items-center gap-4 w-full md:w-2/3">
          <Input
            value={search}
            onChange={(e) => setSearch(e.target.value)}
            placeholder="Search users by name or email..."
            className="w-full md:w-80"
          />
        </div>

        <Select
          onValueChange={(v) => setRole(v === "all" ? "" : v)}
          defaultValue="all"
        >
          <SelectTrigger className="w-40">
            <SelectValue placeholder="Role" />
          </SelectTrigger>
          <SelectContent>
            <SelectItem value="all">All</SelectItem>
            <SelectItem value="admin">Admin</SelectItem>
            <SelectItem value="customer">Customer</SelectItem>
            <SelectItem value="instructor">Instructor</SelectItem>
          </SelectContent>
        </Select>
      </div>

      <Card className="border shadow-sm ">
        <CardContent className="overflow-x-auto p-0">
          {isLoading ? (
            <Loading />
          ) : isError ? (
            <ErrorDialog onRetry={refetch} />
          ) : users.length === 0 ? (
            <div className="py-12 text-center text-gray-500 text-sm">
              No users found{search && ` for “${search}”`}
            </div>
          ) : (
            <>
              {/* Desktop Table */}
              <div className="hidden md:block max-w-8xl w-full">
                <Table>
                  <TableHeader>
                    <TableRow className="bg-muted/40">
                      <TableHead>Avatar</TableHead>
                      <TableHead
                        onClick={() => handleSort("name")}
                        className="cursor-pointer"
                      >
                        <div className="flex justify-center items-center gap-2">
                          Fullname
                          {sort === "name_asc" ? (
                            <ArrowUp className="w-4 h-4 opacity-100" />
                          ) : sort === "name_desc" ? (
                            <ArrowDown className="w-4 h-4 opacity-100" />
                          ) : (
                            <ArrowUp className="w-4 h-4 opacity-30" />
                          )}
                        </div>
                      </TableHead>
                      <TableHead className="text-center">Email</TableHead>
                      <TableHead className="text-center">Role</TableHead>
                      <TableHead
                        onClick={() => handleSort("created")}
                        className="cursor-pointer select-none"
                      >
                        <div className="flex  justify-center items-center gap-2">
                          Joined Since
                          {sort === "latest" ? (
                            <ArrowDown className="w-4 h-4 opacity-100" />
                          ) : sort === "oldest" ? (
                            <ArrowUp className="w-4 h-4 opacity-100" />
                          ) : (
                            <ArrowDown className="w-4 h-4 opacity-30" />
                          )}
                        </div>
                      </TableHead>
                      <TableHead className="text-right">Details</TableHead>
                    </TableRow>
                  </TableHeader>
                  <TableBody className="h-12">
                    {users.map((user) => (
                      <TableRow
                        key={user.id}
                        className="border-t border-border hover:bg-muted transition"
                      >
                        <TableCell>
                          <img
                            src={user.avatar}
                            alt={user.fullname}
                            className="w-10 h-10 rounded-full object-cover border"
                          />
                        </TableCell>
                        <TableCell className="font-medium">
                          {user.fullname}
                        </TableCell>
                        <TableCell className="text-sm text-muted-foreground">
                          {user.email}
                        </TableCell>
                        <TableCell>
                          <Badge
                            variant={
                              user.role === "admin"
                                ? "destructive"
                                : user.role === "instructor"
                                ? "secondary"
                                : "default"
                            }
                          >
                            {user.role}
                          </Badge>
                        </TableCell>
                        <TableCell className="text-sm text-muted-foreground">
                          {formatDateTime(user.createdAt)}
                        </TableCell>
                        <TableCell className="text-right">
                          <UserDetailDialog
                            userId={user.id}
                            trigger={
                              <Button size="icon" variant="ghost">
                                <Eye className="w-5 h-5" />
                              </Button>
                            }
                          />
                        </TableCell>
                      </TableRow>
                    ))}
                  </TableBody>
                </Table>
              </div>

              {/* Mobile Cards */}
              <div className="md:hidden w-full space-y-4 p-4">
                {users.map((user) => (
                  <div
                    key={user.id}
                    className="border rounded-lg p-4 shadow-sm space-y-2"
                  >
                    <div className="flex gap-4 mb-6">
                      <img
                        src={user.avatar}
                        alt={user.fullname}
                        className="w-12 h-12 rounded-full object-cover border"
                      />
                      <div className="flex-1 text-start">
                        <h3 className="text-base font-semibold">
                          {user.fullname}
                        </h3>
                        <p className="text-sm text-muted-foreground">
                          {user.email}
                        </p>
                      </div>

                      <div>
                        <Badge
                          variant={
                            user.role === "admin"
                              ? "destructive"
                              : user.role === "instructor"
                              ? "secondary"
                              : "default"
                          }
                        >
                          {user.role}
                        </Badge>
                      </div>
                    </div>

                    <div className="flex items-center justify-between ">
                      <div className="text-xs text-start">
                        <p className="text-muted-foreground">
                          Joined: {formatDateTime(user.createdAt)}
                        </p>
                      </div>
                      <UserDetailDialog
                        userId={user.id}
                        trigger={
                          <Button size="sm" variant="outline">
                            <Eye className="w-4 h-4 mr-1" /> View
                          </Button>
                        }
                      />
                    </div>
                  </div>
                ))}
              </div>
            </>
          )}
          {users.length > 0 && (
            <Pagination
              page={page}
              limit={limit}
              total={total}
              onPageChange={(p) => setPage(p)}
            />
          )}
        </CardContent>
      </Card>
    </section>
  );
};

export default UsersList;
