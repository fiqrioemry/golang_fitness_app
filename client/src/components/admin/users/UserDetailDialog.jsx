import {
  Dialog,
  DialogTrigger,
  DialogContent,
  DialogHeader,
  DialogTitle,
  DialogClose,
  DialogDescription,
} from "@/components/ui/Dialog";
import { formatDateTime } from "@/lib/utils";
import { Badge } from "@/components/ui/Badge";
import { Button } from "@/components/ui/Button";
import { Skeleton } from "@/components/ui/Skeleton";
import { useUserDetailQuery } from "@/hooks/useUsers";

const UserDetailDialog = ({ userId, trigger }) => {
  const { data, isLoading, isError } = useUserDetailQuery(userId);

  return (
    <Dialog>
      <DialogTrigger asChild>{trigger}</DialogTrigger>
      <DialogContent className="max-w-md">
        <DialogHeader>
          <DialogTitle>User Details</DialogTitle>
          <DialogDescription>
            Full information about the selected user.
          </DialogDescription>
        </DialogHeader>

        {isLoading ? (
          <div className="space-y-3">
            <Skeleton className="w-full h-6" />
            <Skeleton className="w-full h-6" />
            <Skeleton className="w-full h-6" />
          </div>
        ) : isError ? (
          <div className="text-red-500 text-sm">
            Failed to load user details.
          </div>
        ) : (
          <div className="space-y-4">
            <div className="flex items-center gap-4">
              <img
                src={data.avatar}
                alt={data.fullname}
                className="w-16 h-16 rounded-full object-cover border"
              />
              <div>
                <h3 className="text-lg font-semibold">{data.fullname}</h3>
                <p className="text-sm text-muted-foreground">{data.email}</p>
                <Badge variant="outline">{data.role}</Badge>
              </div>
            </div>

            <div className="text-sm space-y-2">
              <p>
                <span className="font-medium">Phone:</span> {data.phone || "-"}
              </p>
              <p>
                <span className="font-medium">Gender:</span>{" "}
                {data.gender || "-"}
              </p>
              <p>
                <span className="font-medium">Bio:</span> {data.bio || "-"}
              </p>
              <p>
                <span className="font-medium">Created:</span>{" "}
                {formatDateTime(data.createdAt)}
              </p>
              <p>
                <span className="font-medium">Updated:</span>{" "}
                {formatDateTime(data.updatedAt)}
              </p>
            </div>
          </div>
        )}

        <DialogClose asChild>
          <Button variant="outline" className="w-full mt-4">
            Close
          </Button>
        </DialogClose>
      </DialogContent>
    </Dialog>
  );
};

export { UserDetailDialog };
