import React from "react";
import { useNavigate } from "react-router-dom";
import { useInstructorsQuery } from "@/hooks/useInstructor";
import { Loading } from "@/components/ui/Loading";
import { ErrorDialog } from "@/components/ui/ErrorDialog";
import { Badge } from "@/components/ui/badge";
import { Button } from "@/components/ui/button";
import { Star, Pencil, Trash2, Plus } from "lucide-react";

const InstructorsList = () => {
  const navigate = useNavigate();
  const {
    data: instructors = [],
    isLoading,
    isError,
    refetch,
  } = useInstructorsQuery();

  if (isLoading) return <Loading />;
  if (isError) return <ErrorDialog onRetry={refetch} />;

  return (
    <section className="max-w-7xl mx-auto px-6 py-10 space-y-6">
      {/* Header + Action */}
      <div className="flex justify-between items-center mb-6">
        <h2 className="text-2xl font-bold">Daftar Instruktur</h2>
        <Button onClick={() => navigate("/admin/instructors/add")}>
          <Plus className="w-4 h-4 mr-2" />
          Tambah Instruktur
        </Button>
      </div>

      {/* Instructors Grid */}
      <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6">
        {instructors.map((inst) => (
          <div
            key={inst.id}
            className="bg-white shadow-sm border rounded-xl overflow-hidden flex flex-col"
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
                  {inst.experience} tahun pengalaman
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
                <span className="font-medium text-gray-800">Spesialisasi:</span>{" "}
                {inst.specialties}
              </p>
              <p>
                <span className="font-medium text-gray-800">Sertifikasi:</span>{" "}
                {inst.certifications}
              </p>
              <p>
                <span className="font-medium text-gray-800">Total Kelas:</span>{" "}
                <Badge variant="outline">{inst.totalClass}</Badge>
              </p>
            </div>

            <div className="border-t p-4 flex justify-end gap-2">
              <Button
                variant="outline"
                size="sm"
                onClick={() => navigate(`/instructors/edit/${inst.id}`)}
              >
                <Pencil className="w-4 h-4 mr-1" />
                Edit
              </Button>
              <Button
                variant="destructive"
                size="sm"
                onClick={() => {
                  // Trigger modal/hook delete here
                  console.log("Delete", inst.id);
                }}
              >
                <Trash2 className="w-4 h-4 mr-1" />
                Delete
              </Button>
            </div>
          </div>
        ))}
      </div>
    </section>
  );
};

export default InstructorsList;
