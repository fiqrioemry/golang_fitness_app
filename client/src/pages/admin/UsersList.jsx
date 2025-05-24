import { useUsersQuery } from "@/hooks/useUsers";
import { Loading } from "@/components/ui/Loading";
import { useDebounce } from "@/hooks/useDebounce";
import { useQueryStore } from "@/store/useQueryStore";
import { Pagination } from "@/components/ui/Pagination";
import { Card, CardContent } from "@/components/ui/Card";
import { ErrorDialog } from "@/components/ui/ErrorDialog";
import { UsersCard } from "@/components/admin/users/UsersCard";
import { SectionTitle } from "@/components/header/SectionTitle";
import { SearchInput } from "@/components/ui/SearchInput";
import { FilterSelection } from "@/components/ui/FilterSelecction";
import { roleOptions } from "@/lib/constant";

const UsersList = () => {
  const { page, limit, q, sort, role, setPage, setQ, setSort, setRole } =
    useQueryStore();
  const debouncedQ = useDebounce(q, 500);
  const { data, isLoading, isError, refetch } = useUsersQuery({
    q: debouncedQ,
    sort,
    role,
    page,
    limit,
  });

  const users = data?.data || [];
  const pagination = data?.pagination;

  return (
    <section className="section">
      <SectionTitle
        title="User List"
        description="Manage all registered users efficiently."
      />

      <div className="flex flex-col md:flex-row md:items-center md:justify-between gap-4 mt-4">
        <SearchInput
          q={q}
          setQ={setQ}
          setPage={setPage}
          placeholder={"search by name or email"}
        />
        <FilterSelection
          value={role}
          onChange={setRole}
          options={roleOptions}
        />
      </div>

      <Card className="border shadow-sm">
        <CardContent className="overflow-x-auto p-0">
          {isLoading ? (
            <Loading />
          ) : isError ? (
            <ErrorDialog onRetry={refetch} />
          ) : users?.length === 0 ? (
            <div className="py-12 text-center text-muted-foreground text-sm">
              No users found{q && ` for “${q}”`}
            </div>
          ) : (
            <UsersCard users={users} sort={sort} setSort={setSort} />
          )}

          {pagination && (
            <Pagination
              page={pagination.page}
              onPageChange={setPage}
              limit={pagination.limit}
              total={pagination.totalRows}
            />
          )}
        </CardContent>
      </Card>
    </section>
  );
};

export default UsersList;
