import { Star } from "lucide-react";
import { Badge } from "@/components/ui/badge";
import { Loading } from "@/components/ui/Loading";
import { ErrorDialog } from "@/components/ui/ErrorDialog";
import { useInstructorsQuery } from "@/hooks/useInstructor";
import { EditInstructor } from "@/components/admin/instructors/EditInstructor";
import { DeleteInstructor } from "@/components/admin/instructors/DeleteInstructor";
import { AddInstructor } from "@/components/admin/instructors/AddInstructor";

const InstructorsList = () => {
  const {
    data: instructors = [],
    isLoading,
    isError,
    refetch,
  } = useInstructorsQuery();

  if (isLoading) return <Loading />;

  if (isError) return <ErrorDialog onRetry={refetch} />;

  return (
    <section className="section">
      {/* Header */}
      <div className="space-y-1 text-center">
        <h2 className="text-2xl font-bold">Instructors List</h2>
        <p className="text-muted-foreground text-sm">
          Manage all instructors registered on the platform, including their
          expertise, certifications, and teaching activity.
        </p>
      </div>

      <div className="flex justify-end">
        <AddInstructor />
      </div>

      {/* Instructors Grid */}
      <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6">
        {instructors.map((inst) => (
          <div
            key={inst.id}
            className="bg-background shadow-sm border rounded-xl overflow-hidden flex flex-col"
          >
            <div className="p-4 flex gap-4 items-start">
              <img
                src={inst.avatar}
                alt={inst.fullname}
                className="w-16 h-16 rounded-full object-cover border"
              />
              <div className="flex-1">
                <h3 className="font-semibold text-lg">{inst.fullname}</h3>
                <p className="text-sm text-muted-foreground">
                  {inst.experience} years of experience
                </p>
                <div className="mt-1 flex items-center gap-1 text-yellow-500 text-sm">
                  {[...Array(inst.rating)].map((_, i) => (
                    <Star key={i} className="w-4 h-4 fill-yellow-500" />
                  ))}
                  <span className="text-xs text-muted-foreground ml-1">
                    ({inst.rating})
                  </span>
                </div>
              </div>
            </div>

            <div className="border-t px-4 py-3 space-y-1 text-sm text-muted-foreground">
              <p>
                <span className="font-medium text-muted-foreground">
                  Specialties :
                </span>{" "}
                {inst.specialties}
              </p>
              <p>
                <span className="font-medium text-muted-foreground">
                  Certifications :
                </span>{" "}
                {inst.certifications}
              </p>
              <p>
                <span className="font-medium text-muted-foreground">
                  Total Classes :
                </span>{" "}
                <Badge variant="outline">{inst.totalClass}</Badge>
              </p>
            </div>

            <div className="border-t p-4 flex justify-end gap-2">
              <EditInstructor instructor={inst} />
              <DeleteInstructor instructor={inst} />
            </div>
          </div>
        ))}
      </div>
    </section>
  );
};

export default InstructorsList;
